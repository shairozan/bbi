package parser

import (
	"testing"

	"github.com/metrumresearchgroup/wrapt"
)

func TestReadExt(tt *testing.T) {
	var tests = []struct {
		name string
		lines   []string
	}{
		{
			name: "ext test",
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Goal Function=MINIMUM VALUE OF OBJECTIVE FUNCTION: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" ITERATION    THETA1       THETA2       THETA3       THETA4       THETA5       THETA6       THETA7       THETA8       THETA9       SIGMA(1,1)   OMEGA(1,1)   OMEGA(2,1)   OMEGA(2,2)   OMEGA(3,1)   OMEGA(3,2)   OMEGA(3,3)   OBJ",
				"		   0  2.60000E+01  2.80000E+02  3.00000E+02  6.00000E+01  1.50000E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.50000E-03  1.00000E-01  0.00000E+00  4.00000E-02  0.00000E+00  0.00000E+00  1.00000E-02   -14341.129441443112",
				"		   5  2.68037E+01  2.86491E+02  2.95392E+02  5.80377E+01  1.52246E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.44324E-03  1.01091E-01  0.00000E+00  3.55313E-02  0.00000E+00  0.00000E+00  1.07235E-02   -14344.687671015432",
				"		  10  2.65561E+01  2.83390E+02  2.96541E+02  5.85919E+01  1.51142E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.43804E-03  1.02209E-01  0.00000E+00  3.55773E-02  0.00000E+00  0.00000E+00  1.14589E-02   -14345.851399127478",
				"		  15  2.64905E+01  2.82616E+02  2.97043E+02  5.87490E+01  1.50950E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.45104E-03  1.00611E-01  0.00000E+00  3.59990E-02  0.00000E+00  0.00000E+00  1.11726E-02   -14346.006029685614",
				"  -1000000000  2.64905E+01  2.82616E+02  2.97043E+02  5.87490E+01  1.50950E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.45104E-03  1.00611E-01  0.00000E+00  3.59990E-02  0.00000E+00  0.00000E+00  1.11726E-02   -14346.006029685614",
				"  -1000000001  6.09080E-01  4.40586E+00  2.32933E+00  1.01296E+00  2.12776E-02  1.00000E+10  1.00000E+10  1.00000E+10  1.00000E+10  7.34053E-05  9.59481E-03  1.00000E+10  3.53287E-03  1.00000E+10  1.00000E+10  1.86447E-03    0.0000000000000000",
				"  -1000000002  2.53088E-01  4.28969E-01  6.18549E-01  7.76208E-01  8.78956E-01  9.30257E-01  1.19431E+00  1.30593E+00  2.61373E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000003  1.03274E+01  2.53088E-01  2.61373E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000004  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  4.95080E-02  3.17192E-01  0.00000E+00  1.89734E-01  0.00000E+00  0.00000E+00  1.05700E-01    0.0000000000000000",
				"  -1000000005  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  7.41349E-04  1.51246E-02  1.00000E+10  9.31006E-03  1.00000E+10  1.00000E+10  8.81961E-03    0.0000000000000000",
				"  -1000000006  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  1.00000E+00  1.00000E+00  1.00000E+00  1.00000E+00  0.00000E+00  0.00000E+00  1.00000E+00  0.00000E+00  1.00000E+00  1.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000007  0.00000E+00  3.70000E+01  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000008 -3.23564E-05  6.43530E-06 -3.23355E-05  5.31805E-04 -1.53493E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  6.02637E-01  3.65275E-02  0.00000E+00  1.60995E-01  0.00000E+00  0.00000E+00  1.32073E-01    0.0000000000000000",

				"TABLE NO.     2: First Order Conditional Estimation with Interaction: Goal Function=MINIMUM VALUE OF OBJECTIVE FUNCTION: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" ITERATION    THETA1       THETA2       THETA3       THETA4       THETA5       THETA6       THETA7       THETA8       THETA9       SIGMA(1,1)   OMEGA(1,1)   OMEGA(2,1)   OMEGA(2,2)   OMEGA(3,1)   OMEGA(3,2)   OMEGA(3,3)   OBJ",
				"		   0  2.60000E+01  2.80000E+02  3.00000E+02  6.00000E+01  1.50000E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.50000E-03  1.00000E-01  0.00000E+00  4.00000E-02  0.00000E+00  0.00000E+00  1.00000E-02   -14341.129441443112",
				"		   5  2.68037E+01  2.86491E+02  2.95392E+02  5.80377E+01  1.52246E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.44324E-03  1.01091E-01  0.00000E+00  3.55313E-02  0.00000E+00  0.00000E+00  1.07235E-02   -14344.687671015432",
				"		  10  2.65561E+01  2.83390E+02  2.96541E+02  5.85919E+01  1.51142E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.43804E-03  1.02209E-01  0.00000E+00  3.55773E-02  0.00000E+00  0.00000E+00  1.14589E-02   -14345.851399127478",
				"		  15  2.64905E+01  2.82616E+02  2.97043E+02  5.87490E+01  1.50950E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.45104E-03  1.00611E-01  0.00000E+00  3.59990E-02  0.00000E+00  0.00000E+00  1.11726E-02   -14346.006029685614",
				"  -1000000000  1.00000E+00  2.82616E+02  2.97043E+02  5.87490E+01  1.50950E+00  7.50000E-01  1.00000E+00  1.00000E+00  7.50000E-01  2.45104E-03  1.00611E-01  0.00000E+00  3.59990E-02  0.00000E+00  0.00000E+00  1.11726E-02   -14346.006029685614",
				"  -1000000001  2.00000E+00  4.40586E+00  2.32933E+00  1.01296E+00  2.12776E-02  1.00000E+10  1.00000E+10  1.00000E+10  1.00000E+10  7.34053E-05  9.59481E-03  1.00000E+10  3.53287E-03  1.00000E+10  1.00000E+10  1.86447E-03    0.0000000000000000",
				"  -1000000002  3.00000E+00  4.28969E-01  6.18549E-01  7.76208E-01  8.78956E-01  9.30257E-01  1.19431E+00  1.30593E+00  2.61373E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000003  4.00000E+00  2.53088E-01  2.61373E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000004  5.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  4.95080E-02  3.17192E-01  0.00000E+00  1.89734E-01  0.00000E+00  0.00000E+00  1.05700E-01    0.0000000000000000",
				"  -1000000005  6.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  7.41349E-04  1.51246E-02  1.00000E+10  9.31006E-03  1.00000E+10  1.00000E+10  8.81961E-03    0.0000000000000000",
				"  -1000000006  7.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  1.00000E+00  1.00000E+00  1.00000E+00  1.00000E+00  0.00000E+00  0.00000E+00  1.00000E+00  0.00000E+00  1.00000E+00  1.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000007  8.00000E+00  3.70000E+01  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00    0.0000000000000000",
				"  -1000000008  9.00000E+00  6.43530E-06 -3.23355E-05  5.31805E-04 -1.53493E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  6.02637E-01  3.65275E-02  0.00000E+00  1.60995E-01  0.00000E+00  0.00000E+00  1.32073E-01    0.0000000000000000",
			},
		},
	}

	for _, test := range tests {
		tt.Run(test.name, func(tt *testing.T) {
			t := wrapt.WrapT(tt)

			ext := ParseExtLines(test.lines)
			t.A.Equal(2, len(ext.EstimationMethods))

			pd, pn := ParseParamsExt(ext)

			t.A.Equal(2, len(pd))
			t.A.Equal(test.lines[0], pd[0].Method)
			t.A.Equal(test.lines[15], pd[1].Method)

			t.A.Equal(9, len(pn.Theta))
			t.A.Equal(26.4905, pd[0].Estimates.Theta[0])
			t.A.Equal(0.75, pd[0].Estimates.Theta[len(pd[0].Estimates.Theta)-1])

			t.A.Equal(6, len(pn.Omega))
			t.A.Equal(0.100611, pd[0].Estimates.Omega[0])
			t.A.Equal(0.0111726, pd[0].Estimates.Omega[len(pd[0].Estimates.Omega)-1])

			t.A.Equal(1, len(pn.Sigma))
			t.A.Equal(0.00245104, pd[0].Estimates.Sigma[0])

			t.A.Equal(1.0, pd[1].Estimates.Theta[0])
			t.A.Equal(2.0, pd[1].StdErr.Theta[0])
			t.A.Equal(0.317192, pd[1].RandomEffectSD.Omega[0])
			t.A.Equal(0.0151246, pd[1].RandomEffectSDSE.Omega[0])
			t.A.Equal(7.0, pd[1].Fixed.Theta[0])
		})
	}
}
