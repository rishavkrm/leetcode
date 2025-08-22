func letterCombinations(digits string) []string {
    dict := map[byte][]string{
        '2' : []string{"a", "b", "c"},
        '3' : []string{"d", "e", "f"},
        '4' : []string{"g", "h", "i"},
        '5' : []string{"j", "k", "l"},
        '6' : []string{"m", "n", "o"},
        '7' : []string{"p", "q", "r", "s"},
        '8' : []string{"t", "u", "v"},
        '9' : []string{"w", "x", "y", "z"},
    }
    result := []string{}
    helper(digits, dict, 0, "", &result)
    return result
}

func helper(digits string, dict map[byte][]string, index int, curr string, result *[]string){
    if len(curr) > 0 && len(curr) == len(digits){
        *result = append(*result, curr)
    }
    if index >= len(digits){
        return
    }
    s := digits[index]
    arr := dict[s]
    for i := range(len(arr)){
        curr += arr[i]
        helper(digits, dict, index+1, curr, result)
        curr = curr[:len(curr)-1]
    }
}
