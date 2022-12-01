# Tools
# Please install tools listed below
- Docker
- Docker Compose
- [Golang Migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

# Important Command
```sh
migrate --database ${POSTGRES_URL} --path db/migrations
migrate create -ext sql  -path db/migrations -seq create_table_users
```

# TODO
- [x] Create reusable error marshalled

# Note
- different between view and table, in simple way table is empty object that consists from row and column, you should fill table manually, while view pretty similiar to table but view generate the content from other query not manually insert data into it.
  
