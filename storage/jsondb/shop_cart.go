package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"app/models"
)

const ShortForm = "2022-10-07 22:51:01"

type timeSlice []models.ShopCart

type ShopCartRepo struct {
	fileName string
	file     *os.File
}

func NewShopCartRepo(fileName string, file *os.File) *ShopCartRepo {
	return &ShopCartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *ShopCartRepo) Create(req *models.CreateShopCart) (*models.ShopCart, error) {

	orders, err := u.read()
	if err != nil {
		return nil, err
	}

	var (
		order = models.ShopCart{
			ProductId: req.ProductId,
			UserId:    req.UserId,
			Count:     req.Count,
			Time:      time.Time.Format(time.Now(), "2022-10-07 22:51:01"),
		}
	)
	orders[order.UserId] = append(orders[order.UserId], order)

	err = u.write(orders)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (u *ShopCartRepo) GetById(req *models.ShopCartPrimaryKey) (*models.ShopCart, error) {

	var order models.ShopCart

	orders, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := orders[req.UserId]; !ok {
		return nil, errors.New("order not found")
	}

	for _, ord := range orders[req.UserId] {
		if ord.ProductId == req.ProductId {
			order = ord
		}
	}

	return &order, nil
}

func (u *ShopCartRepo) GetList(req *models.ShopCartGetListRequest, id string) (*models.ShopCartGetListResponse, error) {

	var resp = &models.ShopCartGetListResponse{}
	resp.Orders = []models.ShopCart{}

	response := []models.ShopCart{}

	ShopCartMap, err := u.read()
	if err != nil {
		return nil, err
	}

	if id != "" {
		for _, val := range ShopCartMap[id] {
			users := val
			resp.Orders = append(resp.Orders, users)
			resp.Count += len(resp.Orders)
		}
	}
	for _, val := range ShopCartMap {
		response = append(response, val...)
	}

	resp.Orders = append(resp.Orders, response...)
	return resp, nil
}

func (u *ShopCartRepo) Update(req *models.UpdateShopCart, id string) (*models.ShopCart, error) {
	var saveUpdate models.ShopCart
	orders, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := orders[id]; !ok {
		return nil, errors.New("user not found")
	}
	for _, ord := range orders[id] {
		if ord.UserId == id {
			ord = models.ShopCart{
				ProductId: req.ProductId,
				Count:     req.Count,
				Time:      time.Now().Format("2022-04-04 04:41:01"),
			}
			saveUpdate = ord
		}
	}

	err = u.write(orders)
	if err != nil {
		return nil, err
	}

	return &saveUpdate, nil
}

func (u *ShopCartRepo) Delete(req *models.ShopCartPrimaryKey) error {

	orders, err := u.read()
	if err != nil {
		return err
	}

	for _, ord := range orders[req.UserId] {
		if ord.ProductId == req.ProductId {
			delete(orders, ord.ProductId)
		}
	}

	err = u.write(orders)
	if err != nil {
		return err
	}

	return nil
}

func (u *ShopCartRepo) DateFilter(req *models.ShopCartGetListRequest, id string) ([]*models.ShopCart, error) {
	var (
		orderDateFilter []*models.ShopCart
	)
	orders, err := u.GetList(req, id)
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

func (u *ShopCartRepo) read() (map[string][]models.ShopCart, error) {
	var (
		orders   []*models.ShopCart
		orderMap = make(map[string][]models.ShopCart)
	)

	data, err := ioutil.ReadFile(u.fileName)
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

func (u *ShopCartRepo) write(ShopCartMap map[string][]models.ShopCart) error {

	var orders []models.ShopCart

	for _, val := range ShopCartMap {
		orders = append(orders, val...)
	}

	body, err := json.MarshalIndent(orders, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
