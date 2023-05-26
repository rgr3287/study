package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	First_name string `gorm:"column:first_name"`
	Last_name  string `gorm:"column:last_name"`
}

func main() {
	db, err := gorm.Open(mysql.Open("아이디:비번@tcp(localhost:3306)/db이름?charset=utf8mb4&parseTime=True&loc=Local")) // db open
	if err != nil {
		panic(err)
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("Rollback")
		}
	}()

	tables, err := tx.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	beforeWord := "SPURLINER"
	afterWord := "SNOWLINER"
	replaceTotalCount := int64(0)
	for _, table := range tables {
		columns, err := tx.Migrator().ColumnTypes(table)
		if err != nil {
			panic(err)
		}

		for _, column := range columns {
			columnType, _ := column.ColumnType()
			columnName := column.Name()
			if strings.Contains(columnType, "varchar") || strings.Contains(columnType, "text") {
				result := tx.Table(table).Where(columnName+" LIKE ?", "%"+beforeWord+"%").Update(columnName, gorm.Expr("REPLACE("+columnName+", ?, ?)", beforeWord, afterWord))
				if result.Error != nil {
					panic(result.Error)
				}
				replaceTotalCount += result.RowsAffected
			}
		}
	}
	tx.Commit()
	fmt.Println(replaceTotalCount)
}
