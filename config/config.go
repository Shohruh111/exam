package config

type Config struct {
	DefaultOffset int
	DefaultLimit  int

	Path string

	UserFileName     string
	ProductFileName  string
	CategoryFileName string
	ShopCartFileName string
}

func Load() Config {
	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Path = "./data"

	cfg.CategoryFileName = "/category.json"
	cfg.ProductFileName = "/product.json"
	cfg.UserFileName = "/user.json"
	cfg.ShopCartFileName = "/shop_cart.json"

	return cfg
}
