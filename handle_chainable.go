package lgorm

import (
	"reflect"
	"regexp"
)

// 实例化
func (db *Db) getChainAbleInstance() *Db {
	ptrType := reflect.TypeOf(db.Statement)
	numField := ptrType.NumField()
	for i := 0; i < numField; i++ {
		db.addChainAble(ptrType.Field(i).Name)
	}
	return db
}

// 添加chainAble_api的方法
func (db *Db) addChainAble(name string) {
	dataList := db.getInputDataList(name)
	if !db.isChainAblePool(dataList) {
		return
	}
	if db.isList(dataList) {
		newDataList := dataList.([]ChainAblePool)
		for _, data := range newDataList {
			if !data.IsCall {
				continue
			}
			var iValue []reflect.Value
			for _, v := range data.Params {
				iValue = append(iValue, reflect.ValueOf(v))
			}
			db.runStructFunc(name, iValue)
		}
	} else {
		newData := dataList.(ChainAblePool)
		if !newData.IsCall {
			return
		}
		var iValue []reflect.Value
		for _, v := range newData.Params {
			iValue = append(iValue, reflect.ValueOf(v))
		}
		db.runStructFunc(name, iValue)
	}
}

// 是否chainAble_api的方法
func (db *Db) isChainAblePool(value interface{}) bool {
	var regChainAblePool = regexp.MustCompile(`^(\[\])*(lgorm\.)*ChainAblePool$`)
	if reflect.ValueOf(value).IsValid() {
		lType := reflect.ValueOf(value).Type().String()
		return regChainAblePool.MatchString(lType)
	}
	return false
}
