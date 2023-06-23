package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsondb"
	"log"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	// users, err := con.ShopCartDateFilter(&models.ShopCartGetListRequest{0, 0, "2022-03-13 20:53:48", "2023-01-22 14:21:38"}, "48097741-22c9-4663-8796-3c9993d88ffe")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, val := range users {
	// 	fmt.Println(val)
	// }
	_, err = con.ShopCartHistory()
	if err != nil {
		log.Println("Error !!!")
		return
	}
	// fmt.Println("ShopCart History Finished!")

	// products := con.FindActiveProducts(10)
	// for key, val := range products {
	// 	fmt.Println(key, ":", val)
	// }
	// fmt.Println("products finished!")

	// con.ProductTotalSellCount()
	// order := con.FindActiveProducts(10)
	// for key, val := range order {
	// 	fmt.Println(key, ":", val)
	// }

	// pasOrder := con.FindPassiveProducts(10)
	// for key, val := range pasOrder {
	// 	fmt.Println(key, ":", val)
	// }
	// user, num := con.UserTotalPrice("27457ac2-74dd-4656-b9b0-0d46b1af10dc")
	// fmt.Println(user, ":", num)
}
