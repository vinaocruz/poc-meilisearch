# POC MeiliSearch
Este repositório tem como objetivo subir container do MeiliSearch e avaliar comportamento com diversos tamanho de base de dados.

## Data Source
Na pasta `data/` encontram-se os seguintes dados em formato json para importação como exemplo:

- Filmes: ~32k disponibilizado pelo MeiliSearch
- Usuários fake: ~478k gerado em [Mock Turtle]
- undefined: ~2M

## Docker
Primeiramente, crie um arquivo `.env` através do modelo `.env.example`, atualizandos os valores. Depois, utilize o docker-compose para criar um container do MeiliSearch local

```sh
docker-compose up -d
```

Para acessar a interface de navegação:
```sh
localhost:7700
```
> Nota: essa interface é indisponível para ambiente de `production`

## Importação de dados

Para importar os arquivos json para base no MeiliSearch, execute na raiz do projeto passando como argumento o index do arquivo (que será o mesmo da base):

```sh
go run script/import.go movies
go run script/import.go users
```

[Mock Turtle]: <https://mockturtle.net>
