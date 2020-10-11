package models

import (
  "regexp"
  "strings"
)

type UserValidate struct {
  Name   string
  Email string
  Password string
  Errors  map[string]string
}

var rxEmail = regexp.MustCompile(".+@.+\\..+")

func (msg *UserValidate) Validate() bool {
  msg.Errors = make(map[string]string)

  if strings.TrimSpace(msg.Name) == "" {
    msg.Errors["Name"] = "Please input Name"
  }

  match := rxEmail.Match([]byte(msg.Email))
  if match == false {
    msg.Errors["Email"] = "Please enter a valid email address"
  }

  if strings.TrimSpace(msg.Password) == "" {
    msg.Errors["Password"] = "Please input Password"
  }

  return len(msg.Errors) == 0
}

type ReturnRes struct {
  Code   string `json:"code"`
  Text string `json:"text"`
}

type LoginRequest struct {
  Email string
  Password string
}
