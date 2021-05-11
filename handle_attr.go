package lgorm

import (
	"reflect"
	"regexp"
)

// 操作attr
func (db *Db) handleAttr(dest interface{}, handleType string) {
	var model interface{}
	if db.Statement.Model.IsCall {
		model = db.getModel(db.Statement.Model.Params[0])
	}
	if model == nil {
		model = db.getModel(dest)
	}
	if model == nil {
		return
	}
	if db.isStruct(dest) {
		db.handleStructAttr(model, dest, handleType)
	} else {
		db.handleInterfaceAttr(model, dest, handleType)
	}
}

// 操作interface类型
func (db *Db) handleInterfaceAttr(model interface{}, dest interface{}, handleType string) {
	if reflect.ValueOf(dest).Kind() == reflect.Ptr {
		dest = reflect.ValueOf(dest).Elem().Interface()
	}
	fieldMethodMap := db.getFieldMethodAttr(model, handleType)
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}
	if db.isList(dest) {
		for _, v := range dest.([]map[string]interface{}) {
			db.runInterfaceAttr(modelValue, v, fieldMethodMap)
		}
	} else {
		db.runInterfaceAttr(modelValue, dest.(map[string]interface{}), fieldMethodMap)
	}
}

func (db *Db) runInterfaceAttr(modelValue reflect.Value, value map[string]interface{}, fieldMethodMap map[string]string) {
	for field, val := range value {
		if fieldMethodMap[field] == "" {
			continue
		}
		if val == nil {
			continue
		}
		rValue := []reflect.Value{reflect.ValueOf(val)}
		callRs := modelValue.MethodByName(fieldMethodMap[field]).Call(rValue)
		value[field] = callRs[0].Interface()
	}
}

// 操作struct
func (db *Db) handleStructAttr(model interface{}, dest interface{}, handleType string) {
	fieldMethodMap := db.getFieldMethodAttr(model, handleType)
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	}
	if value.Kind() == reflect.Struct {
		db.runStructAttr(modelValue, value, fieldMethodMap)
	} else {
		structLen := value.Len()
		for i := 0; i < structLen; i++ {
			db.runStructAttr(modelValue, value.Index(i), fieldMethodMap)
		}
	}
}

func (db *Db) runStructAttr(modelValue reflect.Value, value reflect.Value, fieldMethodMap map[string]string) {
	numField := value.NumField()
	for i := 0; i < numField; i++ {
		field := value.Type().Field(i).Name
		if fieldMethodMap[field] == "" {
			continue
		}
		if value.Field(i).Type() != modelValue.FieldByName(field).Type() {
			continue
		}
		rValue := []reflect.Value{reflect.ValueOf(value.Field(i).Interface())}
		callRs := modelValue.MethodByName(fieldMethodMap[field]).Call(rValue)
		value.Field(i).Set(reflect.ValueOf(callRs[0].Interface()))
	}
}

func (db *Db) getFieldMethodAttr(model interface{}, handleType string) map[string]string {
	result := map[string]string{}
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	numField := value.NumField()
	numMethod := value.NumMethod()
	for i := 0; i < numMethod; i++ {
		reg := regexp.MustCompile(`^` + handleType + `(.+)Attr$`)
		methodName := value.Type().Method(i).Name
		fields := reg.FindStringSubmatch(methodName)
		if len(fields) > 0 {
			field := fields[1]
			for j := 0; j < numField; j++ {
				if field == value.Type().Field(j).Name {
					key := value.Type().Field(j).Tag.Get("json")
					result[key] = methodName
					result[field] = methodName
				}
			}
		}
	}
	return result
}
