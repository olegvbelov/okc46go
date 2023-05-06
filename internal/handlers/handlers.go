package handlers

import (
	fmt "fmt"
	"github.com/go-chi/chi"
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/olegvbelov/okc46go/internal/config"
	"github.com/olegvbelov/okc46go/internal/forms"
	"github.com/olegvbelov/okc46go/internal/models"
	"github.com/olegvbelov/okc46go/internal/render"
	"log"
	"net/http"
	"strconv"
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

func (m *Repository) Page404(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["page_title"] = "ОШИБКА 404"

	// send data to the template
	render.RenderTemplate(w, r, "404.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Details(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatal(err)
	}
	service := models.Services[intId-1]
	stringMap := make(map[string]string)
	stringMap["page_title"] = "Детали"
	stringMap["id"] = id
	isActive := make(map[string]bool)
	isActive["details"] = true
	render.RenderTemplate(w, r, "details.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IsActive:  isActive,
		Service:   service,
	})
}

func (m *Repository) Sitemap(w http.ResponseWriter, r *http.Request) {
	sm := stm.NewSitemap(1)
	sm.SetDefaultHost("https://okc46.ru")

	sm.Create()
	sm.Add(stm.URL{{"loc", "/"}, {"changefreq", "daily"}})
	sm.Add(stm.URL{{"loc", "/services"}, {"changefreq", "daily"}})
	sm.Add(stm.URL{{"loc", "/about"}, {"changefreq", "daily"}})
	sm.Add(stm.URL{{"loc", "/contact"}, {"changefreq", "daily"}})
	//sm.Add(stm.URL{{}})
	w.Write(sm.XMLContent())
	return
}

func (m *Repository) Robot(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "User-agent: *\n")
	fmt.Fprintf(w, "Disallow: /admin\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "User-agent: Yandex\n")
	fmt.Fprintf(w, "Disallow: /admin\n")
	fmt.Fprintf(w, "Disallow: /*?\n")
	fmt.Fprintf(w, "Host: okc46.ru\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "User-agent: Googlebot\n")
	fmt.Fprintf(w, "Allow: *.css\n")
	fmt.Fprintf(w, "Allow: *.js\n")
	fmt.Fprintf(w, "Disallow: /admin\n")
	fmt.Fprintf(w, "Disallow: /*?\n")
	fmt.Fprintf(w, "Sitemap: https://okc46.ru/sitemap.xml\n")
}
