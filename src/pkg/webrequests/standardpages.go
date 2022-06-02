package webrequests

import (
  "net/http"
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
  "html/template"
)

var tmpl = make(map[string]*template.Template)

func returnTemplate(t string) string {
  return "./templates/" + t
}

func errorPage(w http.ResponseWriter, r *http.Request, err string) {
  data := models.PageTemplate{
    Title: "Something went wrong.",
    Data: err,
  }

  templateArray := []string{
    returnTemplate("errorPage.go.html"),
  }

  tmpl["errorPage.html"] = template.Must(template.ParseFiles(templateArray...))
  templateError := tmpl["errorPage.html"].Execute(w, data)
  StandardPageError(templateError)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  data := models.PageTemplate{
    Title: "SimpleClassicBattleGames - Home",
  }

  templateArray := []string{
    returnTemplate("homePage.go.html"),
  }

  tmpl["homePage.html"] = template.Must(template.ParseFiles(templateArray...))

  templateError := tmpl["homePage.html"].Execute(w, data)
  StandardPageError(templateError)
}
