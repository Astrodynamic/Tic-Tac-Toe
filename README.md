# Tic-Tac-Toe Service

A RESTful microservice implementation of the Tic-Tac-Toe game logic with an unbeatable AI opponent. This project serves as a reference for building scalable, thread-safe web services in Go using Clean Architecture principles.

## Technical Implementation

The service demonstrates a production-ready structure without heavy external frameworks:

- **Architecture:** Standard Go Project Layout with strict layer separation.
- **Dependency Injection:** Managed dependency graph and graceful shutdown lifecycle using `uber/fx`.
- **Networking:** Pure `net/http` implementation utilizing Go 1.22+ routing capabilities.
- **Concurrency:** Thread-safe game state management supporting multiple concurrent sessions via `sync.Map`.
- **Algorithm:** Minimax algorithm with alpha-beta pruning for the game AI.

## Getting Started

### Prerequisites
- Go 1.22 or higher
- Make (optional)

### Execution
Run the service locally on port 8080:
```bash
make run
```

## API Reference

**Endpoint:** `POST /game/{uuid}`

Processes a player's move and returns the updated board state after the bot's turn.

**Request:**

* Content-Type: `application/json`
* `0`: Empty, `1`: Player (X), `2`: Bot (O)

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

* `winner`: Included only if the game has ended (`1`, `2`, or `0` for draw).

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
