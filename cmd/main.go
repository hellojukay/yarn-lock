package main

import (
	"github.com/hellojukay/yarn-lock/lock"
	lock "github.com/hellojukay/yarn-lock/pkg"
)

func main() {
	_, err := lock.New("../yarn.lock")
	if err != nil {
		panic(err)
	}
}
