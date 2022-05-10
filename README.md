# Project Name
> animalia-api

## Indíce
* [Informação](#informação)
* [Tecnologias](#tecnologias)
* [Instalação](#instalação)
* [Ambiente](#ambiente)
* [Arquitetura de pastas](#arquitetura-de-pastas)
* [Iniciando](#iniciando)
* [Testes](#testes)

## Informação
Este projeto é reponsável por gerenciar as informações de identificação de cada cidadão pertencente à Federação Animalia

## Tecnologias
* [GoLang](https://golang.org/) - Compilador da linguagem Go
* [Go Mod](https://github.com/golang/mod) - Gerenciador de dependências
* [Gin](https://github.com/gin-gonic/gin) - Framework Go
* [go-redis](https://github.com/go-redis/redis) - Redis client
* [go-redsync](https://github.com/go-redsync) - Implementação de Lock distribuído com Redis em Go
* [postgres](https://github.com/lib/pq) - Postgres driver
* [swag](https://github.com/swaggo/swag) - Documentação OpenAPI
* [SQlmock](https://github.com/DATA-DOG/go-sqlmock) - Sql driver mock para Golang
* [gorm](https://gorm.io/index.html) - Framework para manipulação do banco de dados

## Instalação
Clonar o projeto
``` bash
cd $GOPATH
git clone git@github.com:eduardohslfreire/animalia-api.git
cd animalia-api
```
Instalar as dependências
```bash
$ make dependencies
```

## Ambiente
Configurando as variáveis de ambiente

| Nome | Descrição | Valor Padrão | Obrigatório |
| -- | -- | -- | -- |
| HTTP_READ_TIMEOUT | Tempo de timeout em segundos utilizado para leitura na api | 60 | :white_check_mark: |
| HTTP_REQUEST_TIMEOUT | Tempo de timeout em segundos utilizado nas requests feita pela api | 60 | :white_check_mark: |
| HTTP_WRITE_TIMEOUT | Tempo de timeout em segundos utilizado para escrita na api | 60 | :white_check_mark: |
| APP_PORT | Porta padrão que a API irá subir | 5000 | :white_check_mark: |
| LOG_LEVEL | Determina o log level a ser exibido no stdout (controle da severidade do log): DEBUG, ERROR, FATAL, INFO PANIC ou WARN | INFO | :white_check_mark: |
| DB_HOST | Host da base de dados | | :white_check_mark: |
| DB_PORT | Porta da base de dados | | :white_check_mark: |
| DB_USER | Usuário de aplicação da base de dados | | :white_check_mark: |
| DB_PASSWORD | Senha do usuário da aplicação da base de dados | | :white_check_mark: |
| DB_NAME | Nome da base de dados | | :white_check_mark: |
| DB_TIMEZONE | Time Zone padrão da API | 'America/Sao_Paulo' | :white_check_mark: |
| DB_CONN_MAX_LIFETIME | Tempo máximo para que uma conexão inativa pode ser reutilizada em segundos | 15 | :white_check_mark: |
| DB_MAX_IDLE_CONNS | Número máximo de conexões inativas na base de dados PostgreSQL | 3 | :white_check_mark: |
| DB_MAX_OPEN_CONNS | Número máximo de conexões abertas na base de dados PostgreSQL  | 15 | :white_check_mark: |
| REDIS_HOST | Host do Redis | | :white_check_mark: |
| REDIS_PASSWORD | Senha do usuário da aplicação do Redis | | :white_check_mark: |
| REDIS_EXPIRATION_HOURS | Tempo da expiração das mensagens no Redis em horas | 48 | :white_check_mark: |


## Arquitetura de pastas
### Diretórios
```bash
├── animalia-api
├── api
│   ├── dto
│   ├── errors
│   ├── handler
│   ├── middleware
│   └── validation
│   └── server.go
├── config
│   ├── cache
│   ├── db
│   └── env
├── database
│   ├── init
│   └── migrations
├── dev
├── docker-compose.yml
├── Dockerfile
├── docs
├── entity
├── infrastructure
│   └── repository
├── main.go
├── Makefile
├── pkg
│   ├── cache
│   ├── db
│   ├── env
│   ├── logger
│   └── metric
├── README.md
├── usecase
├── util

```
Uma breve descrição dos diretórios:
* `api` contém artefatos para iniciar a API e servir como delivery para os clientes que a utilizam
* `database` contém os scripts executados no banco de dados. (Inicial e migrações)
* `config` contém as configurações de acesso externo (SO, RabbitMQ, Postgres).
* `dev` contém configurações dos serviços locais que apoiam no desenvolvimento.
* `entity` irá armazenar a estrutura de qualquer objeto (structs, enums, tipos). Esta camada será usada em todas as outras camadas.
* `infrastructure` irá armazenar todos os acessos de banco de dados ou manipuladores de requisições http para outros serviços.
* `pkg` contém todos os pacotes de suporte externo.
* `usecase` contém todas as regras de negócios. Qualquer processo será tratado aqui. Essa camada decidirá qual camada de repositório usará.
* `util` contém as funcões utilitárias que são utilizadas por diversos módulos do sistema.
* `docs` contém a documentação openAPI gerada pela ferramenta de swagger.
* `.gitignore` contém todos os arquivos e diretórios ignorados.
* `Makefile` é usado para construir o projeto, possui utilitários de forma organizada que abstrai a execução de vários comandos do shell.
* `go.mod` contém todos as dependências do projeto.
* `README.md` é uma descrição detalhada do projeto.

## Iniciando
Exportar as variáveis de ambiente
```shell
$ source ./dev/env.dev
```
Subir os serviços locais configurados no docker-compose
```
$ make setup-dev-up
```
Compilar e subir o APP
```
$ make run
```
Executar o migrate para popular a base de dados com os DDL e DMLs
```shell
# Caso não tenha instalado, rodar esse comando antes
$ make migrate-install
# Comando de migrate
$ make migrate-up
```
## Testes
```bash
$ make cover
```

## Swagger
#### Swagger é uma ferramenta para a documentação do contrato das APIs.
Instalação
```shell
make swaggo-install
```
Geração da documentação
```shell
make swaggo-generate
```
Acessar o endereço [local](http://localhost:5000/api/v1/swagger/index.html#/) da aplicação para visualizar a documentação 

### Considerações
``` bash
# certifique-se que efetuou a instalação correta do go na versão 1.17
# IDEs recomendadas: visual studio code
```