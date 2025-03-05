package main

import (
	"fmt"

	"calculate/internal/application"
)

func main() {
	app := application.New()
	fmt.Println("RunServer")
	app.RunServer()
}
