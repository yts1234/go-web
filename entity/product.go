package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() (status string) {

	if p.Stock < 3 {
		status = "Stock almost run out"
	} else if p.Stock < 10 {
		status = "Stock is limited"
	} else {
		status = "Stock is available"
	}
	return
}
