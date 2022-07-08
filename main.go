package main

import (
	"fmt"
	"regexp"

	"me.yusta/benchmark"
	"me.yusta/loop"
	"me.yusta/regex"
)



func main() {

	benchmark.InitBenchmark()
	r, _ := regexp.Compile("(b(a|@)ks(o|0))")

	fmt.Println(r.FindAllString("bakso baks0", 2))
	regex.Test()
	loop.Init()

}