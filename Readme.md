# Netcat Project by Go
# TCP-Chat (net-cat)

A simple TCP chat server written in Go. Multiple clients can connect (up to 10), choose a unique username, and exchange messages in real time. Messages are prefixed with timestamp and username, similar to `netcat`-style chats.

## Features
- Up to 10 concurrent clients
- Username prompt with uniqueness and basic validation
- Message validation (printable ASCII only)
- Join/leave notifications
- In-memory message history broadcasted to newly joined users
- Timestamped message format: `[YYYY-MM-DD HH:MM:SS][username]:`

## Project structure
- `net-cat/main.go` — program entry point (parses port arg and starts server)
- `net-cat/funces/TCPfuntion.go` — TCP listener and connection accept loop (`Connection()`)
- `net-cat/funces/handleconx.go` — per-connection handler, broadcasting, names & messages storage
- `net-cat/funces/MessageForma.go` — message prefix formatting
- `net-cat/funces/valid.go` — validation helpers for names and messages

## Requirements
- Go 1.20+ (any recent Go version should work)

## Build & Run
From the repository root:

- Run directly:
  ```bash
  go run ./net-cat [PORT]
  ```
  Examples:
  - Default port 8989:
    ```bash
    go run .
    ```
  - Custom port 3000:
    ```bash
    go run . 3000
    ```

- Build binary:
  ```bash
  go build -o TCPChat 
  ```
  Then run:
  ```bash
  ./TCPChat             # listens on :8989
  ./TCPChat 3000        # listens on :3000
  ```

Notes:
- If you prefer running inside the module directory, you can also execute:
  ```bash
  (cd net-cat/net-cat && go run . 3000)
  ```

## Usage (client)
Connect using `nc` (netcat) or `telnet` from another terminal:

```bash
nc 127.0.0.1 8989
# or
nc 127.0.0.1 3000
```

You’ll see a welcome banner followed by a prompt:
```
[ENTER YOUR NAME]:
```
Choose a username (max 24 chars, printable ASCII, must be unique). On join, other users see:
```
<name> has joined our chat...
```
Type messages and press Enter to send. Other users receive your message like:
```
[2025-08-22 17:25:00][alice]: Hello!
```
Leaving the session (e.g., closing the terminal) notifies others:
```
<name> has left our chat...
```

## Constraints & behavior
- Maximum 10 connected clients. The 11th connection receives: `The group is full`.
- Name validation:
  - Non-empty, ≤ 24 characters
  - Printable ASCII characters only (32–126)
  - Must be unique among connected users
- Message validation: printable ASCII only (32–126). Empty lines are ignored for broadcast.
- New joiners receive the in-memory history of earlier messages for context.

## Troubleshooting
- Port already in use: choose a different port (e.g., `3000`).
- Firewall rules may block inbound connections; allow the chosen port locally if needed.
- If `go run ./net-cat` fails, ensure you are at the repository root and have a recent Go toolchain installed.


