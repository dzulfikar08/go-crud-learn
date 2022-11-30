# Golang JSON CRUD API

Golang project that features JSON CRUD API from PostgreSQL database using Gin and GORM framework.

## Initialize
To run the project, run following command:
1. `go mod tidy`
2. `go run main`

The web server will run in localhost port 3000.

## Dependencies
As you run `go mod tidy` , dependencies that installed are :
- Gin Web Framework (github.com/gin-gonic/gin) 
- ORM Library (gorm.io/gorm)
- GoDotEnv (https://github.com/joho/godotenv)

## Features
1. JSON CRUD API in Go.
2. Endpoints of GET, POST, UPDATE, DELETE

## Endpoints
- **/posts**

| Method | Header | Params | JSON                                                      | Remark |
| ------ | ------ | ------ | --------------------------------------------------------- | ------- | 
| `POST` | `none` | `none` | Nbr: `int`<br>Name: `string` | Create Post |
| `GET` | `none` | `none` | `none` | Fetch Posts |

- **/posts/:id**

| Method | Header | Params | JSON                                    | Remark |
| ------ | ------ | ------ | --------------------------------------- | ------ |
| `GET` | `none` | `none` | `none` | Fetch Post |
| `PUT` | `none` | `none` | Nbr: `int` | Update Post |
| `DELETE` | `none` | `none` | Nbr: `int` | Delete Post |

