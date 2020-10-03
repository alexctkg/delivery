# delivery

1. Copie a pasta delivery no `$GOPATH/src` 
2. Execute `$ cp .env.example .env`
3. Configure seu .env (acrescente a GIPHY_API_KEY) 
4. Execute `$ docker build -t go-docker .`
5. Execute `$ docker run -d -p 8080:8080 go-docker`

# Documentação
o acesso da documentação é feita através do pacote Swaggo (https://github.com/swaggo/swag). Pode ser acessada através da rota `docs/index.html`
como por exemplo: `http://localhost:8080/docs/index.html`

user:delivery
password:1234

Agora você pode fazer o request através do endpoint:

`http:////localhost:8080/recipes/?i={ingredient_1},{ingredient_2}`