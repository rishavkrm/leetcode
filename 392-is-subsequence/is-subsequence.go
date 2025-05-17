func isSubsequence(s string, t string) bool {
    i := 0
    j := 0
    for i < len(s) && j < len(t){
        if s[i] == t[j]{
            if i == len(s)-1{
                return true
            }
            i += 1
        }
        j += 1
    }
    return i == len(s)
}