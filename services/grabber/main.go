package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Link        string   `json:"link"`
}

func main() {
	url := "http://localhost:8080/api/products/v1/add"
	method := "POST"

	products, err := readProductsFromFile("./result.json")
	if err != nil {
		fmt.Println(err)
	}

	for _, product := range products[10:] {
		jsonData, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(resp.StatusCode)
	}

}

func readProductsFromFile(filename string) ([]Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []Product

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
