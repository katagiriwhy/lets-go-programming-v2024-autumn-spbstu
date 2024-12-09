package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {

	//pathOfCfg := config.ReadFlag()

	//cfgDB := storage.Load("/home/danil/lets-go-programming-v2024-autumn-spbstu/danil.novokhatskiy/task-9/internal/storage/config.yaml")

	//db, err := storage.NewStorage(cfgDB)
	db, err := sql.Open("postgres", "go_db://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//_, err = db.Exec("GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS contacts(id SERIAL PRIMARY KEY, name TEXT, phone TEXT)")

	/*rows, err := db.Query("SELECT * FROM contacts")

	for rows.Next() {
		var contact Contact
		rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		fmt.Println(contact)
	}*/

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query("INSERT INTO contacts (name, phone) VALUES ($1, $2)", "danil", "some numbers")
	if err != nil {
		log.Fatal(err)
	}
}

/*	router := mux.NewRouter()
	router.HandleFunc("/contacts", getContacts(db)).Methods("GET")
	router.HandleFunc("contacts/{id}", getContact(db)).Methods("GET")
	router.HandleFunc("/contacts", createContact(db)).Methods("POST")
	router.HandleFunc("/contacts/{id}", updateContact(db)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", deleteContact(db)).Methods("DELETE")*/

//log.Fatal(http.ListenAndServe("localhost:8080", jsonContentTypeMiddware(router)))

/*cfg := config.LoadConfig(pathOfCfg)

	fmt.Println(cfg)

}

func jsonContentTypeMiddware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT name FROM Contacts")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		contacts := []Contact{}
		for rows.Next() {
			var contact Contact
			if err := rows.Scan(&contact.Name, &contact.Phone); err != nil {
				log.Fatal(err)
			}
			contacts = append(contacts, contact)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(contacts)
	}
}

func getContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		var contact Contact

		err := db.QueryRow("SELECT name FROM Contacts WHERE id = $1", id).Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(contact)
	}
}

func createContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contact Contact
		json.NewDecoder(r.Body).Decode(&contact)

		err := db.QueryRow("INSERT INTO contacts (name, phone) VALUES ($1,$2) RETURNING id", contact.Name, contact.Phone).Scan(&contact.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(contact)
	}
}

func updateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contact Contact
		json.NewDecoder(r.Body).Decode(&contact)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3", contact.Name, contact.Phone, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(contact)
	}
}

func deleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		_, err := db.Exec("DELETE FROM contacts WHERE id = $1", id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode("Contact deleted")
	}
}*/
