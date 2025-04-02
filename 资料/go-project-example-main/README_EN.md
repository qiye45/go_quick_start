# Go Project Example

## Project Introduction

This is a community forum system example project developed in Go, using the Gin framework as the web server, GORM as the ORM framework, and MySQL as the database storage. The project demonstrates a standard Go project structure and best practices, including API handling, business logic, data storage, testing, and concurrent programming.

## Features

- User System: User information management
- Topic System: Topic creation and querying
- Post System: Post publishing and replies
- RESTful API: Standard REST-style interfaces
- Concurrent Programming Examples: Includes goroutine and channel examples
- Go Programming Notes: Examples of string handling, JSON operations, closures, and array processing

## Tech Stack

- [Go](https://golang.org/) - Programming Language
- [Gin](https://github.com/gin-gonic/gin) - Web Framework
- [GORM](https://gorm.io/) - ORM Framework
- [MySQL](https://www.mysql.com/) - Database
- [Zap](https://github.com/uber-go/zap) - Logging Library

## Project Structure


```
.
├── attention/            # Example code for Go programming notes
│   ├── array.go          # Array operation examples
│   ├── closure.go        # Closure examples
│   ├── json.go           # JSON processing examples
│   └── string.go         # String handling examples
├── concurrence/          # Concurrent programming examples
│   ├── channel.go        # Channel examples
│   └── goroutine.go      # Goroutine examples
├── handler/              # HTTP request handling layer
│   ├── publish_post.go   # Post publishing handler
│   └── query_page_info.go # Page information query handler
├── repository/           # Data persistence layer
│   ├── db_init.go        # Database initialization
│   ├── post.go           # Post data operations
│   ├── topic.go          # Topic data operations
│   └── user.go           # User data operations
├── service/              # Business logic layer
│   ├── publish_post.go       # Post publishing business logic
│   ├── publish_post_test.go  # Post publishing tests
│   ├── query_page_info.go    # Page information query business logic
│   └── query_page_info_test.go # Page information query tests
├── util/                 # Utility classes
│   └── logger.go         # Logging utility
├── .gitignore            # Git ignore file
├── example.sql           # Database example SQL
├── go.mod                # Go module definition
├── go.sum                # Go dependency checksum
├── LICENSE               # License file
├── README.md             # Project documentation (Chinese)
├── README_EN.md          # Project documentation (English)
└── sever.go              # Main service entry
```

## Installation and Usage

### Prerequisites

- Go 1.16+
- MySQL 5.7+

### Database Setup

1. Create MySQL database and table structure:

```bash
mysql -u root -p < example.sql
```

### Project Configuration

1. Modify database connection configuration as needed (located in `repository/db_init.go`)

```go
dsn := "root:00000000@tcp(127.0.0.1:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
```

### Running the Project

1. Clone the repository

```bash
git clone https://github.com/Moonlight-Zhao/go-project-example.git
cd go-project-example
```

2. Install dependencies

```bash
go mod download
```

3. Build and run

```bash
go build -o app
./app
```

Or run directly:

```bash
go run sever.go
```

4. Access APIs

```
GET http://localhost:8080/ping               # Health check
GET http://localhost:8080/community/page/get/:id  # Get topic page information by ID
POST http://localhost:8080/community/post/do     # Publish a post
```

## API Documentation

### Get Topic Page Information

```
GET /community/page/get/:id
```

Parameters:
- `id`: Topic ID

Response example:
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "topic": { ... },
    "post_list": [ ... ]
  }
}
```

### Publish Post

```
POST /community/post/do
```

Form parameters:

- `uid`: User ID
- `topic_id`: Topic ID
- `content`: Post content

Response example:
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "post_id": 123
  }
}
```

## Learning Resources

The project includes common Go programming notes and concurrent programming examples that can be used as learning references:

- `attention/`: Contains examples and tests for string handling, JSON operations, closures, and array processing
- `concurrence/`: Contains example code and tests for goroutines and channels

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 
