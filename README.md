# school_ride_backend
An amazing backend service that lets you track and manage school ride routes, ensuring you always know where your children are.

## Installation

Install [`just`](https://github.com/casey/just) by following the instructions in the [official repository](https://github.com/casey/just#installation).

After installing `just`, you can install the Go dependencies:

```sh
just install      # Installs Go dependencies
```

Next, prepare a PostgreSQL 16 database using Docker:

Start a PostgreSQL 16 container with Docker. If you encounter a port conflict error, it may be because your system is already using port 5432 (the default PostgreSQL port). In that case, either stop the service using that port or modify the Docker command to use a different port (e.g., 5434) to avoid conflicts.

```sh
just preparedb
```

This command will run:

```sh
docker run --name school_ride_postgres -e POSTGRES_PASSWORD=yourpassword -p 5432:5432 -d postgres:16
```

Create the database and set up the schema:

```sh
just createdb    # Creates the database
```

To start the development server:

```sh
just startdb      # Starts the backend server in development mode
```

To connect to the database and visualize tables, you can install [DBeaver](https://dbeaver.io/) and create a new connection to `localhost:5432` using the credentials you set above.

## Running the Backend Locally

To run the backend locally, you need to have [`air`](https://github.com/cosmtrek/air) (installed automatically by `just install`). If you encounter any package errors, please consult the [air documentation](https://github.com/cosmtrek/air#installation), as these are most likely related to your system's PATH configuration.

## API Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) to generate interactive API documentation, similar to FastAPI's Swagger UI.

To generate and serve the API docs:

1. Install `swag` CLI if you haven't already:

    ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

2. Generate the documentation:

    ```sh
    swag init # Make sure the current parent have the main.go or just change directory and run it again.
    ```

3. Start the backend server (e.g., with `just startdb`). The Swagger UI will be available at:

    ```
    http://localhost:8080/swagger/index.html
    ```

You can now explore and test all available API endpoints directly from the browser.