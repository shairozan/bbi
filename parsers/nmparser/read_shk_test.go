package parser

import (
	"testing"

	"github.com/metrumresearchgroup/wrapt"
)

// TODO: did not add a test for nm75 with shrinkage type 11 since these table tests are kind of a mess

func TestReadParseShkLines(tt *testing.T) {
	var tests = []struct {
		lines          []string
		expectedTables []string
		context        string
		etaCount       int
		epsCount       int
	}{
		{
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TYPE          SUBPOP       ETA(1)       ETA(2)       ETA(3)       ETA(4)",
				"            1            1 -3.67026E-03 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            2            1  3.19719E-02  2.60947E-02  1.62789E-02  1.89996E-02",
				"            3            1  9.08606E-01  8.80753E-01  6.22835E-01  9.17950E-01",
				"            4            1  3.82499E+00  2.52933E+01  5.46357E+01  3.69109E+01",
				"            5            1  1.32691E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"            6            1  3.79795E+00  2.56533E+01  5.45618E+01  3.71115E+01",
				"            7            1  7.90000E+01  7.90000E+01  7.90000E+01  7.90000E+01",
				"            8            1  7.50368E+00  4.41891E+01  7.94208E+01  6.01976E+01",
				"            9            1  7.45166E+00  4.47258E+01  7.93537E+01  6.04504E+01",
				"            10           1  2.47776E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"TABLE NO.     2: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TYPE          SUBPOP       ETA(1)       ETA(2)       ETA(3)       ETA(4)",
				"            1            1 -3.67026E-03 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            2            1  3.19719E-02  2.60947E-02  1.62789E-02  1.89996E-02",
				"            3            1  9.08606E-01  8.80753E-01  6.22835E-01  9.17950E-01",
				"            4            1  3.82499E+00  2.52933E+01  5.46357E+01  3.69109E+01",
				"            5            1  1.32691E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"            6            1  3.79795E+00  2.56533E+01  5.45618E+01  3.71115E+01",
				"            7            1  7.90000E+01  7.90000E+01  7.90000E+01  7.90000E+01",
				"            8            1  7.50368E+00  4.41891E+01  7.94208E+01  6.01976E+01",
				"            9            1  7.45166E+00  4.47258E+01  7.93537E+01  6.04504E+01",
				"            10            1  2.47776E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"TABLE NO.     3: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TYPE          SUBPOP       ETA(1)       ETA(2)       ETA(3)       ETA(4)",
				"            1            1  1.67026E-03  2.91460E-03  3.00654E-03  4.95727E-03",
				"            2            1  3.19719E-02  2.60947E-02  1.62789E-02  1.89996E-02",
				"            3            1  9.08606E-01  8.80753E-01  6.22835E-01  9.17950E-01",
				"            4            1  3.82499E+00  2.52933E+01  5.46357E+01  3.69109E+01",
				"            5            1  1.32691E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"            6            1  3.79795E+00  2.56533E+01  5.45618E+01  3.71115E+01",
				"            7            1  7.90000E+01  7.90000E+01  7.90000E+01  7.90000E+01",
				"            8            1  7.50368E+00  4.41891E+01  7.94208E+01  6.01976E+01",
				"            9            1  7.45166E+00  4.47258E+01  7.93537E+01  6.04504E+01",
				"            10           1  2.47776E+01  0.00000E+00  0.00000E+00  0.00000E+00",
			},
			expectedTables: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TABLE NO.     2: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TABLE NO.     3: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
			},
			context:  "3 methods",
			etaCount: 4,
			epsCount: 1,
		},
		{
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TYPE          SUBPOP       ETA(1)       ETA(2)       ETA(3)       ETA(4)       ETA(5)       ETA(6)       ETA(7)       ETA(8)       ETA(9)       ETA(10)      ETA(11)      ETA(12)      ETA(13)      ETA(14)      ETA(15)      ETA(16)      ETA(17)      ETA(18)      ETA(19)      ETA(20)      ETA(21)      ETA(22)      ETA(23)      ETA(24)      ETA(25)      ETA(26)      ETA(27)      ETA(28)      ETA(29)      ETA(30)      ETA(31)      ETA(32)      ETA(33)      ETA(34)      ETA(35)      ETA(36)      ETA(37)      ETA(38)      ETA(39)      ETA(40)      ETA(41)      ETA(42)      ETA(43)      ETA(44)      ETA(45)      ETA(46)      ETA(47)      ETA(48)      ETA(49)      ETA(50)      ETA(51)      ETA(52)      ETA(53)      ETA(54)      ETA(55)      ETA(56)      ETA(57)      ETA(58)      ETA(59)      ETA(60)      ETA(61)      ETA(62)      ETA(63)      ETA(64)",
				"            1            1 -6.01005E-04  1.22635E-01 -8.49248E-03  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  9.02849E-03 -8.37981E-03  9.76088E-03 -8.19207E-03 -1.33891E-03 -4.76584E-03 -6.68812E-03  9.12071E-03  2.48323E-03  9.22838E-03 -1.08256E-02 -1.17092E-02 -1.01130E-02 -5.46966E-03  2.60054E-03 -3.25626E-03 -7.62912E-03 -6.57872E-03  1.43729E-02  1.86127E-02  1.80308E-02  4.63292E-02  1.81368E-03  1.25328E-02 -8.24692E-04  2.85221E-03  8.44749E-02  2.11887E-02  4.28626E-03 -2.04620E-02 -6.93946E-03 -2.39934E-03  1.12561E-02  1.37280E-02  2.17852E-02 -6.26904E-03  2.40792E-02 -1.97701E-02 -2.11946E-02  7.93780E-03 -5.86888E-03  1.86654E-02  1.22182E-02  1.04869E-02  2.21485E-02 -1.09223E-02 -9.84011E-03 -1.17376E-02 -5.79249E-04  1.02599E-03  2.26101E-04  2.15617E-02  1.10678E-02  1.12565E-02  5.28936E-03 -1.67686E-03 -2.73584E-02",
				"            2            1  5.81447E-03  1.37570E-02  1.55023E-02  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  1.52296E-03  1.40667E-03  2.50410E-03  3.19302E-03  2.81056E-03  2.45154E-03  2.87449E-03  4.10239E-03  4.37026E-03  5.81395E-03  3.50639E-03  3.03715E-03  3.62095E-03  5.70026E-03  6.54706E-03  5.30739E-03  6.15760E-03  7.94877E-03  5.64031E-03  6.90864E-03  8.51129E-03  1.13786E-02  1.23934E-02  9.34193E-03  9.34076E-03  9.84667E-03  2.32813E-02  2.57969E-02  1.83869E-02  1.12213E-02  1.13861E-02  1.60896E-02  1.72981E-02  2.14628E-02  1.83997E-02  1.60425E-02  7.78718E-03  1.30811E-02  4.02980E-03  4.27898E-03  7.67835E-03  5.39342E-03  4.80673E-03  5.41088E-03  6.93264E-03  1.28827E-02  1.29125E-02  8.96520E-03  4.94963E-03  6.47538E-03  8.15350E-03  1.26088E-02  9.74106E-03  8.49337E-03  7.20186E-03  2.09508E-02  1.14469E-02",
				"            3            1  9.17674E-01  4.97489E-19  5.83816E-01  1.00000E+00  1.00000E+00  1.00000E+00  1.00000E+00  3.07261E-09  2.57488E-09  9.70414E-05  1.02993E-02  6.33799E-01  5.18929E-02  1.99803E-02  2.61976E-02  5.69891E-01  1.12449E-01  2.01938E-03  1.15614E-04  5.22352E-03  3.37284E-01  6.91215E-01  5.39524E-01  2.15354E-01  4.07875E-01  1.08266E-02  7.05758E-03  3.41364E-02  4.67167E-05  8.83651E-01  1.79739E-01  9.29646E-01  7.72074E-01  2.85209E-04  4.11437E-01  8.15673E-01  6.82278E-02  5.42216E-01  8.81456E-01  5.15231E-01  5.22421E-01  2.36415E-01  6.95963E-01  1.98718E-03  1.30699E-01  1.44757E-07  6.35862E-02  4.44664E-01  5.38739E-04  1.10251E-02  5.26084E-02  1.39931E-03  3.96532E-01  4.46022E-01  1.90454E-01  9.06837E-01  8.74106E-01  9.77877E-01  8.72581E-02  2.55871E-01  1.85062E-01  4.62678E-01  9.36207E-01  1.68468E-02",
				"            4            1  5.04329E+00  2.94237E+01  2.25710E+01  1.00000E+02  1.00000E+02  1.00000E+02  1.00000E+02  3.73421E+01  4.24215E+01  2.70029E+01  3.29925E+01  4.14813E+01  4.95677E+01  4.33213E+01  2.82102E+01  2.79458E+01  2.37307E+01  5.83427E+01  6.43597E+01  5.80427E+01  3.39489E+01  2.41367E+01  3.85013E+01  3.33726E+01  5.78667E+01  7.19281E+01  5.70746E+01  4.73863E+01  4.99066E+01  6.07221E+01  7.06253E+01  7.09807E+01  7.06787E+01  3.84726E+01  3.57674E+01  6.35730E+01  7.98670E+01  7.98216E+01  7.18443E+01  6.97295E+01  6.24416E+01  6.78018E+01  7.37850E+01  9.37664E+01  9.01679E+01  5.77602E+01  5.53769E+01  4.29738E+01  7.11638E+01  7.45022E+01  7.16411E+01  6.51735E+01  4.25642E+01  4.57608E+01  7.00367E+01  8.50185E+01  8.06406E+01  7.59297E+01  6.27769E+01  7.12430E+01  7.49264E+01  8.01465E+01  7.17071E+01  8.54853E+01",
				"            5            1  8.18719E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
				"            6            1  5.86549E+00  3.35183E+01  2.13559E+01  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  4.26594E+01  3.44983E+01  3.14350E+01  3.44900E+01  3.56414E+01  3.83864E+01  3.90084E+01  2.64336E+01  2.03490E+01  2.76062E+01  3.90409E+01  4.26471E+01  3.12001E+01  2.17182E+01  2.63142E+01  3.29567E+01  4.71988E+01  4.88425E+01  4.92398E+01  6.01581E+01  5.21242E+01  5.56289E+01  6.08103E+01  6.18964E+01  6.94596E+01  6.75469E+01  4.62640E+01  4.27434E+01  5.84082E+01  6.82727E+01  7.12933E+01  6.21504E+01  5.00404E+01  5.60272E+01  5.87863E+01  8.51840E+01  8.52855E+01  8.70739E+01  6.47360E+01  5.72552E+01  5.81063E+01  6.48603E+01  6.56960E+01  7.08005E+01  7.01443E+01  5.16139E+01  4.63885E+01  5.93020E+01  7.77564E+01  7.72027E+01  7.11554E+01  5.32606E+01  5.94764E+01  6.00998E+01  7.46642E+01  7.58981E+01  7.43579E+01",
				"            7            1  4.91000E+02  4.91000E+02  4.91000E+02  4.91000E+02  4.91000E+02  4.91000E+02  4.91000E+02  4.91000E+02  4.86000E+02  2.46000E+02  1.27000E+02  1.25000E+02  1.22000E+02  1.12000E+02  8.80000E+01  7.80000E+01  4.90000E+01  4.00000E+01  3.90000E+01  3.80000E+01  3.80000E+01  3.80000E+01  3.80000E+01  3.30000E+01  7.00000E+00  6.00000E+00  4.91000E+02  4.86000E+02  2.46000E+02  1.27000E+02  1.25000E+02  1.22000E+02  1.12000E+02  8.80000E+01  7.80000E+01  4.90000E+01  4.00000E+01  3.90000E+01  3.80000E+01  3.80000E+01  3.80000E+01  3.80000E+01  3.30000E+01  7.00000E+00  6.00000E+00  4.91000E+02  4.86000E+02  2.46000E+02  1.27000E+02  1.25000E+02  1.22000E+02  1.12000E+02  8.80000E+01  7.80000E+01  4.90000E+01  4.00000E+01  3.90000E+01  3.80000E+01  3.80000E+01  3.80000E+01  3.80000E+01  3.30000E+01  7.00000E+00  6.00000E+00",
				"            8            1  9.83223E+00  5.01898E+01  4.00474E+01  1.00000E+02  1.00000E+02  1.00000E+02  1.00000E+02  6.07398E+01  6.68472E+01  4.67142E+01  5.51000E+01  6.57556E+01  7.45659E+01  6.78752E+01  4.84623E+01  4.80820E+01  4.18299E+01  8.26467E+01  8.72977E+01  8.23958E+01  5.63725E+01  4.24476E+01  6.21791E+01  5.56079E+01  8.22479E+01  9.21197E+01  8.15741E+01  7.23180E+01  7.49065E+01  8.45725E+01  9.13713E+01  9.15788E+01  9.14026E+01  6.21438E+01  5.87418E+01  8.67308E+01  9.59466E+01  9.59283E+01  9.20726E+01  9.08369E+01  8.58937E+01  8.96328E+01  9.31277E+01  9.96114E+01  9.90333E+01  8.21580E+01  8.00877E+01  6.74801E+01  9.16848E+01  9.34986E+01  9.19577E+01  8.78711E+01  6.70112E+01  7.05811E+01  9.10220E+01  9.77556E+01  9.62521E+01  9.42062E+01  8.61444E+01  9.17304E+01  9.37131E+01  9.60584E+01  9.19951E+01  9.78932E+01",
				"            9            1  1.13869E+01  5.58019E+01  3.81510E+01  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  6.71205E+01  5.70953E+01  5.29884E+01  5.70845E+01  5.85797E+01  6.20376E+01  6.28003E+01  4.58798E+01  3.65572E+01  4.75913E+01  6.28399E+01  6.71065E+01  5.26658E+01  3.87196E+01  4.57041E+01  5.50519E+01  7.21203E+01  7.38291E+01  7.42340E+01  8.41262E+01  7.70791E+01  8.03120E+01  8.46417E+01  8.54812E+01  9.06728E+01  8.94680E+01  7.11244E+01  6.72169E+01  8.27013E+01  8.99338E+01  9.17593E+01  8.56741E+01  7.50404E+01  8.06639E+01  8.30143E+01  9.78049E+01  9.78348E+01  9.83292E+01  8.75645E+01  8.17289E+01  8.24492E+01  8.76520E+01  8.82323E+01  9.14739E+01  9.10864E+01  7.65878E+01  7.12580E+01  8.34367E+01  9.50522E+01  9.48028E+01  9.16799E+01  7.81543E+01  8.35784E+01  8.40797E+01  9.35810E+01  9.41910E+01  9.34248E+01",
				"           10            1  1.57041E+01  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00  0.00000E+00",
			},
			expectedTables: []string{"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0"},
			context:        "wide",
			etaCount:       64,
			epsCount:       1,
		},
	}

	for _, test := range tests {
		tt.Run(test.context, func(tt *testing.T) {
			t := wrapt.WrapT(tt)

			ed := ParseShkLines(test.lines)
			t.R.Equal(test.expectedTables, ed.EstimationMethods)
			t.R.Equal("TYPE", ed.ParameterNames[0])
			t.R.Equal("ETA(4)", ed.ParameterNames[5])

			shk := ParseShrinkage(ed.EstimationLines[0], test.etaCount, test.epsCount)
			t.R.Equal(test.etaCount, len(shk[0].EtaBar))

			sd := ParseShkData(ed, test.etaCount, test.epsCount)

			t.R.Equal(test.epsCount, len(sd[0][0].EpsSD))
			t.R.Equal(test.epsCount, len(sd[0][0].EpsVR))

			t.R.Equal(test.etaCount, len(sd[0][0].EtaBar))
		})
	}
}

