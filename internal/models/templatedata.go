package models

import "github.com/Daviddlh1/bookings/internal/forms"

// TemplateDate holds data from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
