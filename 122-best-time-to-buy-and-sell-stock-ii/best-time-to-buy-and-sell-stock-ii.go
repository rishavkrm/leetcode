func maxProfit(prices []int) int {
    // 7, 1, 5, 3, 6, 4
    curr_min := prices[0]
    total_profit := 0
    n := len(prices)
    for i := range(n){
        if prices[i] - curr_min > 0{
            total_profit += prices[i] - curr_min
            curr_min = prices[i]
        }else{
            curr_min = Min(curr_min, prices[i])
        }
    }
    return total_profit
}

func Min(a int, b int) int{
    if a > b{
        return b
    }
    return a
}