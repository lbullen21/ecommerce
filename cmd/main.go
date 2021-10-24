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
	ID     int     `json:"id"`
	Name   string  `json:"flavor"`
	Photo  string  `json:"photo"`
	Price  float32 `json:"price"`
	Detail string  `json:"detail"`
}

var db *sql.DB

func databaseConnect() {

	//Create products table in MySQL
	create, err := db.Query(`CREATE TABLE IF NOT EXISTS products (
		id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
		flavor TEXT NOT NULL,
		photo TEXT NOT NULL,
		price FLOAT (10) NOT NULL,
		detail TEXT NOT NULL
		)`)
	if err != nil {
		panic(err.Error())
	}

	defer create.Close()

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

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	//testing route
	fmt.Fprintf(w, "Welcome to the Products Page!")
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

			err := rows.Scan(&currentProduct.ID, &currentProduct.Name, &currentProduct.Photo, &currentProduct.Price, &currentProduct.Detail)
			if err != nil {
				fmt.Println()
				return
			}

			products = append(products, currentProduct)
		}
		json.NewEncoder(w).Encode(products)

	}

}

func searchProducts() {

}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//gets all products in db
	myRouter.HandleFunc("/products", productsHandler)

	// Running the Server
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	//Connect to MySQL database
	database, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/fizzyFactory")
	fmt.Println("connected to db")
	if err != nil {
		fmt.Println("Something is not working")
	}
	db = database
	defer db.Close()

	databaseConnect()
	handleRequests()
}
