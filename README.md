# auth

This package is for the microservices dealing with registration and authentication, as well as basic user functions.

## How to run this repo locally:
```
git clone https://github.com/mailwilliams/auth.git
cd auth
go mod download
docker-compose up --build
```

##  Project Structure
```
main.go               //  where the initial values are being set and the http listener lives

Dockerfile            //  the build logic for the "backend" container found in docker-compose.yaml
docker-compose.yaml   //  the configuration file to build the MySQL database, Redis Cache, and http service

src/    
  commands/           //  specific files ran separately from the main http service
    populateUsers.go    //  i.e. populate database with fake users
  
  database/           //  files specifically used for database functionality
    db.go               //  MySQL
    cache.go            //  Redis
    migrate.go          //  migrate structs from models package to MySQL database
    
  handlers/           //  files for configuring and interacting with endpoints
    handler.go          //  parent handler, establishes context has generic helper methods
    routes.go           //  handler method to configure routes and their associated handler method
    middleware.go       //  handler method to parse authenticated users, used in routes.go
    
    //  specific handler methods
    register.go         //  POST    /api/register
    login.go            //  POST    /api/login
    logout.go           //  DELETE  /api/logout
    listUsers.go        //  GET     /api/users
    updateInfo.go       //  PUT     /api/users/me
    
  models/             //  files for structs that will be used in the handler methods as well as database migrations 
    user.go             //  Go struct for users, GORM model (https://gorm.io/docs/models.html) 
```

##  How to add a new endpoint:
  1.  Add new file in `src/handlers`
    - Look at other files to see the general pattern, they are all set up the same
  2.  Add the route in `src/handlers/routes.go`
