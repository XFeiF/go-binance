package common

import (
	"bytes"
	"fmt"
	"math"
)

// AmountToLotSize converts an amount to a lot sized amount
// 将amount转换为一个lot（手数）大小的数量
func AmountToLotSize(lot float64, precision int, amount float64) float64 {
	// math.Floor(amount/lot) 整数倍，向下取整
	// *lot 乘以lot，得到lot大小的数量
	// math.Pow10(precision) 10的precision次方
	// math.Trunc 去掉小数部分
	// 不直接使用math.Floor(amount/lot)*lot的原因：有额外的精度要求
	return math.Trunc(math.Floor(amount/lot)*lot*math.Pow10(precision)) / math.Pow10(precision)
}

// ToJSONList convert v to json list if v is a map
func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}

func ToInt(digit interface{}) (i int, err error) {
	if intVal, ok := digit.(int); ok {
		return int(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}

func ToInt64(digit interface{}) (i int64, err error) {
	if intVal, ok := digit.(int); ok {
		return int64(intVal), nil
	}
	if floatVal, ok := digit.(float64); ok {
		return int64(floatVal), nil
	}
	return 0, fmt.Errorf("unexpected digit: %v", digit)
}
