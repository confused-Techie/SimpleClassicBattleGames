package webrequests

import (
  "net/http"
  "fmt"
  "time"
  users "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/users"
)

func LogInCheck(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // first check for cookie existance
    user, err := r.Cookie("username")
    // while in chrome this works fine, it seems to crash in Firefox, due to an unhandled error here.
    // an error grabbing the cookie will also redirect to Login
    if err != nil {
      fmt.Println("Failed to grab User Cookies. Assumed as not logged in.")
      http.Redirect(w, r, "sign-in", http.StatusSeeOther)
    } else {
      validUser := users.CheckValidUser(user.Value)

      if !validUser {
        fmt.Println("User is not signed in. Redirecting...")
        http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
      } else {
        h.ServeHTTP(w, r)
      }
    }
    //fmt.Println("LogInCheck Cookie Gathering Error: "+err.Error())
    //validUser := users.CheckValidUser(user.Value)

    //if !validUser {
    //  fmt.Println("User is not signed in. Redirecting...")
    //  http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
    //}

    //h.ServeHTTP(w, r)
  })
}

func CreateSignInCookie(user string, w http.ResponseWriter) {
  expiration := time.Now().Add(365 * 24 * time.Hour) // TODO:: Make this a responable time
  cookie := http.Cookie{Name: "username", Value: user, Expires: expiration}
  http.SetCookie(w, &cookie)
}
