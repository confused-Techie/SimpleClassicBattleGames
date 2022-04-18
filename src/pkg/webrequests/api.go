package webrequests

import(
  "encoding/json"
  "fmt"
  "net/http"
  games "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/games"
)

func SubmitProgressHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {

  } else {
    fmt.Println("Non-Post Request submitted to POST only JSON SubmitProgressHandler")
    json.NewEncoder(w).Encode("Non-Post Request submitted to POST only JSON SubmitProgressHandler")
  }
}

func RetreiveProgressHandler(w http.ResponseWriter, r *http.Request) {

  gameID := ObtainGameID(r)

  errFunc := func() {
    errorJSON(w, r, gameID)
  }

  newFunc := func() {
    errorJSON(w, r, "This 'new' game has no progress that can be reported.")
  }

  defaultFunc := func() {
    // now knowing that this is valid game progress, we can query the save file.
    gameProgress := games.GetGameProgress(gameID)

    if gameProgress == nil {
      errorJSON(w, r, "No Progress for game can be found: "+gameID)
    }

    json.NewEncoder(w).Encode(gameProgress)
  }

  DetermineGameState(gameID, errFunc, newFunc, defaultFunc)

}

func errorJSON(w http.ResponseWriter, r *http.Request, err string) {
  json.NewEncoder(w).Encode(err)
}
