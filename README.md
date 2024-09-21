# Simple Bank API

This project implements a simple banking system API with features for account management, money transfers, and user authentication.

## API Documentation

### Accounts

#### Create Account
- **Endpoint**: `POST /accounts`
- **Description**: Creates a new account for a user.
- **Request Body**:
  ```json
  {
    "owner": "string",
    "currency": "USD"
  }
  ```
- **Response**: Returns the created account details.

#### Get Account
- **Endpoint**: `GET /accounts/{id}`
- **Description**: Retrieves account details by ID.
- **Response**: Returns the account details.

#### List Accounts
- **Endpoint**: `GET /accounts`
- **Description**: Lists accounts for the authenticated user.
- **Query Parameters**:
  - `page_id`: int (optional)
  - `page_size`: int (optional)
- **Response**: Returns a list of accounts.

#### Update Account
- **Endpoint**: `PUT /accounts/{id}`
- **Description**: Updates account balance.
- **Request Body**:
  ```json
  {
    "balance": 1000
  }
  ```
- **Response**: Returns the updated account details.

#### Delete Account
- **Endpoint**: `DELETE /accounts/{id}`
- **Description**: Deletes an account by ID.

### Transfers

#### Create Transfer
- **Endpoint**: `POST /transfers`
- **Description**: Creates a new transfer between accounts.
- **Request Body**:
  ```json
  {
    "from_account_id": 1,
    "to_account_id": 2,
    "amount": 100
  }
  ```
- **Response**: Returns the transfer details.

#### Get Transfer
- **Endpoint**: `GET /transfers/{id}`
- **Description**: Retrieves transfer details by ID.
- **Response**: Returns the transfer details.

#### List Transfers
- **Endpoint**: `GET /transfers`
- **Description**: Lists transfers for the authenticated user.
- **Query Parameters**:
  - `from_account_id`: int (optional)
  - `to_account_id`: int (optional)
  - `page_id`: int (optional)
  - `page_size`: int (optional)
- **Response**: Returns a list of transfers.

### Users

#### Create User
- **Endpoint**: `POST /users`
- **Description**: Creates a new user account.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "full_name": "string",
    "email": "string"
  }
  ```
- **Response**: Returns the created user details (excluding password).

#### Login User
- **Endpoint**: `POST /users/login`
- **Description**: Authenticates a user and returns an access token.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response**: Returns an access token for the authenticated user.

## Project Structure

- `db/`: Database-related code
  - `migration/`: SQL migration files
  - `sqlc/`: Generated Go code for database operations
  - `query/`: SQL query files
- `api/`: API handlers and middleware
- `util/`: Utility functions and configurations
- `token/`: JWT token generation and validation

## Setup and Installation

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Set up the database and run migrations
4. Configure environment variables
5. Run the server: `go run main.go`

## Testing

Run tests with: `go test ./...`

## Technologies Used

- Go
- PostgreSQL
- SQLC
- Gin (Web Framework)
- JWT for authentication
- Viper for configuration management

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.