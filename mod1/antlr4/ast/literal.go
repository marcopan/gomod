package ast

type BooleanLiteral struct {
	Value bool
}

func (bl *BooleanLiteral) Evaluate() (interface{}, error) {
	return bl.Value, nil
}

type FloatLiteral struct {
	Value float64
}

func (fl *FloatLiteral) Evaluate() (interface{}, error) {
	return fl.Value, nil
}

type IntegerLiteral struct {
	Value int64
}

func (il *IntegerLiteral) Evaluate() (interface{}, error) {
	return il.Value, nil
}

type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) Evaluate() (interface{}, error) {
	return sl.Value, nil
}

type Operator struct {
	Operator string
}
