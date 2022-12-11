# account-producer-service

## About
Service responsible for account management.

## Technologies
* Golang 1.18

## Development requirements
* Docker Compose
* Visual Studio Code
* DBeaver
* Driver Scylla (https://downloads.datastax.com/jdbc/cql/2.0.11.1012/SimbaCassandraJDBC42-2.0.11.1012.zip)

## Directory Structure
- `build`
    - It has all cloud package, container (Docker), operating system (deb, rpm, pkg) and scripts settings.
- `cmd`
    - It has the `main` function that imports and invokes code from the `/internal` and `/pkg` directories.
- `internal`
    - It has all the code that is not available for import.
- `local-dev`
    - Possui toda configuração do docker.

## Running
- `Docker`
    - Run the following command: docker-compose -f local-dev/docker-compose.yaml --profile infra up -d
- `Config DB`
    - JDBC url: jdbc:cassandra://localhost:9042;AuthMech=1;UID=cassandra;PWD=cassandra
    - Host: localhost
    - Port: 9042
    - Username: cassandra
    - Password: cassandra
    - Run the commands available at: account-consumer-service/build/package/docker/scylla/cql/V001_setup.cql
- `Run the project`
    - Run -> start debugging -> to allow

## Stop running
- `Stop docker`
    - docker-compose -f local-dev/docker-compose.yaml --profile infra down