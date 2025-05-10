func hIndex(citations []int) int {
    if nPositives(citations) == 0{
        return 0
    }
    hIndex := 1
    for i := 1; i < len(citations); i++{
        subtractOne(citations)
        if nPositives(citations) > hIndex{
            hIndex += 1
        }
    }
    return hIndex 
    
}

func nPositives(arr []int) int{
    count := 0
    for _, i := range(arr){
        if i > 0{
            count += 1
        } 
    }
    return count
}

func subtractOne(arr []int){
    for i,_ := range(arr){
        arr[i] = arr[i]-1
    }
}