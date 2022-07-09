package main

import (
	"fmt"

	"me.yusta/benchmark"
	"me.yusta/config"
	"me.yusta/httptest"
	"me.yusta/loop"
	"me.yusta/regex"
)

func main() {
	fmt.Println("Running Test-GO by (@Yusta)")
	config.Init()
	if config.Yaml.Test.Benchmark.All {
		benchmark.InitBenchmark()
	}

	if config.Yaml.Test.HttpTest.All {
		httptest.Init()
	}

	if config.Yaml.Test.Regex.All {
		regex.Test()
	}

	if config.Yaml.Test.Loop.All {
		loop.Init()
	}

}
