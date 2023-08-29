### To Run

- Navigate until folder cmd/desafio_multithreading

go run main.go <cep>

Ex: go run main.go 01311200

# desafio-multithreading

Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json (fora de operação)

https://brasilapi.com.br/api/cep/v1/" + Cep"

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- [x]Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

-[x] O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- [x]Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
