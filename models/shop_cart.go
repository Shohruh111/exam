package models

type CreateShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
}
type ShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type UpdateShopCart struct {
	ProductId string `json:"product_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}
type ShopCartPrimaryKey struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}
type ShopCartGetListRequest struct {
	Offset    int
	Limit     int
	From_date string
	To_date   string
}

type ShopCartGetListResponse struct {
	Count  int
	Orders []ShopCart
}
type ShopCartHistory struct {
	TotalPrice  int
	ProductName string
	Price       int
	Count       int
	Time        string
}
type ShopCartSellDay struct {
	Name  string
	Count int
}
