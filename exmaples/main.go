package main

import (
	"fmt"

	lock "github.com/hellojukay/yarn-lock/pkg"
)

func main() {
	lock, err := lock.New("./yarn.lock")
	if err != nil {
		panic(err)
	}
	components := lock.Component()
	for _, c := range components {
		fmt.Printf("%20s%30s\n", c.Name, c.Version)
	}
}
