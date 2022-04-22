package webrequests

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func StandardPageError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func JSONReturnError(w http.ResponseWriter, err error) {
  if err != nil {
    fmt.Println(err)
    json.NewEncoder(w).Encode(err)
  }
}
