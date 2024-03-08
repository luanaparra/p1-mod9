# prova 1 - módulo 9

## Como executar?
Para a execução dos módulos, é necessário clonar este repositório:

```
git clone https://github.com/luanaparra/p1-mod9
```

Caso precise, instale o mosquitto com o arquivo de configuração e dependências do GO.

```
mosquitto -c mosquitto.conf
go mod tidy
```

Dessa maneira, para rodar o publisher é preciso realizar esses comandos:

```
cd publisher
go run publisher.go
```

Por outro lado, temos o subscriber:

```
cd subscriber
go run subscriber.go
```

Por conseguinte, para rodar os testes, diferentes comandos são necessários:

**Publisher**
```
cd publisher
go test -v
```

**Subscriber**
```
cd subscriber
go run subscriber.go
```

Desse modo, os testes foram adicionados para a avaliação do recebimento das mensagens pelo broker usando o mecanismo de QoS, além do alerta dos sensores. Ademais, os testes asseguram que as mensagens enviadas pelo publisher foram recebidas corretamente pelo subscriber. 

## Demonstração

As imagens do exercício se encontram na pasta \demonstração.



