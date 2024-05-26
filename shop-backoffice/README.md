# Shop-Backoffice

## Rotas Disponíveis:
    POST /backoffice/products
    GET /backoffice/health/status

#### Exemplos de rotas:
    POST http://localhost:8081/backoffice/products
    Content-Type: application/json
    correlationId: "30eb2105-f4cc-49e0-95e8-04da1eca3737"
    traceId: "30eb2105"
    flowId: "99882992"

    {
        "nome": "NEW",
        "departamento": "Informatica",
        "tags": ["Informatica","Notebook","Acessorios","Home-Office", "Promocao"],
        "preco": 1500000,
        "quantidade": 9,
        "ativo": true
    }

    GET http://localhost:8081/backoffice/health/status

#### Curl Rotas:
    curl --request GET \
    --url http://localhost:8081/backoffice/health/status

    curl --request POST \
    --url http://localhost:8081/backoffice/products \
    --header 'Content-Type: application/json' \
    --header 'correlationId: 4e41452c-35e4-44f8-ad02-506328a563bb' \
    --header 'flowId: de2c47b9-84a6-4caa-b540-ecab00e72ab4' \
    --header 'traceId: 0307f376-d0c4-471a-afe1-dca1b39abf4e' \
    --data '{
        "nome": "Celular",
        "departamento": "Informatica",
        "tags": ["Informatica","Notebook","Acessorios","Home-Office", "Promocao"],
        "preco": 1500000,
        "quantidade": 9,
        "ativo": true
    }'

## Como rodar a aplicação:
    Rodar o arquivo docker-compose.yaml, ele será responsável por subir o dynamo e a aplicação na port 8081. Caso a variável de ambiente ENVIRONMENT não tenha sido setada, a aplicação se comunicará com o dynamo iniciado via docker e na inicialização tentará fazer a criação da tabela Products, de forma automática.

## Stack da Aplicação:
    GO
    GIN
    AWS DynamoDB
    Arquitetura Hexagonal/Limpa

## Dependências diretas da Aplicação:
    GIN (github.com/aws/aws-sdk-go)
    AWS (github.com/gin-gonic/gin)
    GOOGLE UUID (github.com/google/uuid)
    TESTIFY (github.com/stretchr/testify)
    MOCKGEN (go.uber.org/mock)
    CONTEXTCORRELATIONHANDLER (github.com/hugovallada/correlationcontexthandler)