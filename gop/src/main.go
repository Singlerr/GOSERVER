package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id string
	ip string
}

func main() {

	db, err := sql.Open("mysql", "mysql51873:6U0VENe7nUntqpf@tcp(127.0.0.1:3306)/Minecraft")
	if err != nil {
		log.Fatal(err)
	}

http.HandleFunc("/authentication", func(w http.ResponseWriter, req *http.Request) {
		len := req.ContentLength
		body := make([]byte, len)
		req.Body.Read(body)
		var data Data
		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}

		var key string

		err = db.QueryRow("SELECT authorization FROM Minecraft WHERE id = '" + data.id + "'").Scan(&key)
		if err != nil {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
			fmt.Fprintf(w, key)
		}

	})

	defer db.Close()

	
	println("서버 오픈")
	http.ListenAndServe(":4500", nil)
}
