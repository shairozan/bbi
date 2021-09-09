package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadParseCovLines(t *testing.T) {
	var tests = []struct {
		lines    []string
		context  string
		nMethods int
		nThetas  int
	}{
		{
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" NAME         THETA1       THETA2       THETA3       THETA4       THETA5       THETA6       THETA7       THETA8       THETA9      ",
				" THETA1        3.70979E-01 -2.14341E-01 -1.50899E-01 -2.41113E-02  1.04312E-03  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA2       -2.14341E-01  1.94116E+01 -3.53392E+00 -1.99270E+00  3.64529E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA3       -1.50899E-01 -3.53392E+00  5.42576E+00  1.44045E+00 -2.83908E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA4       -2.41113E-02 -1.99270E+00  1.44045E+00  1.02609E+00 -1.57357E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA5        1.04312E-03  3.64529E-02 -2.83908E-02 -1.57357E-02  4.52737E-04  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA6        6.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA7        7.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA8        8.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA9        9.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  9.90000E+00",
			},
			context:  "only thetas",
			nMethods: 1,
			nThetas:  9,
		},
		{
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" NAME         THETA1       THETA2       THETA3       THETA4       THETA5       THETA6       THETA7       THETA8       THETA9       SIGMA(1,1)   OMEGA(1,1)   OMEGA(2,1)   OMEGA(2,2)   OMEGA(3,1)   OMEGA(3,2)   OMEGA(3,3)",
				" THETA1        3.70979E-01 -2.14341E-01 -1.50899E-01 -2.41113E-02  1.04312E-03  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.69174E-06  3.87129E-04  0.00000E+00 -2.69021E-04  0.00000E+00  0.00000E+00  1.02250E-04",
				" THETA2       -2.14341E-01  1.94116E+01 -3.53392E+00 -1.99270E+00  3.64529E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  2.72226E-05 -2.38849E-03  0.00000E+00  2.19637E-03  0.00000E+00  0.00000E+00 -2.16771E-04",
				" THETA3       -1.50899E-01 -3.53392E+00  5.42576E+00  1.44045E+00 -2.83908E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -2.66410E-05 -5.20942E-04  0.00000E+00  2.75030E-04  0.00000E+00  0.00000E+00  1.78136E-04",
				" THETA4       -2.41113E-02 -1.99270E+00  1.44045E+00  1.02609E+00 -1.57357E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -8.57217E-06  7.63109E-04  0.00000E+00  7.27702E-05  0.00000E+00  0.00000E+00 -5.31316E-05",
				" THETA5        1.04312E-03  3.64529E-02 -2.83908E-02 -1.57357E-02  4.52737E-04  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  1.50750E-07  6.32715E-06  0.00000E+00  2.54001E-06  0.00000E+00  0.00000E+00  2.49581E-06",
				" THETA6        6.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA7        7.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA8        8.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA9        9.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  9.90000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" SIGMA(1,1)   -1.69174E-06  2.72226E-05 -2.66410E-05 -8.57217E-06  1.50750E-07  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  5.38834E-09 -1.35902E-07  0.00000E+00  2.47065E-09  0.00000E+00  0.00000E+00 -1.05448E-08",
				" OMEGA(1,1)    3.87129E-04 -2.38849E-03 -5.20942E-04  7.63109E-04  6.32715E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.35902E-07  9.20604E-05  0.00000E+00  3.55059E-06  0.00000E+00  0.00000E+00  1.58649E-06",
				" OMEGA(2,1)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(2,2)   -2.69021E-04  2.19637E-03  2.75030E-04  7.27702E-05  2.54001E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  2.47065E-09  3.55059E-06  0.00000E+00  1.24812E-05  0.00000E+00  0.00000E+00 -3.22735E-07",
				" OMEGA(3,1)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(3,2)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(3,3)    1.02250E-04 -2.16771E-04  1.78136E-04 -5.31316E-05  2.49581E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.05448E-08  1.58649E-06  0.00000E+00 -3.22735E-07  0.00000E+00  0.00000E+00  3.47626E-06",

				"TABLE NO.     2: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" NAME         THETA1       THETA2       THETA3       THETA4       THETA5       THETA6       THETA7       THETA8       THETA9       SIGMA(1,1)   OMEGA(1,1)   OMEGA(2,1)   OMEGA(2,2)   OMEGA(3,1)   OMEGA(3,2)   OMEGA(3,3)",
				" THETA1        2.00000E-00 -2.14341E-01 -1.50899E-01 -2.41113E-02  1.04312E-03  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.69174E-06  3.87129E-04  0.00000E+00 -2.69021E-04  0.00000E+00  0.00000E+00  1.02250E-04",
				" THETA2       -2.14341E-01  1.94116E+01 -3.53392E+00 -1.99270E+00  3.64529E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  2.72226E-05 -2.38849E-03  0.00000E+00  2.19637E-03  0.00000E+00  0.00000E+00 -2.16771E-04",
				" THETA3       -1.50899E-01 -3.53392E+00  5.42576E+00  1.44045E+00 -2.83908E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -2.66410E-05 -5.20942E-04  0.00000E+00  2.75030E-04  0.00000E+00  0.00000E+00  1.78136E-04",
				" THETA4       -2.41113E-02 -1.99270E+00  1.44045E+00  1.02609E+00 -1.57357E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -8.57217E-06  7.63109E-04  0.00000E+00  7.27702E-05  0.00000E+00  0.00000E+00 -5.31316E-05",
				" THETA5        1.04312E-03  3.64529E-02 -2.83908E-02 -1.57357E-02  4.52737E-04  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  1.50750E-07  6.32715E-06  0.00000E+00  2.54001E-06  0.00000E+00  0.00000E+00  2.49581E-06",
				" THETA6        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA7        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA8        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA9        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" SIGMA(1,1)   -1.69174E-06  2.72226E-05 -2.66410E-05 -8.57217E-06  1.50750E-07  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  5.38834E-09 -1.35902E-07  0.00000E+00  2.47065E-09  0.00000E+00  0.00000E+00 -1.05448E-08",
				" OMEGA(1,1)    3.87129E-04 -2.38849E-03 -5.20942E-04  7.63109E-04  6.32715E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.35902E-07  9.20604E-05  0.00000E+00  3.55059E-06  0.00000E+00  0.00000E+00  1.58649E-06",
				" OMEGA(2,1)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(2,2)   -2.69021E-04  2.19637E-03  2.75030E-04  7.27702E-05  2.54001E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  2.47065E-09  3.55059E-06  0.00000E+00  1.24812E-05  0.00000E+00  0.00000E+00 -3.22735E-07",
				" OMEGA(3,1)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(3,2)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(3,3)    1.02250E-04 -2.16771E-04  1.78136E-04 -5.31316E-05  2.49581E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.05448E-08  1.58649E-06  0.00000E+00 -3.22735E-07  0.00000E+00  0.00000E+00  3.47626E-06",

				"TABLE NO.     3: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				" NAME         THETA1       THETA2       THETA3       THETA4       THETA5       THETA6       THETA7       THETA8       THETA9       SIGMA(1,1)   OMEGA(1,1)   OMEGA(2,1)   OMEGA(2,2)   OMEGA(3,1)   OMEGA(3,2)   OMEGA(3,3)",
				" THETA1        3.00000E-00 -2.14341E-01 -1.50899E-01 -2.41113E-02  1.04312E-03  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.69174E-06  3.87129E-04  0.00000E+00 -2.69021E-04  0.00000E+00  0.00000E+00  1.02250E-04",
				" THETA2       -2.14341E-01  1.94116E+01 -3.53392E+00 -1.99270E+00  3.64529E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  2.72226E-05 -2.38849E-03  0.00000E+00  2.19637E-03  0.00000E+00  0.00000E+00 -2.16771E-04",
				" THETA3       -1.50899E-01 -3.53392E+00  5.42576E+00  1.44045E+00 -2.83908E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -2.66410E-05 -5.20942E-04  0.00000E+00  2.75030E-04  0.00000E+00  0.00000E+00  1.78136E-04",
				" THETA4       -2.41113E-02 -1.99270E+00  1.44045E+00  1.02609E+00 -1.57357E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -8.57217E-06  7.63109E-04  0.00000E+00  7.27702E-05  0.00000E+00  0.00000E+00 -5.31316E-05",
				" THETA5        1.04312E-03  3.64529E-02 -2.83908E-02 -1.57357E-02  4.52737E-04  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  1.50750E-07  6.32715E-06  0.00000E+00  2.54001E-06  0.00000E+00  0.00000E+00  2.49581E-06",
				" THETA6        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA7        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA8        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" THETA9        0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" SIGMA(1,1)   -1.69174E-06  2.72226E-05 -2.66410E-05 -8.57217E-06  1.50750E-07  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  5.38834E-09 -1.35902E-07  0.00000E+00  2.47065E-09  0.00000E+00  0.00000E+00 -1.05448E-08",
				" OMEGA(1,1)    3.87129E-04 -2.38849E-03 -5.20942E-04  7.63109E-04  6.32715E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.35902E-07  9.20604E-05  0.00000E+00  3.55059E-06  0.00000E+00  0.00000E+00  1.58649E-06",
				" OMEGA(2,1)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(2,2)   -2.69021E-04  2.19637E-03  2.75030E-04  7.27702E-05  2.54001E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  2.47065E-09  3.55059E-06  0.00000E+00  1.24812E-05  0.00000E+00  0.00000E+00 -3.22735E-07",
				" OMEGA(3,1)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(3,2)    0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				" OMEGA(3,3)    1.02250E-04 -2.16771E-04  1.78136E-04 -5.31316E-05  2.49581E-06  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00 -1.05448E-08  1.58649E-06  0.00000E+00 -3.22735E-07  0.00000E+00  0.00000E+00  3.47626E-06",
			},
			context:  "full control stream with multiple methods",
			nMethods: 3,
			nThetas:  9,
		},
	}

	for _, tt := range tests {
		extData := parseCovLines(tt.lines)
		assert.Equal(t, len(extData.EstimationMethods), tt.nMethods, "failed to extract separate tables")
		assert.Equal(t, tt.lines[0], extData.EstimationMethods[0], "Fail :"+tt.context)
		t.Log("heres a log message in EstimationMethods")
		res := GetThetaValues(tt.lines)
		for _, fa := range res {
			assert.Equal(t, -0.214341, fa.Values[9], "Fail :"+tt.context)
			assert.Equal(t, 9, fa.Dim, "Fail :"+tt.context)
		}
	}
}
