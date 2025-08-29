func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	stack := make([]string, 0)
	result := ""
	for i := 0; i < len(paths); i++ {
		if paths[i] == "" || paths[i] == "." {
			continue
		}
		if paths[i] == ".." {
			if len(stack) == 0 {
				continue
			} else {
				stack = stack[:len(stack)-1]
			}
		}else{
            stack = append(stack, paths[i])
        }
	}
    for i := range(len(stack)){
        result += "/"
        result += stack[i]
    }
    if result == ""{
        result = "/"
    }
	return result
}