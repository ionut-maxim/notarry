package templates

import (
	"testing"
)

type Person struct {
	Name string
}

func TestNewSlackTemplate(t *testing.T) {
	content := `- type: header
  text:
    type: plain_text
    text: |
      {{upper .Name}}
      {{add 1 2 3}}
      {{uuidv4}}`

	person := &Person{Name: "John Doe"}

	blocks, err := NewSlackTemplate(person, content)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
		return
	}

	if blocks.BlockSet[0].BlockType() != "header" {
		t.Errorf("Wrong block type, it should be 'header' but the value is '%s'", blocks.BlockSet[0].BlockType())
	}
}
