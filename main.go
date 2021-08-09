package main

import (
	"fmt"
	"github.com/kirklewis/go-template-processing/src/demo/template"
)

func main() {
	// data structure the template will be applied to
	vars := make(map[string]interface{})

	vars["Name"] = "Brienne"
	vars["House"] = "Tarth"
	vars["Traits"] = []string{"Brave", "Loyal"}

	// process a template string
	resultA := template.ProcessString("{{.Name}} of house {{.House}}", vars)

	// process a template file
	resultB := template.ProcessFile("templates/got.tmpl", vars)

	fmt.Println(resultA, "\n")
	fmt.Println(resultB)

	// using a Struct
	type Westerosi struct {
		Name   string
		House  string
		Traits []string
	}

	jorah := Westerosi{"Ser Jorah", "Mormont", []string{"Brave", "Protective"}}
	resultC := template.ProcessFile("templates/got.tmpl", jorah)

	fmt.Println(resultC)
}
