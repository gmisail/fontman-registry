# fontman-registry

## Building

### Dependencies

- Task (build system)
- PostgreSQL (database)
- Goose (database migrations)

For development, it is recommended that you use `docker-compose`. To run with `docker-compose`, first make sure you have both Docker & Docker Compose.
Then, simply run `sudo docker-compose up --build`. Note that running with `sudo` may be optional depending on how it is installed.

## Installation

### Step 1 

First, the `fontman-registry` binary must be built. This can be done by running `task build`. 

### Step 2

Next, make sure that you have a PostgreSQL server running. You can run migrations by running the following:

```
cd migrations
goose postgres <DATABASE URL> up
cd ..
```

Now your database is up-to-date!

### Step 3

Make sure that the `PORT` and `DATABASE_URL` environment variables are set. The former will be which port the server will run on (defaults to 8080),
while the latter is the URL of your database ("postgres://user:pass@database/name").

You are now all set! Simply run `./bin/fontman-registry`.

## Usage

TBD.
