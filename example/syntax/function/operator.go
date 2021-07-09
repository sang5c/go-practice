package function

type opFunction func(int, int) int

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}

}

func add(left, right int) int {
	return left + right
}

func mul(left, right int) int {
	return left * right
}
