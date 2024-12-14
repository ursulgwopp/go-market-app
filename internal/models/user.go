package models

type User struct {
	Id          int     `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Balance     int     `json:"balance,omitempty"`
	ProductList []int64 `json:"productList" db:"product_list"`
}
