package webrequests

import (
  "html/template"
  "net/http"
  "fmt"
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
  games "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/games"
)

var tmpl = make(map[string]*template.Template)

func returnTemplate(t string) string {
  return "./templates/" + t
}

func returnSubTemplate(t string) string {
  return "./templates/components/" + t
}

func errorPage(w http.ResponseWriter, r *http.Request, err string) {

  data := models.PageTemplate{
    Title: "Something went wrong.",
    Data: err,
  }

  templateArray := []string{
    returnTemplate("errorPage.go.html"),
    returnSubTemplate("head.go.html"),
  }

  tmpl["errorPage.html"] = template.Must(template.ParseFiles(templateArray...))
  templateError := tmpl["errorPage.html"].Execute(w, data)
  StandardPageError(templateError)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

  data := models.PageTemplate{
    Title: "SimpleClassicBattleGames - Home",
    Data: games.GetGameCollection(),
  }

  templateArray := []string{
    returnTemplate("homePage.go.html"),
    returnSubTemplate("head.go.html"),
  }

  tmpl["homePage.html"] = template.Must(template.ParseFiles(templateArray...))

  templateError := tmpl["homePage.html"].Execute(w, data)
  StandardPageError(templateError)
}

func SquaresSquares4Handler(w http.ResponseWriter, r *http.Request) {
  gameID := ObtainGameID(r)

  errFunc := func() {
    errorPage(w, r, gameID)
  }

  newFunc := func() {
    newGameID := games.CreateGameID()
    fmt.Println("New Squares & Squares 4 created with: "+newGameID)

    // now to actually create the progress file.
    createRes := games.CreateGameProgress(gameID, "squares-sqaures4", "playerone")
    if createRes != "Success" {
      erroPage(w, r, "Unable to create Game Progress File Successfully: "+createRes)
    }
    
    http.Redirect(w, r, "/squares-squares4?game="+newGameID, http.StatusSeeOther)
  }

  defaultFunc := func() {
    gameEntry := games.GetGameEntry("squares-squares4")

    data := models.PageTemplate{
      Title: gameEntry.Title,
      GameRules: gameEntry.Rules,
    }

    templateArray := []string{
      returnTemplate("game.go.html"),
      returnSubTemplate("head.go.html"),
    }

    tmpl["squares-squares4.html"] = template.Must(template.ParseFiles(templateArray...))

    templateError := tmpl["squares-squares4.html"].Execute(w, data)
    StandardPageError(templateError)

  }

  DetermineGameState(gameID, errFunc, newFunc, defaultFunc)

}
