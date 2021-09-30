package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists devbook")
	exec(db, "use devbook")
	exec(db, "drop table if exists users")
	exec(db, `create table users (
		id integer auto_increment primary key,
		name varchar(50) not null,
    	nick varchar(50) not null unique,
    	email varchar(20) not null unique,
		password varchar(20) not null unique,
    	createdAt timestamp default current_timestamp()
	)`)

}