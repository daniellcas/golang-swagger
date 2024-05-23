# Swagger Com Golang

No código acima conseguimos gerar uma página do swagger em nossa api, trazendo as informações base da aplicação e também informações sobre as rotas cadastradas na api

## Quais pacotes eu usei nesse projeto ?

> [SWAGGO](https://github.com/swaggo/swag)

> [HTTP_SWAGGER](https://github.com/swaggo/http-swagger)

## Etapas de desenvolvimento da aplicação
1. Criação do servidor GO nativo
2. Instalação do swag usando:
    ~~~go
    go install github.com/swaggo/swag/cmd/swag@latest
    ~~~
3. Instalação da dependência http_swagger:
    ~~~go
    go get -u github.com/swaggo/http-swagger
    ~~~
4. Dar o comando: `swag init -g {nome_do_path_main.go}`

    importante dar um export dos binários instalados no go, para rodar o comando init.
5. Assim que esse comando começar a rodar ele irá ler nossas especificações da documentação colocadas no projeto

    - A especificação da aplicação está no [main.go](https://github.com/daniellcas/golang-swagger/blob/main/cmd/webserver/main.go)
    - A especificação das rota estará nos [handlers](https://github.com/daniellcas/golang-swagger/blob/main/internal/infra/webserver/handlers/finance_handlers.go)
6. Passo final foi a criação da rota de acesso a doc que também está no [main.go](https://github.com/daniellcas/golang-swagger/blob/main/cmd/webserver/main.go#L27)
7. Assim quando iniciarmos nosso servidor e acessarmos o caminho definido para o documentação iremos ter essa resposta.

![Imagem](/assets/image.png)

---
### Como rodar ?
1. Foi adicionado um docker compose na aplicação para facilitar o modo de execução.
2. somente rodar um `docker compose up` na raiz do projeto ele irá rodar e executar em sua porta `3000`.

---
### Conclusão

Esse foi um guia de como cheguei ao resultado de documentar aplicações e golang.
