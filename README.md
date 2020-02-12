# Teste Back-end

O teste consiste em escrever 2 microserviços que possibilitam retornar uma lista de produtos com desconto personalizado para cada usuário.

## Restrições

 1. Os serviços desse teste devem ser escritos usando linguagens distintas
 2. Os serviços desse teste devem se comunicar via [gRPC](https://grpc.io/)
 3. Utilize [docker](https://www.docker.com/) para provisionar os serviços
 4. Para facilitar, os serviços podem usar um banco de dados compartilhado

# Executando

## Makefile

Utilize o seguinte comando do Makefile contido na raiz do projeto para construir as imagens dos serviços
```
    make build
```

E então, para criar os containers:
```
    make up
```

Após isso o serviço de produtos estará rodando na porta `:8080` e o de discontos na porta `:50052`.
O endpoint para recuperar os produtos com descontos será `/v1/products`
