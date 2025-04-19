# Go Expert Rate Limiter

Este projeto implementa o desafio do Rate Limiter da pós Go Expert.

O Rate Limiter funciona como um middleware e limita o número de requisições que podem serem feitas em uma determinada janela de tempo. A verificação pode ser feita por IP ou token API_KEY.
O Rate limiter recebe as informações de limite máximo de requests por IP, limite máximo de requests por token e a janela de tempo da verificação em milisegundos.

No servidor de exemplo contido desse projeto, as configurações do Rate Limiter são lidas do arquivo .env através das seguintes variáveis de ambiente:

RATE_LIMITE_IP_MAX_REQUESTS
RATE_LIMITE_TOKEN_MAX_REQUESTS
RATE_LIMITE_TIME_WINDOW

No projeto também foi desenvolvido dois Stores para salvar as informações das requisições, uma simples em Go que guarda as informações em memória e outra para utilizar o Redis

Para executar o projeto é necessário criar o arquivo .env na raiz do projeto, esse arquivo pode ter o conteúdo duplicado de .env.dev

A execução do projeto via docker pode ser feita através do comando:

make start

esse comando irá subir uma instancia do redis e uma do app, o app roda um servidor de exemplo da porta 8080 e rota /hello

# Go Expert Rate Limiter

Este projeto implementa o desafio de Rate Limiting da pós-graduação **Go Expert**.

---

## 📌 Descrição

O Rate Limiter funciona como um **middleware** que limita o número de requisições permitidas em uma janela de tempo configurável.  
A verificação pode ser feita com base no **IP do cliente** ou no **token fornecido via header `API_KEY`**.

Você pode configurar:

- **Máximo de requisições por IP**
- **Máximo de requisições por token**
- **Tamanho da janela de tempo** (em milissegundos)

---

## ⚙️ Configuração

As configurações do Rate Limiter são lidas a partir de variáveis definidas no arquivo `.env`, utilizando:

```env
RATE_LIMIT_IP_MAX_REQUESTS=10
RATE_LIMIT_TOKEN_MAX_REQUESTS=20
RATE_LIMIT_TIME_WINDOW=1000
```
Obs.: Você pode criar o .env manualmente ou copiá-lo de um exemplo:

```bash
cp .env.dev .env
```

## 🗃️ Stores disponíveis

O projeto possui duas implementações de Store para armazenamento e controle de requisições:

- **MemoryStore**: implementado em Go, armazena os dados em memória.

- **RedisStore**: utiliza o Redis como backend, ideal para ambientes distribuídos.

## 🚀 Execução via Docker

O projeto já está configurado com Docker Compose. Para subir o ambiente completo (app + Redis), execute:

```bash
make start
```

Esse comando irá:

- Construir a imagem do app

- Subir o Redis

- Iniciar o servidor Go na porta 8080

## 🔍 Testando o servidor

Após iniciar, o servidor estará disponível em:

GET http://localhost:8080/hello

Você pode testar o Rate Limiter com um loop no terminal, por exemplo:

```bash
for i in {1..20}; do curl -i http://localhost:8080/hello; done
```

Se o limite de requisições for excedido, o servidor irá responder com:

HTTP/1.1 429 Too Many Requests

Você também pode executar o seguinte comando para executar testes automáticos onde será enviado requisições com o header API_KEY e requisições sem o header

```bash
make test
```
