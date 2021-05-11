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
			var iValue []reflect.Value
			var handleParams = []interface{}{}
			for k, v := range data.Params {
				if data.HandleType == "Set" {
					for _, v1 := range data.HandleParamsIndex {
						if k == v1 {
							handleParams = append(handleParams, v)
							if newV := db.handleAttr(v, data.HandleType); newV != nil {
								v = newV
							}
						}
					}
				} else if data.HandleType == "SetOne" {
					fieldNum := data.HandleParamsIndex[0]
					valNum := data.HandleParamsIndex[1]
					field := data.Params[fieldNum].(string)
					val := data.Params[valNum]
					tmp := map[string]interface{}{field: val}
					if newV := db.handleAttr(tmp, data.HandleType); newV != nil {
						v = newV.(map[string]interface{})[field]
					} else {
						v = tmp[field]
					}
					if fieldNum == k {
						v = field
					}
				}
				iValue = append(iValue, reflect.ValueOf(v))
			}
			db.runStructFunc(name, iValue)
			if data.HandleType == "Get" {
				for _, v := range handleParams {
					db.handleAttr(v, data.HandleType)
				}
			}
		}
	} else {
		newData := dataList.(FinisherPool)
		if !newData.IsCall {
			return
		}
		var iValue []reflect.Value
		var handleParams = []interface{}{}
		for k, v := range newData.Params {
			iValue = append(iValue, reflect.ValueOf(v))
			for _, v1 := range newData.HandleParamsIndex {
				if k == v1 {
					handleParams = append(handleParams, v)
				}
			}
		}
		if newData.HandleType == "Set" {
			for _, v := range handleParams {
				db.handleAttr(v, newData.HandleType)
			}
		}
		db.runStructFunc(name, iValue)
		if newData.HandleType == "Get" {
			for _, v := range handleParams {
				db.handleAttr(v, newData.HandleType)
			}
		}
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
