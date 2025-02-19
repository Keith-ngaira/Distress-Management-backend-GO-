# Distress Messages Management System - Frontend

A modern React-based frontend application for managing and tracking distress messages through their lifecycle. This system provides a user-friendly interface for handling distress cases from initial receipt to final resolution, designed specifically for front office personnel, directors, and case officers.

## Project Overview

The Distress Messages Management System frontend is built with React and Material-UI, providing a responsive and intuitive interface for managing distress cases. It implements a comprehensive workflow system that guides users through each stage of case management.

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

## Project Structure

```
distress-management-system/
├── public/
│   └── index.html          # Main HTML file
├── src/
│   ├── components/         # React components
│   │   ├── Auth/            # Authentication components
│   │   │   ├── Login.js     # Login form
│   │   │   └── Register.js  # Registration form
│   │   ├── Cases/           # Case management components
│   │   │   ├── CasesList.js # Cases dashboard
│   │   │   ├── CaseView.js  # Individual case view
│   │   │   └── DistressForm.js # Case creation form
│   │   ├── Layout/          # Layout components
│   │   │   ├── Navbar.js    # Navigation bar
│   │   │   └── Sidebar.js   # Side navigation
│   │   └── common/          # Shared components
│   ├── services/            # API and service layer
│   │   ├── api.js          # API client configuration
│   │   ├── auth.js         # Authentication service
│   │   └── cases.js        # Cases service
│   ├── context/            # React context
│   │   └── AuthContext.js  # Authentication context
│   ├── styles/            # CSS and styling
│   │   └── index.css      # Global styles
│   └── utils/             # Utility functions
│       ├── constants.js   # Constants and enums
│       └── helpers.js     # Helper functions
├── .gitignore             # Git ignore configuration
└── package.json           # Project dependencies and scripts
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

## Getting Started

### Prerequisites

- Node.js (v14 or higher)
- npm (v6 or higher)
- Git (for version control)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Keith-ngaira/Distress-Management-frontend-.git
   ```

2. Navigate to the project directory:
   ```bash
   cd Distress-Management-frontend-
   ```

3. Install dependencies:
   ```bash
   npm install
   ```

4. Start the development server:
   ```bash
   npm start
   ```

The application will open in your default browser at `http://localhost:3000`.

### Project Portability

The project is configured for maximum portability across different development environments:

1. **Git Configuration**
   - Comprehensive `.gitignore` file excludes:
     - Node modules and dependencies
     - Build artifacts and cache
     - Environment-specific files
     - IDE configurations
     - System-generated files

2. **Environment Independence**
   - Environment-agnostic configuration
   - No hard-coded variables
   - Cross-platform compatibility
   - Flexible deployment options

3. **Development Setup**
   - Streamlined npm-based workflow
   - Minimal configuration requirements
   - Consistent development experience
   - Hot-reloading support

## Workflow Stages

1. **Front Office Receipt**
   - Digital case registration
   - Document digitization
   - Initial categorization

2. **Director Review**
   - Priority assessment
   - Officer assignment
   - Instruction provision

3. **Cadet Assignment**
   - Case acceptance
   - Resource planning
   - Initial response

4. **Case Investigation**
   - Information gathering
   - Progress documentation
   - Status updates

5. **Case Resolution**
   - Solution implementation
   - Outcome recording
   - Client communication

6. **Final Review**
   - Quality assurance
   - Compliance check
   - Director approval

7. **Archived**
   - Digital archival
   - Record finalization
   - Accessibility maintenance

## User Interface Features

### Theme Switching
- Seamless light/dark mode toggle
- Context-aware color adjustments
- Persistent user preferences
- High-contrast accessibility

### Responsive Design
- Mobile-first approach
- Fluid layouts
- Touch-friendly interfaces
- Cross-browser compatibility

## Technology Stack

- **Frontend Framework**: React 18
- **UI Library**: Material-UI v5
- **Routing**: React Router v6
- **State Management**: React Hooks & Context API
- **Data Grid**: MUI X-Data Grid
- **Theming**: Material-UI Theme Provider
- **Version Control**: Git

## Authentication
- JWT-based authentication
- Role-based access control
- Session management
- Secure password handling
- Remember me functionality

## Case Management
- Create and track cases
- File attachments
- Status updates
- Progress notes
- Reference number generation
- Case history tracking

## Data Management
- Real-time updates
- Pagination
- Sorting
- Filtering
- Search functionality

## State Management
- React Context API
- Local storage
- Session management
- Form state handling
- Cache management

## API Integration
- RESTful API consumption
- Token-based authentication
- Error handling
- Loading states
- Retry mechanisms

## Development Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Configure environment:
   Create a .env file:
   ```env
   REACT_APP_API_URL=http://localhost:8080/api
   REACT_APP_ENV=development
   ```

3. Start development server:
   ```bash
   npm start
   ```

4. Build for production:
   ```bash
   npm run build
   ```

## Available Scripts

- `npm start` - Start development server
- `npm test` - Run tests
- `npm run build` - Build for production
- `npm run eject` - Eject from Create React App
- `npm run lint` - Run ESLint
- `npm run format` - Format code with Prettier

## Dependencies

### Core
- React
- React Router
- Axios
- Material-UI
- @mui/x-data-grid

### Development
- ESLint
- Prettier
- Jest
- React Testing Library
- TypeScript

## Browser Support
- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Security Features
- JWT token handling
- XSS protection
- CSRF protection
- Secure password handling
- Input sanitization

## Error Handling
- Global error boundary
- API error handling
- Form validation errors
- Network error handling
- Fallback UI components

## Performance Optimization
- Code splitting
- Lazy loading
- Image optimization
- Caching strategies
- Bundle size optimization

## Accessibility
- ARIA labels
- Keyboard navigation
- Screen reader support
- Color contrast
- Focus management

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Testing
- Unit tests
- Integration tests
- End-to-end tests
- Component testing
- Snapshot testing
#   D i s t r e s s - M a n a g e m e n t - b a c k e n d - G O -  
 #   D i s t r e s s - M a n a g e m e n t - b a c k e n d - G O -  
 