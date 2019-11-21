# Hash Teste Back-end
Implementação do [desafio de back-end](https://github.com/hashlab/hiring/blob/master/challenges/pt-br/back-challenge.md) da [Hash](https://www.hash.com.br).

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


# Stack

## Go
Para implementar o serviço de produtos eu escolhi Go.
Como as chamadas gRPC para calcular os discontos devem ser realizadas individualmente para cada produto, percebi aqui uma boa oportunidade para utilizar concorrência. Também é a linguagem com a qual eu possuo maior afinidade no momento.

## Javascript (Node)
Para o serviço de descontos eu utilizei Node.js.
Por se tratar de um serviço um tanto simples e pequeno, descartei a possibilidade de usar Typescript e parti para algo mais básico. Acredito que essa não foi uma escolha muito boa, por não poder gerar o código a partir do `.proto` antes do build da aplicação, o que torna o código um pouco frágil na minha opinião. De qualquer forma, ficou o aprendizado.