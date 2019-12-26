package idcard_test

import (
  "testing"

  "github.com/axetroy/idcard.go"
  "github.com/stretchr/testify/assert"
)

func Test_Parser(t *testing.T) {
  id, err := idcard.Parse("11012219820101101x")

  assert.Nil(t, err)
  assert.Equal(t, id.Gender, idcard.Male)
  assert.Equal(t, id.ProvinceCode, "11")
  assert.Equal(t, id.CityCode, "01")
  assert.Equal(t, id.CountyCode, "22")
  assert.Equal(t, id.BornYear, "1982")
  assert.Equal(t, id.BornMonth, "01")
  assert.Equal(t, id.BornDay, "01")
  assert.Equal(t, id.Code, "101x")
}
