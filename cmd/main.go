package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func database() {

	//Connect to MySQL database
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/fizzy_factory")
	fmt.Println("connected to db")
	if err != nil {
		fmt.Println("Something is not working")
	}
	defer db.Close()

	//Create products table in MySQL
	create, err := db.Query(`CREATE TABLE IF NOT EXISTS products (
		id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
		flavor TEXT NOT NULL,
		photo TEXT NOT NULL,
		price FLOAT (10) NOT NULL
		)`)
	if err != nil {
		panic(err.Error())
	}

	defer create.Close()

	//Add products to db
	insert, err := db.Query(`INSERT INTO products (flavor, photo, price)
	VALUES
	    ('Limoncello', './images/limoncello.png', '5.99'),
	    ('Grapefruit', './images/sdGrapefruit.png', '5.99'),
	    ('Key Lime', './images/keyLime.png', '5.99'),
	    ('Lemon', './images/sdLemon.png', '5.99'),
	    ('Lime', './images/sdLime.png', '5.99'),
	    ('Tangerine', './images/tangerine.png', '5.99'),
	    ('Hibiscus', './images/hibiscus.png', '5.99'),
	    ('Variety Pack', './images/assorted.png', '12.99')`)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	//testing route
	fmt.Fprintf(w, "Welcome to the Products Page!")
	fmt.Println("Endpoint Hit: Products Page")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//gets all products in db
	myRouter.HandleFunc("/products", getAllProducts)

	// Running the Server
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	database()
	handleRequests()
}
