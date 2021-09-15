package findpeaks

const (
	L = iota // Low
	H = iota // High
	S = iota // Stationary
)

type Type int
type LowAndHigh struct {
	Value float64
	Type  Type
}

type Sequence []LowAndHigh

func (s Sequence) Float64() []float64 {
	res := []float64{}
	for _, v := range s {
		res = append(res, v.Value)
	}
	return res
}

// FindPeaks returns peaks indexes of a float64 sequence
func FindPeaks(sequence []float64) []int {
	var peak []int
	for i := range sequence {
		if (i == 0 || sequence[i-1] < sequence[i]) &&
			(i == len(sequence)-1 || sequence[i+1] < sequence[i]) {
			peak = append(peak, i)
		}
	}

	return peak
}

// ReverseSequence reverse a sequence sign. Useful to find low peaks of a sequence
func ReverseSequence(sequence []float64) []float64 {
	var reversed []float64
	for _, s := range sequence {
		reversed = append(reversed, -s)
	}
	return reversed
}

// Returns the input sequence with the additional value of the
// point type. If it's a peak, a Type H is set, if it's a low point an,
// L type is set. S is set to stationary in case it's neither a low or a high point.
func LowAndHighs(sq []float64) []LowAndHigh {
	res := []LowAndHigh{}

	highs := FindPeaks(sq)
	lows := FindPeaks(ReverseSequence(sq))

	for i, v := range sq {
		currentHigh, currentLow := 0, 0
		if len(highs) > 0 {
			currentHigh = highs[0]
		}
		if len(lows) > 0 {
			currentLow = lows[0]
		}

		var t Type
		switch i {
		case currentHigh:
			t = Type(H)
			highs = highs[1:]
		case currentLow:
			t = Type(L)
			lows = lows[1:]
		default:
			t = Type(S)
		}

		res = append(res, LowAndHigh{
			Value: v,
			Type:  t,
		})
	}

	return res
}
