package models

import "github.com/olegvbelov/okc46go/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	Error      string
	Form       *forms.Form
	IsActive   map[string]bool
	Services   []OkcService
	Categories []Category
	Contacts   []Contact
	Phone      string
	Shortphone string
	Service    OkcService
}
