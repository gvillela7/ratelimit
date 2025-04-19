# ratelimit
Desafio  fullcycle rate limit

Para configurar a API, edite o arquivo config.toml e modifique os parâmetros conforme necessário
```
rate_limit_request = 5 --> Quantidade de requests
rate_limit_time_second = 10 --> Tempo em segundos 
rate_limit_time_block_second = 120 ## 2 minutes --> Tempo de bloqueio
```
Com a configuração acima, a API vai aceitar 5 requisições em um intervalo de 10 segundos.
Caso uma sexta requisição seja enviada, o IP ou o token será bloqueado por 2 minutos.

Para rodar o desafio basta executar o seguinte comando:
```
docker compose up -d --build

```
```
 curl -X GET http://localhost:8080
 Response: 200
 {
    "StatusCode":200,
    "message":"Success",
    "data":[
      {"message":"Ok"}
    ]
  }
  Response: 429
 {
    "StatusCode":429,
    "message":"To Many Request",
    "data":[
       {
         "message": "you have reached the maximum number of requests or actions allowed within a certain time frame."
       }
    ]
 }
 ```
 
## Test
```
 docker compose run app go test test/ratelimit_test.go
 docker compose run app go test test/ratelimitBlock_test.go
```