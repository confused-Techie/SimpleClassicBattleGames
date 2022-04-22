package users

import (
  models "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/models"
  "fmt"
  "os"
  "time"
  "io/ioutil"
  "encoding/json"
  "golang.org/x/crypto/bcrypt"
)

func checkError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func ReadUserList() (au *models.UserCollection){
  file, err := os.OpenFile("./external_data/users.json", os.O_RDWR|os.O_APPEND, 0666)
  checkError(err)
  b, err := ioutil.ReadAll(file)
  var usrs models.UserCollection
  json.Unmarshal(b, &usrs.UserEntry)
  checkError(err)
  return &usrs
}

func GetUserDetails(username string) (au *models.UserEntry) {
  userCollection := ReadUserList()

  for _, entry := range userCollection.UserEntry {
    if entry.UserName == username {
      return entry
    }
  }

  return nil
}

func WriteUserEntry(user string, pass string, created string) (bool, error) {
  userCollection := ReadUserList()

  // create new entry to append
  var newUser models.UserEntry
  newUser.UserName = user
  newUser.PassHash = pass
  newUser.CreatedAt = created

  // append the new entry onto the original
  userCollection.UserEntry = append(userCollection.UserEntry, &newUser)
  newUserCollection, err := json.MarshalIndent(&userCollection.UserEntry, "  ", "  ")
  if err != nil {
    return false, err
  }

  ioutil.WriteFile("./external_data/users.json", newUserCollection, 0666)
  return true, nil
}


func CheckValidUser(username string) bool {
  userDetails := GetUserDetails(username)

  if userDetails == nil {
    return false
  }
  return true
}

func AttemptSignIn(user string, pass string) (bool, string) {
  // first we verify that its a valid user
  validUser := CheckValidUser(user)

  if validUser {
    // then we can gather their details
    userDetails := GetUserDetails(user)

    // because the hash's salt will always be different we actually need to check the validity of plaintext to
    // generate a given hash. So we pass the plaintext password first.
    match := CheckPasswordHash(pass, userDetails.PassHash)

    if match {
      fmt.Println("User: ", user)
      fmt.Println("Match: ", match)
      return true, ""
    }
    fmt.Println("User:", user)
    fmt.Println("Match: ", match)
    return false, "Login Failed."
  } else {
    return false, "User does not exist."
  }
}

func CreateUser(user string, pass string) (bool, string) {
  // first make sure a user doesn't already exist with this name.
  validUser := CheckValidUser(user)

  if validUser {
    return false, "User already exists."
  } else {
    hash, err := HashPassword(pass)
    if err != nil {
      return false, err.Error()
    }

    now := time.Now()

    fmt.Println("Creating user...")
    fmt.Println("User: "+user)

    successWrite, err := WriteUserEntry(user, hash, now.String())

    if successWrite {
      return true, ""
    } else {
      return false, err.Error()
    }
  }
}

func HashPassword(pass string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 15)
  return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
