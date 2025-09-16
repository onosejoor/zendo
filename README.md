# Zendo: Streamlined Task & Project Management

Organize your work, track progress, and boost productivity with Zendo, a powerful and intuitive task management application. Built with a robust Go backend and a modern Next.js frontend, Zendo provides a seamless experience for individuals and teams to manage tasks and projects efficiently.

---

## ✨ Features

-   **User Authentication**: Secure user registration, login, and session management using JWT.
-   **Email Verification**: Ensures account security and authenticity through email verification links.
-   **Google OAuth2 Integration**: Seamless sign-in experience using Google accounts.
-   **Task Management**: Create, view, update, and delete tasks. Add and manage subtasks within larger tasks. Set due dates and track task status (pending, in-progress, completed). Receive automated email reminders for upcoming task due dates.
-   **Project Management**: Organize tasks into projects for better structuring. View all tasks associated with a specific project. Create, update, and delete projects.
-   **Team Collaboration**: Create teams, invite members with specific roles (owner, admin, member), and manage team-specific tasks and projects.
-   **Dashboard & Analytics**: Overview of total tasks, total projects, completion rate, and tasks due today.
-   **User Profile Management**: Update username and avatar.
-   **Data Caching**: Utilizes Redis for efficient data retrieval and improved performance.
-   **Atomic Operations**: Employs MongoDB transactions for reliable data consistency (e.g., when creating/deleting tasks within projects).
-   **Prometheus Metrics**: Exposes metrics for monitoring HTTP requests, database operations, Redis interactions, cron jobs, and user/project/task creation counts.

---

## 🛠️ Technologies Used

