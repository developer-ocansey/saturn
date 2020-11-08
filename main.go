package main
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main(){
	conn := "dbname=<db connection> sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err) //exit program
	}
	err = db.Ping()

	if err != nil {
		panic(err) //exit program
	}

	NewStore(&dbStore{db: db})
}

// Improve into library system (concurrency and authentication) Add logging and metrics