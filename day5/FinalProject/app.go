package main

import (
	"FinalProject/Handlers"
	"FinalProject/functions"
	"log"
	"net/http"
	"time"
)

func startServer() {
	// Author-related routes
	http.HandleFunc("/api/authors", Handlers.HandleCreateAuthor)
	http.HandleFunc("/api/get/author/", Handlers.HandleGetAuthorByID)
	http.HandleFunc("/api/update/author/", Handlers.HandleUpdateAuthor)
	http.HandleFunc("/api/delete/author/", Handlers.HandleDeleteAuthor)
	http.HandleFunc("/api/authors/list", Handlers.HandleListAllAuthors)

	// Book-related routes
	http.HandleFunc("/api/books", Handlers.HandleCreateBook)
	http.HandleFunc("/api/get/book/", Handlers.HandleGetBookByID)
	http.HandleFunc("/api/update/book/", Handlers.HandleUpdateBook)
	http.HandleFunc("/api/delete/book/", Handlers.HandleDeleteBook)
	http.HandleFunc("/api/books/list", Handlers.HandleListAllBooks)

	// Customer-related routes
	http.HandleFunc("/api/customers", Handlers.HandleCreateCustomer)
	http.HandleFunc("/api/get/customer/", Handlers.HandleGetCustomerByID)
	http.HandleFunc("/api/update/customer/", Handlers.HandleUpdateCustomer)
	http.HandleFunc("/api/delete/customer/", Handlers.HandleDeleteCustomer)
	http.HandleFunc("/api/customers/list", Handlers.HandleListAllCustomers)

	// Order-related routes
	http.HandleFunc("/api/orders", Handlers.HandleCreateOrder)
	http.HandleFunc("/api/get/order/", Handlers.HandleGetOrderByID)
	http.HandleFunc("/api/update/order/", Handlers.HandleUpdateOrder)
	http.HandleFunc("/api/delete/order/", Handlers.HandleDeleteOrder)
	http.HandleFunc("/api/orders/list", Handlers.HandleListAllOrders)

	// Sales Report API
	http.HandleFunc("/api/reports", Handlers.HandleListSalesReports)

	go scheduleDailySalesReportGeneration()

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Periodic task for generating sales reports
func scheduleDailySalesReportGeneration() {
	ticker := time.NewTicker(24 * time.Hour)
	for {
		<-ticker.C
		go func() {
			if err := functions.GenerateDailySalesReport(); err != nil {
				log.Printf("Error generating sales report: %v", err)
			} else {
				log.Println("Daily sales report generated successfully")
			}
		}()
	}
}

func main() {
	go startServer()
	select {}
}
