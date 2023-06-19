package main

import "github.com/a4anthony/go-commerce/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
