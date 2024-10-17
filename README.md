# AWS Cognito Go SDK Example

This project demonstrates how to use the AWS SDK in Go to interact with AWS Cognito for user authentication and management.

## Installation

First, ensure you have Go installed on your machine. Then, run the following commands to tidy your module and run the application:

```bash
go mod tidy
go run main.go
```

## API Endpoints

### Example API Usage

- **Base URL**: `http://localhost:8181`

### Endpoints

#### Sign Up

**POST** `/user`

##### Request

```http
POST http://localhost:8181/user
Content-Type: application/json

{
  "name": "John",
  "email": "user@example.com",
  "password": "Pass@1234"
}
```

#### Confirm Account

**POST** `/user/confirmation`

##### Request

```http
POST http://localhost:8181/user/confirmation
Content-Type: application/json

{
  "email": "user@example.com",
  "code": "018219" // Replace with the actual confirmation code received via email
}
```

#### Sign In

**POST** `/user/login`

##### Request

```http
POST http://localhost:8181/user/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "Pass@1234"
}
```

##### Response

```json
{
  "token": "eyJraWQiOiJNOTFKVVVxRVBaa1lBS1wvUXBhS29Ld292ZkVYQXV4Snl0TFBFWkk0Z3Bjbz0iLCJhbGciOiJSUzI1NiJ9..."
}
```

**Note**: The token is an OAuth access token that allows you to authenticate future requests to the API.

#### Get User By Token

**GET** `/user`

##### Request

```http
GET http://localhost:8181/user
Content-Type: application/json
Authorization: Bearer <your-access-token>
```

#### Update Password

**PATCH** `/user/password`

##### Request

```http
PATCH http://localhost:8181/user/password
Content-Type: application/json
Authorization: Bearer <your-access-token>

{
  "email": "user@example.com",
  "password": "Wvj0104**99" // Replace with the new password
}
```

## Notes

- The AWS Cognito setup is configured to require names and email confirmations before users can sign in. Make sure the user confirms their email address after signing up to enable successful login.
- Ensure you replace `user@example.com` and `Pass@1234` with your actual test email and password.
