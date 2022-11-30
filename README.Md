# Tools
# Please install tools listed below
- Docker
- Docker Compose
- [Golang Migrate](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

# Important Command
```sh
migrate -database ${POSTGRES_URL} --path db/migrations
migrate create -ext sql -seq -dir db/migrations create_table_users
```

# TODO
- [x] Create reusable error marshalled
  
