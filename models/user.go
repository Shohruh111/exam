package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Balance   int    `json:"balance"`
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Balance   int    `json:"balance"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Balance   int    `json:"balance"`
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}

type UserHistory struct {
	ProductName string `json:"name"`
	Count       int    `json:"count"`
	TotalCount  int    `json:"total_count"`
}
