package model

import (
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      Author    `json:"author"`
	Genres      []string  `json:"genres"`
	PublishedAt time.Time `json:"published_at"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
}

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreayedAt time.Time `json:"created_at"`
}

type Order struct {
	ID         int         `json:"id"`
	Customer   Customer    `json:"customer"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	Status     string      `json:"status"`
}

type OrderItem struct {
	Book     Book `json:"book"`
	Quantity int  `json:"quantity"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type SalesReport struct {
	Timestamp       time.Time   `json:"timestamp"`
	TotalRevenue    float64     `json:"total_revenue"`
	TotalOrders     int         `json:"total_orders"`
	TopSellingBooks []BookSales `json:"top_selling_books"`
}

type BookSales struct {
	Book     Book `json:"book"`
	Quantity int  `json:"quantity_sold"`
}

type BookStore interface {
	CreateBook(book Book) (Book, error)
	GetBook(id int) (Book, error)
	UpdateBook(id int, book Book) (Book, error)
	DeleteBook(id int) error
}

type AuthorStore interface {
	CreateAuthor(author Author) (Author, error)
	GetAuthor(id int) (Author, error)
	UpdateAuthor(id int, book Book) (Book, error)
	DeleteAuthor(id int) error
}

type OrderStore interface {
	CreateOrder(order Order) (Order, error)
	GetOrder(id int, order Order) (Order, error)
	UpdateOrder(id int, order Order) (Order, error)
	DeleteOrder(id int) error
}
