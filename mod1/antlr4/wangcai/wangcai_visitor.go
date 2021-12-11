// Code generated from /Users/pan/goworkspace/gomod/mod1/wangcai.g4 by ANTLR 4.9.2. DO NOT EDIT.

package wangcai // wangcai
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by wangcaiParser.
type wangcaiVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by wangcaiParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by wangcaiParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by wangcaiParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by wangcaiParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by wangcaiParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by wangcaiParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by wangcaiParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by wangcaiParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by wangcaiParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}
}
