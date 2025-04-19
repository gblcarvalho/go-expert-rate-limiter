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

Você pode testar o comportamento do Rate Limiter manualmente utilizando curl:

```bash
for i in {1..20}; do curl -i http://localhost:8080/hello; done
```

Se o número de requisições ultrapassar o limite configurado, o servidor responderá com:

HTTP/1.1 429 Too Many Requests


## 🧪 Testes automatizados com e sem API_KEY

Para facilitar, o projeto também fornece um comando para rodar testes automáticos que disparam múltiplas requisições:

```bash
make test
```

Esse comando executa dois loops:

- Requisições sem header API_KEY (baseadas apenas no IP do cliente)

- Requisições com o header API_KEY

Durante a execução, o script exibe o número da requisição, o loop correspondente, e as respostas completas (headers + corpo), permitindo visualizar facilmente quando o limite é atingido.
