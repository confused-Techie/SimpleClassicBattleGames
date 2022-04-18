package webrequests

import (
  "fmt"
)

func StandardPageError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
