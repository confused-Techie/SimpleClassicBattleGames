package games

import (
  logger "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/logger"
  "github.com/satori/go.uuid"
)

func checkError(err error) {
  if err != nil {
    logger.ErrorLogger.Println(err)
  }
}

func GetGameCollection() () {

}

func GetGameProgressCollection() () {

}

func SetGameProgressCollection(au []byte) {

}

func GetGameEntry() () {

}

func CreateGameID() (string) {
  u4 := uuid.NewV4()
  return u4.String()
}