func TestReadParseShkLines2(tt *testing.T) {
	var tests = []struct {
		lines     []string
		expected1 string
		expected2 string
		expected3 string
		context   string
		etaCount  int
		epsCount  int
	}{
		{
			lines: []string{
				"TABLE NO.     1: First Order Conditional Estimation with Interaction: Problem=1 Subproblem=0 Superproblem1=0 Iteration1=0 Superproblem2=0 Iteration2=0",
				"TYPE          SUBPOP       ETA(1)       ETA(2)       ETA(3)       ETA(4)",
				"            1            1  11.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            1            2  12.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            1            3  13.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            2            1  21.9719E-02  2.60947E-02  1.62789E-02  1.89996E-02",
				"            2            2  22.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            2            3  33.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            3            1  31.8606E-01  8.80753E-01  6.22835E-01  9.17950E-01",
				"            3            2  32.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            3            3  33.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            4            1  31.2499E+00  2.52933E+01  5.46357E+01  3.69109E+01",
				"            4            2  32.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            4            3  33.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            5            1  51.2691E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"            5            2  52.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            5            3  53.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            6            1  61.9795E+00  2.56533E+01  5.45618E+01  3.71115E+01",
				"            6            2  62.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            6            3  63.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            7            1  71.0000E+01  7.90000E+01  7.90000E+01  7.90000E+01",
				"            7            2  72.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            7            3  73.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            8            1  81.5368E+00  4.41891E+01  7.94208E+01  6.01976E+01",
				"            8            2  82.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            8            3  83.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            9            1  91.4166E+00  4.47258E+01  7.93537E+01  6.04504E+01",
				"            9            2  92.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            9            3  93.0000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",

				"            10           1  101.776E+01  0.00000E+00  0.00000E+00  0.00000E+00",
				"            10           2  102.000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
				"            10           3  103.000E-00 -3.91460E-03 -8.00654E-03 -1.95727E-03",
			},
			context:  "subpop",
			etaCount: 4,
			epsCount: 1,
		},
	}

	for _, test := range tests {
		tt.Run(test.context, func(tt *testing.T) {
			t := wrapt.WrapT(tt)

			ed := ParseShkLines(test.lines)

			shk := ParseShrinkage(ed.EstimationLines[0], test.etaCount, test.epsCount)
			t.R.Equal(3, len(shk))

			for i := range shk {
				t.R.Equal(int64(i+1), shk[i].SubPop)
			}

			sd := ParseShkData(ed, test.etaCount, test.epsCount)
			t.R.Equal(float64(11), sd[0][0].EtaBar[0])
			t.R.Equal(float64(92), sd[0][1].EbvVR[0])
			t.R.Equal(float64(33), sd[0][2].EtaSD[0])
		})
	}
}
