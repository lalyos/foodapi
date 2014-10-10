# Foodapi [![Build Status](https://drone.io/github.com/lalyos/foodapi/status.png)](https://drone.io/github.com/lalyos/foodapi/latest)

This is a small webapp which can be used as a `microserice`. It is an imaginary
food list REST api.

## postgres in a can

Its using postgres to store the food list. Start postgres in a Docker container:
```
docker run --name food-psql -d postgres
```

## Start the webapp

The easiest way is to start the webapp in docker. You need to link the postgres
container:

```
docker run -d --name foodweb --link food-psql:db food
```

## Configuration

You can configure the http listen port
```
export PORT=9090
```
## NoSQL

I mean if there is **no** postgre**sql** db available, an dummy in-memory repo
will be used.

## DB url manually

The webapp looks for the postgres url under the `DBURL` env variable:
```
export DBURL="postgres://postgres:@$(docker inspect -f "{{.NetworkSettings.IPAddress}}" food-psql)/postgres?sslmode=disable"
```

## psql in cli

If you want to play around with the db from cli:
```
docker run -it --rm \
  --link food-psql:db \
  postgres \
  sh -c 'psql -h $DB_PORT_5432_TCP_ADDR -U postgres postgres'
```
## Heroku

Create golang heroku app
```
heroku apps:create -b https://github.com/kr/heroku-buildpack-go.git --region eu myfoodapi
git push heroku master
```

Add postgres
```
heroku addons:add heroku-postgresql
```
