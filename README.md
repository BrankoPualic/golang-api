## Project structure
```
tiago_tutorials_api/
┣ bin/
┃ ┗ main.exe
┣ cmd/
┃ ┣ api/
┃ ┃ ┣ api.go
┃ ┃ ┣ health.go
┃ ┃ ┗ main.go
┃ ┗ migrate/
┃   ┗ migrations/
┃ ┃   ┣ 000001_create_users.down.sql
┃ ┃   ┣ 000001_create_users.up.sql
┣ docs/
┣ internal/
┃ ┣ database/
┃ ┃ ┗ database.go
┃ ┣ env/
┃ ┃ ┗ env.go
┃ ┣ model/
┃ ┃ ┣ post.go
┃ ┃ ┗ user.go
┃ ┗ store/
┃   ┣ posts.go
┃   ┣ storage.go
┃   ┗ users.go
┣ scripts/
┃ ┗ db_init.sql
┣ web/
┣ go.mod
┗ go.sum
```

`bin/`      Contains compiled executable.

`cmd/`      Holds the entry points for your applications.

`docs/`     Documentation related to the project. API specifications, architecture notes, diagrams, ADRs, etc.

`web/`      Static frontend files served by the API (if any)

`scripts/`  Utility scripts

`internal/` Private application code that **should not be imported by external modules.**
