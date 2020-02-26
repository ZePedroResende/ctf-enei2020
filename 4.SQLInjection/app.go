package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type Env struct {
	db *sql.DB
}

func (env *Env) timeQuery(query string) []string {
	error := []string{""}
	if env.db == nil {
		log.Println("killed the db")
		return error
	}

	rows, err := env.db.Query("EXPLAIN ANALYZE " + query)

	if err != nil {
		log.Println(err)
		return error
	}

	defer rows.Close()
	response := []string{}
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
		}

		response = append(response, s)
		fmt.Println(s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return error
	}

	return response

}

func (env *Env) getHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func (env *Env) postHandler(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	query := r.FormValue("query")
	response := env.timeQuery(query)
	fmt.Fprintf(w, strings.Join(response, "\n"))
}

func (env *Env) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		env.getHandler(w, r)
	case "POST":
		env.postHandler(w, r)
	}
}

func main() {

	db, err := sql.Open("postgres", "host= 0.0.0.0 user=docker password=docker dbname=enei2020 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	deleteTable := `
	DROP TABLE enei_users;
	`

	_, err = db.Exec(deleteTable)
	if err != nil {
		log.Println("no problem")
	}

	createTable := `
CREATE TABLE enei_users(
 user_id SERIAL PRIMARY KEY,
 name VARCHAR (100) NOT NULL,
 secret VARCHAR (100) NOT NULL
);

INSERT INTO enei_users(name, secret)
VALUES('enei','afs12d');
`
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}
	insertUser := `
INSERT INTO enei_users(name, secret)
VALUES('enei','adsfa9q9i0vadsojoj2h3501h230hfhashoa0casdfieqw003r12142dsa09xnvzzo');
	`
	_, err = db.Exec(insertUser)
	if err != nil {
		panic(err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	env := &Env{db: db}

	http.HandleFunc("/adsfa9q9i0vadsojoj2h3501h230hfhashoa0casdfieqw003r12142dsa09xnvzzo", func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "./adsfa9q9i0vadsojoj2h3501h230hfhashoa0casdfieqw003r12142dsa09xnvzzo")
					})

	http.HandleFunc("/", env.index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}
