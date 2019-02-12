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

// ProcessForTarget searches the target directory for .new.yml and processes it
func ProcessForTarget(targetDirectory string, fetchUserInput func(string, string) string) {
	const separator = string(os.PathSeparator)

	yamlFilePath := path.Join(targetDirectory, ".new.yml")

	config := readYamlFile(yamlFilePath)

	var replacements []replace.Replacement
	for _, x := range config.Replace.Strings {
		with := fetchUserInput(x.Match, x.Description)
		if x.Match == with {
			fmt.Printf("Skipped replacement for %s\n", x.Match)
			return
		}
		replacements = append(replacements, replace.Replacement{
			Match:    x.Match,
			FilePath: targetDirectory,
			With:     with,
		})
	}
	replace.StartReplacementProcess(replacements, targetDirectory)
}
