package main

import (
  "net/http"
  "github.com/spf13/viper"
  "os"
  logger "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/logger"
  webrequests "github.com/confused-Techie/SimpleClassicBattleGames/src/pkg/webrequests"
)

func main() {
  // setup viper
  viper.SetConfigName("app")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")
  viper.ReadInConfig()

  logger.InfoLogger.Println("Setup app.yaml arguments...")

  logger.InfoLogger.Println("SimpleClassicBattleGames Starting...")

  // since we know we only will have our variables set as env variables in production
  production := os.Getenv("PRODUCTION")

  if production != "" {
    logger.InfoLogger.Printf("Running in Production Environment")
  } else {
    logger.InfoLogger.Printf("Running in Development Environment")
  }

  mux := http.NewServeMux()

  // =================== Standard Page Endpoints ==================
  mux.Handle("/", http.HandlerFunc(webrequests.HomeHandler))

  // =================== Game Endpoints ===========================
  //mux.Handle("/squares-squares4", http.HandlerFunc(webrequests.SquaresSquares4Handler))
  mux.Handle("/square-squares4", http.HandlerFunc(webrequests.GameHandler("squares-squares4")))

  // =================== Asset Endpoints ==========================

  // =================== API Endpoints ============================

  // =================== Listen ===================================

  port := os.Getenv("PORT")

  if port == "" {
    port = viper.GetString("env_variables.PORT")
    logger.WarningLogger.Println("Port not available via Environment Variables. Falling back to Config File")
  }

  logger.InfoLogger.Printf("Listening on %v...", port)
  logger.ErrorLogger.Fatal(http.ListenAndServe(":"+port, mux))
}
