package webrequests

import (
  "net/http"
)

func ManifestHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "./assets/static/manifest.json")
}
