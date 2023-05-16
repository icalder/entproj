# entproj

Create a test database in Postgres, to be used when running ent tests.
Assume there is a way tests can create whole schema then roll everything back.

Using pgadmin4:

Create login role (ensure 'can login' toggle is selected or you will get a group role)

Create database 'enttest', set owner to the enttest user

Connection parameters:
 host: opti
 port: 30529
 user: enttest
 db: enttest

https://entgo.io/docs/tutorial-setup

```
go mod init github.com/icalder/entproj
go get -d entgo.io/ent/cmd/ent
go run -mod=mod entgo.io/ent/cmd/ent new Registry

go generate ./...
``` 
