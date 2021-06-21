package main

func main() {
	Calculate(1, 2, "")
}

func Calculate(left int, right int, s string) int {
	return NewOperator(s).Calc(left, right)
}
