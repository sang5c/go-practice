package main

type Operator struct {
	Symbol string
}

func NewOperator(symbol string) *Operator {
	operator := &Operator{Symbol: symbol}
	return operator
}

// Calc 자바의 enum처럼 상태와 행위를 한곳에서 관리할 방법이 필요함.
func (o *Operator) Calc(left, right int) int {
	if o.Symbol == "+" {
		return left + right
	} else {
		return left - right
	}
}
