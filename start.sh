#!/bin/bash
: <<USAGE
docker run -d --name foodweb --link food-psql:db food
USAGE

[ -n "$DB_PORT_5432_TCP_ADDR" ] && export DATABASE_URL=postgres://postgres:@${DB_PORT_5432_TCP_ADDR}/postgres?sslmode=disable

/foodapi
