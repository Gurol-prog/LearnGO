package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type product struct {
	Id          int     `json:"id,omitempty"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}
type category struct {
	Id           int    `json:id`
	CategoryName string `json:categoryName`
}

//tüm ürünleri listeyelen fonksiyon

func GetAllProduct() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println("Hata var")
	}
	defer response.Body.Close()

	BodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []product
	json.Unmarshal(BodyBytes, &products)
	fmt.Println(products)

}

func AddProduct() {
	product1 := product{ProductName: "Kalem", CategoryId: 2, UnitPrice: 10.0}
	jsonProduct, err := json.Marshal(product1)
	if err != nil {
		fmt.Println("Hata var")
	}

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println("Hata var")
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse product
	json.Unmarshal(bodyBytes, &productResponse)

	fmt.Println("Kaydedildi: ", productResponse)
}

func GetAllProductRefactor() ([]product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	BodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []product
	json.Unmarshal(BodyBytes, &products)
	return products, nil

}

func AddProductRefactor() (product, error) {
	product1 := product{ProductName: "Kalem", CategoryId: 2, UnitPrice: 10.0}
	jsonProduct, _ := json.Marshal(product1)

	response, _ := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse product
	json.Unmarshal(bodyBytes, &productResponse)

	return productResponse, nil
}
