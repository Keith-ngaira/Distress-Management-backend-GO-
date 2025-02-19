# Distress Messages Management System

A comprehensive system for managing and tracking distress messages through their lifecycle. This system provides a user-friendly interface for handling distress cases from initial receipt to final resolution, designed specifically for front office personnel, directors, and case officers.

## Project Overview

The Distress Messages Management System consists of two main parts:
1. A React-based frontend for user interaction
2. A Go-based backend API for data management and business logic

The system implements a comprehensive workflow that guides users through each stage of case management while ensuring data integrity and security.

### Key Features

- Modern, intuitive user interface
- Case creation and registration system
- Document attachment capabilities
- Real-time case status tracking
- Progress note management
- Automated case assignment workflow
- Comprehensive case details view
- Interactive dashboard for case overview
- Dark/Light theme support with persistent settings
- Fully responsive design for all devices
- Cross-platform compatibility
- Secure file storage and management
- RESTful API endpoints
- Structured data management

## Frontend Documentation

### Project Structure

```
frontend/
├── public/
│   └── index.html          # Main HTML file
├── src/
│   ├── components/         # React components
│   │   ├── CaseDetails.js  # Individual case view component
│   │   ├── CasesList.js    # Cases dashboard component
│   │   ├── DistressForm.js # Case creation form
│   │   └── Layout.js       # Main layout component
│   ├── contexts/           # React contexts
│   │   └── ThemeContext.js # Theme management context
│   ├── services/           # API and service layer
│   │   └── api.js         # API client configuration
│   └── App.js             # Main application component
└── package.json           # Frontend dependencies and scripts
```

### Component Details

#### Layout.js
- Main application shell with responsive design
- Persistent navigation drawer with auto-hide on mobile
- Dynamic theme-aware top app bar
- Integrated theme toggle functionality

#### DistressForm.js
- Comprehensive case registration form with:
  - Automatic reference number generation
  - File attachment handling
  - Form validation
  - Progress tracking
  - Real-time updates

#### CasesList.js
- Interactive data grid with:
  - Sortable columns
  - Search functionality
  - Status indicators
  - Quick actions
  - Pagination support

#### CaseDetails.js
- Detailed case view featuring:
  - Case progress timeline
  - Document management
  - Status updates
  - Officer assignments
  - Progress note system

### Context Details

#### ThemeContext.js
- Global theme management with:
  - Light/Dark mode toggle
  - Persistent theme preferences
  - Custom color palettes
  - Automatic contrast optimization

## Backend Documentation

### Project Structure

```
backend/
├── cmd/
│   └── db/                # Database initialization and migrations
├── handlers/
│   ├── app.go            # Application setup and middleware
│   ├── case_handler.go   # Case management endpoints
│   ├── dashboard_handler.go # Dashboard data endpoints
│   ├── document_handler.go  # Document management endpoints
│   ├── progress_notes_handler.go # Progress notes endpoints
│   └── response_utils.go    # Common response utilities
├── models/
│   ├── case.go           # Case data model
│   ├── document.go       # Document data model
│   ├── progress_note.go  # Progress notes model
│   └── user.go           # User data model
├── uploads/              # File storage directory
├── main.go              # Application entry point
├── go.mod               # Go module definition
└── go.sum               # Go module checksums
```

### Handler Details

#### app.go
- Application initialization
- Middleware configuration
- Route setup
- CORS configuration
- Error handling

#### case_handler.go
- Case management endpoints:
  - Case creation
  - Case updates
  - Case retrieval
  - Case listing
  - Case deletion

#### document_handler.go
- Document management:
  - File uploads
  - File downloads
  - Document metadata
  - Document deletion

#### progress_notes_handler.go
- Progress notes functionality:
  - Note creation
  - Note updates
  - Note retrieval
  - Note listing

#### dashboard_handler.go
- Dashboard data endpoints:
  - Case statistics
  - Recent activities
  - Status summaries

### Model Details

#### case.go
- Case data structure
- Case status management
- Case validation
- Database interactions

#### document.go
- Document metadata
- File storage management
- Document validation

#### progress_note.go
- Progress note structure
- Note validation
- Timestamp management

#### user.go
- User data structure
- Authentication data
- User permissions

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- Go (v1.16 or higher)
- PostgreSQL database

### Installation

1. Clone the repository
2. Set up the frontend:
   ```bash
   npm install
   npm start
   ```
3. Set up the backend:
   ```bash
   cd backend
   go mod download
   go run main.go
   ```

### Environment Configuration

Create a `.env` file in the backend directory with the following variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=distress_management
JWT_SECRET=your_jwt_secret
UPLOAD_DIR=./uploads
```

## API Documentation

The backend provides a RESTful API with the following main endpoints:

### Cases
- `GET /api/cases` - List all cases
- `POST /api/cases` - Create new case
- `GET /api/cases/:id` - Get case details
- `PUT /api/cases/:id` - Update case
- `DELETE /api/cases/:id` - Delete case

### Documents
- `POST /api/documents` - Upload document
- `GET /api/documents/:id` - Download document
- `DELETE /api/documents/:id` - Delete document

### Progress Notes
- `GET /api/cases/:id/notes` - List case notes
- `POST /api/cases/:id/notes` - Add note
- `PUT /api/notes/:id` - Update note
- `DELETE /api/notes/:id` - Delete note

### Dashboard
- `GET /api/dashboard/stats` - Get dashboard statistics
- `GET /api/dashboard/recent` - Get recent activities
