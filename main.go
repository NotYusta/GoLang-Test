package main

import (
	"fmt"
	"os"

	"me.yusta/benchmark"
	"me.yusta/config"
	"me.yusta/httptest"
	"me.yusta/loop"
	"me.yusta/regex"
)

func main() {
	fmt.Println("Running Test-GO by (@Yusta)")
	config.Init()

	

	if *config.Flag.ShowHelp {
		println()
		println("--------------------------")
		println("      Argument List")
		println("--------------------------")
		println()
		println("> help - Shows the argument list")
		println("> init-config-example - Generate the config example")
		println()
		os.Exit(0)
	}
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


	fmt.Println("Test has been completed.")
}
