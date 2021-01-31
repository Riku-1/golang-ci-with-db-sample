package main

import (
	"errors"
	"fmt"
	"golang-ci-with-db-sample/config"
	"golang-ci-with-db-sample/domain"
	"gorm.io/gorm"
	"log"
)

func main() {
	c, err := config.GetDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.GetGormDB(&c)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.Product{})

	r := domain.NewProductRepository(db)

	product, _ := r.Create("first_price", 100)
	fmt.Println("create")
	fmt.Printf("%+v\n", product)

	productGet, _ := r.Get(product.ID)
	fmt.Println("get")
	fmt.Printf("%+v\n", productGet)

	productUpdate, _ := r.Update(product.ID, domain.Product{Code: "updated_price", Price: 10000})
	fmt.Println("update")
	fmt.Printf("%+v\n", productUpdate)

	r.Delete(product.ID)
	_, err = r.Get(product.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("succeeded to delete")
	}
}
