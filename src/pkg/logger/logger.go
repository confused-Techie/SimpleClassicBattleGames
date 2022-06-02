package logger


import (
  "log"
  "os"
)

var (
  // Proper comments for this file can be found:
  // https://github.com/confused-Techie/Quotle/blob/main/src/pkg/logger/logger.go
  WarningLogger *log.Logger
  InfoLogger *log.Logger
  ErrorLogger *log.Logger
)

func init() {
  production := os.Getenv("PRODUCTION")

  if production != "" {

    InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
  } else {
    // TODO: local file write
    InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
  }
}

func main() {
  InfoLogger.Println("Starting Logger...")
}
