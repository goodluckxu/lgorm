package lgorm

import (
	"reflect"
	"regexp"
)

func (db *Db) RunFinisher() {
	db.getChainAbleInstance()
	db.getFinisherInstance()
	db.initStatement()
}

// 实例化
func (db *Db) getFinisherInstance() *Db {
	value := reflect.ValueOf(&db.Statement).Elem()
	numField := value.NumField()
	for i := 0; i < numField; i++ {
		db.addFinisher(value.Type().Field(i).Name)
	}
	return db
}

// 添加chainAble_api的方法
func (db *Db) addFinisher(name string) {
	dataList := db.getInputDataList(name)
	if !db.isFinisherPool(dataList) {
		return
	}
	value := reflect.ValueOf(db.Statement)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if db.isList(dataList) {
		newDataList := dataList.([]FinisherPool)
		for _, data := range newDataList {
			if !data.IsCall {
				continue
			}
			iValue, handleParams := db.handleSetAttr(data)
			db.runStructFunc(name, iValue)
			db.handleGetAttr(data, handleParams)
		}
	} else {
		newData := dataList.(FinisherPool)
		if !newData.IsCall {
			return
		}
		iValue, handleParams := db.handleSetAttr(newData)
		db.runStructFunc(name, iValue)
		db.handleGetAttr(newData, handleParams)
	}
}

func (db *Db) handleSetAttr(data FinisherPool) (iValue []reflect.Value, handleParams []interface{}) {
	handleType := data.HandleType
	if data.HandleType == "GetOrSet" {
		handleType = "Set"
	}
	switch handleType {
	case "Set":
		for k, v := range data.Params {
			for _, v1 := range data.HandleParamsIndex {
				if k != v1 {
					continue
				}
				if newV := db.handleAttr(v, handleType); newV != nil {
					v = newV
				}
			}
		}
	case "SetOne":
		fieldNum := data.HandleParamsIndex[0]
		valNum := data.HandleParamsIndex[1]
		field := data.Params[fieldNum].(string)
		val := data.Params[valNum]
		tmp := map[string]interface{}{field: val}
		if newV := db.handleAttr(tmp, handleType); newV != nil {
			val = newV.(map[string]interface{})[field]
		} else {
			val = tmp[field]
		}
		data.Params[valNum] = val
	}
	for _, v := range data.Params {
		iValue = append(iValue, reflect.ValueOf(v))
		handleParams = append(handleParams, v)
	}
	return
}

func (db *Db) handleGetAttr(data FinisherPool, handleParams []interface{}) {
	handleType := data.HandleType
	if data.HandleType == "GetOrSet" {
		handleType = "Get"
	}
	switch handleType {
	case "Get":
		for _, v := range handleParams {
			db.handleAttr(v, handleType)
		}
	case "GetOne":
		fieldNum := data.HandleParamsIndex[0]
		valNum := data.HandleParamsIndex[1]
		field := data.Params[fieldNum].(string)
		val := data.Params[valNum]
		tmp := map[string]interface{}{field: val}
		if newV := db.handleAttr(tmp, handleType); newV != nil {
			val = newV.(map[string]interface{})[field]
		} else {
			val = tmp[field]
		}
		data.Params[valNum] = val
	}
}

// 是否chainAble_api的方法
func (db *Db) isFinisherPool(value interface{}) bool {
	var regChainAblePool = regexp.MustCompile(`^(\[\])*(lgorm\.)*FinisherPool$`)
	if reflect.ValueOf(value).IsValid() {
		lType := reflect.ValueOf(value).Type().String()
		return regChainAblePool.MatchString(lType)
	}
	return false
}
