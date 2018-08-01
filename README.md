# Basic Golang Bootstrapper Project

### Go getter
  `go get github.com/minhajuddinkhan/todogo/cmd/todogo`

### Usage
```


### Initialize Database
  `todogo db`
  
### Credentials to get todos and stuff (that are seeded lol)
```go
username: Rameez
password 123
```
Heres the sample environment variables you need to set. 
By default it takes `sqlite3` as dialect and `/tmp/todo.db` as path to dump data.
### env
```env
JWTSECRET=ILOVEPAKISTAN    
DB_HOST=localhost
DB_PORT=5432
DB_NAME=tododb
DB_USER=admin
DB_PASSWORD=admin
SVR_PORT=3000
DB_VOULME_PATH=/tmp/todo.db
DB_DIALECT=sqlite```

