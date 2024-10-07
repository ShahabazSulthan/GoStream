# GoStream

GoStream is a microservice-based video streaming platform designed to provide efficient video management and delivery. Built with Go, Gin, PostgreSQL, and gRPC, it leverages an API Gateway to handle routing and authentication, ensuring smooth communication between services.

## Features

- **Video Streaming**: Supports segmented video streaming using HLS (HTTP Live Streaming).
- **Microservice Architecture**: Includes two primary services:
  - **API Gateway**: Manages requests, routes them to the appropriate service, and handles authentication.
  - **Video Service**: Manages video storage, processing, and delivery.
- **gRPC Communication**: Efficient inter-service communication using gRPC.
- **PostgreSQL Database**: Stores video metadata and user-related information.
- **REST API**: Exposes video streaming and management functionalities via a RESTful interface.
  
## Tech Stack

- **Go**: The primary language used for building the services.
- **Gin**: A web framework for handling HTTP requests in the API Gateway.
- **gRPC**: Used for inter-service communication between the API Gateway and the Video Service.
- **PostgreSQL**: The database for storing video metadata and user information.
- **HLS (HTTP Live Streaming)**: For efficient video streaming in .ts format.
  
## Project Structure

```
GoStream
│
├── api_gateway
│   ├── cmd
│   │   └── main.go                  # Main entry point for the API Gateway
│   ├── pkg
│   │   ├── api
│   │   │   ├── handler
│   │   │   │   └── handler.go        # API route handlers
│   │   │   └── server.go             # Server initialization
│   │   ├── client
│   │   │   ├── interfaces
│   │   │   │   └── client.go         # Interface for gRPC client
│   │   │   └── client.go             # gRPC client implementation
│   │   ├── config
│   │   │   └── config.go             # Configuration for the API Gateway
│   │   ├── di
│   │   │   └── wire.go               # Dependency injection setup
│   │   ├── pb
│   │   │   ├── video.pb.go           # Generated protobuf files for gRPC
│   │   │   └── video_grpc.pb.go      # gRPC service implementations
│   │   ├── routes
│   │   │   └── routes.go             # API routes definition
│   ├── static
│   │   └── index.js                  # Frontend assets (if any)
│   ├── template
│   │   ├── index.html                # Frontend template
│   │   └── upload.html               # Video upload template
│   ├── .env                          # Environment variables
│   ├── go.mod                        # Go module definitions
│   ├── go.sum                        # Go dependencies checksum
│
├── video_service
│   ├── cmd
│   │   └── main.go                   # Main entry point for the Video Service
│   ├── pkg
│   │   ├── api
│   │   │   ├── service
│   │   │   │   └── video.go          # Business logic for video handling
│   │   │   └── server.go             # gRPC server setup for video service
│   │   ├── config
│   │   │   └── config.go             # Configuration for the Video Service
│   │   ├── db
│   │   │   └── db.go                 # Database connection setup
│   │   ├── di
│   │   │   └── wire.go               # Dependency injection setup
│   │   ├── domain
│   │   │   └── video.go              # Domain model for video
│   │   ├── pb
│   │   │   ├── video.pb.go           # Generated protobuf files for gRPC
│   │   │   └── video_grpc.pb.go      # gRPC service implementations
│   │   ├── repository
│   │   │   └── video.go              # Video repository implementation
│   ├── storage
│   │   ├── [video_id]
│   │   │   ├── video.mp4             # Uploaded videos
│   │   │   ├── playlist.m3u8         # HLS playlist files
│   │   │   ├── playlistX.ts          # Video segments
│   ├── .env                          # Environment variables
│   ├── go.mod                        # Go module definitions
│   ├── go.sum                        # Go dependencies checksum
│
└── .gitignore                        # Git ignore file for sensitive data and binaries
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gostream.git
   cd GoStream
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Setup environment variables:
   - Copy the `.env.example` file and rename it to `.env`.
   - Update the necessary configurations like PostgreSQL connection, API Gateway settings, etc.

4. Run the services:
   - Start the API Gateway:
     ```bash
     cd api_gateway/cmd
     go run main.go
     ```
   - Start the Video Service:
     ```bash
     cd video_service/cmd
     go run main.go
     ```

## Usage

- **Upload Videos**: Access the `/upload` endpoint through the API Gateway to upload videos.
- **Stream Videos**: Stream the uploaded videos via HLS by accessing the `/videos/{id}` endpoint.
- **gRPC Communication**: The Video Service communicates with the API Gateway using gRPC for faster and efficient inter-service calls.

## Contributing

Feel free to submit issues, fork the repository, and send pull requests. Please make sure to update tests as appropriate.
