package antlr4

type ttt interface {
	test()
}

type tests struct {
	ttt
}

func test() {
	t := tests{}
	t.test()
}
