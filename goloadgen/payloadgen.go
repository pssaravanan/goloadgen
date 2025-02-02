package goloadgen

import (
	"fmt"
	"math/rand"
	"github.com/google/uuid"
	"text/template"
)

// PayloadGenParams holds the parameters for generating payloads
type PayloadGenParams struct {
	TemplateStr string
	SessionVar  map[string]string
	Vars        map[string]string
}

// TemplateWriter is used to capture the output of the template execution
type TemplateWriter struct {
	body string
}

func (writer *TemplateWriter) Write(p []byte) (n int, err error) {
	writer.body += string(p)
	return len(p), nil
}

// Variable to hold the rand.Intn function, so it can be mocked in tests
var randIntn = rand.Intn

// Variable to hold the UUID generation function, so it can be mocked in tests
var randUUID = func() string {
	return uuid.New().String()
}

// GeneratePayload generates a payload based on the provided parameters
func GeneratePayload(params PayloadGenParams) string {
	writer := &TemplateWriter{}
	funcMap := template.FuncMap{
		"randInt":  randIntn,
		"randUUID": randUUID,
	}
	tmpl, err := template.New("").Funcs(funcMap).Parse(params.TemplateStr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	tmpl.Execute(writer, nil)
	return writer.body
}