package main

import (
	"fmt"

	"expenses/ui/optionlist"
	"expenses/ui/textinput"
)

func main() {
	//Get user option
	selection := optionlist.GetOption()
	fmt.Println(selection)

	//Get user text input
	choice := textinput.GetInput("Enter some text", "Placeholder")
	fmt.Println(choice)

}
