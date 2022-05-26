package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
	// ambil semua data product
	records, err := u.db.Load("products")
	if err != nil {
		// jika error maka bikin kan data product baru
		records = [][]string{
			{"category", "product_name", "price"},
		}

		// save
		err = u.db.Save("products", records)
		if err != nil {
			return nil, err
		}
	}

	// jika ada data product maka load semua
	result := make([]Product, 0) // variable untuk menampung data product
	for i := 1; i < len(records); i++ {
		// convert string harga ke int
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		// isi product
		product := Product{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
		}

		// append
		result = append(result, product)
	}

	return result, nil
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	// load semua product menggunakan method LoadOrCreate
	products, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	// return semua product
	return products, nil
}
