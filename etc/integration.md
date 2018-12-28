# Running integration tests

In order to run integration tests for some of the handlers a local development environment must be configured.
Unfortunately it's a little bit more complicated than a "fire and forget" command,
but most of it can be done in the CLI.

The first part is as easy as executing a series of commands:

```bash
cp docker-compose.override.yml.dist docker-compose.override.yml
docker-compose up -d
docker-compose run --rm errbit rake db:seed
docker-compose run --rm sentry upgrade --noinput
docker-compose run --rm sentry createuser --email admin@example.com --password admin --superuser --no-input
```

Go to the Sentry dashboard:
```bash
open http://localhost:32622
```

Login with `admin@example.com` and `admin` credentials.

Complete the setup wizard and add a new test Go project.

Create a new `.env.test` file:
```bash
cp .env.test.dist .env.test
```
