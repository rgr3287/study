package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
	"time"
)

const (
	DB_USER     = "grooo-spice"
	DB_PASSWORD = "grooo~!"
	DB_NAME     = "spice"
)

func mai2n() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	var wg sync.WaitGroup
	var count = uint(5000)
	wg.Add(5000)
	for ; count <= 5000; count-- {
		go CreateAge(tx, &wg, count)
		if count == 1 {
			break
		}
	}
	wg.Wait()

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	ShowAge(db)
}

func CreateAge(tx *sql.Tx, wg *sync.WaitGroup, age uint) {
	fmt.Println("CreateAge", age)
	defer wg.Done()
	time.Sleep(time.Second * 5)
	_, err := tx.Exec("INSERT INTO test (age) VALUES($1)", age)
	if err != nil {
		panic(err)
	}
}

func ShowAge(db *sql.DB) {
	rows, err := db.Query("SELECT id,age FROM test")
	if err != nil {
		panic(err)
	}
	var id uint
	var age string
	for rows.Next() {
		err := rows.Scan(&id, &age)
		if err != nil {
			panic(err)
		}
		fmt.Println("id : ", id)
		fmt.Println("age : ", age)
		fmt.Println("------------")
	}
}
