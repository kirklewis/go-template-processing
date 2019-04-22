package main

import (
	"demo/template"
	"fmt"
)

func main() {
	vars := make(map[string]string)

	vars["Name"] = "Brienne"
	vars["House"] = "Tarth"

	// process a template string
	resultA := template.ProcessString("{{.Name}} of house {{.House}}", vars)

	// process a template file
	resultB := template.ProcessFile("templates/got.tmpl", vars)

	fmt.Println(resultA, "\n")
	fmt.Println(resultB, "\n")

	// using a Struct
	type Westerosi struct {
		Name  string
		House string
	}

	jorah := Westerosi{"Ser Jorah", "Mormont"}
	resultC := template.ProcessFile("templates/got.tmpl", jorah)
	fmt.Println(resultC)
}
