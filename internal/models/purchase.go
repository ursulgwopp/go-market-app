package models

type Purchase struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	ProductId int    `json:"productId"`
	Quantity  int    `json:"quantity"`
	Timestamp string `json:"timestamp"`
}
