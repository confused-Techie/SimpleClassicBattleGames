package webrequests

import (
  "html/template"
  "net/http"
  "fmt"
  games "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/games"
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
)

func SquaresSquares4Handler(w http.ResponseWriter, r *http.Request) {
  gameID := ObtainGameID(r)

  errFunc := func() {
    errorPage(w, r, gameID)
  }

  newFunc := func() {
    newGameID := games.CreateGameID()
    fmt.Println("New Squares & Squares 4 created with: "+newGameID)

    // now to actually create the progress file.
    createRes := games.CreateGameProgress(newGameID, "squares-sqaures4", GetSignInUser(r))
    if createRes != "Success" {
      errorPage(w, r, "Unable to create Game Progress File Successfully: "+createRes)
    }

    http.Redirect(w, r, "/squares-squares4?game="+newGameID, http.StatusSeeOther)
  }

  defaultFunc := func() {
    // then if thhis is player two joining the game we want to update the second player value of the game progress.

    gameProgress := games.GetGameProgress(gameID)
    curPlayer := GetSignInUser(r)

    if gameProgress.PlayerOne != curPlayer && gameProgress.PlayerTwo == "" {
      games.UpdateGameProgressPlayerTwo(gameID, curPlayer)
    }

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
