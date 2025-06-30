package main

import (
	"os"
	"text/template"
)

func main() {
	const tmpl = `
Names:
{{range .}}
- {{.}}
{{end}}`

	t, err := template.New("list").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	names := []string{"Alice", "Bob", "Charlie"}
	t.Execute(os.Stdout, names)
}
