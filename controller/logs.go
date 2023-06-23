package controller

import (
	"app/config"
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"time"
)

var (
	History                 = make(map[string][]*models.ShopCartHistory)
	TotalCountOfSellProduct = make(map[string]int)
	MostSellInDay           = make(map[time.Time]models.ShopCartSellDay)
)

func (c *Controller) ShopCartHistory() (*models.ShopCartGetListResponse, error) {

	var (
		userId = "27457ac2-74dd-4656-b9b0-0d46b1af10dc"
	)

	cfg := config.Load()

	orders, err := read(cfg.ShopCartFileName)
	if err != nil {
		return nil, err
	}
	for _, val := range orders[userId] {
		userName, err := c.Strg.User().GetById(&models.UserPrimaryKey{userId})
		if err != nil {
			return nil, err
		}
		// fmt.Println(userName.FirstName)
		prodName, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{val.ProductId})
		if err != nil {
			return nil, err
		}
		if val.Status {
			order := models.ShopCartHistory{
				ProductName: prodName.Name,
				Price:       prodName.Price,
				Count:       val.Count,
				Time:        val.Time,
				TotalPrice:  val.Count * prodName.Price,
			}
			History[userName.FirstName+" "+userName.LastName] = append(History[userName.FirstName+" "+userName.LastName], &order)
		}
	}

	for key, val := range History {
		fmt.Println("Name: ", key)
		c.findProductsWithinUser(val)
	}
	return nil, nil

}
func (c *Controller) Sort(req *models.ShopCartGetListRequest) ([]*models.ShopCart, error) {
	var orderDateFilter []*models.ShopCart
	getorder, err := c.ShopCartGetList(req, "")
	if err != nil {
		return nil, err
	}
	for _, ord := range getorder.Orders {
		orderDateFilter = append(orderDateFilter, &ord)

	}
	sort.Slice(orderDateFilter, func(i, j int) bool {
		return orderDateFilter[i].Time < orderDateFilter[j].Time
	})
	for _, v := range orderDateFilter {
		fmt.Println(v)
	}
	return orderDateFilter, nil
}
func (c *Controller) DateFilter(req *models.ShopCartGetListRequest, id string) ([]*models.ShopCart, error) {
	var (
		orderDateFilter []*models.ShopCart
	)
	orders, err := c.Strg.ShopCart().GetList(req, id)
	if err != nil {
		return nil, err
	}
	for _, ord := range orders.Orders {
		if ord.Time >= req.From_date && ord.Time < req.To_date {
			orderDateFilter = append(orderDateFilter, &ord)
		}
	}

	return orderDateFilter, nil

}
func (c *Controller) UserTotalPrice(req string) (string, int) {
	var (
		totalPrice = 0
	)
	name, _ := c.Strg.User().GetById(&models.UserPrimaryKey{req})
	for _, val := range History[name.FirstName+" "+name.LastName] {
		totalPrice += val.Price
	}

	return name.FirstName, totalPrice
}
func (c *Controller) ProductTotalSellCount() {
	var (
	// totalCount = map[string]int{}
	)
	data, err := read("/shop_cart.json")
	if err != nil {
		log.Printf("Error while read shopcart!!!")
		return
	}
	for _, val := range data {
		c.countTotalProds(val)
	}
	fmt.Println(TotalCountOfSellProduct)

}
func (c *Controller) FindActiveProducts(num int) map[string]int {
	keys := make([]string, 0, len(TotalCountOfSellProduct))
	activeProducts := map[string]int{}

	for k := range TotalCountOfSellProduct {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		// fmt.Println(k, TotalCountOfSellProduct[k])
		activeProducts[k] = TotalCountOfSellProduct[k]
		num--
		if num < 0 {
			break
		}
	}
	return activeProducts
}
func (c *Controller) FindPassiveProducts(num int) map[string]int {
	var (
		count               = 1
		passiveElementsList = []string{}
		countProduct        = make(map[string]int)
	)
	keys := make([]string, 0, len(TotalCountOfSellProduct))

	for k := range TotalCountOfSellProduct {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for i := len(keys) - 1; i > len(keys)-1-num; i-- {
		passiveElementsList = append(passiveElementsList, keys[i])
	}
	for _, k := range passiveElementsList {
		// fmt.Println(k, TotalCountOfSellProduct[k])
		countProduct[k] = TotalCountOfSellProduct[k]
		if num < count {
			break
		}
		count++
	}
	return countProduct
}
func (c *Controller) ActiveUser() (string, error) {
	users := make(map[string]int)
	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{}, "")
	if err != nil {
		return "", err
	}
	for _, value := range getorder.Orders {
		if value.Status == true {
			getproduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: value.ProductId})
			if err != nil {
				return "", err
			}
			users[value.UserId] += value.Count * getproduct.Price
		}
	}
	user, sum := "", 0
	for key, value := range users {
		if sum < value {
			user = key
			sum = value
		}
	}
	getuser, err := c.GetByIdUser(&models.UserPrimaryKey{
		Id: user,
	})
	if err != nil {
		return "", err
	}
	return getuser.FirstName, nil
}

func (c *Controller) TopTime() (*models.DateHistory, error) {
	getorder, err := c.ShopCartGetList(&models.ShopCartGetListRequest{}, "")
	if err != nil {
		return nil, err
	}

	var (
		countcalc = 0
		data      = ""
		count     = []int{}
		date      = []string{}
	)

	for _, value := range getorder.Orders {
		if value.Status == true {
			if value.Count >= countcalc {
				countcalc = value.Count
				data = value.Time
			}
		}
	}
	for _, value := range getorder.Orders {
		if value.Status == true {
			if value.Count == countcalc {
				count = append(count, countcalc)
				date = append(date, data)
			}
		}
	}
	result := models.DateHistory{
		Count: count,
		Date:  date,
	}

	return &result, nil
}

func (c *Controller) countTotalProds(req []models.ShopCart) {
	for _, val := range req {
		name, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{val.ProductId})
		if err != nil {
			return
		}
		TotalCountOfSellProduct[name.Name] += val.Count
	}
}

func (c *Controller) findProductsWithinUser(req []*models.ShopCartHistory) {
	for _, val := range req {
		fmt.Printf("Name: %v\t\tPrice:%v\t\tCount:%v\t\tTotal:%v\t\tTime:%v \n", val.ProductName, val.Price, val.Count, val.TotalPrice, val.Time)
	}
}

func read(fileName string) (map[string][]models.ShopCart, error) {
	var (
		orders   []*models.ShopCart
		orderMap = make(map[string][]models.ShopCart)
	)
	cfg := config.Load()

	data, err := ioutil.ReadFile(cfg.Path + fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, order := range orders {
		orderMap[order.UserId] = append(orderMap[order.UserId], *order)
	}

	return orderMap, nil
}
