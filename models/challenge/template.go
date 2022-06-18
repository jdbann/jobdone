package challenge

import "text/template"

const md = `# Challenge #{{ .Number }}: {{ .Title }}

---

{{ .Description }}

## Objectives`

var mdTemplate = template.Must(template.New("challenge").Parse(md))
