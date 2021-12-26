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
	AlertClass    string            `json:"alert_class"`
	Options       []SelectOption    `json:"select_options"`
	IsMultSelect  bool              `json:"is_multi_select"`
	IsHidden      bool              `json:"is_hidden"`
	IsRequired    bool              `json:"is_required"`
	Value         string            `json:"value"`
	Placeholder   string            `json:"placeholder"`
	AlertMessages []string          `json:"alert_messages"`
	CheckBoxValue string            `json:"checkbox_value"`
	RadioLabel1   string            `json:"radio_label_1"`
	RadioValue1   string            `json:"radio_value_1"`
	RadioLabel2   string            `json:"radio_label_2"`
	RadioValue2   string            `json:"radio_value_2"`
	HTMLTemplate  tmpl.Template     `json:"html_template"`
}

// ParseElement is a method on FormElement that parses a select statement
func (elm *FormElement) ParseElement() (string, error) {
	var tplOut bytes.Buffer
	err = HTMLTemplates[elm.Type].ExecuteTemplate(&tplOut, elm.Type, elm)
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

// HTMLTemplateStrings maps form element types to a default Twitter Boostrap HTML template string
var HTMLTemplateStrings = map[string]string{
	//
	// alert messages
	//
	"alert_messages": `
	<div class="alert {{ .AlertClass }}" role="alert">
		{{range .AlertMessages}}
			{{.}}
		{{end}}
	</div>
	`,
	//
	// select element
	//
	"select_element": `
	<div class="mb-3">
		<label for="{{ .ID }}">{{ .Label }}</label>
		<select name="{{ .Name }}" id="{{ .ID }}" class="form-select">
			<option value="not_selected" selected>-- Select --</option>
			{{range .Options}}
			<option value="{{ .Value }}">{{ .Display }}</option> 
			{{end}}
		</select>
		{{ if .NotEmpty .HelpText }}
		<small class="form-text text-muted">{{ .HelpText }}</small>
		{{ end }}
	<div class="mb-3">
	`,
	//
	// textarea element
	//
	"textarea": `
	<div class="mb-3">
		<label for="{{ .ID }}">{{ .Label }}</label>
		<textarea class="form-control" id="{{ .ID }}" rows="4" readonly>{{ .Value }}</textarea>
	</div>
	`,
	//
	// pdf file element
	//
	"pdf_file": `
	<div class="mb-3">
		<div id="{{ .ID }}_pdf"></div>
	</div>
	`,
	//
	// horizonal line (hr) element
	//
	"hr": `<hr>`,
}

// HTMLTemplates maps form element types to default Twitter Bootstrap HTML templates
var HTMLTemplates = make(map[string]*tmpl.Template)

var err error

func init() {
	HTMLTemplates["alert_messages"], err = tmpl.New("alert_messages").Parse(HTMLTemplateStrings["alert_messages"])
	if err != nil {
		fmt.Printf("error parsing the alert_messages: %v template; see: %v\n", len(HTMLTemplateStrings["alert_messages"]), err)
		os.Exit(1)
	}
	HTMLTemplates["select_element"], err = tmpl.New("select_element").Parse(HTMLTemplateStrings["select_element"])
	if err != nil {
		fmt.Printf("error parsing the select_element: %v template; see: %v\n", len(HTMLTemplateStrings["select_element"]), err)
		os.Exit(1)
	}
	HTMLTemplates["textarea"], err = tmpl.New("textarea").Parse(HTMLTemplateStrings["textarea"])
	if err != nil {
		fmt.Printf("error parsing the textarea: %v template; see: %v\n", len(HTMLTemplateStrings["textarea"]), err)
		os.Exit(1)
	}
	HTMLTemplates["pdf_file"], err = tmpl.New("pdf_file").Parse(HTMLTemplateStrings["pdf_file"])
	if err != nil {
		fmt.Printf("error parsing the pdf_file: %v template; see: %v\n", len(HTMLTemplateStrings["pdf_file"]), err)
		os.Exit(1)
	}
	HTMLTemplates["hr"], err = tmpl.New("hr").Parse(HTMLTemplateStrings["hr"])
	if err != nil {
		fmt.Printf("error parsing the hr: %v template; see: %v\n", len(HTMLTemplateStrings["hr"]), err)
		os.Exit(1)
	}
}
