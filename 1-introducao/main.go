package main

import (
	"github.com/google/uuid"
)

//subir o docker: docker-compose up -d
//criando uma tabela no banco:
//1- docker-compose exec db bash
//2- mysql -uuser -p db
//3- password
//4- create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key(id));

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

}
