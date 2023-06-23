package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type StoreJSON struct {
	user     *UserRepo
	category *CategoryRepo
	product  *ProductRepo
	shopCart *ShopCartRepo
}

func NewConnectionJSON(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	categoryFile, err := os.Open(cfg.Path + cfg.CategoryFileName)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	shopCartFile, err := os.Open(cfg.Path + cfg.ShopCartFileName)
	if err != nil {
		return nil, err
	}

	return &StoreJSON{
		user:     NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		category: NewCategoryRepo(cfg.Path+cfg.CategoryFileName, categoryFile),
		product:  NewProductRepo(cfg.Path+cfg.ProductFileName, productFile),
		shopCart: NewShopCartRepo(cfg.Path+cfg.ShopCartFileName, shopCartFile),
	}, nil
}

func (u *StoreJSON) User() storage.UserRepoI {
	return u.user
}

func (u *StoreJSON) Category() storage.CategoryRepoI {
	return u.category
}

func (p *StoreJSON) Product() storage.ProductRepoI {
	return p.product
}

func (w *StoreJSON) ShopCart() storage.ShopCartRepoI {
	return w.shopCart
}
