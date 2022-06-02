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
