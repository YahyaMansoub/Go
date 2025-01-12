
# API Documentation for FinalProject

## How to Run the Application

### 1. Install Dependencies
Ensure that you have all the necessary dependencies installed by running:
```bash
go mod tidy
2. Run the Application
To start the server, run:

bash
Copier le code
go run main.go
The application will start a web server on port 8080.

3. Testing the API
You can use tools like Postman or curl to make requests to the endpoints.

Author Routes
Create Author
Endpoint: POST /api/authors
Description: Creates a new author.
Request Body:
json
Copier le code
{
  "first_name": "John",
  "last_name": "Doe",
  "bio": "Experienced author with expertise in technology."
}
Response:
json
Copier le code
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "bio": "Experienced author with expertise in technology."
}
Get Author by ID
Endpoint: GET /api/get/author/{id}
Description: Retrieves an author by their ID.
Example Request: /api/get/author/1
Response:
json
Copier le code
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "bio": "Experienced author with expertise in technology."
}
Update Author
Endpoint: PUT /api/update/author/{id}
Description: Updates an author's information.
Request Body:
json
Copier le code
{
  "first_name": "John",
  "last_name": "Doe Updated",
  "bio": "Updated bio with new information."
}
Response:
json
Copier le code
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe Updated",
  "bio": "Updated bio with new information."
}
Delete Author
Endpoint: DELETE /api/delete/author/{id}
Description: Deletes an author by their ID.
Example Request: /api/delete/author/1
Response:
json
Copier le code
{
  "message": "Author deleted successfully"
}
List All Authors
Endpoint: GET /api/authors/list
Description: Retrieves a list of all authors.
Response:
json
Copier le code
[
  {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Experienced author."
  },
  {
    "id": 2,
    "first_name": "Jane",
    "last_name": "Smith",
    "bio": "Another experienced author."
  }
]
Book Routes
Create Book
Endpoint: POST /api/books
Description: Creates a new book.
Request Body:
json
Copier le code
{
  "title": "Go Programming",
  "author_id": 1,
  "genres": ["Technology", "Programming"],
  "published_at": "2025-01-01T00:00:00Z",
  "price": 29.99,
  "stock": 100
}
Response:
json
Copier le code
{
  "id": 1,
  "title": "Go Programming",
  "author": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Experienced author."
  },
  "genres": ["Technology", "Programming"],
  "published_at": "2025-01-01T00:00:00Z",
  "price": 29.99,
  "stock": 100
}
Get Book by ID
Endpoint: GET /api/get/book/{id}
Example Request: /api/get/book/1
Response:
json
Copier le code
{
  "id": 1,
  "title": "Go Programming",
  "author": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Experienced author."
  },
  "genres": ["Technology", "Programming"],
  "published_at": "2025-01-01T00:00:00Z",
  "price": 29.99,
  "stock": 100
}
Update Book
Endpoint: PUT /api/update/book/{id}
Description: Updates a book's information.
Request Body:
json
Copier le code
{
  "title": "Advanced Go Programming",
  "author_id": 1,
  "genres": ["Technology", "Programming", "Advanced"],
  "published_at": "2025-01-10T00:00:00Z",
  "price": 39.99,
  "stock": 120
}
Response:
json
Copier le code
{
  "id": 1,
  "title": "Advanced Go Programming",
  "author": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Experienced author."
  },
  "genres": ["Technology", "Programming", "Advanced"],
  "published_at": "2025-01-10T00:00:00Z",
  "price": 39.99,
  "stock": 120
}
Delete Book
Endpoint: DELETE /api/delete/book/{id}
Example Request: /api/delete/book/1
Response:
json
Copier le code
{
  "message": "Book deleted successfully"
}
List All Books
Endpoint: GET /api/books/list
Response:
json
Copier le code
[
  {
    "id": 1,
    "title": "Go Programming",
    "author": {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "bio": "Experienced author."
    },
    "genres": ["Technology", "Programming"],
    "published_at": "2025-01-01T00:00:00Z",
    "price": 29.99,
    "stock": 100
  }
]
Customer Routes
Create Customer
Endpoint: POST /api/customers
Description: Creates a new customer.
Request Body:
json
Copier le code
{
  "name": "Alice",
  "email": "alice@example.com",
  "address": "123 Main St, Springfield, IL",
  "created_at": "2025-01-10T00:00:00Z"
}
Response:
json
Copier le code
{
  "id": 1,
  "name": "Alice",
  "email": "alice@example.com",
  "address": "123 Main St, Springfield, IL",
  "created_at": "2025-01-10T00:00:00Z"
}
Get Customer by ID
Endpoint: GET /api/get/customer/{id}
Example Request: /api/get/customer/1
Response:
json
Copier le code
{
  "id": 1,
  "name": "Alice",
  "email": "alice@example.com",
  "address": "123 Main St, Springfield, IL",
  "created_at": "2025-01-10T00:00:00Z"
}
Update Customer
Endpoint: PUT /api/update/customer/{id}
Description: Updates a customer's information.
Request Body:
json
Copier le code
{
  "name": "Alice Cooper",
  "email": "alice.cooper@example.com",
  "address": "123 Main St, Springfield, IL",
  "created_at": "2025-01-10T00:00:00Z"
}
Response:
json
Copier le code
{
  "id": 1,
  "name": "Alice Cooper",
  "email": "alice.cooper@example.com",
  "address": "123 Main St, Springfield, IL",
  "created_at": "2025-01-10T00:00:00Z"
}
Delete Customer
Endpoint: DELETE /api/delete/customer/{id}
Response:
json
Copier le code
{
  "message": "Customer deleted successfully"
}
Order Routes
Create Order
Endpoint: POST /api/orders
Request Body:
json
Copier le code
{
  "customer_id": 1,
  "items": [
    {
      "book_id": 1,
      "quantity": 2
    }
  ],
  "total_price": 59.98,
  "status": "pending",
  "created_at": "2025-01-12T00:00:00Z"
}
Response:
json
Copier le code
{
  "id": 1,
  "customer": {
    "id": 1,
    "name": "Alice Cooper",
    "email": "alice.cooper@example.com",
    "address": "123 Main St, Springfield, IL"
  },
  "items": [
    {
      "book": {
        "id": 1,
        "title": "Go Programming"
      },
      "quantity": 2
    }
  ],
  "total_price": 59.98,
  "status": "pending",
  "created_at": "2025-01-12T00:00:00Z"
}
Get Order by ID
Endpoint: GET /api/get/order/{id}
Response:
json
Copier le code
{
  "id": 1,
  "customer": {
    "id": 1,
    "name": "Alice Cooper",
    "email": "alice.cooper@example.com",
    "address": "123 Main St, Springfield, IL"
  },
  "items": [
    {
      "book": {
        "id": 1,
        "title": "Go Programming"
      },
      "quantity": 2
    }
  ],
  "total_price": 59.98,
  "status": "pending",
  "created_at": "2025-01-12T00:00:00Z"
}
Sales Report Routes
Get Sales Reports
Endpoint: GET /api/reports
Response:
json
Copier le code
[
  {
    "date": "2025-01-10T00:00:00Z",
    "total_revenue": 2000.50,
    "total_orders": 150,
    "total_books_sold": 250,
    "top_selling_books": ["Go Programming", "Advanced Go Programming"]
  }
]
