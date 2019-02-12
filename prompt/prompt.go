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
		fmt.Printf("readYamlFile error: %v\n", err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		fmt.Printf("Couldnt unmarshall: %s", err.Error())
	}

	return &config
}

type replacement struct {
	match           string
	filePath        string
	replacementText string
}

// ProcessForTarget searches the target directory for .new.yml and processes it
func ProcessForTarget(target string, fetchUserInput func(string, string) string) {
	const separator = string(os.PathSeparator)

	yamlFilePath := path.Join(target, ".new.yml")

	config := readYamlFile(yamlFilePath)

	var replacements []replacement
	for _, x := range config.Replace.Strings {
		with := fetchUserInput(x.Match, x.Description)
		if x.Match == with {
			fmt.Printf("Skipped replacement for %s\n", x.Match)
			return
		}
		replacements = append(replacements, replacement{
			match:           x.Match,
			filePath:        target,
			replacementText: with,
		})
	}

	for _, r := range replacements {
		fmt.Printf("Replacing string %v with %v...\n", r.match, r.replacementText)
		replace.StartProcessWithString(r.match, r.filePath, r.replacementText)
	}
}
