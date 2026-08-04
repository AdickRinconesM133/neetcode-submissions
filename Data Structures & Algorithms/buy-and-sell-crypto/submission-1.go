func maxProfit(prices []int) int {
    l := prices[0]
    best := 0

    for r, priceLen := 1, len(prices); r < priceLen; r++ {
        l = min(l, prices[r])
        best = max(best, prices[r] - l)
    }

    return best
}
