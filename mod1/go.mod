module example.com/mod1/v1

go 1.17


require (
	test.com/mod2 v0.0.1
)
replace (
	test.com/mod2 => /Users/pan/goworkspace/gomod/mod2
)
