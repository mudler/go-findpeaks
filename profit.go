package findpeaks

// MaxProfit returns the maximum profit in a sequence with the given amount
func MaxProfit(seq []float64, amount float64) float64 {
	currentBalance := 0.0
	lh := LowAndHighs(seq)
	for i, s := range seq {
		currentAmount := (amount * s)
		switch lh[i].Type {
		case H:
			currentBalance = currentBalance + currentAmount
		case L:
			currentBalance = currentBalance - currentAmount
		}
	}
	return currentBalance
}
