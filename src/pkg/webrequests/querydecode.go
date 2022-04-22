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

func GetSignInMsg(r *http.Request) string {
  keys, ok := r.URL.Query()["sign-in-error"]

  if !ok || len(keys[0]) < 1 {
    return ""
  }

  key := keys[0]
  return string(key)
}

func GetCreateUserMsg(r *http.Request) string {
  keys, ok := r.URL.Query()["create-user-error"]

  if !ok || len(keys[0]) < 1 {
    return ""
  }

  key := keys[0]
  return string(key)
}

func GetSignInUser(r *http.Request) string {
  user, _ := r.Cookie("username")

  return user.Value
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
