package fake_join

import (
	"reflect"
	"strconv"
)

// FormatNumber 将数字类型转成string类型
//   i 需要转的类型类型
//     1. 可转类型: int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64
//     2. string类型原样输出
//     3. 其他类型返回空字符串
func FormatNumber(i interface{}) string {
	var int64I int64
	switch reflect.TypeOf(i).Kind().String() {
	case "int":
		int64I = int64(i.(int))
	case "int8":
		int64I = int64(i.(int8))
	case "int16":
		int64I = int64(i.(int16))
	case "int32":
		int64I = int64(i.(int32))
	case "int64":
		int64I = i.(int64)
	case "uint":
		int64I = int64(i.(uint))
	case "uint8":
		int64I = int64(i.(uint8))
	case "uint16":
		int64I = int64(i.(uint16))
	case "uint32":
		int64I = int64(i.(uint32))
	case "uint64":
		int64I = int64(i.(uint64))
	case "float32":
		int64I = int64(i.(float32))
	case "float64":
		int64I = int64(i.(float64))
	case "string":
		return i.(string)
	default:
		return ""
	}
	return strconv.FormatInt(int64I, 10)
}
