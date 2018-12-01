package id_parser

import (
  "errors"
  "regexp"
  "strconv"
  "strings"
)

type Gender int32

const (
  Male   Gender = 0
  Female Gender = 1
)

type Id struct {
  ProvinceCode string // 1-2位代表省份
  CityCode     string // 3-4位代表城市
  CountyCode   string // 5-6县城/城区
  BornYear     string // 7-10出生年份
  BornMonth    string // 11-12出生月份
  BornDay      string // 13-14出生日期
  Code         string // 14-18最后4位码
  Gender       Gender // 性别
}

var idCardRegexp = regexp.MustCompile(`^(\d{2})(\d{2})(\d{2})(\d{4})(\d{2})(\d{2})(\d{4})`)

func Parse(id string) (entity Id, err error) {
  if len(id) != 18 {
    err = errors.New("id Number must be 18 length")
    return
  }

  result := idCardRegexp.FindStringSubmatch(id)

  if len(result) < 8 {
    err = errors.New("invalid Id")
    return
  }

  entity.ProvinceCode = result[1]
  entity.CityCode = result[2]
  entity.CountyCode = result[3]
  entity.BornYear = result[4]
  entity.BornMonth = result[5]
  entity.BornDay = result[6]
  entity.Code = result[7]

  var idSlice = strings.Split(id, "")
  var index = idSlice[16]
  var gender int64
  if gender, err = strconv.ParseInt(index, 10, 10); err != nil {
    return
  }

  if gender == 1 {
    entity.Gender = Male
  } else {
    entity.Gender = Female
  }

  return
}
