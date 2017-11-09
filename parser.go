package id_parser

import (
  "errors"
  "regexp"
)

// 区号数据来源: http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201703/t20170310_1471429.html

type Id struct {
  province     string
  provinceCode string // 1-2位代表省份
  city         string
  cityCode     string // 3-4位代表城市
  county       string
  countyCode   string // 5-6县城/城区
  bornYear     string // 7-10出生年份
  bornMonth    string // 11-12出生月份
  bornDay      string // 13-14出生日期
  code         string // 14-18最后4位码
}

var idCardRegexp = regexp.MustCompile(`^(\d{2})(\d{2})(\d{2})(\d{4})(\d{2})(\d{2})(\d{4})`)

/**
解析身份证号
 */
func Parse(id string) (entity Id, err error) {
  if len(id) != 18 {
    err = errors.New("Id Number must be 18 length")
    return
  }

  result := idCardRegexp.FindStringSubmatch(id)

  if len(result) < 8 {
    err = errors.New("Invalid Id")
    return
  }

  entity.provinceCode = result[1]
  entity.cityCode = result[2]
  entity.countyCode = result[3]
  entity.bornYear = result[4]
  entity.bornMonth = result[5]
  entity.bornDay = result[6]
  entity.code = result[7]

  // TODO: 校验出生年月是否正确

  return
}

func init() {
  Parse("450122199206013017")
}
