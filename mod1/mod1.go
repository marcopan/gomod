package main

import (
	"example.com/mod1/v1/util"
	"fmt"
	util2 "test.com/mod2/util"
	//util2 "test.com/mod2/util"
)

func main() {
	fmt.Println(util.Util())
	//fmt.Println(util2.Mod2())
	fmt.Println(util2.Mod2())
}
