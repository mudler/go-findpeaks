package findpeaks

// FindPeaks returns peaks indexes of a float64 sequence
func FindPeaks(sequence []float64) []int {
	var peak []int
	for i := range sequence {
		if (i == 0 || sequence[i-1] <= sequence[i]) &&
			(i == len(sequence)-1 || sequence[i+1] <= sequence[i]) {
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
