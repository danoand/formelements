package formelements

import (
	"bytes"
	"fmt"
	tmpl "html/template"
	"os"
)

// FormElement describes an HTML element that can be included in a form
type FormElement struct {
	ID            string            `json:"docid_form_element"`
	IDS           map[string]string `json:"docids"`
	CampaignID    string            `json:"docid_campaign"`
	Type          string            `json:"type"`
	Order         int               `json:"order"`
	Name          string            `json:"name"`
	Label         string            `json:"label"`
	HelpText      string            `json:"help_text"`
	Classes       []string          `json:"string"`
	Options       []SelectOption    `json:"select_options"`
	IsMultSelect  bool              `json:"is_multi_select"`
	IsHidden      bool              `json:"is_hidden"`
	IsRequired    bool              `json:"is_required"`
	Value         string            `json:"value"`
	Placeholder   string            `json:"placeholder"`
	CheckBoxValue string            `json:"checkbox_value"`
	RadioLabel1   string            `json:"radio_label_1"`
	RadioValue1   string            `json:"radio_value_1"`
	RadioLabel2   string            `json:"radio_label_2"`
	RadioValue2   string            `json:"radio_value_2"`
	HTMLTemplate  tmpl.Template     `json:"html_template"`
}

// ParseSelect is a method on FormElement that parses a select statement
func (elm *FormElement) ParseSelect() (string, error) {
	var tplOut bytes.Buffer
	err = t.ExecuteTemplate(&tplOut, "select_element", elm)
	if err != nil {
		return "", err
	}

	return tplOut.String(), nil
}

// NotEmpty is a method on FormElement that determines if a passed string is empty (length == 0)
func (elm *FormElement) NotEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}

	return false
}

// SelectOption describes an HTML select option that is displayed in a dropdown
type SelectOption struct {
	Order   int    `json:"order"`
	Value   string `json:"value"`
	Display string `json:"display"`
}

// HTMLTemplateMap maps form element types to a default Twitter Boostrap HTML template string
var HTMLTemplateMap = make(map[string]string)

var err error
var t *tmpl.Template

func init() {
	HTMLTemplateMap["select_element"] = `
	<label"" for="{{ .ID }}">{{ .Label }}</label>
	<select name="{{ .Name }}" id="{{ .ID }}" class="form-select">
	  <option value="not_selected" selected>-- Select --</option>
	  {{range .Options}}
		<option value="{{ .Value }}">{{ .Display }}</option> 
	  {{end}}
	</select>
	{{ if .FormElement.NotEmpty .HelpText }}
	<small class="form-text text-muted">We'll never share your email with anyone else.</small>
	{{ .end }}
	`
	t, err = tmpl.New("select_element").Parse(HTMLTemplateMap["select_element"])
	if err != nil {
		fmt.Printf("hey man! error parsing the select_element template; see: %v\n", err)
		os.Exit(1)
	}
}
