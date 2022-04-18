package webrequests

import(
  "net/http"
)

func ObtainGameID(r *http.Request) (string) {
  keys, ok := r.URL.Query()["game"]

  if !ok || len(keys[0]) < 1 {
    return "Url Param 'game' is missing"
  }

  key := keys[0]

  return string(key)
}

func DetermineGameState(gameID string, errF func(), newF func(), continueF func()) {
  switch(gameID) {
    case "Url Param 'game' is missing":
      errF()
    case "new":
      newF()
    default:
      continueF()
  }
}
