package main

import (
	"fmt"
	"os"

	"github.com/samsamisamsam/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cfg.SetUser("sam")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
