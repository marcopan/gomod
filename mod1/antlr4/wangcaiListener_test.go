package antlr4

import (
	"example.com/mod1/v1/antlr4/wangcai"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"testing"
)

func TestLexer_test(t *testing.T) {
	data := "FooBar(324,343,432,Sub(\"232\"+343))"
	is := antlr.NewInputStream(data)
	lexer := wangcai.NewwangcaiLexer(is)
	for {
		token := lexer.NextToken()
		fmt.Println(token.GetText())
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
}

func TestExpression_test(t *testing.T) {
	s := "1234+4321-1111"
	ss := string(s)
	stream := antlr.NewInputStream(ss)
	lexer := wangcai.NewwangcaiLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := NewWangcaiListener()
	parser := wangcai.NewwangcaiParser(tokenStream)
	parser.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Expression())

	root := listener.Root
	evaluate, err := root.Evaluate()
	if err != nil {
		t.Errorf(err.Error())
	}
	if reflect.ValueOf(evaluate).Kind() != reflect.Int64 {
		t.Errorf("not int 64")
	}

	if evaluate.(int64) != 4444 {
		t.Errorf("incorrect result")
	}

}

func TestFunctionExpression_test(t *testing.T) {
	data := "ToUpper(\"Thisisstring\"+1234)"
	sdata := string(data)
	inStream := antlr.NewInputStream(sdata)
	lexer := wangcai.NewwangcaiLexer(inStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := NewWangcaiListener()
	parser := wangcai.NewwangcaiParser(stream)
	parser.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Expression())

	root := listener.Root
	evaluate, err := root.Evaluate()
	if err != nil {
		t.Errorf(err.Error())
	}
	if reflect.ValueOf(evaluate).Kind() != reflect.String {
		t.Errorf("not an int64")
	}

	if evaluate.(string) != "THISISSTRING1234" {
		t.Errorf("incorrect result")
	}
}
