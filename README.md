A RESTful microservice implementing Tic-Tac-Toe logic with a Minimax-based AI opponent. Designed as a reference for Clean Architecture and Dependency Injection in Go.

## Features

- **Architecture:** Clean Architecture with strict layer separation.
- **DI:** `uber/fx` for dependency graph management and graceful shutdown.
- **Networking:** Standard `net/http` using Go 1.22+ routing.
- **Concurrency:** Thread-safe state management via `sync.Map`.
- **AI:** Minimax algorithm.

## Run

```bash
make run
```

Server will be running on port `8080`.

## API Reference

**Endpoint:** `POST /game/{uuid}`

Processes a player's move and returns the updated board state after the bot's turn.

**Request:**

* Content-Type: `application/json`
* `0`: Empty
* `1`: Player (X)
* `2`: Bot (O)

```json
{
  "board": [
    [0, 0, 0],
    [0, 0, 0],
    [0, 0, 0]
  ]
}

```

**Response:**

* `winner`: Included only if the game has ended.

```json
{
  "uuid": "session-id",
  "board": [
    [1, 0, 0],
    [0, 2, 0],
    [0, 0, 0]
  ],
  "winner": 2
}

```
