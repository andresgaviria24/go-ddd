*IMPORTANT*
This basic proyect is for show ddd (https://www.paradigmadigital.com/dev/ddd-dominio-implica-crecer-fuerte/) with golang, remember, ddd can be implemented in many ways depending on the author!

# Web Service Restaurant
WS for get food and orders.

# Start 🚀
    1. Clone this project -> https://github.com/andresgaviria2020/go-ddd.git
    2. Make sure port 8080 is not busy.
    3. go get all 
    4. go run main.go

# Pre-requirements 📋
It is necessary to install -> https://golang.org/ 

# Dependencies 🤝
**IMPORTAT:** All dependencies will register and will be usable all dependency, however all dependency that is not used and is imported they should delete from de mod file

- [github.com/gin-gonic/gin] Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [github.com/gin-contrib/pprof] Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.
- [gorm.io/gorm] This is an orm framework and is used to be connected to sql server database
- [github.com/swaggo/gin-swagger] [github.com/swaggo/files] This is a open api for documentation a some tests
- 

# Project Structure 🧱

```
go-ddd
├─ .DS_Store
├─ BD_Restaurant.sql
├─ README.md
└─ WS_Restaurant
   ├─ .DS_Store
   ├─ .env
   ├─ application
   │  └─ RestaurantController.go
   ├─ docs
   │  ├─ docs.go
   │  ├─ swagger.json
   │  └─ swagger.yaml
   ├─ domain
   │  ├─ dto
   │  │  ├─ FoodDto.go
   │  │  ├─ Response.go
   │  │  └─ UserDto.go
   │  ├─ entity
   │  │  ├─ Food.go
   │  │  ├─ Order.go
   │  │  └─ Users.go
   │  └─ service
   │     ├─ RestaurantService.go
   │     └─ RestaurantServiceImpl.go
   ├─ go.mod
   ├─ go.sum
   ├─ infrastructure
   │  ├─ persistence
   │  │  ├─ DbHelper.go
   │  │  └─ FoodRepositoryImpl.go
   │  └─ repository
   │     └─ FoodRepository.go
   ├─ interfaces
   │  └─ middleware
   │     ├─ CORSMiddleware.go
   │     └─ server
   │        ├─ Server.go
   │        └─ ServerImpl.go
   └─ main.go

```

# Built with 🛠️
    - Visual Studio Code
    - Goland

# Endpoints
    - GET /food

# Authors
Andrés Gaviria
Semi Senior III Developer
andres_felipe_gaviria28@hotmail.com