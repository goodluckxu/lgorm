package lgorm

import (
	"fmt"
	"gorm.io/datatypes"
	"time"
)

var (
	handleAttrType handleAttrTypeInterface
	attrTypeMap    = map[string]func(value interface{}) interface{}{
		"datatypes.JSON": handleAttrType.datatypesJson,
		"datatypes.Date": handleAttrType.datatypesDate,
	}
)

func (db *Db) changeAttrType(attrType string, value interface{}) interface{} {
	if attrTypeMap[attrType] == nil {
		return value
	}
	return attrTypeMap[attrType](value)
}

type handleAttrTypeInterface struct {
}

func (h handleAttrTypeInterface) datatypesJson(value interface{}) interface{} {
	newJson := datatypes.JSON{}
	_ = newJson.UnmarshalJSON([]byte(value.(string)))
	return newJson
}

func (h handleAttrTypeInterface) datatypesDate(value interface{}) interface{} {
	newDate := datatypes.Date{}
	timeByte,er := value.(time.Time).MarshalJSON()
	fmt.Println(string(timeByte),er)
	_ = newDate.UnmarshalJSON(timeByte)
	return newDate
}
