package models

import (

)

type PageTemplate struct {
  Title string
  GameRules string
  Theme string
  CSS []string
  JS []string
  Data interface{}
  TargetStrings map[string]string
  DefaultStrings map[string]string
  TargetLanguage string
}

type GameEntry struct {
  Title string `json:"title"`
  ID string `json:"id"`
  Link string `json:"link"`
  Rules string `json:"rules"`
}

type GameCollection struct {
  GameEntry []*GameEntry
}

type GameProgressEntry struct {
  ID string `json:"id"`
  Game string `json:"game"`
  PlayerOne string `json:"playerone"`
  PlayerTwo string `json:"playertwo"`
  Turn string `json:"turn"`
  Board interface{} `json:"board"`
}

type GameProgressCollection struct {
  GameProgressEntry []*GameProgressEntry
}
