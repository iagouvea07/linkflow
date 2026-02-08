#!/bin/sh
ls migrations
migrate -path /migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:3306)/${DB_SCHEMA}" up
