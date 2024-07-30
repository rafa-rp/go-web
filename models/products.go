package models

import "go-web/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func FindAllProducts() []Product {

	db := db.DbConection()

	allProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func SaveProduct(name, description string, price float64, quantity int) {
	db := db.DbConection()

	prepareInsert, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	prepareInsert.Exec(name, description, price, quantity)
	defer db.Close()
}
