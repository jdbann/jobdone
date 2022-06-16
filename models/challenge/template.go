package challenge

import "text/template"

const md = `# Challenge #{{ .Number }}: {{ .Title }}

---

{{ .Description }}

{{ range .Objectives }}
- [{{ if .Complete }}x{{ else }} {{ end }}] {{ .Description }}
{{ end}}`

var mdTemplate = template.Must(template.New("challenge").Parse(md))
