package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/TiagoSilvaLourenco/desafio-multithreading/internal/dto"
)

const (
	BrasilApi = "BrasilApiCep"
	ViaCepApi = "ViaCepApi"
)

func request(url, nameApi string, ch chan interface{}) {

	// c := http.Client{Timeout: time.Second}
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data interface{}

	if nameApi == BrasilApi {
		var brasilApiCep dto.BrasilApiCep
		data = &brasilApiCep
		data.(*dto.BrasilApiCep).NameApi = "Retornado por Brasil Cep API"
	} else if nameApi == ViaCepApi {
		var viaCep dto.ViaCep
		data = &viaCep
		data.(*dto.ViaCep).NameApi = "Retornado por ViaCep API"
	} else {
		fmt.Printf("API desconhecida")
	}

	err = json.Unmarshal(res, data)
	if err != nil {
		panic(err)
	}

	ch <- data
}

func main() {
	chBrasilApi := make(chan interface{})
	chViaCepApi := make(chan interface{})

	for _, cep := range os.Args[1:] {
		urlViaCep := "http://viacep.com.br/ws/" + cep + "/json/"
		urlBrasilCep := "https://brasilapi.com.br/api/cep/v1/" + cep

		go request(urlBrasilCep, BrasilApi, chBrasilApi)
		go request(urlViaCep, ViaCepApi, chViaCepApi)

		select {
		case resBrasilApi := <-chBrasilApi:
			fmt.Println(resBrasilApi)

		case resViaCepApi := <-chViaCepApi:
			fmt.Println(resViaCepApi)

		case <-time.After(time.Second):
			println("Timeout")
		}
	}
}
