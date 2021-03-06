package main

func generateParenthesis(n int) []string {
	res := make([]string, 0, 0)
	// put "(" at first
	backtrack("(", n-1, n, &res)
	return res
}

// remaining left "(" and remaining right ")"
func backtrack(str string, left int, right int, res *[]string) {
	// if "right" == 0, put it into the array
	if right == 0 {
		*res = append(*res, str)
		return
	}

	if right == left && left != 0 {
		// We don't allow ")" more than "("
		backtrack(str+"(", left-1, right, res)
	} else if left != 0 && right != 0 {
		// take a step
		backtrack(str+"(", left-1, right, res)
		// take another step
		backtrack(str+")", left, right-1, res)
	} else if left == 0 && right != 0 {
		// we don't have "(" any more
		backtrack(str+")", left, right-1, res)
	}

	return

}
