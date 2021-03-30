package application

import "errors"

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

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

// golang implements the interfaces automatically
type Product struct {
	Id     string
	Name   string
	Status string
	Price  float32
}

func (p *Product) IsValid(bool, error) {

}

func (p *Product) Enable() error {

	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the produce price must be greater than zero to enabled it")
}

//func (p *Product) Disable() error {
//
//}

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
