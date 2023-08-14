package main

import (
	"fmt"
	"github.com/vkolev/locmock"
)

func main() {
	config, err := locmock.LoadConfig("/Users/vladi/GolandProjects/locmock/locmock.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	locmock.Run(config)
}
