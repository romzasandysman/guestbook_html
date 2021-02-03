package main

import (
	"os"
	"text/template"
)

func executeTemplate(text string, data interface{}){
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

type Part struct {
	Name string
	Count int
	Active bool
}

func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{if.}} true {{end}}\n", true)
	executeTemplate("Dot is: {{if.}} true {{end}}\n", false)

	templateText := "Before loop: {{.}} \n{{range .}} In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})

	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "text", Count: 2})

	templateText = "Name: {{.Name}}\n{{if.Active}}Count: {{.Count}}\n{{end}}"
	executeTemplate(templateText, Part{Name: "text", Count: 2, Active: true})
	executeTemplate(templateText, Part{Name: "text", Count: 2, Active: false})
}
