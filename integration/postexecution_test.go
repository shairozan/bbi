package bbitest

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/metrumresearchgroup/bbi/utils"

	"github.com/metrumresearchgroup/wrapt"
	"github.com/spf13/afero"
)

// Will need to set a custom env for execution.
const postExecutionScriptString string = `#!/bin/bash

env > ${BBI_ROOT_EXECUTION_DIR}/working/${BBI_SCENARIO}/${BBI_MODEL_FILENAME}.out
`

func generatePostWorkEnvsString(content map[string]string) (string, error) {
	// "--additional_post_work_envs=\"SCENARIO=" + v.identifier + "\",ROOT_EXECUTION_DIR=\"" + ROOT_EXECUTION_DIR + "\"",

	const templateString string = `--additional_post_work_envs="{{ range $key, $value := . }}{{$key}}={{$value}},{{end}}"`

	t, err := template.New("pieces").Parse(templateString)

	if err != nil {
		return "", err
	}

	outBuffer := new(bytes.Buffer)

	err = t.Execute(outBuffer, content)

	if err != nil {
		return "", err
	}

	stringResult := outBuffer.String()

	stringResult = strings.Replace(stringResult, `,"`, `"`, 1)

	return stringResult, nil
}

func TestKVPExpansion(tt *testing.T) {
	testId := "INT-POSTEX-001"
	tt.Run(utils.AddTestId("", testId), func(tt *testing.T) {
		t := wrapt.WrapT(tt)

		mapdata := make(map[string]string)
		mapdata["SCENARIO"] = "240"
		mapdata["ROOT_EXECUTION_DIR"] = "/data/one"

		generated, err := generatePostWorkEnvsString(mapdata)

		t.R.NoError(err)
		t.R.NotNil(generated)
	})
}

