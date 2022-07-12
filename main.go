package main

import (
	"fmt"

	"me.yusta/benchmark"
	"me.yusta/config"
	"me.yusta/httptest"
	"me.yusta/loop"
	"me.yusta/midtrans"
	"me.yusta/regex"
	"me.yusta/number"
)

func main() {
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
		return
	}

	fmt.Println("Running Test-GO by (@Yusta)")

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

	if config.Yaml.Test.Midtrans.All {
		midtrans.InitMidtrans()
	}

	if config.Yaml.Test.Number.All {
		number.Init()
	}


	fmt.Println("Test has been completed.")
}
