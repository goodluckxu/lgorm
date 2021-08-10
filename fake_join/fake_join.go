package fake_join

import (
	"encoding/json"
	"github.com/goodluckxu/go-lib/handle_interface"
	"github.com/goodluckxu/lgorm"
	"reflect"
	"strings"
)

// 多表联查
//   mainData 数据集合
//   otherDb 联查数据，例如：db.Model(&User{}).Where("is_del = 0")
//   where 联查条件
//     1.key:value格式, key为mainData里面的字段，value需要使用key字段去查找otherDb里面的字段value
//     2.key里面的.表示层级，*表示数组，如：a.*.b.c表示，a的map下面有数组，数组里的map有b，b的map有c
//   alias 别名
//   joinType 拼接类型，one表示查询单个，more表示查询多个
func MoreJoinTable(mainData interface{}, otherDb *lgorm.Db, where map[string]string, alias string, joinType string) interface{} {
	mainByte, _ := json.Marshal(mainData)
	var newMainData interface{}
	_ = json.Unmarshal(mainByte, &newMainData)
	whereList := []string{}
	whereIn := map[string]interface{}{
		"key":  "",
		"list": []interface{}{},
	}
	whereLen := 0
	aliasMKey := ""
	index := -10
	for mKey, _ := range where {
		starIndex := -1
		ruleKeyList := strings.Split(mKey, ".")
		for key, rule := range ruleKeyList {
			if rule == "*" {
				starIndex = key
			}
		}
		if starIndex+1 == len(ruleKeyList) {
			starIndex--
		}
		if starIndex == -1 {
			starIndex++
		}
		if index == -10 {
			index = starIndex
			aliasMKey = strings.Join(ruleKeyList[0:index+1], ".") + "." + alias + alias
		} else {
			if index < starIndex {
				index = starIndex
				aliasMKey = strings.Join(ruleKeyList[0:index+1], ".") + "." + alias + alias
			}
		}
		whereLen++
	}
	if whereLen == 0 {
		return newMainData
	}
	reMain := reorganizingMainData(newMainData, aliasMKey, where)
	for _, oKey := range where {
		keyList := getMKeyDataList(aliasMKey+"."+oKey, reMain)
		whereIn["key"] = oKey + " in (?)"
		whereIn["list"] = removeDuplicateElement(keyList)
		for k, v := range keyList {
			vStr := FormatNumber(v)
			if len(whereList) > k {
				whereList[k] += " and " + oKey + " = " + vStr
			} else {
				whereList = append(whereList, "("+oKey+" = "+vStr)
			}
		}
	}
	var list []map[string]interface{}
	if whereLen == 1 {
		otherDb.Where(whereIn["key"], whereIn["list"]).Find(&list)
	} else {
		whereStr := strings.Join(whereList, ") or ") + ")"
		otherDb.Where(whereStr).Find(&list)
	}
	newMainData = packageMainData(reMain, list, aliasMKey, alias, joinType)
	newMainData = handle_interface.UpdateInterface(newMainData, []handle_interface.Rule{
		{FindField: aliasMKey, Type: "_"},
	})
	return newMainData
}

func reorganizingMainData(
	mainData interface{},
	aliasMKey string,
	where map[string]string,
) interface{} {
	rules := []handle_interface.Rule{
		{FindField: aliasMKey, UpdateValue: map[string]interface{}{}},
	}
	for mKey, oKey := range where {
		rules = append(rules, handle_interface.Rule{
			FindField:   aliasMKey + "." + oKey,
			UpdateValue: mKey,
			Type:        "*",
		})
	}
	mainData = handle_interface.UpdateInterface(mainData, rules)
	return mainData
}

func packageMainData(
	mainData interface{},
	list []map[string]interface{},
	aliasMKey string,
	alias string,
	joinType string,
) interface{} {
	aliasMKeyList := strings.Split(aliasMKey, ".")
	newAliasMKey := aliasMKeyList[0]
	otherAliasMKey := strings.Join(aliasMKeyList[1:], ".")
	if len(aliasMKeyList) > 1 {
		if newAliasMKey == "*" {
			newMainDataList := []interface{}{}
			newMainData, _ := mainData.([]interface{})
			for _, v := range newMainData {
				newMainDataList = append(newMainDataList, packageMainData(v, list, otherAliasMKey, alias, joinType))
			}
			mainData = newMainDataList
		} else {
			newMainDataMap, _ := mainData.(map[string]interface{})
			newMainDataMap[newAliasMKey] = packageMainData(newMainDataMap[newAliasMKey], list, otherAliasMKey, alias, joinType)
			mainData = newMainDataMap
		}
	} else {
		newMainDataMap, _ := mainData.(map[string]interface{})
		newAliasData, _ := newMainDataMap[newAliasMKey].(map[string]interface{})
		newList := []interface{}{}
		for _, v := range list {
			isEq := true
			for key, val := range newAliasData {
				if v[key] == nil {
					isEq = false
					break
				}
				if !jsonTypeEqual(v[key], val) {
					isEq = false
					break
				}
			}
			if isEq {
				newList = append(newList, v)
			}
		}
		if joinType == "one" {
			if len(newList) > 0 {
				newMainDataMap[alias] = newList[0]
			} else {
				newMainDataMap[alias] = map[string]interface{}{}
			}
		} else if joinType == "more" {
			newMainDataMap[alias] = newList
		}
		mainData = newMainDataMap
	}
	return mainData
}

func getMKeyDataList(mKey string, mainData interface{}) []interface{} {
	mKeyList := strings.Split(mKey, ".")
	newMKey := mKeyList[0]
	otherMKey := strings.Join(mKeyList[1:], ".")
	if len(mKeyList) > 1 {
		if newMKey == "*" {
			newMainDataList := []interface{}{}
			newMainData, _ := mainData.([]interface{})
			for _, v := range newMainData {
				newMainDataList = append(newMainDataList, getMKeyDataList(otherMKey, v)...)
			}
			return newMainDataList
		} else {
			newMainDataMap, _ := mainData.(map[string]interface{})
			return getMKeyDataList(otherMKey, newMainDataMap[newMKey])
		}
	} else {
		newMainDataMap, _ := mainData.(map[string]interface{})
		return []interface{}{newMainDataMap[newMKey]}
	}
}

func removeDuplicateElement(addrs []interface{}) []interface{} {
	result := make([]interface{}, 0, len(addrs))
	temp := map[interface{}]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func jsonTypeEqual(one interface{}, two interface{}) bool {
	if one == nil || two == nil {
		return one == two
	}
	oneValue := reflect.ValueOf(one)
	if oneValue.Kind() == reflect.Ptr {
		oneValue = oneValue.Elem()
	}
	twoValue := reflect.ValueOf(two)
	if twoValue.Kind() == reflect.Ptr {
		twoValue = twoValue.Elem()
	}
	if oneValue.Type() == twoValue.Type() {
		return one == two
	}
	if oneValue.Kind() == reflect.Int {
		one = float64(one.(int))
	}
	if twoValue.Kind() == reflect.Int {
		two = float64(two.(int))
	}
	return one == two
}
