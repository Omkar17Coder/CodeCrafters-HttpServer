# Go HTTP Server

A simple, modular HTTP/1.1 server written in Go.  
This project demonstrates how to build a basic HTTP server, route requests using a trie-based router, and organize code for clarity and maintainability.

---

## Features

- Handles multiple routes (root, echo, user-agent, file serving)
- Trie-based URL router for efficient path matching
- Modular handler structure for easy extension
- Clean, idiomatic Go code

---

## Supported Routes

The server currently supports the following HTTP endpoints:

- `GET /` - Root endpoint that returns a simple response
- `GET /echo/<text>` - Returns the provided text in the response
- `GET /user-agent` - Returns the User-Agent header from the request
- `GET /files/<filename>` - Serves files from the specified directory
- `POST /files/<filename>` - Accepts file uploads to the specified directory 

Each route is handled by its own dedicated handler package, following clean architecture principles.
Additonal task to work on is  using tmp package of golang to create temprorary files before content actually  updated to desired files since multiple  work on same file and lead to race conndtion that needs to be handeled

---

## Folder Structure

```
codecrafters-http-server-go/
│
├── app/
│   ├── http/                  # Core HTTP server logic and router
│   │   ├── handlers/          # Each handler in its own subfolder
│   │   │   ├── echo/
│   │   │   │   └── echo_handler.go
│   │   │   ├── files/
│   │   │   │   └── files_handler.go

│   │   │   ├── root/
│   │   │   │   └── root_handler.go
│   │   │   └── user_agent/
│   │   │       └── user_agent_handler.go
│   │   ├── router.go          # Trie-based router implementation
│   │   └── server.go          # Server entry point
│   ├── pkg/
│   │   └── config/            # Constants and configuration
│   ├── types/                 # Shared types (Request, Response, Context)
│   └── utils/                 # Utility functions (parsing, response writing)
│
└── ...
```

---

## Trie-based URL Router

The router uses a trie (prefix tree) to efficiently match incoming request paths to registered handlers.  
- Each node in the trie represents a segment of the URL path.
- Handlers are registered by splitting the path and inserting each segment into the trie.
- When a request comes in, the path is split and traversed through the trie to find the best matching handler.
- This allows for fast lookups and supports wildcard routes (e.g., `/echo/*`).

**Example:**
- Registering `/echo/*` creates a branch in the trie for `/echo` and a wildcard child.
- A request to `/echo/hello` will match the `/echo/*` handler.

---

## How to Run

1. **Build:**
   ```sh
    ./your_program.sh --directory /tmp
   ```

---

## Adding New Routes

1. Create a new handler in `app/http/handlers/<your_route>/`.
2. Register it in `server.go`:
   ```go
   server.router.Register("GET", "/your-route", your_route.Handle)
   ```

---

## License

MIT

---

## Acknowledgments

A special thanks to [CodeCrafters](https://app.codecrafters.io/catalog) for their amazing programming challenges. This project started as their "Build Your Own HTTP Server" challenge, which provided a solid foundation in HTTP server implementation.

Beyond the basic requirements, this implementation includes several advanced features:

- **Trie-based URL Router**: A custom implementation for efficient path matching and routing
- **Modular Architecture**: Clean separation of concerns with a well-organized folder structure
- **Handler Organization**: Each route handler in its own package for better maintainability
- **Production-grade Features**: Implementation of best practices and clean architecture principles

The challenge structure provided an excellent starting point, while the advanced features demonstrate how to take a basic implementation to a production-ready level.
