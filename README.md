# Go Gin API Server

A RESTful API server built with Go and Gin framework featuring JWT authentication, role-based access control, and email functionality.

## Features

- 🔐 JWT-based authentication
- 👥 Role-based access control (User/Admin)
- 🔒 Password hashing with bcrypt
- 📧 Email sending functionality
- 🛡️ Custom middleware (Logger, Auth, Admin-only)
- 🌐 CORS enabled

## Tech Stack

- **Go** - Programming language
- **Gin** - Web framework
- **JWT** - Authentication tokens
- **bcrypt** - Password hashing
- **SMTP** - Email sending

## Prerequisites

- Go 1.16 or higher
- Gmail account (for SMTP email sending)

## Installation

1. Clone the repository
```bash
git clone <your-repo-url>
cd back
```

2. Install dependencies
```bash
go mod download
```

3. Create `.env` file
```bash
cp .env.example .env
```

4. Configure your `.env` file with your SMTP credentials:
```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASS=your-app-password
```

**Note:** For Gmail, you need to generate an [App Password](https://support.google.com/accounts/answer/185833)

## Running the Server

```bash
go run .
```

Server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/welcome` | Welcome message |
| GET | `/send-email` | Send test email |
| POST | `/signup` | Create new user account |
| POST | `/login` | Login and get JWT token |

### Protected Endpoints (Requires JWT Token)

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| GET | `/me` | Get current user info | User/Admin |
| GET | `/admin` | Admin dashboard | Admin only |

## API Usage Examples

### 1. Signup
```bash
POST http://localhost:8080/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### 2. Login
```bash
POST http://localhost:8080/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "message": "login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Access Protected Route
```bash
GET http://localhost:8080/me
Authorization: Bearer YOUR_JWT_TOKEN
```

## Testing

Use the provided `api-tests.http` file with VS Code REST Client extension:

1. Install **REST Client** extension
2. Open `api-tests.http`
3. Click "Send Request" above each endpoint

## Project Structure

```
back/
├── main.go           # Entry point, route definitions
├── auth.go           # Authentication handlers (login, signup)
├── middleware.go     # Custom middleware (Logger, Auth, AdminOnly)
├── handlers.go       # Request handlers
├── email.go          # Email functionality
├── api-tests.http    # API testing file
├── .env              # Environment variables (not tracked)
├── .env.example      # Environment template
├── go.mod            # Go dependencies
└── README.md         # This file
```

## Security Notes

- Passwords are hashed using bcrypt
- JWT tokens expire after 24 hours
- Default user role is "user"
- Admin routes require "admin" role
- Never commit `.env` file to Git

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| SMTP_HOST | SMTP server host | smtp.gmail.com |
| SMTP_PORT | SMTP server port | 587 |
| SMTP_USER | Email address | your-email@gmail.com |
| SMTP_PASS | Email app password | xxxx xxxx xxxx xxxx |

## Troubleshooting

**"No .env file found"**
- Create `.env` file with required variables

**Email not sending**
- Verify SMTP credentials
- For Gmail, enable 2FA and generate App Password

**"invalid token"**
- Token may be expired (24h validity)
- Login again to get new token

**"admin access only"**
- Your user has "user" role
- Modify code to create admin user for testing

## License

MIT
