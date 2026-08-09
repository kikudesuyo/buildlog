#!/bin/sh

export TZ="UTC"
export POSTGRES_PORT="11162"
export POSTGRES_DB="buildlog"
export POSTGRES_USER="buildlog"
export POSTGRES_PASSWORD="buildlog_dev_password"
export DATABASE_URL="postgres://buildlog:buildlog_dev_password@127.0.0.1:11162/buildlog?sslmode=disable"
export ADMIN_PASSWORD="replace-with-a-local-password"
export ADMIN_SESSION_SECRET="replace-with-a-long-random-secret"
