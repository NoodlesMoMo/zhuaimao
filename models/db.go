package models

import (
	"sync"

	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	_db_instance           *gorm.DB
	_db_once, _db_del_once = sync.Once{}, sync.Once{}
)

var (
	dbConfig = mysql.Config{
		Addr:                 "localhost:3306",
		User:                 "root",
		Passwd:               "123",
		DBName:               "ime_admin",
		Net:                  "tcp4",
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset": "utf8",
		},
	}
)

func GetDBInstance() *gorm.DB {
	if _db_instance != nil {
		return _db_instance
	}

	_db_once.Do(func() {
		var err error
		_db_instance, err = gorm.Open("mysql", dbConfig.FormatDSN())
		if err != nil {
			panic(err)
		}

		_db_instance.DB().SetMaxIdleConns(4)
		_db_instance.DB().SetMaxOpenConns(8)

		_db_instance.LogMode(true)
	})

	return _db_instance
}

func DestroyDBInstance() {
	if _db_instance != nil {
		_db_del_once.Do(func() {
			_db_instance.Close()
			_db_instance = nil
			fmt.Println(">>>> destroy database instance.")
		})
	}
}
