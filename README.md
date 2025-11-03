# Simple HTTP API

Simple Go API that checks if name starts with A-M or N-Z.

## Run

```bash
go run main.go
```

Server runs on `http://localhost:8080`

## Test

```bash
go test -v
```

## Endpoint

GET `/hello-world?name=<name>`

- A-M: returns 200 with message
- N-Z or empty: returns 400 with error

Examples:
- `GET /hello-world?name=Alice` → `{"message":"Hello Alice"}`
- `GET /hello-world?name=Zane` → `{"error":"Invalid Input"}`

## Assumptions

1. Only the first letter is checked (A-M valid, N-Z invalid)
2. Case-insensitive (Alice = alice)
3. Empty or whitespace-only names return 400 error
4. Non-alphabetic first characters result in 400 error
5. Server runs on port 8080
