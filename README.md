# Golang Tech Test

As part of the recruitment process we want to know how you think, code and structure your work.
In order to do that, we're going to ask you to complete this coding challenge.

**Please do not spend more than 5 hours on this task**, as this would not be
respectful of your time.

## Task

In this repo is a small application with the following endpoints already implemented:

1. `POST /buff`

    That creates a question ('buff') and it's answers in the system
    ```
    $ curl -X POST -d '{"question":"whos the man?", "answers": ["you", "me"]}' localhost:8080/buff -D -
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 07 Oct 2020 13:05:53 GMT
    Content-Length: 59
    
    {"id":1,"question":"whos the man?","answers":["you","me"]}
    ```

2. `GET /buff/{buff_id}`

    That returns a question ('buff') by ID
    ```
    $ curl localhost:8080/buff/1 -D -
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 07 Oct 2020 13:16:24 GMT
    Content-Length: 59
    
    {"id":1,"question":"whos the man?","answers":["you","me"]}
    ```

**We need you to add some more endpoints to this application**

1. `POST /stream`

    That creates a stream in the system.
    
    Maybe structured like:
    ```json
    {"stream_name": "super cool stream"}
    ```


2. `GET /stream`

    That returns a paginated list of streams in the system.
    
    Maybe structured like:
    ```json
    [
        {
            "id": 1,
            "stream_name": "super cool stream"
        },
        {
            "id": 2,
            "stream_name": "super awesome stream"
        }
    ]
    ```

3. `GET /stream/{stream_id}`

    That returns a single stream along with the buffs associated with that stream.
    
    Maybe structured like:
    ```json
    {
        "id": 1,
        "stream_name": "super cool stream",
        "buffs": [1,3,12,14]
    }
    ```

To do this you will have to:
* Add `Stream` and `StreamCreate` types to the `cmd/server/internal/handlers/api.go` file
* Add `GetStream` and `SetStream` methods to the `handlers.Store` interface
* Add the new routes and handlers to `cmd/server/internal/handlers` package

The endpoints you add should also have test coverage added

Additionaly, the api is currently served from a memory backed store implementation.
For production useage we need a relational DB implementation of the `handlers.Store` interface.

To add this you will have to:
* Setup a docker-compose file to spin up your DB of choice
* Use whatever tooling you are comfortable with to manage migrating the schema
* Add a database package that prodives a DB backed implementation of the `handlers.Store` interface
    and use that in `cmd/server/main.go` in place of the in-memory store


## How to impress us

There are a few optional tasks you can complete if you really want to show off.

1. There is no logging in the application

    Structured logging is an important part of a production ready app

2. There is no config management in the application.

    Ideally things like DB creds and the app's port should be configure-able

3. The pre-existing endpoints do not have test coverage


## How to setup the current app

1. Start the server app
    ```bash
    make run
    ```

2. In another terminal, run a few curl commands like:
    ```bash
    curl -X POST -d '{"question":"whos the man?", "answers": ["you", "me"]}' localhost:8080/buff -D -
    ```
