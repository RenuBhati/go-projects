

A RESTful API built with Go for managing a movie collection, featuring CRUD operations and director information.

## 📑 Table of Contents
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Getting Started](#-getting-started)
- [API Reference](#-api-reference)
- [Data Models](#-data-models)
- [Examples](#-examples)
- [Contributing](#-contributing)

## 🚀 Features
- Complete CRUD operations for movies
- Director information management
- JSON responses
- RESTful architecture
- Built with gorilla/mux router
- In-memory data storage

## 💻 Tech Stack
- gorilla/mux for routing
- Standard library's encoding/json
- golang.org/x/exp for random number generation


### Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/movies` | Retrieve all movies |
| GET | `/movies/{id}` | Get movie by ID |
| POST | `/movies` | Create new movie |
| PUT | `/movies/{id}` | Update movie |
| DELETE | `/movies/{id}` | Delete movie |

## 📊 Data Models

### Movie
```json
{
    "id": "string",
    "isbn": "string",
    "title": "string",
    "director": {
        "firstname": "string",
        "lastname": "string"
    }
}
```

### Director
```json
{
    "firstname": "string",
    "lastname": "string"
}
```

## 📝 Examples

### Get All Movies
```bash
curl -X GET http://localhost:8000/movies
```

### Get Movie by ID
```bash
curl -X GET http://localhost:8000/movies/1
```

### Create Movie
```bash
curl -X POST http://localhost:8000/movies \
-H "Content-Type: application/json" \
-d '{
    "isbn": "553234",
    "title": "Inception",
    "director": {
        "firstname": "Christopher",
        "lastname": "Nolan"
    }
}'
```

### Update Movie
```bash
curl -X PUT http://localhost:8000/movies/1 \
-H "Content-Type: application/json" \
-d '{
    "isbn": "553234",
    "title": "Inception 2",
    "director": {
        "firstname": "Christopher",
        "lastname": "Nolan"
    }
}'
```

### Delete Movie
```bash
curl -X DELETE http://localhost:8000/movies/1
```

## 🔍 Response Examples

### Successful GET Response
```json
{
    "id": "1",
    "isbn": "553234",
    "title": "Inception",
    "director": {
        "firstname": "Christopher",
        "lastname": "Nolan"
    }
}
```

### Successful List Response
```json
[
    {
        "id": "1",
        "isbn": "23323",
        "title": "One Movie",
        "director": {
            "firstname": "John",
            "lastname": "Doe"
        }
    },
    {
        "id": "2",
        "isbn": "4223",
        "title": "Two Movie",
        "director": {
            "firstname": "Ram",
            "lastname": "Banshali"
        }
    }
]
```

## 💡 Implementation Highlights

### Router Setup
```go
r := mux.NewRouter()
r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
r.HandleFunc("/movies", createMovie).Methods("POST")
r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
```

### Data Storage
The API uses in-memory storage with a slice of Movie structs:
```go
var movies []Movie
```


