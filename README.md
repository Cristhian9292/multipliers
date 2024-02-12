# Multipliers - Description

The multipliers is an API created to check multipliers of 3, 5 or both and return its type (Type 1, Type 2 or Type 3) or just the number if it is not one of them with numbers between 1-100

## Endpoints

The API consists of 3 main endpoints:

1. `GET /multipliers?number={number}`: Get the value of a number.
2. `POST /multipliers`: Add a number to a users collection.
3. `GET /multipliers/collection`: Retrieves the users collection.
4. `GET /docs/index.html`: Swagger documentation to interact with the services.
5. `GET /healthcheck`: Checks if service is up.

## Project Setup

### With Docker and Docker Compose

1. Clone the repository:

    ```bash
    git clone https://github.com/Cristhian9292/multipliers.git
    cd multipliers
    ```
2. Install Docker:

    ```bash
    sudo apt install docker.io
    ```
3. Run the project with Docker Compose:

    ```bash
    docker-compose up --build
    ```

### Locally
1. Clone the repository:

    ```bash
    git clone https://github.com/Cristhian9292/multipliers.git
    cd multipliers
    ```
2. Install Golang on Ubuntu:

    ```bash
    sudo apt update
    sudo apt install golang-go
    ```

3. Install Docker:

    ```bash
    sudo apt install docker.io
    ```

4. Install Docker Compose:

    ```bash
    sudo apt install docker-compose
    ```

5. Download dependencies:

    ```bash
    go get -d -v ./...
    ```

6. Initialize Swagger documentation:

    ```bash
    swag init
    ```

7. Run the application:

    ```bash
    go run main.go
    ```

## Documentation and Testing

The services are currently allocated in AWS at:
(ec2-18-116-235-91.us-east-2.compute.amazonaws.com:8080/)
- Swagger Documentation: [Swagger UI](http://ec2-18-116-235-91.us-east-2.compute.amazonaws.com:8080/docs/index.html)
- Download Insomnia(Also works on Postman) Collection: [Postman Collection](https://multipliers-collection.s3.us-east-2.amazonaws.com/multipliers.json)

## Testing Roadmap

1. Make a `POST` request to `/multipliers` with a number to start the collection.
2. You can check the collection with a `GET` request to `/multipliers/collection`.
3. Retrieve a value of a number between 1-100 with a `GET` request to `/multipliers`.
4. Explore and test functionalities using the Swagger documentation.

## Contact Information

For more information, contact:

- Name: [Cristhian Cotes]
- Email: [cristhian9292@gmail.com]
- Phone: [+57-300-6587228]
