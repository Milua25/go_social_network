package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type User struct {
	Name string
}

func main() {
	//tmpl := template.New("example")
	//
	//tmpl, err := tmpl.Parse("Welcome, {{.Name}}! How are things?")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = tmpl.Execute(os.Stdout, User{Name: "Bob"})
	//if err != nil {
	//	panic(err)
	//}

	//tmpl := template.Must(template.New("example").Parse("Welcome, {{.Name}}! How are things?"))
	//
	//err := tmpl.Execute(os.Stdout, User{Name: "Bob"})
	//fmt.Println()
	//if err != nil {
	//	panic(err)
	//}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}!",
		"notification": "{{.name}}, you have a new messages: {{.notification}}",
		"error":        "{{.name}}, something went wrong: {{.errorMessage}}",
	}

	// parse and store templates in a map
	parsedTemplates := make(map[string]*template.Template, len(templates))

	for key, tmplStr := range templates {
		parsedTemplates[key] = template.Must(template.New(key).Parse(tmplStr))
	}

	for {
		// show menu
		fmt.Println("\nMenu:")
		fmt.Println("\n1. Join:")
		fmt.Println("\n2. Get notifications:")
		fmt.Println("\n3. Get Error:")
		fmt.Println("\n4. Exit:")

		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}

		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Print("Enter notification: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"name": name, "notification": notification}
		case "3":
			fmt.Print("Enter error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"name": name, "errorMessage": errorMessage}
		case "4":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
			continue
		}

		// Render a template and print it to stdout
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
