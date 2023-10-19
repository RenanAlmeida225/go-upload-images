# Go upload images

| Methods | Urls           | Actions                                               |
| :------ | :------------- | :---------------------------------------------------- |
| POST    | /images        | upload a image                                        |
| GET     | /images        | get a list of images (id, title, description and url) |
| GET     | /images/{data} | search image by title or description                  |
| GET     | /images/:id    | get a image by id                                     |
| PUT     | /images/:id    | update a image by id (title and description)          |
| DELETE  | /images/:id    | delete a image by id                                  |

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/username/repo-name.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`
