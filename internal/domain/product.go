package domain

// Product represents a product in the domain logic.
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
}

// ProductRepository provides an interface for the storage of products.
type ProductRepository interface {
	FindByID(id string) (*Product, error)
	FindAll() ([]*Product, error)
	Save(product *Product) error
	Delete(id string) error
}