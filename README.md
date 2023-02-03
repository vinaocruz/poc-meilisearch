# POC MeiliSearch
Este repositório tem como objetivo subir container do MeiliSearch e avaliar comportamento com diversos tamanho de base de dados. Utilizei o golang para criar o script de importação dos dados

## Data Source
Na pasta `data/` encontram-se os seguintes dados em formato ndjson para importação como exemplo:

- Filmes: ~32k disponibilizado pelo MeiliSearch
- Usuários (fake): ~478k gerados com [Mock Turtle]
- Empregos (fake): ~1.1M gerados com [Mock Turtle]

> Nota: a quantidade de linhas dos arquivos podem divergir da quantidade de documentos indexados, pois os dados foram gerados automaticamente, sem garantia de ter `id` efetivamente único

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

Para importar os arquivos ndjson para o MeiliSearch, execute na raiz do projeto passando como argumento o index do arquivo (que será o mesmo da base):

```sh
go run script/import.go movies
go run script/import.go users
go run script/import.go jobs
```

> Nota: atualmente está configurado um limite para dataset de 250Mb. Caso deseja enviar um arquivo maior, configure o `MEILI_HTTP_PAYLOAD_SIZE_LIMIT` no .env

[Mock Turtle]: <https://mockturtle.net>
