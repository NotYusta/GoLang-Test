package main

import (
	"fmt"
	"regexp"

	"me.yusta/regex"
)



func main() {

	r, _ := regexp.Compile("(b(a|@)ks(o|0))")

	fmt.Println(r.FindAllString("bakso baks0", 2))
	regex.Test()
}