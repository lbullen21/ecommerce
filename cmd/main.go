package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/* Product struct */

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"flavor"`
	Photo    string  `json:"photo"`
	Price    float32 `json:"price"`
	Detail   string  `json:"detail"`
	Category string  `json:"category"`
}

var db *sql.DB

/* Allows access to the password in env file */
// func envVariables() {
// 	err := godotenv.Load("auth.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

/* Was used to create db */

// func databaseConnect() {

// 	//Create products table in MySQL
// 	create, err := db.Query(`CREATE TABLE IF NOT EXISTS products (
// 		id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
// 		flavor TEXT NOT NULL,
// 		photo TEXT NOT NULL,
// 		price FLOAT (10) NOT NULL,
// 		detail TEXT NOT NULL
// 		)`)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer create.Close()

//Add products to db (only run once)
// insert, err := db.Query(`INSERT INTO products (flavor, photo, price, detail)
// VALUES
//     ('Limoncello', './images/limoncello.png', '5.99', 'pack of 6'),
//     ('Grapefruit', './images/sdGrapefruit.png', '5.99', 'pack of 6'),
//     ('Key Lime', './images/keyLime.png', '5.99', 'pack of 6'),
//     ('Lemon', './images/sdLemon.png', '5.99', 'pack of 6'),
//     ('Lime', './images/sdLime.png', '5.99', 'pack of 6'),
//     ('Tangerine', './images/tangerine.png', '5.99', 'pack of 6'),
//     ('Hibiscus', './images/hibiscus.png', '5.99', 'pack of 6'),
//     ('Variety Pack', './images/assorted.png', '12.99', '3 packs of 6')`)
// if err != nil {
// 	panic(err.Error())
// }

// defer insert.Close()

// }

func productsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// fmt.Fprintf(w, "Welcome to the Products Page!")
	fmt.Println("Endpoint Hit: Products Page")

	//Get method for all products
	if r.Method == http.MethodGet {
		var products []Product
		query := `SELECT * FROM products;`

		rows, err := db.Query(query)
		if err != nil {
			fmt.Println()
			return
		}

		for rows.Next() {
			var currentProduct Product

			err := rows.Scan(&currentProduct.ID, &currentProduct.Name, &currentProduct.Photo, &currentProduct.Price, &currentProduct.Detail, &currentProduct.Category)
			if err != nil {
				fmt.Println()
				return
			}

			products = append(products, currentProduct)
		}

		json.NewEncoder(w).Encode(products)

	}

}

func searchProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// fmt.Fprintf(w, "You have hit the search page!")

	if r.Method == http.MethodGet {
		var products []Product
		var query string

		searchTerm := r.URL.Query().Get("term")
		term := "%" + searchTerm + "%"
		order := r.URL.Query().Get("order")

		//this switch order will allow the filter to work
		switch order {
		case "lowToHigh":
			query = `SELECT * FROM products WHERE flavor LIKE ? ORDER BY price ASC`
		case "highToLow":
			query = `SELECT * FROM products WHERE flavor LIKE ? ORDER BY price DESC`
		default:
			query = `SELECT * FROM products WHERE flavor LIKE ?`
		}

		rows, err := db.Query(query, term)
		if err != nil {
			fmt.Println(err)
			return
		}
		for rows.Next() {
			var currentProduct Product

			err := rows.Scan(&currentProduct.ID, &currentProduct.Name, &currentProduct.Photo, &currentProduct.Price, &currentProduct.Detail, &currentProduct.Category)
			if err != nil {
				fmt.Println(err)
			}

			products = append(products, currentProduct)
		}
		json.NewEncoder(w).Encode(products)
	}

}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//gets all products in db
	myRouter.HandleFunc("/products", productsHandler)
	myRouter.HandleFunc("/search", searchProducts)

	// Running the Server
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {

	// envVariables()
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")

	//Connect to MySQL database

	//change to fizzy_db before 3306 127.0.0.1
	database, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/fizzyFactory")
	fmt.Println("connected to db")
	if err != nil {
		fmt.Println("Something is not working")
	}
	db = database
	defer db.Close()

	// databaseConnect()
	handleRequests()
}
