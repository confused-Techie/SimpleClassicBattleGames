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
  fmt.Println("GetGameProgressCollection called")
  file, err := os.OpenFile("./external_data/game_progress.json", os.O_RDWR|os.O_APPEND, 0666)
  checkError(err)
  b, err := ioutil.ReadAll(file)
  var gmpc models.GameProgressCollection
  json.Unmarshal(b, &gmpc.GameProgressEntry)
  checkError(err)
  return &gmpc
}

func SetGameProgressCollection(au []byte) () {
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
    fmt.Println("Checking for Matching GameID: "+entry.ID)
    if entry.ID == gameid {
      fmt.Println("Game progress found")
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

func UpdateGameProgressPlayerTwo(gameid string, player string) {
  gameProgressCollection := GetGameProgressCollection()

  for _, entry := range gameProgressCollection.GameProgressEntry {
    if entry.ID == gameid {
      entry.PlayerTwo = player
      if entry.Turn == "" {
        entry.Turn = player
      }
    }
  }

  newGameProgressCollection, err := json.MarshalIndent(&gameProgressCollection, "  ", "  ")
  checkError(err)
  SetGameProgressCollection(newGameProgressCollection)
}

func UpdateGameProgressBoard(gameid string, board interface{}) {
  gameProgressCollection := GetGameProgressCollection()

  for _, entry := range gameProgressCollection.GameProgressEntry {
    if entry.ID == gameid {
      entry.Board = board
    }
  }

  newGameProgressCollection, err := json.MarshalIndent(&gameProgressCollection, "  ", "  ")
  checkError(err)
  SetGameProgressCollection(newGameProgressCollection)
}

func UpdateGameProgressTurn(gameid string, turn string) {
  gameProgressCollection := GetGameProgressCollection()

  for _, entry := range gameProgressCollection.GameProgressEntry {
    if entry.ID == gameid {
      entry.Turn = turn
    }
  }

  newGameProgressCollection, err := json.MarshalIndent(&gameProgressCollection, "  ", "  ")
  checkError(err)
  SetGameProgressCollection(newGameProgressCollection)
}

func UpdateGameProgressWinner(gameid string, winner string) {
  gameProgressCollection := GetGameProgressCollection()

  for _, entry := range gameProgressCollection.GameProgressEntry {
    if entry.ID == gameid {
      entry.Winner = winner
    }
  }

  newGameProgressCollection, err := json.MarshalIndent(&gameProgressCollection, "  ", "  ")
  checkError(err)
  SetGameProgressCollection(newGameProgressCollection)
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
  gameProgressCollectionNew, err := json.MarshalIndent(&gameProgressCollectionOriginal.GameProgressEntry, "  ", "  ")
  checkError(err)

  // then write this file back
  SetGameProgressCollection(gameProgressCollectionNew)

  // once done we can exit successfully.
  return "Success"
}
