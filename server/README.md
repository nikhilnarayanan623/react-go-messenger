
# Setup Instructions
To use and test the react-go-messenger project, please follow these steps:

### Prerequisites
Make sure you have the following installed on your system:
- Go (https://golang.org/doc/install)
- PostgreSQL (https://www.postgresql.org/download/)

### 1. Clone the Repository
Clone the react-go-messenger repository to your local system:
```bash
git clone https://github.com/nikhilnarayanan623/react-go-messenger.git
cd  react-go-messenger && server
```
### 2. Install Dependencies
Install the required dependencies using either the provided Makefile command or Go's built-in module management:
```bash
# Using Makefile
make deps
# OR using Go
go mod tidy
```
### 3. Configure Environment Variables
details provided at the end of file
### 4. Make Swagger Files (For Swagger API Documentation)
```bash
make swag
```
# To Run The Application
```bash
make run
```
### To See The API Documentation
1. visit [swagger] ***http://localhost:8000/swagger/index.html***

# To Test The Application
### 1. Generate Mock Files
```bash
make generate
```
### 2. Run The Test Functions
```bash
make test
```

# Set up Environment Variables
Set up the necessary environment variables in a .env file at the project's root directory. Below are the variables required:
### PostgreSQL database details
1. DB_HOST="```your database host name```"
2. DB_NAME="```your database name```"
3. DB_USER="```your database user name```"
4. DB_PASSWORD="```your database owner password```"
5. DB_PORT="```your database running port number```"
### JWT
1. JWT_KEY="```secret key for signing JWT token```"
### Google Auth
1. GOAUTH_CLIENT_ID="```your Google Auth client ID```"
2. GOAUTH_CLIENT_SECRET="```your Google Auth secret key```"
3. GOAUTH_CALL_BACK_URL="```your registered callback URL for Google Auth```"

