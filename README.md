# lgorm
将gorm封装了一层，可以实现getAttr和setAttr

## 用法
可以完全按照gorm的方法去实现功能

优化了链式调用时，需要重新赋值，如：db = db.Where("id = ?", id)，避免多次调用时的冲突

Get{structFieldName}Attr可以在获取数据时处理输出的数据

Set{structFieldName}Attr可以在添加和修改的时候处理添加的数据

两个方法传入的参数类型和返回类型都应该是结构体field的类型

## 未实现attr的方法
Transaction,FindInBatches 由于内部需要传*gorm.DB类型的值，内部的*gorm.DB调用的方法无法处理attr方法

## 用法实例
~~~
package model

import (
	"time"
)

type Bank struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	ShortName    string    `json:"short_name"`
	EnglishName  string    `json:"english_name"`
	EnglishAbbr  string    `json:"english_abbr"`
	Remark       string    `json:"remark"`
	OptionStatus int       `json:"option_status"`
	Sort         int       `json:"sort"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (b Bank) GetRemarkAttr(value string) string {
	return value + "abc"
}

func (b Bank) GetCreatedAtAttr(value time.Time) time.Time {
	return time.Now()
}

func (b Bank) SetNameAttr(value string) string {
	return value + "_abc"
}
~~~
