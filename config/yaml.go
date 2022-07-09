package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Yaml *YamlStruct

type YamlStruct struct {
	Test struct {
		Benchmark struct {
			All bool `yaml:"all"`
		} `yaml:"benchmark"`
		HttpTest struct {
			All bool `yaml:"all"`
		} `yaml:"httpTest"`
		Loop struct {
			All bool `yaml:"all"`
		} `yaml:"loop"`
		Regex struct {
			All bool `yaml:"all"`
		} `yaml:"regex"`
	}
	Midtrans struct {
		ClientKey      string `yaml:"clientKey"`
		ServerKey      string `yaml:"serverKey"`
		ProductionMode bool   `yaml:"productionMode"`
	} `yaml:"midtrans"`
}

func initYaml() *YamlStruct {
	if Flag.InitConfigExample {
		yamlRes, _ := yaml.Marshal(&YamlStruct{})
		os.WriteFile(".config.example.yml", yamlRes, 2)
	}

	_, err := os.Stat(".config.yml")
	if os.IsNotExist(err) {
		fmt.Println("A configuration file was not found, a configuration has been generated.")
		yamlRes, _ := yaml.Marshal(&YamlStruct{})
		os.WriteFile(".config.yml", yamlRes, 2)

		os.Exit(1)
	}

	res, err := os.ReadFile(".config.yml")
	if err != nil {
		fmt.Println("An error has occurred, " + err.Error())
		os.Exit(1)
	}

	var config YamlStruct
	err = yaml.Unmarshal(res, &config)

	if err != nil {
		fmt.Println("An error has occurred, " + err.Error())
		os.Exit(1)
	}

	return &config
}
