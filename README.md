# Prerequisites
- Golang 1.23.4
- Docker 27.5.1 + Docker Compose

# Running locally
```shell
templ generate --watch --proxy="http://localhost:80" --cmd="go run ."
```