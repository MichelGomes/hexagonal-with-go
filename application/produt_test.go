package application_test

import (
	"testing"

	"github.com/MichelGomes/hexagonal-with-go/application"
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
