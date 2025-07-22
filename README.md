# 🧹 ExpireDex

**ExpireDex** is a lightweight Go microservice designed to perform scheduled cleanup of expired keys from a high-performance key-value database like **Aerospike**. It targets keys indexed by expiration date and removes them efficiently using a structured, observable cron-based engine.

---

## 🚀 Why ExpireDex?

In systems that generate short-lived data like OTPs, session tokens, or retry queues, relying solely on native TTLs may lead to leftover data or require periodic garbage collection.

**ExpireDex solves this by:**
- Organizing expirable keys under a cleanup index (e.g. `delete_on:{YYYY-MM-DD}`)
- Performing **daily deletions** of expired entries
- Logging cleanups with structured Go logs
- Operating as a plug-and-play background microservice

---

## 🔧 Architecture Overview

```
                   +--------------------+
                   |  Cleanup Scheduler |
                   +---------+----------+
                             |
                             v
      +-----------------------------+       +-------------------+
      |   delete_on:{2025-07-22}    |<----->|   Aerospike DB    |
      | [key1, key2, key3, ...]     |       +-------------------+
      +-----------------------------+
                             |
                             v
                +-------------------------+
                |  Structured Logger      |
                +-------------------------+
```

---

## 🛠️ Tech Stack

- **Language**: Go 1.22+
- **Database**: Aerospike (via official Go client)
- **Scheduler**: Custom cron-like runner
- **Logging**: JSON-structured logs via custom logger
- **Structure**: Modular with `cmd/`, `internal/`, `config/`

---

## 🗂️ Folder Structure

```
expiredex/
├── cmd/expiredex/main.go         # Entry point
├── config/config.go              # Env/config loader
├── internal/
│   ├── cleanup/                  # Cleanup job logic
│   ├── db/                       # Aerospike wrappers
│   └── utils/logger.go           # Structured logger
├── .env                          # Config vars
├── go.mod
└── README.md
```

---

## ⚙️ Setup & Run

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/expiredex.git
cd expiredex
```

### 2. Create `.env` file

```env
AEROSPIKE_HOST=127.0.0.1
AEROSPIKE_PORT=3000
AEROSPIKE_NAMESPACE=test
LOG_LEVEL=INFO
```

### 3. Run the cleanup job manually

```bash
go run cmd/expiredex/main.go
```

---

## 🧪 Example Use Case: OTP Cleanup

- OTPs are stored with keys like `otp:user1234`
- Each OTP key is added to a set: `delete_on:2025-07-22`
- At midnight, ExpireDex reads `delete_on:2025-07-22` and deletes the associated keys

---

## 🧠 Designed For

- Engineers who want **modular, testable cleanup utilities**
- Teams using **Aerospike/Redis** and need **custom TTL enforcement**
- Projects requiring **decoupled deletion logic** from main services

---

## 📌 Future Enhancements

- Dry-run and metrics mode
- Webhook or Slack notifications on failures
- Optional Web API control layer
- Generic backend interface to support Redis, DynamoDB

---

## 👨‍💻 Author

**Hari Aneesh Siddhartha**  
Backend Engineer | Aerospike, Redis, Golang | System Design Enthusiast

---

## 🛡 License

MIT License
