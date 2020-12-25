package helper

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"
)

// GenerateResetPasswordHTML generates reset password html using template and required data
func GenerateResetPasswordHTML(name, resetLink string) (result string, err error) {
	emailTemplate, err := getTemplate("reset_passowrd")
	if err != nil {
		Logger.Error(err)
		return
	}
	var data = map[string]interface{}{
		"name": name,
		"link": resetLink,
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
	filepath := path.Join(root, fmt.Sprintf("/html-templates/%s.html", templateName))
	emailTemplate, err = template.ParseFiles(filepath)
	if err != nil {
		Logger.Error(err)
		return
	}
	return
}
