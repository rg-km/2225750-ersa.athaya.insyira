package repository

type TransactionRepository struct {
	cartItemRepository CartItemRepository
	salesRepository    SalesRepository
}

func NewTransactionRepository(cartItemRepository CartItemRepository, salesRepository SalesRepository) TransactionRepository {
	return TransactionRepository{cartItemRepository, salesRepository}
}

func (u *TransactionRepository) Pay(cartItems []CartItem, amount int) (int, error) {
	// ambil total price
	totalPrice, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}

	// perhitungan
	totalPay := amount - totalPrice

	// tambahakan ke tabel sales
	err = u.salesRepository.Add(cartItems)
	if err != nil {
		return 0, err
	}

	// return totalPay
	return totalPay, nil
}
