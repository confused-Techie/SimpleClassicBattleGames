package games

import (
  "fmt"
  "os"
  "io/ioutil"
  "encoding/json"
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
  "github.com/satori/go.uuid"
)

func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func GetGameCollection() (au *models.GameCollection){
  file, err := os.OpenFile("./internal_data/available_games.json", os.O_RDWR|os.O_APPEND, 0666)
  checkError(err)
  b, err := ioutil.ReadAll(file)
  var gms models.GameCollection
  json.Unmarshal(b, &gms.GameEntry)
  checkError(err)
  return &gms
}

func GetGameProgressCollection() (au *models.GameProgressCollection) {
  file, err := os.OpenFile("./external_data/game_progress.json", os.O_RDWR|os.O_APPEND, 0666)
  checkError(err)
  b, err := ioutil.ReadAll(file)
  var gmpc models.GameProgressCollection
  json.Unmarshal(b, &gmpc.GameProgressEntry)
  checkError(err)
  return &gmpc
}

func SetGameProgressCollection(au *models.GameProgressCollection) () {
  ioutil.WriteFile("./external_data/game_progress.json", au, 0666)
}

func GetGameEntry(gameid string) (au *models.GameEntry) {
  gameCollection := GetGameCollection()

  for _, entry := range gameCollection.GameEntry {
    if entry.ID == gameid {
      return entry
    }
  }

  return nil
}

func CreateGameID() (string) {
  u4 := uuid.NewV4()
  return u4.String()
}

func GetGameProgress(gameid string) (au *models.GameProgressEntry) {
  gameProgressCollection := GetGameProgressCollection()

  for _, entry := range gameProgressCollection.GameProgressEntry {
    if entry.ID == gameid {
      return entry
    }
  }

  return nil
}

func UpdateGameProgress(gameid string) {
  gameProgressCollection := GetGameProgressCollection()

  for _, entry := range gameProgressCollection.GameProgressEntry {
    if entry.ID == gameid {
      // no we want to replace all fields of the progress entry. And save.
    }
  }
}

func CreateGameProgress(gameid string, gametype string, playerone string) (string) {
  gameProgressCollectionOriginal := GetGameProgressCollection()

  // create the new entry to append.
  var newEntryData models.GameProgressEntry
  newEntryData.ID = gameid
  newEntryData.Game = gametype
  newEntryData.PlayerOne = playerone
  newEntryData.Turn = playerone

  // append the new entry onto the original.
  gameProgressCollectionOriginal.GameProgressEntry = append(gameProgressCollectionOriginal.GameProgressEntry, &newEntryData)
  gameProgressCollectionNew, err := json.MarshalIndent(&gameProgressCollectionOriginal.GameProgressEntry, "", "")
  checkError(err)

  // then write this file back
  SetGameProgressCollection(gameProgressCollectionNew)

  // once done we can exit successfully.
  return "Success"
}
