#!/bin/sh

export TZ="UTC"
export POSTGRES_PORT="11162"
export POSTGRES_DB="buildlog"
export POSTGRES_USER="buildlog"
export POSTGRES_PASSWORD="buildlog_dev_password"
export DATABASE_URL="postgres://buildlog:buildlog_dev_password@localhost:11162/buildlog?sslmode=disable"
