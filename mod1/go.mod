module example.com/mod1/v1

go 1.17

require test.com/mod2 v0.0.1

require (
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211208004053-14916ba2a2ce // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace test.com/mod2 => /Users/pan/goworkspace/gomod/mod2
