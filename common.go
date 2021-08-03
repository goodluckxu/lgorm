package lgorm

import (
	"encoding/json"
	"gorm.io/gorm"
	"reflect"
)

// 是否是slice|array或slice|array指针
func (db *Db) isList(dest interface{}) bool {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() == reflect.Slice ||
		value.Kind() == reflect.Array {
		return true
	}
	return false
}

// 是否是结构体或slice|array结构体或结构体指针或slice|array结构体指针
func (db *Db) isStruct(dest interface{}) bool {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() == reflect.Struct {
		return true
	}
	if value.Kind() == reflect.Slice ||
		value.Kind() == reflect.Array {
		if value.Len() > 0 {
			if value.Index(0).Type().Kind() == reflect.Struct {
				return true
			}
		}
	}
	return false
}

// 获取model
func (db *Db) getModel(dest interface{}) interface{} {
	value := reflect.ValueOf(dest)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() == reflect.Struct {
		return dest
	}
	if value.Kind() == reflect.Slice ||
		value.Kind() == reflect.Array {
		if value.Len() == 0 {
			nilJson := `[{}]`
			err := json.Unmarshal([]byte(nilJson), &dest)
			if err != nil {
				return nil
			}
		}
		if value.Index(0).Type().Kind() == reflect.Struct {
			return value.Index(0).Interface()
		}
	}
	return nil
}

// 获取传入的参数
func (db *Db) getInputDataList(name string) interface{} {
	value := reflect.ValueOf(db.Statement)
	numField := value.NumField()
	for i := 0; i < numField; i++ {
		if value.Type().Field(i).Name == name {
			return value.Field(i).Interface()
		}
	}
	return nil
}

// 执行结构体的方法
func (db *Db) runStructFunc(name string, value []reflect.Value) {
	newValue := []reflect.Value{}
	for _, v := range value {
		newValue = append(newValue, db.handleReflectValue(v))
	}
	dbValue := reflect.ValueOf(db.DB)
	callRs := dbValue.MethodByName(name).Call(newValue)
	if callRs[0].Interface() != nil {
		if reflect.ValueOf(callRs[0].Interface()).Type().String() == "*gorm.DB" {
			db.DB = callRs[0].Interface().(*gorm.DB)
			return
		}
	}
	for _, rs := range callRs {
		db.otherReturn = append(db.otherReturn, rs.Interface())
	}
}

func (db *Db) handleReflectValue(value reflect.Value) reflect.Value {
	if !value.IsValid() {
		return reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())
	}
	switch value.Type().String() {
	case "*lgorm.Db":
		valueDb := value.Interface().(*Db)
		valueDb.getChainAbleInstance()
		return reflect.ValueOf(valueDb.DB)
	case "map[string]interface {}":
		newMap := map[string]interface{}{}
		for k, v := range value.Interface().(map[string]interface{}) {
			newValue := db.handleReflectValue(reflect.ValueOf(v))
			newMap[k] = newValue.Interface()
		}
		return reflect.ValueOf(newMap)
	}
	return value
}

// 初始化statement
func (db *Db) initStatement() {
	db.Statement = *new(Statement)
}
