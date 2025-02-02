package goloadgen

import (
	"fmt"
	"math/rand"
	"text/template"
)

type TemplateWriter struct {
	body string
}

func (writer *TemplateWriter) Write(p []byte) (n int, err error) {
	// fmt.Println("inside write", string(p))
	writer.body = writer.body + string(p)
	return len(p), nil
}

// Variable to hold the rand.Intn function, so it can be mocked in tests
var randIntn = rand.Intn

func GeneratePayload(templatestr string) string {
	writer := &TemplateWriter{}
	funcMap := template.FuncMap{
		"randInt": randIntn,
	}
	tmpl, err := template.New("").Funcs(funcMap).Parse(templatestr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	tmpl.Execute(writer, nil)
	return writer.body
}