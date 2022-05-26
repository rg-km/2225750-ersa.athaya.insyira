package main

type PrimaryKey int

type InvoiceRow struct {
	ID               PrimaryKey
	SubscriptionCode string
	CustomerName     string
	Address          string
	Phone            string
}

type InvoiceDB struct {
	m map[PrimaryKey]InvoiceRow
}

func NewInvoice() *InvoiceDB {
	return &InvoiceDB{
		m: make(map[PrimaryKey]InvoiceRow),
	}
}

func (db *InvoiceDB) Insert(code string, name string, address string, phone string) {
	db.m[PrimaryKey(len(db.m)+1)] = InvoiceRow{
		ID:               PrimaryKey(len(db.m)) + 1,
		SubscriptionCode: code,
		CustomerName:     name,
		Address:          address,
		Phone:            phone,
	}
}

func (db *InvoiceDB) Where(id PrimaryKey) *InvoiceRow {
	// cek apakah id ada di db
	if _, ok := db.m[id]; ok {
		// return data di db
		return &InvoiceRow{
			ID:               id,
			SubscriptionCode: db.m[id].SubscriptionCode,
			CustomerName:     db.m[id].CustomerName,
			Address:          db.m[id].Address,
			Phone:            db.m[id].Phone,
		}
	}

	return nil
}

func (db *InvoiceDB) Update(id PrimaryKey, code string, name string, address string, phone string) (*InvoiceRow, error) {
	// cek apakah id ada di db
	if _, ok := db.m[id]; ok {
		// update data di db
		db.m[id] = InvoiceRow{
			ID:               id,
			SubscriptionCode: code,
			CustomerName:     name,
			Address:          address,
			Phone:            phone,
		}

		return &InvoiceRow{
			ID:               id,
			SubscriptionCode: code,
			CustomerName:     name,
			Address:          address,
			Phone:            phone,
		}, nil
	}

	return nil, nil
}
