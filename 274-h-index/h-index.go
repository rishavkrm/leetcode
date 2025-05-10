func hIndex(citations []int) int {
    // 0, 1, 3, 5, 6
    sort.Ints(citations)
    n := len(citations)
    if n == 1 && citations[0] > 0{
        return 1
    }
    hIndex := 0
    for i := 0; i < n; i++{
        if citations[n-i-1] >= hIndex+1{
            hIndex += 1
        }else{
            break
        }
    }
    return hIndex
}