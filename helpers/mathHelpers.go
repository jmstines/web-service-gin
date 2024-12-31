package helpers

func CalculateSum(prices []int64) int64 {
	var total int64 = 0
	for _, price := range prices {
		total += price
	}
	return total
}
