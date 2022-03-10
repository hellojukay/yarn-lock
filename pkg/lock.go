package lock

import (
	"io/ioutil"
	"os"
	"strings"
)

type Dependency struct {
	Name    string
	Version string
}
type Component struct {
	Name         string
	Version      string
	Resolved     string
	Integrity    string
	Dependencies []Dependency
}
type Yarn []Component

func (y Yarn) Component() []Component {
	return []Component(y)
}
func FromFile(file string) (*Yarn, error) {
	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}
	return New(string(content))
}

func New(content string) (*Yarn, error) {
	lines := strings.Split(content, "\n")
	var cs []Component
	var c Component
	var dependencies []Dependency
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		if len(line) == 0 && len(c.Name) > 0 {
			c.Dependencies = dependencies
			dependencies = nil
			cs = append(cs, c)
			continue
		}
		if !strings.HasPrefix(line, " ") {
			name := line
			c.Name = clarComponentName(name)
			continue
		}
		if strings.HasPrefix(line, "  ") && !strings.HasPrefix(line, "    ") {
			line = strings.Trim(line, " ")
			array := strings.Split(line, " ")
			switch array[0] {
			case "version":
				c.Version = strings.Trim(array[1], "\"")
			case "resolved":
				c.Resolved = strings.Trim(array[1], "\"")
			case "integrity":
				c.Integrity = strings.Trim(array[1], "\"")
			}
			continue
		}
		if strings.HasPrefix(line, "    ") {
			line = strings.Trim(line, " ")
			array := strings.Split(line, " ")
			dependencies = append(dependencies, Dependency{
				Name:    strings.Trim(array[0], "\""),
				Version: strings.Trim(array[1], "\""),
			})
		}
	}
	y := Yarn(cs)
	return &y, nil
}

func clarComponentName(name string) string {
	index := 0
	result := ""
	if strings.HasPrefix(name, "\"") {
		name = strings.Trim(name, "\"")
		result = string(name[0])
		index = 1
	}
	for index < len(name) {
		if rune(name[index]) == rune('@') {
			return result
		}
		result = result + string(name[index])
		index = index + 1
	}
	return result
}
