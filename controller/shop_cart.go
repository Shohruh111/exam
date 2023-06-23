package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) ShopCartCreate(req *models.CreateShopCart) (*models.ShopCart, error) {

	log.Printf("User create req: %+v\n", req)

	resp, err := c.Strg.ShopCart().Create(req)
	if err != nil {
		log.Printf("error while user Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) GetById(req *models.ShopCartPrimaryKey) (*models.ShopCart, error) {

	resp, err := c.Strg.ShopCart().GetById(req)
	if err != nil {
		log.Printf("error while user GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ShopCartGetList(req *models.ShopCartGetListRequest, id string) (*models.ShopCartGetListResponse, error) {

	strg, err := c.Strg.ShopCart().GetList(req, id)
	if err != nil {
		log.Printf("error while user GetList: %+v\n", err)
		return nil, err
	}

	return strg, nil
}

func (c *Controller) ShopCartUpdate(req *models.UpdateShopCart, id string) (*models.ShopCart, error) {

	resp, err := c.Strg.ShopCart().Update(req, id)
	if err != nil {
		log.Printf("error while user Update: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) ShopCartDelete(req *models.ShopCartPrimaryKey) error {

	err := c.Strg.ShopCart().Delete(req)
	if err != nil {
		log.Printf("error while user Delete: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) ShopCartDateFilter(req *models.ShopCartGetListRequest, id string) ([]*models.ShopCart, error) {
	resp, err := c.Strg.ShopCart().DateFilter(req, id)
	if err != nil {
		log.Printf("error while shopCartDateFilter: %v\n", req)
		return nil, err
	}
	return resp, nil
}
