package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

//Person ...
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//GetAll ...
func (p Person) GetAll() (persons []Person, err error) {
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var t Person
		rows.Scan(&t.ID, &t.Name)
		persons = append(persons, t)
	}

	return
}
func main() {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/study?parseTime=true")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := gin.Default()
	r.GET("/person", func(c *gin.Context) {
		//lo := binding.Login{}
		p := Person{}
		persons, err := p.GetAll()
		fmt.Println(persons)
		if err != nil {
			log.Fatal(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"records": persons,
			"count":   len(persons),
		})
	})

	http.ListenAndServe(":8800", r)
}
