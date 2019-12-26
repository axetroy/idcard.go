package idcard

import (
  "errors"
  "fmt"
  "regexp"
  "strconv"
)

type Gender int32

// 123456 12345678 123 x
// 地区码{6}   生日编码{8} 顺序码{3} 校验码{1}
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

var (
  // 身份证的正则表达式
  idCardRegexp = regexp.MustCompile(`^(\d{2})(\d{2})(\d{2})(\d{4})(\d{2})(\d{2})(\d{3}(\d|x))$`)
  // 每个余数对应的校验码
  parity = []string{"1", "0", "x", "9", "8", "7", "6", "5", "4", "3", "2"}
  // 身份证前 17 未数的权重
  weights = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
)

const (
  Male   Gender = 0 // 男性
  Female Gender = 1 // 女性
)

// 根据身份证号生成校验码
func generateSumCodeFromId(id string) (string, error) {

  // 把身份证号对应的数字 x 对应的系数, 然后计算总和
  var sum = 0
  for i := 0; i < 17; i++ {
    number, err := strconv.ParseInt(string(id[i]), 10, 10)
    if err != nil {
      return "", err
    }
    sum += int(number) * weights[i]
  }

  mod := sum % 11
  sumCode := parity[mod] // 通过计算出来的校验码

  return sumCode, nil
}

// 计算身份证的校验码
func sumId(id string) error {
  // 身份证后的最后以为校验码
  checkCode := string(id[len(id)-1])

  sumCode, err := generateSumCodeFromId(id)

  if err != nil {
    return err
  }

  // 如果校验码不正确，则是无效身份证
  if sumCode != checkCode {
    return fmt.Errorf("expect sum code `%s` but found `%s`", sumCode, checkCode)
  }

  return nil
}

// 解析身份证号
func Parse(id string) (entity Id, err error) {
  if len(id) != 18 {
    err = errors.New("id number must be 18 length")
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

  // 奇数是男性，偶数是女性
  if entity.Code[0]%2 == 1 {
    entity.Gender = Male
  } else {
    entity.Gender = Female
  }

  // 校验身份证的校验码
  err = sumId(id)

  return
}
