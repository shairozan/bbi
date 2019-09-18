package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadParseGrdLines(t *testing.T) {
	var tests = []struct {
		lines   []string
		context string
	}{
		{
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" ITERATION    GRD(1)       GRD(2)       GRD(3)       GRD(4)       GRD(5)       GRD(6)       GRD(7)       GRD(8)       GRD(9)",
				"		   0 -3.85613E+01  4.33258E+01  2.33602E+02  2.12304E+02  2.12964E+02 -5.05978E+00  2.86950E+01 -1.58137E+01  6.81739E+01",
				"   	   5  8.77273E+01  1.21734E+02 -7.66232E+01 -3.84937E+01  1.10997E+01  1.49415E+00 -7.95547E-01 -7.83447E+00 -1.75762E+01",
				"		  10  4.84466E+01  5.77126E+01 -3.75199E+01 -9.45384E+00 -3.79759E+00  6.11768E+00 -3.10524E+00  3.68255E+00 -2.02148E+01",
				"		  15  1.01544E-02  1.25236E-03  1.93608E-02 -6.25415E-02  4.45540E-02 -1.47003E-02 -2.31826E-02 -5.90237E-03 -5.86988E-03",
				"",
				"TABLE NO.     2: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" ITERATION    GRD(1)       GRD(2)       GRD(3)       GRD(4)       GRD(5)       GRD(6)       GRD(7)       GRD(8)       GRD(9)",
				"		   0 -3.85613E+01  4.33258E+01  2.33602E+02  2.12304E+02  2.12964E+02 -5.05978E+00  2.86950E+01 -1.58137E+01  6.81739E+01",
				"   	   5  8.77273E+01  1.21734E+02 -7.66232E+01 -3.84937E+01  1.10997E+01  1.49415E+00 -7.95547E-01 -7.83447E+00 -1.75762E+01",
				"		  10  4.84466E+01  5.77126E+01 -3.75199E+01 -9.45384E+00 -3.79759E+00  6.11768E+00 -3.10524E+00  3.68255E+00 -2.02148E+01",
				"		  15  1.01544E-02  1.25236E-03  1.93608E-02 -6.25415E-02  4.45540E-02 -1.47003E-02 -2.31826E-02 -5.90237E-03 -5.86988E-03",
				"",
				"TABLE NO.     3: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" ITERATION    GRD(1)       GRD(2)       GRD(3)       GRD(4)       GRD(5)       GRD(6)       GRD(7)       GRD(8)       GRD(9)",
				"		   0 -3.85613E+01  4.33258E+01  2.33602E+02  2.12304E+02  2.12964E+02 -5.05978E+00  2.86950E+01 -1.58137E+01  6.81739E+01",
				"   	   5  8.77273E+01  1.21734E+02 -7.66232E+01 -3.84937E+01  1.10997E+01  1.49415E+00 -7.95547E-01 -7.83447E+00 -1.75762E+01",
				"		  10  4.84466E+01  5.77126E+01 -3.75199E+01 -9.45384E+00 -3.79759E+00  6.11768E+00 -3.10524E+00  3.68255E+00 -2.02148E+01",
				"		  15  1.01544E-02  1.25236E-03  1.93608E-02 -6.25415E-02  4.45540E-02 -1.47003E-02 -2.31826E-02 -5.90237E-03 -5.86988E-03",
			},

			context: "no zero",
		},
	}

	for _, tt := range tests {
		extData := ParseGrdLines(tt.lines)
		assert.Equal(t, tt.lines[0], extData.EstimationMethods[0], "Fail :"+tt.context)
		assert.Equal(t, "GRD(1)", extData.ParameterNames[1], "Fail :"+tt.context)
		assert.Equal(t, strings.Trim(tt.lines[2], "\t "), extData.EstimationLines[0][0], "Fail :"+tt.context)

		parametersData, parameterNames := ParseGrdData(extData)
		assert.Equal(t, tt.lines[0], parametersData[0].Method, "Fail :"+tt.context)
		assert.Equal(t, "GRD(1)", parameterNames.Theta[0], "Fail :"+tt.context)
	}
}