| Category | Technology | Description | Link |
| :------- | :----------------------------------------------------------------------------------------------------------------------- | :--------------------------------------------------------------------------- | :--------------------------------------------------------------------------------- |
| **Backend** | ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) | Primary server-side language for robust API development. | [go.dev](https://go.dev/) |
| | ![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=fiber&logoColor=white) | Fast, Express-inspired web framework for Go. | [gofiber.io](https://gofiber.io/) |
| | ![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white) | NoSQL database for flexible and scalable data storage. | [mongodb.com](https://www.mongodb.com/) |
| | ![MongoDB Go Driver](https://img.shields.io/badge/MongoDB_Go_Driver-47A248?style=for-the-badge&logo=go&logoColor=white) | Official MongoDB driver for Go. | [github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver) |
| | ![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white) | An in-memory data structure store, used as a cache and message broker. | [redis.io](https://redis.io/) |
| | ![go-redis/v9](https://img.shields.io/badge/go--redis%2Fv9-DC382D?style=for-the-badge&logo=go&logoColor=white) | Popular Redis client for Go. | [github.com/redis/go-redis](https://github.com/redis/go-redis) |
| | ![GoCron](https://img.shields.io/badge/GoCron-018786?style=for-the-badge&logo=go&logoColor=white) | Library for scheduling recurring tasks (e.g., email reminders). | [github.com/go-co-op/gocron](https://github.com/go-co-op/gocron) |
| | ![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white) | JSON Web Tokens for secure authentication and authorization. | [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt) |
| | ![Bcrypt](https://img.shields.io/badge/Bcrypt-000000?style=for-the-badge&logo=coda&logoColor=white) | Password hashing for user security. | [pkg.go.dev/golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) |
| | ![Cloudinary](https://img.shields.io/badge/Cloudinary-3448C5?style=for-the-badge&logo=cloudinary&logoColor=white) | Cloud-based image and video management for avatar uploads. | [github.com/cloudinary/cloudinary-go](https://github.com/cloudinary/cloudinary-go/v2) |
| | ![GoDotEnv](https://img.shields.io/badge/DotEnv-ECD53F?style=for-the-badge&logo=dot-env&logoColor=white) | Loads environment variables from `.env` files. | [github.com/joho/godotenv](https://github.com/joho/godotenv) |
| | ![go-playground/validator](https://img.shields.io/badge/Validator-blue?style=for-the-badge&logo=go&logoColor=white) | Struct and field validation for Go. | [github.com/go-playground/validator](https://github.com/go-playground/validator) |
| | ![go-mail/mail](https://img.shields.io/badge/go--mail%2Fmail-red?style=for-the-badge&logo=go&logoColor=white) | Go package for sending emails. | [github.com/go-mail/mail](https://github.com/go-mail/mail) |
| | ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white) | Open-source monitoring system for metrics collection and alerting. | [prometheus.io](https://prometheus.io/) |
| **Frontend** | ![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=next.js&logoColor=white) | React framework for server-side rendering and static site generation. | [nextjs.org](https://nextjs.org/) |
| | ![React](https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black) | JavaScript library for building interactive user interfaces. | [react.dev](https://react.dev/) |
| | ![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white) | Superset of JavaScript for type safety and improved developer experience. | [typescriptlang.org](https://www.typescriptlang.org/) |
| | ![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwind-css&logoColor=white) | Utility-first CSS framework for rapid UI development. | [tailwindcss.com](https://tailwindcss.com/) |
| | ![SWR](https://img.shields.io/badge/SWR-000000?style=for-the-badge&logo=swr&logoColor=white) | React Hooks for data fetching with caching, revalidation, and focus-on-mount. | [swr.vercel.app](https://swr.vercel.app/) |
| | ![shadcn/ui](https://img.shields.io/badge/shadcn%2Fui-000000?style=for-the-badge&logo=data-transfer&logoColor=white) | Reusable UI components for consistent design. | [ui.shadcn.com](https://ui.shadcn.com/) |
| | ![Axios](https://img.shields.io/badge/axios-6710E5?style=for-the-badge&logo=axios&logoColor=white) | Promise-based HTTP client for making API requests. | [axios-http.com](https://axios-http.com/) |
| | ![Lucide React](https://img.shields.io/badge/Lucide_React-000000?style=for-the-badge&logo=lucide&logoColor=white) | A library of simply beautiful open-source icons. | [lucide.dev](https://lucide.dev/) |

---

## 🚀 Getting Started

Follow these steps to set up and run Zendo on your local machine.

### Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.22 or higher.
*   **Node.js**: (v18 or higher) and npm/yarn for the frontend.
*   **MongoDB**: A running MongoDB instance. You can use a local installation or a cloud-based service like MongoDB Atlas.
*   **Redis**: A running Redis instance. For development, a local setup is fine; for production, consider a managed service like Upstash.
*   **Cloudinary Account**: Required for image uploads (e.g., user avatars).
*   **Gmail Account**: For sending email verification and task reminder emails. You'll need an App Password if using Gmail directly.

### Installation

1.  **Clone the Repository**:

    ```bash
    git clone https://github.com/onosejoor/zendo.git
    cd zendo
    ```

2.  **Navigate to the Server Directory and Install Dependencies**:

    ```bash
    cd server
    go mod tidy
    ```

3.  **Navigate to the Client Directory and Install Dependencies**:
    ```bash
    cd ../client
    npm install # or yarn install
    ```

### Environment Variables

Before running the application, set up the following environment variables in a `.env` file in both the `server` and `client` directories.

**For `server/.env`:**

```
PORT=8080
DATABASE=zendo_db_name
MONGO_URI="mongodb+srv://<username>:<password>@<cluster-url>/<db-name>?retryWrites=true&w=majority"
REDIS_UPSTASH_URL="rediss://<username>:<password>@<host>:<port>"
ACCESS_SECRET="your_jwt_access_secret"
JWT_SECRET="your_jwt_refresh_secret" # This is used for refresh token and email secret
EMAIL="your_email@gmail.com"
APP_PASSWORD="your_gmail_app_password" # Generated from Google Account Security
FRONTEND_URL="http://localhost:3000" # Or your deployed frontend URL
CLIENT_URL="http://localhost:3000" # Or your deployed frontend URL
ENVIRONMENT="development" # or "production"

# Optional: Cloudinary for avatar uploads
CLOUDINARY_URL="cloudinary://<api_key>:<api_secret>@<cloud_name>"
UPLOAD_PRESET="your_cloudinary_upload_preset"

# Optional: Google OAuth
G_CLIENT_ID="your_google_client_id"
G_CLIENT_SECRET="your_google_client_secret"
G_REDIRECT="http://localhost:8080/auth/callback"

# Optional: Prometheus Basic Auth
METRICS_USERNAME="metricsuser"
METRICS_PASSWORD="metricspassword"
```

**For `client/.env.local`:**

```
NEXT_PUBLIC_SERVER_URL="http://localhost:8080" # Your backend server URL
NEXT_PUBLIC_BACKUP_SERVER_URL="http://localhost:8080" # Backup backend server URL (can be same as primary for local)
```

### Running the Application

1.  **Start the Backend Server**:
    Navigate to the `server` directory and run:

    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:8080` (or your specified `PORT`).

2.  **Start the Frontend Client**:
    Open a new terminal, navigate to the `client` directory and run:
    ```bash
    npm run dev # or yarn dev
    ```
    The client application will typically start on `http://localhost:3000`.

---

## 💡 Usage

The Zendo application provides a clean and intuitive interface for managing your tasks and projects. The backend exposes a RESTful API that can be consumed by any client application.

### Frontend Usage

Once the frontend is running:

1.  **Sign Up/Sign In**: Register a new account or log in using your email/password or Google OAuth.
2.  **Dashboard**: The main dashboard gives you an overview of your tasks, projects, and overall productivity statistics.
3.  **Manage Tasks**:
    *   Create new tasks with titles, descriptions, due dates, and statuses.
    *   Add sub-tasks to break down larger tasks.
    *   Mark tasks as pending, in-progress, or completed.
    *   View task details and receive email reminders for upcoming due dates.
4.  **Organize Projects**:
    *   Create projects to categorize and group related tasks.
    *   View all tasks associated with a specific project.
    *   Update project details and track their progress.
5.  **Collaborate with Teams (New Feature!)**:
    *   Create teams and invite other users to join.
    *   Assign tasks to team members within a project.
    *   Collaborate on shared tasks and track team progress.
6.  **User Settings**: Update your profile information, including username and avatar. Access the "Danger Zone" to delete all your tasks or projects (use with caution!).

### Backend API Overview

The Zendo API is a high-performance backend service built with **Go** and the **Fiber** web framework. It leverages **MongoDB** for persistent data storage and **Redis** for efficient caching, providing a robust and scalable solution for task and project management.

#### Base URL

The API base URL is dependent on your environment configuration.

-   **Local Development**: `http://localhost:8080`
-   **Production**: Refer to your deployed server's base URL.

#### API Endpoints

Below is a detailed list of the API endpoints, their functionalities, and expected interactions.

##### General Endpoints

*   **GET /health**: Checks the health status of the API.
*   **GET /stats**: Retrieves statistics for the authenticated user (total tasks, projects, completion rate, tasks due today).

##### Authentication Endpoints (`/auth`)

*   **GET /auth/user**: Retrieves the profile information of the authenticated user.
*   **PUT /auth/user**: Updates the profile information for the authenticated user (username, avatar).
*   **GET /auth/verify_email**: Verifies a user's email using a magic link token.
*   **POST /auth/verify_email**: Sends a new email verification link to the authenticated user's email address.
*   **GET /auth/refresh-token**: Refreshes the user's access token using the refresh token stored in the `zendo_session_token` cookie.
*   **POST /auth/signup**: Registers a new user account.
*   **POST /auth/signin**: Authenticates a user and creates a session (sets access and refresh token cookies).
*   **GET /auth/oauth/google**: Initiates Google OAuth flow.
*   **GET /auth/oauth/callback**: Google OAuth callback URL.
*   **POST /auth/oauth/exchange**: Exchanges Google OAuth code for user data and creates a session.

##### Task Endpoints (`/tasks`)

*   **GET /tasks**: Retrieves all tasks for the authenticated user, sorted by creation date.
*   **GET /tasks/search**: Searches for tasks based on a query string (title, description).
*   **GET /tasks/:id**: Retrieves a specific task by its ID for the authenticated user.
*   **POST /tasks/new**: Creates a new task for the authenticated user. Supports linking to a project/team and adding subtasks/assignees.
*   **PUT /tasks/:id**: Updates an existing task for the authenticated user.
*   **PUT /tasks/:id/subtask/:subTaskId**: Updates the completion status of a specific subtask within a task.
*   **DELETE /tasks/:id/subtask/:subTaskId**: Deletes a specific subtask from a task.
*   **DELETE /tasks/all**: Deletes all tasks for the authenticated user.
*   **DELETE /tasks/:id**: Deletes a specific task by its ID.

##### Project Endpoints (`/projects`)

*   **GET /projects**: Retrieves all projects for the authenticated user, sorted by creation date.
*   **GET /projects/search**: Searches for projects based on a query string (name, description).
*   **GET /projects/:id/tasks**: Retrieves all tasks associated with a specific project.
*   **GET /projects/:id**: Retrieves a specific project by its ID.
*   **POST /projects/new**: Creates a new project for the authenticated user.
*   **PUT /projects/:id**: Updates an existing project for the authenticated user.
*   **DELETE /projects/all**: Deletes all projects for the authenticated user, including all tasks within those projects.
*   **DELETE /projects/:id**: Deletes a specific project by its ID, and all associated tasks.

##### Team Endpoints (`/teams`)

*   **GET /teams**: Retrieves all teams the authenticated user is a member of.
*   **GET /teams/stats**: Retrieves overall team statistics for the authenticated user.
*   **POST /teams/new**: Creates a new team.
*   **GET /teams/members/invite**: Accepts a team invitation using a token.

##### Team-Specific Endpoints (`/teams/:teamId`)

These endpoints require the user to be a member of the specified team and often include role-based authorization.

*   **GET /teams/:teamId**: Retrieves details for a specific team.
*   **GET /teams/:teamId/stats**: Retrieves statistics for a specific team.
*   **GET /teams/:teamId/invites**: (Owner Only) Retrieves all pending invites for the team.
*   **PUT /teams/:teamId**: (Owner Only) Updates a team's details.
*   **DELETE /teams/:teamId/invites/:inviteId**: (Owner Only) Cancels a pending team invitation.
*   **DELETE /teams/:teamId**: (Owner Only) Deletes a team and all its associated data (tasks, members).

##### Team Members Endpoints (`/teams/:teamId/members`)

*   **GET /teams/:teamId/members**: Retrieves all members of a specific team.
*   **POST /teams/:teamId/members/invite**: (Admin/Owner Only) Sends an invitation email to a new team member.
*   **DELETE /teams/:teamId/members/:memberId**: Removes a member from a team.

##### Team Task Endpoints (`/teams/:teamId/tasks`)

*   **GET /teams/:teamId/tasks**: Retrieves all tasks associated with a specific team.
*   **GET /teams/:teamId/tasks/:id**: Retrieves a specific task within a team.
*   **DELETE /teams/:teamId/tasks/:taskId**: (Owner Only) Deletes a task from a team.

---

## 🤝 Contributing

I welcome contributions to the Zendo project! Whether you want to fix a bug, add a new feature, or improve documentation, your input is highly appreciated.

1.  **Fork the Repository**: Start by forking the `zendo` repository to your GitHub account.
2.  **Create a New Branch**: Choose a descriptive name for your branch (e.g., `feature/add-reminders`, `bugfix/auth-issue`).
    ```bash
    git checkout -b feature/your-feature-name
    ```
3.  **Make Your Changes**: Implement your features or bug fixes. Write clean, maintainable, and well-commented code.
4.  **Write Tests (If applicable)**: Ensure your changes are robust by adding unit or integration tests where appropriate.
5.  **Lint Your Code**: Make sure your code adheres to Go (backend) and JavaScript/TypeScript (frontend) best practices and styling.
    *   **Backend**: `go fmt ./...` and `go vet ./...`
    *   **Frontend**: `npm run lint` or `yarn lint`
6.  **Commit Your Changes**: Write clear and concise commit messages.
7.  **Push to Your Fork**:
    ```bash
    git push origin feature/your-feature-name
    ```
8.  **Create a Pull Request**: Open a pull request from your forked repository to the `main` branch of the original `zendo` repository. Describe your changes thoroughly and link to any relevant issues.

I'll review your contribution as soon as possible. Thank you for making Zendo better!

---

## 📄 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

**Onos Ejoor**
A passionate software developer dedicated to building robust and user-friendly applications.

-   🌐 [Portfolio](https://onos-ejoor.vercel.app/)
-   🐦 [Twitter](https://x.com/DevText16)

---

[![Readme was generated by Readmit](https://img.shields.io/badge/Readme%20was%20generated%20by-Readmit-brightred)](https://readmit.vercel.app)
