package templates

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/slack-go/slack"
	"sigs.k8s.io/yaml"
)

func NewSlackTemplate(i interface{}, content string) (*slack.Blocks, error) {
	var blocks slack.Blocks
	var templateContent bytes.Buffer

	t, err := template.New("").Funcs(sprig.TxtFuncMap()).Parse(string(content))
	if err != nil {
		return nil, err
	}

	if err := t.Execute(&templateContent, i); err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(templateContent.Bytes(), &blocks)
	if err != nil {
		return nil, err
	}

	return &blocks, nil
}
