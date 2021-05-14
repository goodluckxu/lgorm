# lgorm
将gorm封装了一层，可以实现getAttr和setAttr

gorm地址 https://github.com/go-gorm/gorm

## 用法
可以完全按照gorm的方法去实现功能

优化了链式调用时，需要重新赋值，如：db = db.Where("id = ?", id)，避免多次调用时的冲突

Get{structFieldName}Attr可以在获取数据时处理输出的数据

Set{structFieldName}Attr可以在添加和修改的时候处理添加的数据

两个方法传入的参数类型和返回类型都应该是结构体field的类型

## 未实现attr的方法
Transaction,FindInBatches 由于内部需要传*gorm.DB类型的值，内部的*gorm.DB调用的方法无法处理attr方法

## 用法实例
~~~go
package main

import (
	"github.com/goodluckxu/lgorm"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := lgorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/backend_api?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	var bank model.Bank
	db.Model(&bank).Where("id = 80").Update("name", "aaa")
	db.First(&bank)
	fmt.Println(bank)
}
~~~
~~~go
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
控制台输出：
~~~shell script
2021/05/12 09:03:08 C:/Users/luckyxu/sdk/go1.16.3/src/reflect/value.go:476
[10.841ms] [rows:1] UPDATE `bank` SET `name`='aaa_abc',`updated_at`='2021-05-12 09:03:08.748' WHERE id = 80

2021/05/12 09:03:08 C:/Users/luckyxu/sdk/go1.16.3/src/reflect/value.go:476
[0.503ms] [rows:1] SELECT * FROM `bank` ORDER BY `bank`.`id` LIMIT 1
{63 测试添加 测试 你好 hello 描述abc 0 0 2021-05-12 09:03:08.7611426 +0800 CST m=+0.021322101 2021-05-11 09:49:16 +0800 CST}
~~~
