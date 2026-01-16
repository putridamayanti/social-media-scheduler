# Social Media Post Scheduler

A simple social media post scheduler built with Go, Next.js, PostgreSQL, and Redis.
Users can schedule posts for future publishing; a background worker reliably publishes them when due.


## Features

* User registration & login (email + password)
* Secure session-based authentication (http-only cookies)
* Create, edit, delete scheduled posts
* View upcoming scheduled posts
* View published post history
* Background worker for reliable publishing
* Redis-based time scheduling (sorted set)
* Dockerized local development
* CI with GitHub Actions
* Redis locking for idempotent publishing 
* Time-zone safe scheduling (UTC storage)
* Clean separation of API & worker 
* SSR-friendly Next.js frontend

## Design Decision

### API & Worker
* API and worker are separate binaries (cmd/api, cmd/worker)
* Avoids runtime mode flags
* Clear responsibility separation


### Redis Sorted Set
* Posts are stored in a Redis ZSET:

member → post_id

score → scheduled_at (Unix timestamp)

* Worker polls due posts using ZRANGEBYSCORE


### Locking
* Redis SETNX lock (lock:post:{id}) prevents duplicate publishing
* DB update guarded by status = 'scheduled'
* Safe for multiple workers


### Authentication
* Session-based auth using http-only cookies
* No JWT stored in frontend
* Backend identifies user via session middleware


### Frontend
* Next.js App Router
* SSR for protected pages
* Server Actions / Route Handlers proxy API calls
* Cookies forwarded automatically



## Running Locally (Docker)

### Start the stack

`docker-compose up --build`

### Services started:

Go API → http://localhost:8080

Next.js Frontend → http://localhost:3000

PostgreSQL

Redis

Background Worker



## Tests

### Service Test

`go test internal/services/tests`

### Queue Test

`go test internal/queue`


## Demo Video

### Setup and demo video

https://youtu.be/xVHf_lFjDeE


# 🙌 Final Notes

This project prioritizes correctness, reliability, and clarity over feature breadth.
All core requirements are implemented with room for safe scaling and extension.

Thanks for reviewing — I’m happy to walk through any part of the implementation.
