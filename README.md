# Basic Golang Bootstrapper Project

### Go getter
  `$ go get github.com/minhajuddinkhan/todogo/cmd/todogo`

### Usage
```
NAME:
   TODOGO - just a dummy todo application. nothing fancy - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     todos, todogo todos    Handles Your todos
     login, todogo login    Starts session from your command line
     logout, todogo logout  Ends session from your command line
     serve, todogo serve    Starts serve from your command line
     db, todogo db          Initiates DB from your command line
     help, h                Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

### Initialize Database

Migrates and seed the database

`$ todogo db` 


### Either spin up a server, or just use the cli

###### Server

```$ todogo serve ```

###### CLI

``` 
COMMANDS:
     todos, todogo todos    Handles Your todos
     login, todogo login    Starts session from your command line
     logout, todogo logout  Ends session from your command line
     serve, todogo serve    Starts serve from your command line
     db, todogo db          Initiates DB from your command line
 ```

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

