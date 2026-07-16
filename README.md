# Gator — RSS Feed Aggregator

A CLI tool for aggregating and browsing RSS feeds, built with Go and PostgreSQL.

## Prerequisites

- [Go](https://go.dev/dl/) 1.26 or later
- [PostgreSQL](https://www.postgresql.org/download/) running locally

## Installation

```bash
go install github.com/Hamadn/gator@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Setup

### 1. Create the database

```bash
createdb gator
```

### 2. Run migrations

Install [goose](https://github.com/pressly/goose) if you haven't already:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Then apply the schema migrations:

```bash
goose postgres "postgres://postgres:@localhost:5432/gator?sslmode=disable" up
```

### 3. Configure the CLI

Create `~/.gatorconfig.json` with your database connection string:

```json
{"db_url":"postgres://postgres:@localhost:5432/gator?sslmode=disable","current_user_name":""}
```

The `current_user_name` field will be updated automatically when you use the `login` or `register` commands.

## Usage

### Register and log in

```bash
# Create a new account
gator register alice

# Log in as an existing user
gator login alice
```

### Manage feeds

```bash
# Add an RSS feed (you must be logged in)
gator addfeed "Hacker News" https://hnrss.org/frontpage

# Browse all feeds in the system
gator feeds

# Follow a feed (you must be logged in)
gator follow https://hnrss.org/frontpage

# View feeds you follow
gator following

# Unfollow a feed
gator unfollow https://hnrss.org/frontpage
```

### Scrape and browse

```bash
# Start the aggregator — fetches posts from feeds you follow every 60s
gator agg 60s

# Browse recent posts (default limit: 2)
gator browse

# Browse with a custom limit
gator browse 10
```

### Other commands

```bash
# List all registered users
gator users

# Delete all users and start fresh
gator reset
```

## Commands

| Command | Args | Auth | Description |
|---|---|---|---|
| `register` | `<name>` | No | Create a new user |
| `login` | `<name>` | No | Log in as a user |
| `users` | — | No | List all users |
| `reset` | — | No | Delete all users |
| `addfeed` | `<name> <url>` | Yes | Add a new RSS feed |
| `feeds` | — | No | List all feeds |
| `follow` | `<url>` | Yes | Follow a feed |
| `following` | — | Yes | List feeds you follow |
| `unfollow` | `<url>` | Yes | Unfollow a feed |
| `agg` | `<duration>` | No | Start scraping RSS feeds |
| `browse` | `[limit]` | Yes | Browse recent posts |

## Tech Stack

- **Language:** Go 1.26
- **Database:** PostgreSQL
- **SQL layer:** sqlc (type-safe generated Go queries)
- **DB driver:** lib/pq
- **Migrations:** goose
- **RSS parsing:** net/http + encoding/xml