func TestPostExecutionSucceeds(tt *testing.T) {
	var scenarios []*Scenario
	testId := "INT-POSTEX-002"

	func() {
		tt.Run(utils.AddTestId("", testId), func(tt *testing.T) {
			t := wrapt.WrapT(tt)

			var err error
			// Skip the test if the flag isn't enabled
			if !FeatureEnabled("POST_EXECUTION") {
				t.Skip("Post execution not enabled as far as testing is concerned")
			}

			for _, name := range []string{
				"acop",
			} {
				scenario := InitializeScenario(t, name)
				scenarios = append(scenarios, scenario)
			}

			err = ioutil.WriteFile(filepath.Join(ROOT_EXECUTION_DIR, "post.sh"), []byte(postExecutionScriptString), 0755)
			t.R.NoError(err)
		})
	}()

	for _, scenario := range scenarios {
		tt.Run(utils.AddTestId(scenario.identifier, testId), func(tt *testing.T) {
			t := wrapt.WrapT(tt)

			scenario.Prepare(t, context.Background())

			arguments := []string{
				"-d",
				"nonmem",
				"--nm_version",
				os.Getenv("NMVERSION"),
				"run",
				"local",
				"--overwrite=true",
				"--post_work_executable",
				filepath.Join(ROOT_EXECUTION_DIR, "post.sh"),
				"--additional_post_work_envs=\"BBI_ROOT_EXECUTION_DIR=" + ROOT_EXECUTION_DIR + " BBI_SCENARIO=" + scenario.identifier + "\"",
			}

			// Do the actual execution
			for _, m := range scenario.models {
				tt.Run(scenario.identifier+"_post_execution", func(tt *testing.T) {
					t := wrapt.WrapT(tt)

					var output string
					output, err := m.Execute(scenario, arguments...)
					t.R.NoError(err)

					nmd := NonMemTestingDetails{
						OutputDir: filepath.Join(scenario.Workpath, m.identifier),
						Model:     m,
						Output:    output,
					}

					AssertNonMemCompleted(t, nmd)
					AssertNonMemCreatedOutputFiles(t, nmd)
					AssertNonMemCleanedUpFiles(t, nmd)

					var exists bool
					exists, err = afero.Exists(afero.NewOsFs(), filepath.Join(ROOT_EXECUTION_DIR, "working", scenario.identifier, m.identifier+".out"))
					t.R.NoError(err)
					t.R.True(exists)

					lines := func() []string {
						// Does the file contain the expected Details:
						// SCENARIO (Additional provided value)
						var file *os.File
						file, err = os.Open(filepath.Join(ROOT_EXECUTION_DIR, "working", scenario.identifier, m.identifier+".out"))
						t.R.NoError(err)
						defer file.Close()

						scanner := bufio.NewScanner(file)
						// scanner.Split(bufio.ScanLines)

						var lines []string
						for scanner.Scan() {
							lines = append(lines, scanner.Text())
						}

						return lines
					}()

					t.R.True(doesOutputFileContainKeyWithValue(lines, "BBI_MODEL", m.filename))
					t.R.True(doesOutputFileContainKeyWithValue(lines, "BBI_MODEL_FILENAME", m.identifier))
					t.R.True(doesOutputFileContainKeyWithValue(lines, "BBI_MODEL_EXT", strings.Replace(m.extension, ".", "", 1)))
					t.R.True(doesOutputFileContainKeyWithValue(lines, "BBI_SUCCESSFUL", "true"))
					t.R.True(doesOutputFileContainKeyWithValue(lines, "BBI_ERROR", ""))
				})
			}
		})
	}

	// Test a scenario for the first scenario where we force failure. Model is deleted (not found)
	tt.Run(utils.AddTestId("verify_failure_results", testId), func(tt *testing.T) {
		t := wrapt.WrapT(tt)

		scenario := scenarios[0]
		scenario.Prepare(t, context.Background())

		arguments := []string{
			"nonmem",
			"--nm_version",
			os.Getenv("NMVERSION"),
			"run",
			"local",
			"--post_work_executable",
			filepath.Join(ROOT_EXECUTION_DIR, "post.sh"),
			"--overwrite=false",
			// `--additional_post_work_envs "BBI_SCENARIO=` + scenario.identifier + ` BBI_ROOT_EXECUTION_DIR=` + ROOT_EXECUTION_DIR  + `"`,
			// "--additional_post_work_envs BBI_ROOT_EXECUTION_DIR=" + ROOT_EXECUTION_DIR,
		}

		// Removing the model won't do anything. Execute with overwrite = false?
		for _, v := range scenario.models {
			tt.Run(utils.AddTestId(v.identifier, testId), func(tt *testing.T) {
				t := wrapt.WrapT(tt)

				var err error

				err = os.Setenv("BBI_ADDITIONAL_POST_WORK_ENVS", `BBI_SCENARIO=`+scenario.identifier+` BBI_ROOT_EXECUTION_DIR=`+ROOT_EXECUTION_DIR)
				t.R.NoError(err)

				err = os.Remove(filepath.Join(scenario.Workpath, v.identifier+".out"))
				t.R.NoError(err)

				var output string
				output, err = v.Execute(scenario, arguments...)
				t.R.Error(err)

				lines := func() []string {
					var lines []string
					// Does the file contain the expected Details:
					// SCENARIO (Additional provided value)
					var file *os.File
					file, err = os.Open(filepath.Join(ROOT_EXECUTION_DIR, "working", scenario.identifier, v.identifier+".out"))
					t.R.NoError(err)
					defer file.Close()

					scanner := bufio.NewScanner(file)
					// scanner.Split(bufio.ScanLines)

					for scanner.Scan() {
						lines = append(lines, scanner.Text())
					}

					return lines
				}()

				t.R.True(doesOutputFileContainKeyWithValue(lines, "BBI_SUCCESSFUL", "false"))
				if err != nil {
					t.R.True(doesExecutionOutputContainErrorString(err.Error(), output))
				}
			})
		}
	})
}

func doesOutputFileContainKeyWithValue(lines []string, key string, value string) bool {
	for _, v := range lines {
		if strings.Contains(v, key+"=") {
			components := strings.Split(v, "=")

			return components[0] == key && components[1] == value
		}
	}

	return false
}

func doesExecutionOutputContainErrorString(line string, output string) bool {
	lines := strings.Split(output, "\n")

	for _, v := range lines {
		if strings.Contains(line, v) {
			// We have a match
			return true
		}
	}

	return false
}
