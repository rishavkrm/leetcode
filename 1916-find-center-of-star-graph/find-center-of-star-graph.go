func findCenter(edges [][]int) int {
    hash := map[int]int{}
    for _, edge := range edges{
        for _, node := range edge{
            _, ok := hash[node]
            if ok{
                hash[node] = hash[node] + 1
            }else{
                hash[node] = 1
            }
        }
    }
    for key,v := range hash{
        if v >= len(hash)-1{
            return key 
        }
    }
    return -1
}