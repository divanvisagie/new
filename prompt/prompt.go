package prompt

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/divanvisagie/new/replace"
	"gopkg.in/yaml.v2"
)

// NewConfig represents the whole config file
type NewConfig struct {
	Replace Replace `yaml:"replace"`
}

// Replace is the replacement object
type Replace struct {
	Strings []ReplacementString `yaml:"strings"`
}

// ReplacementString represents a string configuration
type ReplacementString struct {
	Match       string `yaml:"match"`
	Description string `yaml:"description"`
}

func readYamlFile(configFilePath string) *NewConfig {
	config := NewConfig{}

	b, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("xxxxxxxxxxxxxxxxxxxxx %v\n", err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("xxxxxxxxxxxxxx Couldnt unmarshall")
	}

	return &config
}

// ProcessForTarget searches the target directory for .new.yml and processes it
func ProcessForTarget(target string, fetchUserInput func(string, string) string) {
	const separator = string(os.PathSeparator)

	yamlFilePath := path.Join(target, ".new.yml")

	config := readYamlFile(yamlFilePath)

	for _, x := range config.Replace.Strings {
		with := fetchUserInput(x.Match, x.Description)
		fmt.Printf("Replacing string %v with %v...", x.Match, with)
		replace.StartProcessWithString(x.Match, target, with)
	}
}
