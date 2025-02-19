# Distress Management System Backend

A robust Go-based backend service for managing distress cases with secure authentication and role-based access control.

## Project Structure

```
backend/
├── auth/                  # Authentication related packages
│   ├── context.go        # Authentication context and middleware
│   ├── jwt.go            # JWT token generation and validation
│   └── password.go       # Password hashing and validation
├── handlers/             # HTTP request handlers
│   ├── app.go           # Main application handler setup
│   ├── auth_handler.go  # Authentication endpoints (login, register)
│   ├── case_handler.go  # Case management endpoints
│   └── user_handler.go  # User management endpoints
├── models/              # Database models and operations
│   ├── case.go         # Case model and database operations
│   ├── db.go           # Database connection and initialization
│   └── user.go         # User model and database operations
├── .env                # Environment configuration
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
├── main.go            # Application entry point
└── schema.sql         # Database schema
```

## Key Components

### Main Application (main.go)
- Entry point of the application
- Sets up database connection
- Configures and starts the HTTP server
- Initializes routing and middleware

### Authentication (auth/)

#### context.go
- Implements authentication middleware
- Manages user context throughout requests
- Handles token validation in requests

#### jwt.go
- JWT token generation and validation
- Token signing and verification
- Token claims management

#### password.go
- Password hashing using bcrypt
- Password validation and comparison
- Secure password storage

### Handlers (handlers/)

#### app.go
- Main application handler setup
- CORS configuration
- Global middleware application
- Database connection management

#### auth_handler.go
- Login endpoint (/api/auth/login)
- Registration endpoint (/api/auth/register)
- Token generation and validation
- User authentication logic

#### case_handler.go
- Case creation (/api/cases)
- Case retrieval (/api/cases/:id)
- Case updates (/api/cases/:id)
- Case listing with pagination
- Case status updates
- Progress note management

#### user_handler.go
- User management endpoints
- User profile updates
- User listing and retrieval
- Role-based access control

### Models (models/)

#### case.go
- Case data structure
- Case database operations
- Case validation logic
- Reference number generation
- Case status management

#### db.go
- Database connection setup
- Connection pool management
- Database initialization
- Migration handling

#### user.go
- User data structure
- User database operations
- Password hashing integration
- Role management
- User validation

### Configuration (.env)
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=distress_management
SERVER_PORT=8080
JWT_SECRET=your_jwt_secret
```

## API Endpoints

### Authentication
- POST /api/auth/login - User login
- POST /api/auth/register - User registration
- POST /api/auth/logout - User logout

### Cases
- GET /api/cases - List all cases (with pagination)
- GET /api/cases/:id - Get specific case
- POST /api/cases - Create new case
- PUT /api/cases/:id - Update case
- PATCH /api/cases/:id/status - Update case status
- POST /api/cases/:id/progress-notes - Add progress note

### Users
- GET /api/users - List all users
- GET /api/users/:id - Get specific user
- PUT /api/users/:id - Update user
- DELETE /api/users/:id - Delete user

## Database Schema (schema.sql)
- Users table - Stores user information and credentials
- Cases table - Stores case information and metadata
- Progress_Notes table - Stores case progress updates
- Roles table - Stores user roles and permissions

## Security Features
- JWT-based authentication
- Password hashing with bcrypt
- Role-based access control
- Request validation
- CORS protection
- SQL injection prevention
- XSS protection

## Development Setup

1. Install Go (1.16 or later)
2. Set up MySQL database
3. Create database using schema.sql
4. Configure .env file
5. Install dependencies:
   ```bash
   go mod download
   ```
6. Run the server:
   ```bash
   go run main.go
   ```

## Error Handling
- Structured error responses
- Detailed logging
- HTTP status codes
- Validation error messages
- Database error handling
- Authentication error handling

## Logging
- Request logging
- Error logging
- Authentication events
- Case status changes
- Database operations
- Server status
