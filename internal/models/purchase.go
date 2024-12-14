package models

type Purchase struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	ProductId int    `json:"productId"`
	Cost      int    `json:"cost"`
	Quantity  int    `json:"quantity"`
	Timestamp string `json:"timestamp"`
}
