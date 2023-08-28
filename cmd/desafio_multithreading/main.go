package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/TiagoSilvaLourenco/desafio-multithreading/internal/dto"
)

const (
	BrasilApi = "BrasilApiCep"
	ViaCepApi = "ViaCepApi"
)

func request(url, nameApi string, wg *sync.WaitGroup) {

	c := http.Client{Timeout: time.Microsecond}
	req, err := c.Get(url)
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
	wg.Done()
	log.Println(data)
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	for _, cep := range os.Args[1:] {
		urlViaCep := "http://viacep.com.br/ws/" + cep + "/json/"
		urlBrasilCep := "https://brasilapi.com.br/api/cep/v1/" + cep

		go request(urlBrasilCep, BrasilApi, &waitGroup)
		go request(urlViaCep, ViaCepApi, &waitGroup)
	}

	waitGroup.Wait()

}
