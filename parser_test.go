package id_parser

import (
  "testing"
  "fmt"
)

func Test_Parser(t *testing.T) {
  entity, err := Parse("450101199801014321")

  if err != nil {
    t.Error("Parse fail")
    return
  }

  fmt.Println(entity)
  t.Log("success.")
}
