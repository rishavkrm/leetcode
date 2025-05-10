func canCompleteCircuit(gas []int, cost []int) int {
    if Sum(gas) < Sum(cost){
        return -1
    }
    min := 10001
    res := 0
    curr := 0
    for i := range(len(cost)){
        curr += (gas[i] - cost[i])
        fmt.Println(curr, min)
        if curr < min{
            res = i
            min = curr
        }
    }
    return (res+1)%len(cost)
}

func Sum(arr []int)int{
    s := 0      
    for i := range(len(arr)){
        s += arr[i]
    }
    return s
}
// -2 -2 -2 3 3
// -2 -4 -6 0 0


