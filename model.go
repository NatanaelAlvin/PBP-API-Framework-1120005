package Model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"Address"`
	UserType int    `json:"type"`
}
type Users struct {
	Data []User `json:"data"`
}
type Address struct {
	ID     int    `json:"id"`
	Street string `json:"street"`
	UserID int    `json:"user_id"`
}
type UserAddresses struct {
	User       User      `json:"user"`
	Addressess []Address `json:"addresses"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Transaction struct {
	ID        int `json:"status"`
	UserID    int `json:"message"`
	ProductID int `json:"data"`
	Quantity  int `json:"quantity"`
}
type TransactionResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
