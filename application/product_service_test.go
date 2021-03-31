package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	//Prolonga a execução até que os demais comandos subquesentes terminem
	defer ctrl.Finish()

	//product := mock
}
