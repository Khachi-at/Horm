package main

import (
	horm "Horm"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := horm.NewEngine("sqlite3", "h.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXIST user;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
