# Concurrent JSON Database Engine

A lightweight, file-based JSON database engine built in Go — with concurrency-safe CRUD operations, per-collection mutex locking, and atomic writes for crash-safe persistence. No external database required.

---

## Why I Built This

Most embedded databases are either too heavy or abstract away the internals. This project explores how a real database handles concurrent access and durable writes at the systems level — using only Go's standard library and filesystem primitives.

---

## Features

- **Concurrent-safe reads and writes** — per-collection `sync.Mutex` prevents data races across goroutines
- **Atomic file writes** — data is written to a `.tmp` file first, then renamed, ensuring no partial writes on crash
- **Collection-based storage** — records are organized as `collection/resource.json` on disk
- **Full CRUD** — `Write`, `Read`, `ReadAll`, `Delete`
- **Pluggable logger** — inject any logger implementing the `Logger` interface; defaults to console

---

## Project Structure

```
.
├── main.go       # Driver implementation + usage example
├── go.mod
├── go.sum
└── users/        # Auto-created collection directory (generated at runtime)
```

---

## How It Works

### Storage Layout
Each record is stored as a JSON file on disk:
```
<dir>/<collection>/<resource>.json
```
Example: `./users/John.json`

### Concurrency Model
- A **global mutex** protects the `mutexes` map itself
- Each **collection gets its own mutex**, so writes to `users` never block writes to `products`
- Reads are unlocked (file reads are atomic at the OS level for single files)

### Atomic Writes
```
Write data → temp file (.tmp)
             ↓
         os.Rename() → final file   ← atomic on all POSIX systems
```
If the process crashes mid-write, the original file is untouched.

---

## Getting Started

```bash
git clone https://github.com/SahilPatel8826/golang
cd golang/Database
go mod tidy
go run main.go
```

---

## Usage

```go
// Initialize the database (creates directory if not exists)
db, err := New("./mydb", nil)

// Write a record
db.Write("users", "john", User{Name: "John", Age: "25"})

// Read a single record
var u User
db.Read("users", "john", &u)

// Read all records in a collection
records, _ := db.ReadAll("users")

// Delete a record
db.Delete("users", "john")
```

---

## Tech Stack

- **Language:** Go
- **Concurrency:** `sync.Mutex` (global + per-collection)
- **Persistence:** `os.Rename` atomic writes, `encoding/json`
- **Logging:** [`lumber`](https://github.com/jcelliott/lumber) (swappable via interface)

---

## Key Concepts Demonstrated

| Concept | Implementation |
|---|---|
| Concurrency safety | Per-collection mutex map with global guard |
| Crash-safe persistence | Write-to-temp + atomic rename |
| Interface-driven design | Pluggable `Logger` interface |
| Zero external DB dependency | Pure filesystem + JSON |
