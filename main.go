package main

import (
	"fmt"
	app "main/app"
)

func main(){

	system := app.GetSystem()

	for {
		fmt.Println("Enter input: (CLOSE to exit)")
		stop := system.ProcessInput()
		if stop{
			break
		}
	}
}