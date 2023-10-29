# todolist-api
# How to Setup
- clone this repo
- edit db connection on config.json
  "db-postgres": {
        "username":"username",
        "password":"password",
        "host":"127.0.0.1",
        "port":5432,
        "database":"yourdbname",
        "max_open_conns": 0,
        "max_idle_conns": 10,
        "conn_max_lifetime": 0
    }
  
**Note: If doesn't have a postgres please install postgres DB on your local computer**
-  please install golang-migrate on your local if you don't have and please refer to this [link](https://pkg.go.dev/github.com/golang-migrate/migrate/cli#section-readme)
-  after install run make postgres-migrate-up
**Note: if the migration not working properly run manual script on resources to create tables**
- run go mod tidy
- run go run cmd/main.go

to testing a file you can import the postman api collection on resources
  
