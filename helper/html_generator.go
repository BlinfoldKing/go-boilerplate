package helper

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"
)

// GenerateHTMLTemplate generate html template
func GenerateHTMLTemplate(template string, data map[string]interface{}) (result string, err error) {
	emailTemplate, err := getTemplate(template)
	if err != nil {
		Logger.Error(err)
		return
	}

	var temp bytes.Buffer
	err = emailTemplate.Execute(&temp, data)
	if err != nil {
		Logger.Error(err)
		return
	}

	result = temp.String()
	return
}

func getTemplate(templateName string) (emailTemplate *template.Template, err error) {
	root, _ := os.Getwd()
	filepath := path.Join(root, fmt.Sprintf("/templates/%s.html", templateName))
	emailTemplate, err = template.ParseFiles(filepath)
	if err != nil {
		Logger.Error(err)
		return
	}
	return
}
