
## postgres in a can

start postgres in a Docker container:
```
docker run --name gofood-psql -d postgres
```


## DB url

```
export DBURL="postgres://postgres:@$(docker inspect -f "{{.NetworkSettings.IPAddress}}" gofood-psql)/postgres?sslmode=disable"
```

## psql in cli
```
psql -h $(docker inspect -f "{{.NetworkSettings.IPAddress}}" gofood-psql) -U postgres
```
