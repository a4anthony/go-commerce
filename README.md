# go-commerce

A Golang e-commerce backend using Fiber, MySQL, Makefile and Swagger

### Installing

1. Install extra packages:

     ```go install github.com/cosmtrek/air@latest```

    ```go install github.com/swaggo/swag/cmd/swag@latest```
2. Clone the repo
3. Create your own .env file
4. ```make dev```
5. view docs at http://localhost:8080/swagger
6. Run migration ```make db-migrate```


### Scripts

- ```make dev``` - runs the server in development mode

- ```make swagger``` - generates the swagger docs

