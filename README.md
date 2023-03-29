# ZepterUserService

<br>
The service provides gRPC api and can be tested through postman.
<br><br>
The service is written in Go and uses Postgres as a database.
<br> Db config can be seen in docker-compose.yml file if you want to run it via docker (listens on port 4002).
<br> If you want to run it locally, you need to have postgres installed and running on port 5432.
<br><br>

## Starting manually
```
go run .
```

## Starting with docker
```
docker compose up --build
```

<br>
DeleteUser should change some flag only to false, not removing from db permanently.
<br><br><br>

## Deployment
If we want to deploy our application on some cloud provider (for example AWS) we will push docker image to dockerhub and then copy only docker-compose.yaml on AWS instance and start application via docker.

