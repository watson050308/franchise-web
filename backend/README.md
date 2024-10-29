# Backend

## hot reload

```Command
$ go install github.com/cosmtrek/air@latest
$ air -v
  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.48.0, built with Go go1.22.2
```

## go migration

```Command
# Install the migrate CLI with PostgreSQL support
$ brew install golang-migrate  # macOS

$ go install -tags 'postgres sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

$ mkdir -p db/migrations
$ migrate create -ext sql -dir db/migrations create_user_table
$ migrate -path db/migrations -database "postgres://username:passname@localhost:5432/database_name?sslmode=disable" up
```

## API doc

How to see api doc?
<http://localhost:8080/swagger/index.html#/>
