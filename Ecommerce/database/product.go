package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(prodID int) *Product {
	for _, prod := range productList {
		if prod.ID == prodID {
			return &prod
		}
	}
	return nil
}

func Update(product Product) {
	for idx, prod := range productList {
		if prod.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(product Product) {
	var tempList []Product
	for idx, prod := range productList {
		if prod.ID != product.ID {
			tempList[idx] = prod
		}
	}
	productList = tempList
}

func init() {
	products := []Product{
		{
			ID:          1,
			Title:       "Laptop",
			Description: "A good laptop for daily work",
			Price:       75000,
			ImgUrl:      "https://example.com/images/laptop.jpg",
		},
		{
			ID:          2,
			Title:       "Smartphone",
			Description: "A smartphone with a clear display",
			Price:       25000,
			ImgUrl:      "https://example.com/images/smartphone.jpg",
		},
		{
			ID:          3,
			Title:       "Headphone",
			Description: "Wireless headphone with good sound",
			Price:       3500,
			ImgUrl:      "https://example.com/images/headphone.jpg",
		},
		{
			ID:          4,
			Title:       "Keyboard",
			Description: "Mechanical keyboard for typing",
			Price:       4200,
			ImgUrl:      "https://example.com/images/keyboard.jpg",
		},
		{
			ID:          5,
			Title:       "Mouse",
			Description: "Wireless mouse for office use",
			Price:       1200,
			ImgUrl:      "https://example.com/images/mouse.jpg",
		},
	}

	for _, product := range products {
		Store(product)
	}
}
