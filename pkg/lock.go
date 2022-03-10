package lock

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Dependency string
type component struct {
	Version      string
	Resolved     string
	Integrity    string
	Dependencies []Dependency
}
type Yarn []component

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
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		n := countPrefix(content, ' ')
		fmt.Printf("%d %s", n, line)
	}
	return nil, nil
}

func countPrefix(s string, c rune) int {
	var n = 0
	for _, ch := range s {
		if ch == c {
			n = 1 + n
		} else {
			return n
		}
	}
	return n
}
