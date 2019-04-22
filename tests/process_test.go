package main

import (
	"testing"

	"demo/template"
)

var vars = map[string]string{
	"Name":  "John",
	"Count": "10",
	"Item":  "apples",
}

func TestProcessString(t *testing.T) {
	str := "{{.Name}} has {{.Count}} {{.Item}}"

	result := template.ProcessString(str, vars)
	expected := "John has 10 apples"

	if result != expected {
		t.Errorf("Got %v expected %v", result, expected)
	}
}

func TestProcessFile(t *testing.T) {
	result := template.ProcessFile("./templates/test.tmpl", vars)
	expected := "John has 10 apples"

	if result != expected {
		t.Errorf("Got %v expected %v", result, expected)
	}
}
