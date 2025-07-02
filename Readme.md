# 📦 Package Handler

A simple yet powerful Go web service called **Package Handler** – designed to optimize item packing and prevent over-shipping with inefficient packaging.
The code is structured based on common golang practices of writing web serves. This might be a bit of 'over-engineering' considering that this is the initial implementation.

---

## What does it do?

Given an order amount number and a set of available pack sizes, this service calculates:
- The **smallest total number of items** (≥ order) the company would ship.
- The **fewest number of packs** to achieve the desired number items.
- The breakdown of packages to use in the warehouse.

It uses a dynamic programming solution to determine the best combination of packs minimizing both the total number of items shipped and the number of packs used. This approach guarantees the most cost-effective and space-efficient packing, even for large and complex orders. The decision to use dynamic programming for solving this problem was inspired by the famous book *"Grokking   Algorithms"* and its solution to the knapsack problem.

---

## How to run

### 1. **Run locally (requires Go 1.21+)**

```sh
go run main.go
```
Server will start at [http://localhost:8080](http://localhost:8080).

### 2. **Run with Docker - recommended**

Final image size is reduced in the most optimal way using stage during the build by disabling CGO so that final binary produced does not require any external C library.
This reduced the image size from 25mg - 7mb

Build the image:
```sh
make build
```

Run the container:
```sh
make run
```

Stop the container:
```sh
make stop
```

---

## 🧪 How to test

Run unit tests (table-driven, of course!):

```sh
go test ./logic
```

---

## Example API usage

Send a POST request to `/calculate-packs`:

```sh
curl -X POST http://localhost:8080/calculate-packs \
  -H "Content-Type: application/json" \
  -d '{"order": 4200, "pack_sizes": [250, 500, 1000, 2000, 5000]}'
```

**Response:**
```json
{
  "total_items": 4250,
  "packs": {
    "2000": 2,
    "250": 1
  }
}
```

---

## Project structure

```
package-handler/
├── handlers/      # HTTP handlers
├── logic/         # Packing algorithm (dynamic programming)
├── models/        # Request/response structs
├── main.go        # Entry point
├── Dockerfile     # Containerization
└── readme.md      # This file!
```

---

## Notes

- **Table-driven tests** for robust, idiomatic Go testing.
- **Multi-stage Docker build** for a low size, production-ready image.
