# SICONV API

> O SICONV disponibiliza acesso livre às informações de Transferências Voluntárias da União com o objetivo de facilitar o acesso aos dados do sistema para a sociedade e a outras esferas de Governo. O usuário poderá baixar as principais informações de Convênios para realizar análises e cruzamentos que desejar a partir desses [dados](http://portal.convenios.gov.br/download-de-dados).

The goal of this project is to provide a REST API around the [CSV files](http://portal.convenios.gov.br/images/docs/CGSIS/csv/siconv.zip) provided by SICONV.

## Run with Docker Compose
    
    # Run containers in background
    $ docker-compose up -d
    
    # Create database schema
    $ docker run -it --rm --link siconv_postgres:postgres -v $PWD/migrations:/tmp/migrations --network siconv_default postgres bash -c "export PGPASSWORD=secret && psql -h postgres -U postgres siconv < /tmp/migrations/1525840555_initial.up.sql" 

## Run with Docker

    $ SICONV_TOKEN=$(openssl rand -base64 32)
    $ docker run -d --restart always --name siconv -e "SICONV_TOKEN=$SICONV_TOKEN" -e "SICONV_DATABASE_URL=postgres://<user>:<password>@<host>:5432/siconv_db" -p 8080:8080 insighted4/siconv
    
## Restore Database

    $ wget --continue http://portal.convenios.gov.br/images/docs/CGSIS/csv/siconv.zip
    $ docker run -t --rm \
      	-v $PWD/siconv.zip:/data/siconv.zip \
      	insighted4/siconv siconv restore --truncate \
          	--file /data/siconv.zip \
          	--database-url postgres://<user>:<password>@<host>:5432/siconv_db

## Restore Database Schema

    psql -U username siconv_db < 1525840555_initial.up.sql

## Endpoints

### Annonymous Access (Read-Only)
        
    GET /consorcios
    GET /consorcios/:id
    GET /convenios
    GET /convenios/:id
    GET /desembolsos
    GET /desembolsos/:id
    GET /emendas
    GET /emendas/:id
    GET /empenhos
    GET /empenhos/:id
    GET /etapa-crono-fisicos
    GET /etapa-crono-fisicos/:id
    GET /historico-situacoes
    GET /historico-situacoes/:id
    GET /ingresso-contrapartidas
    GET /ingresso-contrapartidas/:id
    GET /meta-crono-fisicos
    GET /meta-crono-fisicos/:id
    GET /obtv-convenentes
    GET /obtv-convenentes/:id
    GET /pagamentos
    GET /pagamentos/:id
    GET /plano-aplicacao-detalhados
    GET /plano-aplicacao-detalhados/:id
    GET /programas
    GET /programas/:id
    GET /programa-propostas
    GET /programa-propostas/:id
    GET /proponentes
    GET /proponentes/:id
    GET /propostas
    GET /propostas/:id
    GET /prorroga-oficios
    GET /prorroga-oficios/:id
    GET /termo-aditivos
    GET /termo-aditivos/:id


### Authorized Access (Bearer Token Required)
            
    POST /consorcios
    POST /convenios
    POST /desembolsos
    POST /emendas
    POST /empenhos
    POST /etapa-crono-fisicos
    POST /historico-situacoes
    POST /ingresso-contrapartidas
    POST /meta-crono-fisicos
    POST /obtv-convenentes
    POST /pagamentos
    POST /plano-aplicacao-detalhados
    POST /programas
    POST /programa-propostas
    POST /proponentes
    POST /propostas
    POST /prorroga-oficios
    POST /termo-aditivos


## License

[MIT](LICENSE)

