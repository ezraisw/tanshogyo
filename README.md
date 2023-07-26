# Tanshōgyō (単商業)

Simple example for multi go project monorepo using microservice paradigm.

## Major libraries used

-   Compile-time DI: [wire](https://github.com/google/wire)
-   Configuration: [viper](https://github.com/spf13/viper)
-   Inter-service communication: [gRPC](https://google.golang.org/grpc) and [protobuf](https://google.golang.org/protobuf)
-   DB ORM: [GORM](https://gorm.io/gorm)
-   Testing: [Ginkgo](https://github.com/onsi/ginkgo) and [Gomega](https://github.com/onsi/gomega)
-   HTTP Routing: [chi](https://github.com/go-chi/chi)
-   Validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)

## How to Run

### Windows

1. Use Docker, Docker Desktop, or WSL
2. Make sure these ports are free:
    - 7080
    - 7880
    - 8080
    - 8880
    - 9080
    - 9880
3. Run `make runall`
   (Run `cd deployments && docker compose up --build` if make does not exist)
4. Use postman to test out APIs

### Linux

1. Install Docker
2. Make sure these ports are free:
    - 7080
    - 7880
    - 8080
    - 8880
    - 9080
    - 9880
3. Run `make runall`
   (Run `cd deployments && docker compose up --build` if make does not exist)
4. Use postman to test out APIs

## Use case

### All

-   List products
-   See product detail

### Users

-   Register
-   Login
-   Create seller account
-   Add to cart
    -   Create transaction
-   List transactions (buyer)

### Seller

-   Create, update, delete product

## Services

### User service

-   Login
-   Register
-   Token verification

### Product service

-   Seller
    -   Create seller account
-   Product
    -   Find all products
    -   Find product
    -   Create
    -   Update
    -   Delete

### Transaction service

-   Transaction
    -   Add to cart
    -   Create transaction

## Testing

Testing can be done through the included postman collection. Some requests that require authentication needs the header `X-Auth-Token` obtained from User Login.
