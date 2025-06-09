func candy(ratings []int) int {
    n := len(ratings)
    candies := make([]int, n)
    for i := 0; i < n; i++{
        candies[i] = 1
    }
    // left to right
    for i := 1; i < n; i++{
        if ratings[i] > ratings[i-1]{
            candies[i] = candies[i-1] + 1
        }
    }
    // right to left
    for i := n-2; i >= 0; i--{
        if ratings[i] > ratings[i+1]{
            // candies[i+1] according to right to left but candies[i] according to the last one
            // so max of both
            candies[i] = max(candies[i], candies[i+1] + 1)
        }
    }
    // 
    fmt.Println(candies)

    sum := 0
    for i := 0; i < n; i++{
        sum += candies[i]
    }
    return sum
}