func evalRPN(tokens []string) int {
	stack := []int{}

	for i := 0; i < len(tokens); i++ {
		ch := tokens[i]

			if ch == "+" {
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b + a)
			} else if ch == "-" {
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b - a)
			} else if ch == "*" {
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b * a)
			} else if ch == "/" {
				a := stack[len(stack)-1]
				b := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, b / a)
			} else {
				intConv, _ := strconv.Atoi(ch)
				stack = append(stack, intConv) 
			}
	}

	return stack[len(stack)-1]
}
