package repository

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type SalesRepository struct {
	db db.DB
}

func NewSalesRepository(db db.DB) SalesRepository {
	return SalesRepository{db}
}

func (u *SalesRepository) LoadOrCreate() ([]Sales, error) {
	records, err := u.db.Load("sales")
	if err != nil {
		records = [][]string{
			{"date", "product", "quantity", "price", "total"},
		}
		if err := u.db.Save("sales", records); err != nil {
			return nil, err
		}
	}

	result := make([]Sales, 0)
	for i := 1; i < len(records); i++ {
		date, err := time.Parse("2006-01-02 15:04:05", records[i][0])
		if err != nil {
			return nil, err
		}
		category := records[i][1]
		product := records[i][2]
		quantity, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		price, err := strconv.Atoi(records[i][4])
		if err != nil {
			return nil, err
		}

		total, err := strconv.Atoi(records[i][5])
		if err != nil {
			return nil, err
		}

		// append ke result
		result = append(result, Sales{
			Date:        date,
			Category:    category,
			ProductName: product,
			Quantity:    quantity,
			Price:       price,
			Total:       total,
		})
	}

	return result, nil
}

func (u *SalesRepository) Save(sales []Sales) error {
	records := [][]string{
		{"date", "category", "product", "quantity", "price", "total"},
	}
	for i := 0; i < len(sales); i++ {
		// append ke records
		records = append(records, []string{
			sales[i].Date.Format("2006-01-02 15:04:05"),
			sales[i].Category,
			sales[i].ProductName,
			strconv.Itoa(sales[i].Quantity),
			strconv.Itoa(sales[i].Price),
			strconv.Itoa(sales[i].Total),
		})
	}
	return u.db.Save("sales", records)
}

func (u *SalesRepository) Add(cartItems []CartItem) error {
	sales, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for _, item := range cartItems {
		// TODO: answer here
		total := item.Quantity * item.Price
		sales = append(sales, Sales{
			Date:        time.Now(),
			Category:    item.Category,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Total:       total,
		})

	}

	return u.Save(sales)
}

func (u *SalesRepository) Get(request GetSalesRequest) ([]Sales, error) {
	sales, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	if isEmptyRequest := request.IsEmptyRequest(); isEmptyRequest {
		return sales, nil
	}

	if isValidRequest := request.IsValidRequest(); !isValidRequest {
		return nil, fmt.Errorf("bad request")
	}

	if request.StartPeriod != nil && request.EndPeriod != nil && request.ProductName == "" {
		return GetTimePeriodSales(sales, request.StartPeriod, request.EndPeriod), nil
	}

	if request.StartPeriod == nil && request.EndPeriod == nil && request.ProductName != "" {
		return GetProductNameSales(sales, request.ProductName), nil
	}

	// return keduanya
	return GetProductNameTimePeriodSales(sales, request.ProductName, request.StartPeriod, request.EndPeriod), nil
}

func GetProductNameSales(sales []Sales, productName string) []Sales {
	var productSales []Sales
	for _, product := range sales {
		// cek apakah product name sama dengan product name yang dikirimkan
		if product.ProductName == productName {
			// jika iya maka append
			// productSales = append(productSales, product)
			newProduct := Sales{
				Date:        product.Date,
				Category:    product.Category,
				ProductName: product.ProductName,
				Price:       product.Price,
				Quantity:    product.Quantity,
				Total:       product.Total,
			}
			productSales = append(productSales, newProduct)
		}
	}

	return productSales
}

func GetTimePeriodSales(sales []Sales, startPeriod *time.Time, endPeriod *time.Time) []Sales {
	endOfDay := time.Date(endPeriod.Year(), endPeriod.Month(), endPeriod.Day(), 23, 59, 59, 0, time.UTC)
	var productSales []Sales
	log.Println(endOfDay, startPeriod)
	for _, product := range sales {
		isBeforeEnd := product.Date.Before(endOfDay)

		if isBeforeEnd {
			newProduct := Sales{
				Date:        product.Date,
				Category:    product.Category,
				ProductName: product.ProductName,
				Price:       product.Price,
				Quantity:    product.Quantity,
				Total:       product.Total,
			}
			productSales = append(productSales, newProduct)
		}
	}
	return productSales
}

func GetProductNameTimePeriodSales(sales []Sales, productName string, startPeriod *time.Time, endPeriod *time.Time) []Sales {
	var productSales []Sales
	endOfDay := time.Date(endPeriod.Year(), endPeriod.Month(), endPeriod.Day(), 23, 59, 59, 0, time.UTC)
	for _, product := range sales {
		isBeforeEnd := product.Date.Before(endOfDay)

		if isBeforeEnd && product.ProductName == productName {
			newProduct := Sales{
				Date:        product.Date,
				Category:    product.Category,
				ProductName: product.ProductName,
				Price:       product.Price,
				Quantity:    product.Quantity,
				Total:       product.Total,
			}
			productSales = append(productSales, newProduct)
		}

	}
	return productSales
}
