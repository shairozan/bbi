package parser

import (
	"testing"

	"github.com/metrumresearchgroup/bbi/utils"

	"github.com/metrumresearchgroup/wrapt"
)

var _ = /* RunDetails02 */ "../../testdata/2.lst"

var _ = /* RunDetails02Results */ RunDetails{
	Version:             "7.4.3",
	RunStart:            "Fri Jul 12 09:27:14 EDT 2019",
	RunEnd:              "Fri Jul 12 09:27:20 EDT 2019",
	EstimationTime:      []float64{0.68},
	CovarianceTime:      []float64{0.02},
	CpuTime:             10.5, // this is made up, not for an actual run output
	FunctionEvaluations: 178,
	SignificantDigits:   3.1,
	ProblemText:         "1 model, 1 comp",
	EstimationMethods:   []string{"First Order Conditional Estimation with Interaction"},
	DataSet:             "../../derived/mock1.csv",
	NumberOfSubjects:    50,
	NumberOfObs:         442,
	NumberOfDataRecords: 492,
	OutputTables:        []string{""},
	OutputFilesUsed:     []string{""},
}

func TestParseRunDetails(tt *testing.T) {
	baseInput := []string{
		"Days until program expires : 122",
		"1NONLINEAR MIXED EFFECTS MODEL PROGRAM (NONMEM) VERSION 7.2.0",
		" ORIGINALLY DEVELOPED BY STUART BEAL, LEWIS SHEINER, AND ALISON BOECKMANN",
		" #TERM:",
		"0MINIMIZATION SUCCESSFUL",
		" NO. OF FUNCTION EVALUATIONS USED:      352",
		" NO. OF SIG. DIGITS IN FINAL EST.:  3.4",
		"",
		"#TERE:",
		"Elapsed estimation  time in seconds:     6.84",
		"Elapsed covariance  time in seconds:     3.34",
		"Elapsed postprocess time in seconds:     0.0",
		"This file was created using /opt/NONMEM/nm72g/run/nmfe72",
		"Started  Tue Dec 17 18:10:55 2013",
		"Finished Tue Dec 17 18:11:32 2013",
		"$PROB 3.mod, double inital estimates",
		"",
		"#METH: First Order Conditional Estimation with Interaction",
		"$DATA ../../derived/mock1.csv IGNORE=C",
		"TOT. NO. OF INDIVIDUALS:       50",
		"TOT. NO. OF OBS RECS:      442",
		"NO. OF DATA RECS IN DATA SET:      492",
		"$TABLE NOPRINT ONEHEADER FILE=./1.tab",
	}

	infInput := make([]string, len(baseInput))
	copy(infInput, baseInput)
	infInput[21] = "TOT. NO. OF DATA RECS:      492"

	fullProbInput := make([]string, len(baseInput))
	copy(fullProbInput, baseInput)
	fullProbInput[15] = "$PROBLEM 3.mod, double inital estimates"

	RunDetails01Results := RunDetails{
		Version:             "7.2.0",
		RunStart:            "Tue Dec 17 18:10:55 2013",
		RunEnd:              "Tue Dec 17 18:11:32 2013",
		EstimationTime:      []float64{6.84},
		CovarianceTime:      []float64{3.34},
		PostprocessTime:     0.0,
		CpuTime:             DefaultFloat64, // not specified in test RunDetails01
		FunctionEvaluations: 352,
		SignificantDigits:   3.4,
		ProblemText:         "3.mod, double inital estimates",
		ModFile:             "-999999999",
		EstimationMethods:   []string{"First Order Conditional Estimation with Interaction"},
		DataSet:             "../../derived/mock1.csv",
		NumberOfSubjects:    50,
		NumberOfObs:         442,
		NumberOfDataRecords: 492,
		OutputTables:        []string{},
		OutputFilesUsed:     []string{},
	}

	RunDetails02Results := NewRunDetails()
	RunDetails02Results.ProblemText = "Simulation"
	RunDetails02Results.OnlySim = true

	tests := []struct {
		name     string
		input    []string
		expected RunDetails
	}{
		{
			name:     "RunDetails01",
			input:    baseInput,
			expected: RunDetails01Results,
		},
		{
			name:     "RunDetailsInfn",
			input:    infInput,
			expected: RunDetails01Results,
		},
		{
			name:     "RunDetailFullProb",
			input:    fullProbInput,
			expected: RunDetails01Results,
		},
		{
			name: "RunDetailsOnlySimShort",
			input: []string{
				"$PROB Simulation",
				"$SIMULATION ONLY (123456)",
			},
			expected: RunDetails02Results,
		},
		{
			name: "RunDetailsOnlySimMid",
			input: []string{
				"$PROB Simulation",
				"$SIMULATION ONLYSIM (123456)",
			},
			expected: RunDetails02Results,
		},
		{
			name: "RunDetailsOnlySimLong",
			input: []string{
				"$PROB Simulation",
				"$SIMULATION ONLYSIMULATION (123456)",
			},
			expected: RunDetails02Results,
		},
	}

	testId := "UNIT-NMP-035"
	for _, test := range tests {
		tt.Run(utils.AddTestId(test.name, testId), func(tt *testing.T) {
			t := wrapt.WrapT(tt)

			got := ParseRunDetails(test.input)

			t.R.Equal(test.expected, got)
		})
	}
}
