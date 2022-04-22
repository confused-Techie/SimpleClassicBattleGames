package webrequests

import(
  "encoding/json"
  "fmt"
  "net/http"
  games "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/games"
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
  "io/ioutil"
)

func SubmitProgressHandler(w http.ResponseWriter, r *http.Request) {
  // this will take in a post json object with all possible modifiable values, and check if they are specified,
  // calling the function within games to change that specific value.
  // Rather than a function in games having to handle all of this alone.
  if r.Method == "POST" {
    body, err := ioutil.ReadAll(r.Body)
    JSONReturnError(w, err)

    var updatedGameEntry models.GameProgressEntry
    err = json.Unmarshal(body, &updatedGameEntry)
    JSONReturnError(w, err)

    // now after knowing the data can be unmarshalled into the struct, we can marshall it properly
    origGameEntry := games.GetGameProgress(updatedGameEntry.ID)

    // Now to go through each modifiable element, and update it.
    if origGameEntry.Turn != updatedGameEntry.Turn {
      games.UpdateGameProgressTurn(updatedGameEntry.ID, updatedGameEntry.Turn)
    }

    if origGameEntry.Winner != updatedGameEntry.Winner {
      games.UpdateGameProgressWinner(updatedGameEntry.ID, updatedGameEntry.Winner)
    }

    if origGameEntry.Board != updatedGameEntry.Board {
      games.UpdateGameProgressBoard(updatedGameEntry.ID, updatedGameEntry.Board)
    }

    // now once everything has been checked and updated, we can return success.
    json.NewEncoder(w).Encode("Success")

  } else {
    fmt.Println("Non-Post Request submitted to POST Only JSON APIUpdateGameHandler.")
    json.NewEncoder(w).Encode("Non-Post Request submitted to POST Only JSON APIUpdateGameHandler.")
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
