package main

import(
	"fmt"
)


func main() {
	var colors = map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
	}

	colors["jsdlkfj"] = "aksjd;flkj"
	fmt.Println(colors)
}

