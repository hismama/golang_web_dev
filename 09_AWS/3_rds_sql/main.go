package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:C1Tad3l!@tcp(mydbinstance-1.c9qegi82g6wh.us-east-1.rds.amazonaws.com:3306)/go_web?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}
func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully completed.")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT fname FROM amigos")
	check(err)

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("CREATE TABLE customer (name VARCHAR(20));")
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "CREATED TABLE customer:", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("INSERT INTO customer (name) VALUES ('James');")
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM customer;")
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Println(name)
		fmt.Fprintln(w, "RETRIEVED RECORD: ", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("UPDATE customer SET name = 'Jimmy' WHERE name='James';")
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("DELETE FROM customer WHERE name='Jimmy';")
	check(err)
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare("DROP TABLE customer;")
	check(err)
	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROP TABLE customer:")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
