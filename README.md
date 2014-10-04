
## postgres in a can

start postgres in a Docker container:
```
docker run --name gofood-psql -d postgres 
```

## psql in cli
```
psql -h $(docker inspect -f "{{.NetworkSettings.IPAddress}}" gofood-psql) -U postgres
```
