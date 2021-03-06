package application_test

import (
	"testing"

	"github.com/MichelGomes/hexagonal-with-go/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestEnable_Product_With_Price_Greater_Than_Zero(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}

func TestNot_Possible_To_Enable_Product_With_Price_Equal_Zero(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Enable()
	require.Equal(t, "the produce price must be greater than zero to enabled it", err.Error())
}

func TestNot_Possible_To_Disable_Product_With_Price_Equal_Zero(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "the produce price must be zero to disabled it", err.Error())
}

func TestPossible_To_Disable_Product_With_Price_Equal_Zero(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

}

func TestIsValid_UUID(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

}

func TestIsValid_invalid_status(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = "INVALID"
	product.Price = 10

	_, err := product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

}

func TestIsValid_price_less_than_zero(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = -5

	_, err := product.IsValid()
	require.Equal(t, "the product price must be greater than or equal do zero", err.Error())

}
