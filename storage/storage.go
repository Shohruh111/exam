package storage

import (
	"app/models"
)

type StorageI interface {
	User() UserRepoI
	Product() ProductRepoI
	Category() CategoryRepoI
	ShopCart() ShopCartRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (*models.User, error)
	GetById(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(*models.UpdateUser) (*models.User, error)
	Delete(*models.UserPrimaryKey) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (*models.Category, error)
	GetById(*models.CategoryPrimaryKey) (*models.Category, error)
	GetList(*models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(*models.UpdateCategory) (*models.Category, error)
	Delete(*models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (*models.Product, error)
	GetById(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(*models.UpdateProduct) (*models.Product, error)
	Delete(*models.ProductPrimaryKey) error
}
type ShopCartRepoI interface {
	Create(*models.CreateShopCart) (*models.ShopCart, error)
	GetById(*models.ShopCartPrimaryKey) (*models.ShopCart, error)
	GetList(*models.ShopCartGetListRequest, string) (*models.ShopCartGetListResponse, error)
	Update(*models.UpdateShopCart, string) (*models.ShopCart, error)
	Delete(*models.ShopCartPrimaryKey) error
	DateFilter(*models.ShopCartGetListRequest, string) ([]*models.ShopCart, error)
}
