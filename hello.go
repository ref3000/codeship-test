package main

import (
	"fmt"

	"github.com/ref3000/codeship-test/a"
	"github.com/ref3000/codeship-test/b"
	"github.com/ref3000/codeship-test/calc"
)

func main() {
	fmt.Printf("Hello World%d\n", calc.Sum(114, 514))
	fmt.Printf("%s\n", a.Afunc("hoge"))
	fmt.Printf("%s\n", b.Bfunc("hoge"))
}
