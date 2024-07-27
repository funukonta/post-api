# Post API

## How to run the app?
Clone the Repo First!
``` bash
$ git clone https://github.com/funukonta/post-api
$ cd post-api/
```

Then run the app ! 
### using docker compose
```bash
$ docker compose up -d
```

### using go run
1. run the postgres DB (local or using docker)
```bash
$ docker run --name post-posgres -p 5432:5432 -e POSTGRES_PASSWORD=password123 -d postgres
```

2. uncomment ```DB_HOST``` inside ```.env``` file

3. then casually run 
``` bash 
$ go run ./cmd/main.go
```
or if you prefer to build it first
``` bash
$ go build -o ./bin/post-api ./cmd/main.go && ./bin/post-api
```