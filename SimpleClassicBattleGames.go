package main

import (
  "net/http"
  "log"
  webrequests "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/webrequests"
)

func main() {

  mux := http.NewServeMux()

  // ========== Standard Page Endpoints =============
  mux.Handle("/", webrequests.LogInCheck(http.HandlerFunc(webrequests.HomePageHandler)))
  mux.Handle("/squares-squares4", http.HandlerFunc(webrequests.SquaresSquares4Handler)) // connect four

  // ======= API endpoints ==========
  mux.Handle("/submit_progress", http.HandlerFunc(webrequests.SubmitProgressHandler))
  mux.Handle("/retreive_progress", http.HandlerFunc(webrequests.RetreiveProgressHandler))

  // we are wrapping the listen in log.Fatal since it will only ever return an error
  log.Fatal(http.ListenAndServe(":8080", mux))
}
