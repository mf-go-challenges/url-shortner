# URL Shortener

This project is a URL shortener. You will finish it in four steps. Each step gives a working version. You add new features each time. It should take a few weeks of work.

---

## Milestone 1: Basic Shortener (In-Memory)

**Goal:** Make a simple API that shortens URLs and redirects.

* Create `POST /shorten` endpoint.

  * Request JSON: `{ "url": "https://..." }`
  * Response JSON: `{ "code": "abc123" }`
* Create `GET /{code}` endpoint.

  * Redirect to the original URL with HTTP 302.
* Store data in a Go map (`map[string]string`).
* Make the map safe for concurrent access (use `sync.Mutex` or `sync.Map`).
* Validate URLs (allow only `http://` or `https://`).
* Generate 6-character code using `crypto/rand`.

**Deliverable:**

* A `main.go` file. Running `go run main.go` starts the server.
* A README with examples using `curl`.

---

## Milestone 2: Database & Bulk Upload

**Goal:** Save data in a database and add file upload.

* Use SQLite (with `modernc.org/sqlite`) or Postgres.

  * Create a table `links(code TEXT PRIMARY KEY, url TEXT NOT NULL, created_at TIMESTAMP)`.
* Add `POST /bulk` endpoint for file upload.

  * Accept `multipart/form-data` with a text file.
  * File has one URL per line.
  * For each URL, make a code and save in one database transaction.
  * Return JSON array of `{ "url": "...", "code": "..." }`.
* Handle errors: if one URL is bad, skip or report it.
* Implement graceful shutdown with `http.Server.Shutdown`.

**Deliverable:**

* A `Dockerfile` to build the app.
* A SQL file `schema.sql`.
* Instructions in README to run with Docker.

---

## Milestone 3: Security & Logging

**Goal:** Add user login and rate limit. Add logging.

* Add user login with JWT.

  * Store users in the database with hashed passwords (`bcrypt`).
  * Add `POST /login` endpoint to get JWT.
  * Protect `/shorten`, `/bulk`, and `/stats` endpoints with JWT.
* Add rate limit per IP and per user.

  * Use a token bucket in memory or Redis.
* Add simple logging with a library like `zap` or `zerolog`.

  * Log each request: method, path, user (if any), status.

**Deliverable:**

* Docker Compose file to start app and Redis (if you use Redis).
* README update with login and rate limit info.

---

## Milestone 4: Deployment & Web Interface

**Goal:** Deploy the service on a Linux server using Docker. Add a web page.

* Deploy your Docker image on a Linux server (e.g. a VPS).
* Use a real domain name and HTTPS (Let's Encrypt or Caddy).
* Add a web page “My links”:

  * Show list of shortened links for logged-in user.
  * Show click count and creation date.
  * Build this page as a simple SPA with modern JS.
* (Optional) Add more features:

  * QR code endpoint: `GET /qr/{code}.png`.
  * Delete expired links job with `robfig/cron`.
  * CSV report of link usage.

**Deliverable:**

* Instructions in README to deploy on VPS.
* Working domain with HTTPS.
* Link to the web page.

