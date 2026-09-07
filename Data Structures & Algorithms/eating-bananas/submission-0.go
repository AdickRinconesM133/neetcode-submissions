func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	l, r := 1, maxPile

	for l < r {
		k := l + (r-l)/2

		if canFinish(piles, h, k) {
			r = k
		} else {
			l = k + 1
		}
	}

	return l
}

func canFinish(piles []int, h int, k int) bool {
	hours := 0

	for _, p := range piles {
		hours += (p + k - 1) / k
	}

	return hours <= h
}