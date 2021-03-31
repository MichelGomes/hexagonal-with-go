package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// O Metodo comecando com Letra maiscula são metodos publicos, ou seja, poderão ser acessados fora do pacote.
// O Metodo comecando com Letra minuscula são metodos privados, ou seja, NÃO poderão ser acessados fora do pacote.
// Não existem métodos protected no golang
type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	getId() string
	getName() string
	getStatus() string
	getPrice() float32
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float32) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

//Interface Composition
type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

// golang implements the interfaces automatically
//Using tags to check the value
type Product struct {
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float32 `valid:"float,optional"`
}

//* in golang is a pointer
//& in golang location of value in memory
func NewProduct() *Product {
	product := Product{
		Id:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product

}

func (p *Product) IsValid() (bool, error) {

	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the product price must be greater than or equal do zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (p *Product) Enable() error {

	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the produce price must be greater than zero to enabled it")
}

func (p *Product) Disable() error {

	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the produce price must be zero to disabled it")

}

func (p *Product) getId() string {
	return p.Id
}

func (p *Product) getName() string {
	return p.Name
}

func (p *Product) getStatus() string {
	return p.Status
}

func (p *Product) getPrice() float32 {
	return p.Price
}
