package lgorm

import (
	"reflect"
	"regexp"
)

// 操作attr
func (db *Db) handleAttr(dest interface{}, handleType string) (newDest interface{}) {
	if !db.isAttr(dest) {
		return
	}
	if handleType == "SetOne" {
		handleType = "Set"
	}
	if handleType == "GetOne" {
		handleType = "Get"
	}
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
		newDest = db.handleStructAttr(model, dest, handleType)
	} else {
		db.handleInterfaceAttr(model, dest, handleType)
	}
	return
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
		findMethod := modelValue.MethodByName(fieldMethodMap[field])
		funInTypeInterface := findMethod.Type().In(0)
		funInType := funInTypeInterface.String()
		valueType := reflect.TypeOf(val).String()
		reg := regexp.MustCompile(`^\**\[\]` + funInType + `$`)
		if reg.MatchString(valueType) {
			dataValue := reflect.ValueOf(val)
			if dataValue.Kind() == reflect.Ptr {
				dataValue = dataValue.Elem()
			}
			dataNum := dataValue.Len()
			for i := 0; i < dataNum; i++ {
				var newVal interface{}
				if funInTypeInterface ==
					reflect.TypeOf(dataValue.Index(i).Interface()) {
					newVal = dataValue.Index(i).Interface()
				} else {
					newVal = db.changeAttrType(funInType,
						dataValue.Index(i).Interface())
				}
				rValue := []reflect.Value{reflect.ValueOf(newVal)}
				callRs := findMethod.Call(rValue)
				reflect.ValueOf(val).Elem().Index(i).Set(reflect.ValueOf(callRs[0].Interface()))
			}
		} else {
			var newVal interface{}
			if funInTypeInterface == reflect.TypeOf(val) {
				newVal = val
			} else {
				newVal = db.changeAttrType(funInType, val)
			}
			rValue := []reflect.Value{reflect.ValueOf(newVal)}
			callRs := findMethod.Call(rValue)
			value[field] = callRs[0].Interface()
		}
	}
}

// 操作struct
func (db *Db) handleStructAttr(model interface{}, dest interface{}, handleType string) (newDest interface{}) {
	fieldMethodMap := db.getFieldMethodAttr(model, handleType)
	value := reflect.ValueOf(dest)
	if value.Kind() != reflect.Ptr {
		newValue := reflect.New(value.Type())
		newValue.Elem().Set(value)
		value = newValue
		newDest = value.Interface()
	}
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
	return
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

// 是否可执行attr内容
func (db *Db) isAttr(dest interface{}) bool {
	if dest == nil {
		return false
	}
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	// 空列表不处理
	if value.Kind() == reflect.Slice ||
		value.Kind() == reflect.Array {
		total := value.Len()
		if total == 0 {
			return false
		}
	}
	// 空结构体不处理
	if value.Kind() == reflect.Struct {
		if reflect.DeepEqual(value.Interface(), reflect.New(value.Type()).Elem().Interface()) {
			return false
		}
	}
	// 不等于类型
	if !db.isType(dest) {
		return false
	}
	return true
}

// 是否去掉排除类型
func (db *Db) isType(dest interface{}) bool {
	if dest == nil {
		return true
	}
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	notTypeList := []string{"clause.Expr", "lgorm.Db", "gorm.DB"}
	for _, notType := range notTypeList {
		if value.Type().String() == notType {
			return false
		}
	}
	if value.Type().String() == "map[string]interface {}" {
		for _, destMap := range value.Interface().(map[string]interface{}) {
			if !db.isType(destMap) {
				return false
			}
		}
	} else if value.Type().String() == "[]interface {}" {
		for _, destMap := range value.Interface().([]interface{}) {
			if !db.isType(destMap) {
				return false
			}
		}
	}
	return true
}
