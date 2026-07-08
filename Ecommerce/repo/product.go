package repo

import "errors"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Store(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() []Product
}

type productRepo struct {
	productList []Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)

	return repo
}

func (p *productRepo) Store(product Product) (*Product, error) {
	product.ID = len(p.productList) + 1
	p.productList = append(p.productList, product)
	return &p.productList[len(p.productList)-1], nil
}

func (p *productRepo) List() []Product {
	return p.productList
}

func (p *productRepo) Get(productID int) (*Product, error) {
	for idx := range p.productList {
		if p.productList[idx].ID == productID {
			return &p.productList[idx], nil
		}
	}

	return nil, errors.New("product not found")
}

func generateInitialProducts(p *productRepo) {
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
		p.productList = append(p.productList, product)
	}
}
