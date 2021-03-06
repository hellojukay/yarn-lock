package main

import (
	"fmt"

	lock "github.com/hellojukay/yarn-lock"
)

func main() {
	lock, err := lock.FromFile("./yarn.lock")
	if err != nil {
		panic(err)
	}
	components := lock.Component()
	for _, c := range components {
		fmt.Printf("%20s%30s\n", c.Name, c.Version)
	}
}
