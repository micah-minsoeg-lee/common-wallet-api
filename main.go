package main

import (
	"fmt"

	"github.com/micah-minsoeg-lee/common-wallet-api/app"
)

func main() {
	if app, err := app.NewApp(); err != nil {
		fmt.Println(fmt.Sprintf("Running common-wallet-api fail: %v", err))
		return
	} else {
		app.Run()
	}
}
