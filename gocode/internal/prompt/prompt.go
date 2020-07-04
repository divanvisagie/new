package prompt

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/divanvisagie/new/internal/replace"
	"gopkg.in/yaml.v2"
)

// Verbose determines if we should print more stuff
var Verbose = false

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

func readConfigFile(path string) *NewConfig {
	config := NewConfig{}

	b, err := ioutil.ReadFile(path)
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
func ProcessForTarget(targetDirectory string, getUserInput func(match string, description string) string) {
	const s = string(os.PathSeparator)

	f := path.Join(targetDirectory, ".new.yml")
	config := readConfigFile(f)

	var r []replace.Replacement
	for _, x := range config.Replace.Strings {
		with := getUserInput(x.Match, x.Description)
		if x.Match == with {
			fmt.Printf("Skipped replacement for %s\n", x.Match)
			return
		}
		r = append(r, replace.Replacement{
			Match: x.Match,
			With:  with,
		})
	}
	replace.Verbose = Verbose
	replace.StartReplacementProcess(&r, targetDirectory)
}
