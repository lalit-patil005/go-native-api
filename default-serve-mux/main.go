package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Id          int     `json:"id"`
	Sku         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
type ErrorResponse struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

type ProductRes struct {
	RawData json.RawMessage `json:"-"` // Raw JSON data
}

var products = []*Product{
	{
		Id:          1,
		Sku:         "sku1",
		Name:        "Test Product Name 1",
		Description: " Test Product Description 1",
		Price:       100.77,
		Stock:       10,
	},
	{
		Id:          2,
		Sku:         "sku2",
		Name:        "Test Product Name 2",
		Description: " Test Product Description 2",
		Price:       100.77,
		Stock:       20,
	},
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("allProducts API hit DefaultServeMux.")
	json.NewEncoder(w).Encode(products)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page Endpoint DefaultServeMux.")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/products/fetch/")
	productID, _ := strconv.Atoi(id)
	var filteredProduct Product
	for _, product := range products {
		if product.Id == productID {
			filteredProduct = *product
			break
		}
	}
	if filteredProduct == (Product{}) {
		errorRes := ErrorResponse{
			Message: "Product not found",
			Error:   true,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	json.NewEncoder(w).Encode(filteredProduct)
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		errorRes := ErrorResponse{
			Message: "Method not allowed",
			Error:   true,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/products/update/")
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Error parsing request body test", http.StatusBadRequest)
		return
	}
	var needToupdateProduct Product
	productID, _ := strconv.Atoi(id)
	for _, product := range products {
		if product.Id == productID {
			needToupdateProduct = *product
			break
		}
	}
	if needToupdateProduct == (Product{}) {
		errorRes := ErrorResponse{
			Message: "Product not found",
			Error:   true,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	if price, ok := requestBody["price"].(float64); ok {
		needToupdateProduct.Price = price
	}
	if name, ok := requestBody["name"].(string); ok {
		needToupdateProduct.Name = name
	}
	if description, ok := requestBody["description"].(string); ok {
		needToupdateProduct.Description = description
	}
	if stock, ok := requestBody["stock"].(float64); ok {
		needToupdateProduct.Stock = int(stock)
	}
	json.NewEncoder(w).Encode(needToupdateProduct)
	// productString := fmt.Sprintf("%+v", needToupdateProduct)
	// fmt.Fprintf(w, "Update product of given ID. %s", productString)
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a product")
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a product")
}

func handleRequests() {
	http.HandleFunc("/products/fetch/", getProduct)
	http.HandleFunc("/products/update/", updateProduct)
	http.HandleFunc("/products/delete/", deleteProduct)
	http.HandleFunc("/products/create/", createProduct)
	http.HandleFunc("/products", allProducts)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func main() {
	fmt.Println("go apis with DefaultServeMux")
	handleRequests()
}
