package webrequests

import (
  "html/template"
  "net/http"
  logger "github.com/confused-Techie/SimpleClassicBattleGames/pkg/src/logger"
  models "github.com/confused-Techie/SimpleClassicBattleGames/pkg/src/models"
  games "github.com/confused-Techie/SimpleClassicBattleGames/pkg/src/games"
)

func SquaresSquares4Handler(w http.ResponseWriter, r *http.Request) {
  gameID := ObtainGameID(r)

  errFunc := func() {
    errorPage(w, r, gameID)
  }

  newFunc := func() {
    newGameID := games.CreateGameID()
    logger.GameLog.Println("New Squares & Squares 4 created with: "+newGameID)

    // create the progress file
    createRes := games.CreateGameProgress(newGameID, "squares-squares4")
    if createRes != "Success" {
      errorPage(w, r, "Unable to create Game Progress File Successfully: "+createRes)
    }

    http.Redirect(w, r, "/squares-squares4?game="+newGameID, http.StatusSeeOther)
  }

  defaultFunc := func() {

    gameProgress := games.GetGameProgress(gameID)

    gameEntry := games.GetGameEntry("squares-squares4")

    data := models.PageTemplate{
      Title: gameEntry.Title,
      GameRules: gameEntry.Rules,
    }

    templateArray := []string {
      returnTemplate("game.go.html"),
    }

    tmpl["squares-squares4.html"] = template.Must(template.ParseFiles(templateArray...))

    templateError := tmpl["squares-squares4.html"].Execute(w, data)
    StandardPageError(templateError)
  }

  DetermineGameState(gameID, errFunc, newFunc, defaultFunc)
}

func GameHandler(w http.ResponseWriter, r *http.Request, id string) {
  gameID := ObtainGameID(r)

  errFunc := func() {
    errorPage(w, r, gameID)
  }

  newFunc := func() {
    newGameID := games.CreateGameID()
    logger.GameLog.Println("New "+id+" created with: "+newGameID)

    // create progress file
    createRes := games.CreateGameProgress(newGameID, id)
    if createRes != "Success" {
      errorPage(w, r, "Unable to create Game Progress File Successfully: "+createRes)
    }

    http.Redirect(w, r, "/"+id+"?game="+newGameID, http.StatusSeeOther)
  }

  defaultFunc := func() {
    gameProgress := games.GetGameProgress(gameID)

    gameEntry := games.GetGameEntry(id)

    data := models.PageTemplate{
      Title: gameEntry.Title,
      GameRules: gameEntry.Rules,
    }

    templateArray := []string {
      returnTemplate("game.go.html"),
    }

    tmpl["game.html"] = template.Must(template.ParseFiles(templateArray...))

    templateError := tmpl["game.html"].Execute(w, data)
    StandardPageError(templateError)
  }

  DetermineGameState(gameID, errFunc, newFunc, defaultFunc)
}
