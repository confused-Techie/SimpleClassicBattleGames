package webrequests

import (
  logger "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/logger"
  "net/http"
  "encoding/json"
)

func StandardPageError(err error) {
  if err != nil {
    logger.ErrorLogger.Println(err)
  }
}

func JSONReturnError(w http.ResponseWriter, err error) {
  if err != nil {
    logger.ErrorLogger.Println(err)
    json.NewEncoder(w).Encode(err)
  }
}
