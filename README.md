# Go upload images

## EndPoints

| Methods | Urls                    | Actions              |
| :------ | :---------------------- | :------------------- |
| POST    | /images                 | upload a image       |
| GET     | /images?page=1&limit=10 | get a list of images |
| GET     | /images/:id             | get a image by id    |
| PUT     | /images/:id             | update a image by id |
| DELETE  | /images/:id             | delete a image by id |

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/RenanAlmeida225/go-upload-images.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [Postgres](https://github.com/go-gorm/postgres) as the database
- [SDK aws](https://github.com/aws/aws-sdk-go-v2)for conection with aws s3
- [Godotenv](https://github.com/joho/godotenv) for environment variables
