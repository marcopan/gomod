// Code generated from /Users/pan/goworkspace/gomod/mod1/wangcai.g4 by ANTLR 4.9.2. DO NOT EDIT.

package wangcai // wangcai
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasewangcaiListener is a complete listener for a parse tree produced by wangcaiParser.
type BasewangcaiListener struct{}

var _ wangcaiListener = &BasewangcaiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasewangcaiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasewangcaiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasewangcaiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasewangcaiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasewangcaiListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasewangcaiListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunction is called when production function is entered.
func (s *BasewangcaiListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasewangcaiListener) ExitFunction(ctx *FunctionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasewangcaiListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasewangcaiListener) ExitConstant(ctx *ConstantContext) {}

// EnterOperator is called when production operator is entered.
func (s *BasewangcaiListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BasewangcaiListener) ExitOperator(ctx *OperatorContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BasewangcaiListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BasewangcaiListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasewangcaiListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasewangcaiListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BasewangcaiListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BasewangcaiListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BasewangcaiListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BasewangcaiListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BasewangcaiListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BasewangcaiListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}
