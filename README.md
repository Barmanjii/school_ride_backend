# school_ride_backend
An amazing backend service that lets you track and manage school ride routes, ensuring you always know where your children are.

## Installation

Install [`just`](https://github.com/casey/just) by following the instructions in the [official repository](https://github.com/casey/just#installation).

After installing `just`, you can install the Go dependencies:

```sh
just install      # Installs Go dependencies
```

Next, prepare a PostgreSQL 16 database using Docker:

Start a PostgreSQL 16 container with Docker:

```sh
just prepare
```

This command will run:

```sh
docker run --name school_ride_postgres -e POSTGRES_PASSWORD=yourpassword -p 5432:5432 -d postgres:16
```

Create the database and set up the schema:

```sh
just create-db    # Creates the database
```

To start the development server:

```sh
just startdb      # Starts the backend server in development mode
```

To connect to the database and visualize tables, you can install [DBeaver](https://dbeaver.io/) and create a new connection to `localhost:5432` using the credentials you set above.

## Running the Backend Locally

To run the backend locally, you need to have [`air`](https://github.com/cosmtrek/air) (installed automatically by `just install`). If you encounter any package errors, please consult the [air documentation](https://github.com/cosmtrek/air#installation), as these are most likely related to your system's PATH configuration.