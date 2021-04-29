package handlers

import (
	"github.com/olegvbelov/okc46go/internal/config"
	"github.com/olegvbelov/okc46go/internal/forms"
	"github.com/olegvbelov/okc46go/internal/models"
	"github.com/olegvbelov/okc46go/internal/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	isActive := make(map[string]bool)
	isActive["home"] = true

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		IsActive: isActive,
		Services: models.Services,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["page_title"] = "О НАС"

	isActive := make(map[string]bool)
	isActive["about"] = true

	// send data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IsActive:  isActive,
	})
}

func (m *Repository) Services(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["page_title"] = "УСЛУГИ"

	isActive := make(map[string]bool)
	isActive["services"] = true

	// send data to the template
	render.RenderTemplate(w, r, "services.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IsActive:  isActive,
		Contacts:  models.Offices,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	var emptySiteMessage models.SiteMessage
	data := make(map[string]interface{})
	data["siteMessage"] = emptySiteMessage

	stringMap := make(map[string]string)
	stringMap["page_title"] = "КОНТАКТЫ"

	isActive := make(map[string]bool)
	isActive["contact"] = true

	// send data to the template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
		IsActive:  isActive,
		Contacts:  models.Offices,
	})
}

func (m *Repository) SendEmail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	siteMessage := models.SiteMessage{
		Name:    r.Form.Get("name"),
		Email:   r.Form.Get("email"),
		Phone:   r.Form.Get("phone"),
		Message: r.Form.Get("message"),
	}

	form := forms.New(r.PostForm)

	form.Required("name", "email", "message")
	form.MinLength("name", 3, r)
	form.IsMail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		siteMessage.Email = ""
		data["siteMessage"] = siteMessage
		render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
			Form:     form,
			Data:     data,
			Contacts: models.Offices,
		})
		return
	}
}
