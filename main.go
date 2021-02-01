package main

import (
	"fmt"
	ellipse "workflowcore/sdk/ellipse"
)

func main() {
	//Initalise the Init function with value of A,B
	e := ellipse.Init{
		9, 2,
	}
	//this will give answer as 0.9749960430435691
	fmt.Println(e.GetEccentricity())
}
