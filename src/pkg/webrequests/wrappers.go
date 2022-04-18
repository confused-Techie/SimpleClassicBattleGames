package webrequests

import (
  "net/http"
  "fmt"
)

func LogInCheck(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Log In Check sorta")
    // as a test lets create a permenant cookie
    expiration := time.Now().Add(365 * 24 * time.Hour)
    cookie := http.Cookie{Name: "username", Value: "confused-Techie", Expires: expiration}
    http.SetCookie(w, &cookie)
    // first check for cookie existance
    for _, c := range r.Cookies() {
      fmt.Println(c)
    }
    h.ServeHTTP(w, r)
  })
}
