package config

import "flag"

var Flag *FlagStruct

type FlagStruct struct {
	InitConfigExample *bool
	ShowHelp *bool
}

func initFlag() (flagData *FlagStruct) {
	flagData = &FlagStruct{
		InitConfigExample: flag.Bool("init-config-example", true, "Generate the config example"),
		ShowHelp: flag.Bool("help", false, "Shows the help page"),
	}

	flag.Parse()

	return
}