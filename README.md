# hexagonal-with-go

CREATE MOCKS USING MOCKGEN

INSTALL MOCKGEN
go install github.com/golang/mock/mockgen@v1.5.0

CREATE MOCKS 
mockgen -destination=application/mocks/application -source=application/product.go application.go