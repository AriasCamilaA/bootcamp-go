/*
	Crear un programa que cumpla los siguiente puntos:

	1. Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
	2. Tener un slice global de Product llamado Products instanciado con valores.
	3. 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar
		el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll()
		deberá imprimir todos los productos guardados en el slice Products.
	4. Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto
		correspondiente al parámetro pasado.
	5. Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products []Product = []Product{
	{ID: 1, Name: "Laptop", Price: 1500.50, Description: "Laptop HP", Category: "Computers"},
	{ID: 2, Name: "Mouse", Price: 50.50, Description: "Mouse Logitech", Category: "Computers"},
	{ID: 3, Name: "Keyboard", Price: 100.50, Description: "Keyboard Razer", Category: "Computers"},
	{ID: 4, Name: "Monitor", Price: 500.50, Description: "Monitor Samsung", Category: "Computers"},
}

func (p Product) Save(_product Product) {
	Products = append(Products, _product)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func (p Product) getById(_id int) *Product {
	for _, product := range Products {
		if product.ID == _id {
			return &product
		}
	}
	return nil
}

func main() {

	newProduct := Product{ID: 5, Name: "Tablet", Price: 300.00, Description: "Tablet Apple", Category: "Electronics"}

	newProduct.Save(newProduct)

	newProduct.GetAll()

	product := newProduct.getById(3)

	if product != nil {
		fmt.Printf("Product found: ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	} else {
		fmt.Println("Product not found")
	}
}
