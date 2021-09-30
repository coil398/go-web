#!/bin/bash

source "./tools/$ENV_FILE"

mysql -u${MYSQL_USER} -h${MYSQL_HOST} -p${MYSQL_PASSWORD} <<EOF
DROP DATABASE IF EXISTS ${MYSQL_DATABASE};
CREATE DATABASE ${MYSQL_DATABASE};
EOF

pwd

mysql -u${MYSQL_USER} -h${MYSQL_HOST} -p${MYSQL_PASSWORD} --database ${MYSQL_DATABASE} < ../db/mysql/migrations/init.sql

go test -tags integration ./... -v
