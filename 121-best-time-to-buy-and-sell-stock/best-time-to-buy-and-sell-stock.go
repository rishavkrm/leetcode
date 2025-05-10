
func maxProfit(prices []int) int {
    x := prices[0]
    maxProfit := 0
    for i := range(len(prices)){
        x = Min(x, prices[i])
        maxProfit = Max(maxProfit, prices[i]- x)
    }
    return maxProfit
}

func Min(x int, y int) int{
    if x < y{
        return x
    }
    return y
}

func Max( x int, y int) int{
    if x > y{
        return x
    }
    return y
}