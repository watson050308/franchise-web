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

## Start Go server

### Step 1: Create the Directory for Docker Data

First, create a `docker_data` folder in the root directory. When you start docker-compose.data, the data will be saved locally, even if you delete docker-compose.data.

```command
mkdir ~/docker_date
```

### Step 2: Create a Docker Network

Create a custom Docker network that your containers can use to communicate with each other.

```command
docker network create franchise_web_network
```

### Step 3: Create Docker Compose Files

You’ll need two Docker Compose files to manage your application and database services. Now, create two Docker Compose files: `docker-compose.data` and `docker-compose.app`.

```command
docker-compose -f docker-compose.data.yaml -p franchise_database up -d
docker-compose -f docker-compose.app.yaml -p franchise_web up -d
```

### Step 4: Contecct to Dasebase

If you use Docker to connect to the web app, please connect to the production database.
pdAdmin: <http://localhost:8000/browser/>
Name:
![Name](docs/name.png)
Host Name
![Host name](docs/host_name.png)
