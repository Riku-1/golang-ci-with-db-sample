package domain

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-ci-with-db-sample/config"

	"gorm.io/gorm"
	"log"
	"testing"
)

var db *gorm.DB

func setup() {
	c, err := config.GetDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err = config.GetGormDB(&c)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Product{})
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestGet(t *testing.T) {
	tx := db.Begin()
	defer func() {
		tx.Rollback()
	}()
	r := NewProductRepository(tx)
	product := Product{Price: 100, Code: "test_get"}
	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	res, err := r.Get(product.ID)
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	assert.Equal(t, uint(100), res.Price)
	assert.Equal(t, "test_get", res.Code)
}

func TestCreate(t *testing.T) {
	tx := db.Begin()
	r := NewProductRepository(tx)
	res, err := r.Create("test_create", 200)

	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	// assert response
	assert.Equal(t, "test_create", res.Code)
	assert.Equal(t, uint(200), res.Price)
	tx.Commit()

	// confirm record created
	var product Product
	db.First(&product, res.ID)
	assert.Equal(t, "test_create", product.Code)
	assert.Equal(t, uint(200), product.Price)

	// delete record
	db.Delete(&Product{}, product.ID)
}

func TestDelete(t *testing.T) {
	tx := db.Begin()
	defer func() {
		tx.Rollback()
	}()
	r := NewProductRepository(tx)
	product := Product{Code: "test_delete", Price: 300}
	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	err = r.Delete(product.ID)
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	var dProduct Product
	err = db.First(&dProduct, product.ID).Error
	assert.True(t, errors.Is(gorm.ErrRecordNotFound, err)) // confirm record deleted
}

func TestUpdate(t *testing.T) {
	tx := db.Begin()
	defer func() {
		tx.Rollback()
	}()
	r := NewProductRepository(tx)
	product := Product{Price: 400, Code: "not_updated"}
	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	res, err := r.Update(product.ID, Product{Price: 500, Code: "updated"})
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}

	// assert response
	assert.Equal(t, "updated", res.Code)
	assert.Equal(t, uint(500), res.Price)
	tx.Commit()

	// confirm record created
	var uProduct Product
	db.First(&uProduct, product.ID)
	assert.Equal(t, "updated", uProduct.Code)
	assert.Equal(t, uint(500), uProduct.Price)

	// delete record
	db.Delete(&Product{}, product.ID)
}
