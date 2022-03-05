#!/bin/bash

# DDLでテーブルを作成する
mysql -u root -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < "/docker-entrypoint-initdb.d/sql/create.sql"

# データを流し込む
mysql -u root -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < "/docker-entrypoint-initdb.d/sql/insert.sql"