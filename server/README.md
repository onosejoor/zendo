# Zendo Task Management Backend

Zendo is a robust and efficient backend API built with Go, designed to power a modern task and project management application. This system provides a comprehensive set of functionalities, from secure user authentication and detailed task tracking to project organization and automated reminders, ensuring a seamless and productive user experience. It's engineered for performance and scalability, leveraging powerful technologies like MongoDB for data persistence and Redis for intelligent caching.

## 🚀 Installation

Getting Zendo up and running on your local machine is straightforward. Follow these steps to set up the backend server:

### Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.22 or higher.
*   **MongoDB**: A running MongoDB instance. You can use a local installation or a cloud-based service like MongoDB Atlas.
*   **Redis**: A running Redis instance. For development, a local setup is fine; for production, consider a managed service like Upstash.
*   **Cloudinary Account**: Required for image uploads (e.g., user avatars).
*   **Gmail Account**: For sending email verification and task reminder emails. You'll need an App Password if using Gmail directly.

### Clone the Repository

Start by cloning the project to your local machine:

```bash
git clone https://github.com/onosejoor/zendo.git
cd zendo/server
```

### Install Dependencies

Navigate into the `server` directory and install the Go modules:

```bash
go mod tidy
```

This command will download all necessary dependencies as listed in `go.mod`.

### Environment Variables

Zendo relies on environment variables for configuration. Create a `.env` file in the `server` directory and populate it with your specific credentials and settings. Here's an example of the variables you'll need:

```ini
PORT=8080
DATABASE=zendo_db
MONGODB_URL="mongodb://localhost:27017"
ACCESS_SECRET="your_access_token_secret"
JWT_SECRET="your_refresh_token_secret"
EMAIL_SECRET="your_email_verification_secret"
REDIS_UPSTASH_URL="redis://localhost:6379" # Or your Upstash URL
CLIENT_URL="http://localhost:3000" # Your frontend URL
FRONTEND_URL="http://localhost:3000" # Your frontend URL (if different)
ENVIRONMENT="development" # or "production"

# OAuth (Optional, uncomment if using Google OAuth)
# G_CLIENT_ID="your_google_client_id"
# G_CLIENT_SECRET="your_google_client_secret"
# G_REDIRECT="http://localhost:8080/auth/callback" # Your OAuth callback URL

# Email (Gmail Example)
EMAIL="your_gmail_account@gmail.com"
APP_PASSWORD="your_gmail_app_password" # Generated from Google Account Security

# Cloudinary (for avatar uploads)
CLOUDINARY_URL="cloudinary://your_api_key:your_api_secret@your_cloud_name"
UPLOAD_PRESET="your_cloudinary_upload_preset"

# Prometheus (Optional, for basic authentication to /metrics endpoint)
# METRICS_USERNAME="metricsuser"
# METRICS_PASSWORD="metricspassword"
```

### Run the Server

Once the dependencies are installed and the environment variables are set, you can start the server:

```bash
go run main.go
```

The server will typically start on `http://localhost:8080` (or the `PORT` you configured in your `.env` file). You should see a log message indicating that the "Server listening on port: ..."

## 💡 Usage

The Zendo backend exposes a RESTful API that can be consumed by any client application (e.g., a web frontend, mobile app, or another service). Once the server is running, you can interact with it using tools like Postman, Insomnia, or by integrating it directly into your frontend application.

Here’s a glimpse of the API structure and how to interact with it:

### Health Check

To verify the server is running:

```http
GET /health
```

### Metrics (Optional Basic Auth)

To access Prometheus metrics (if `METRICS_USERNAME` and `METRICS_PASSWORD` are set in `.env`):

```http
GET /metrics
Authorization: Basic <base64_encoded_username:password>
```

### Authentication Endpoints (`/auth`)

*   **User Registration:**
    ```http
    POST /auth/signup
    Content-Type: application/json

    {
        "username": "YourUsername",
        "email": "your.email@example.com",
        "password": "yourStrongPassword"
    }
    ```
*   **User Sign-in:**
    ```http
    POST /auth/signin
    Content-Type: application/json

    {
        "email": "your.email@example.com",
        "password": "yourStrongPassword"
    }
    ```
*   **Get User Profile (Authenticated):**
    ```http
    GET /auth/user
    Cookie: zendo_access_token=yourAccessToken
    ```
*   **Update User Profile (Authenticated):**
    ```http
    PUT /auth/user
    Content-Type: multipart/form-data
    Cookie: zendo_access_token=yourAccessToken

    # Form Data:
    # username: NewUsername
    # avatar: http://example.com/new_avatar.png (or attach avatarFile)
    # avatarFile: <your_image_file>
    ```
*   **Refresh Access Token:**
    ```http
    GET /auth/refresh-token
    Cookie: zendo_session_token=yourRefreshToken
    ```
*   **Verify Email:** (After receiving a magic link in email)
    ```http
    GET /auth/verify_email?token=yourEmailVerificationToken
    Cookie: zendo_access_token=yourAccessToken
    ```
*   **Resend Email Verification Link:**
    ```http
    POST /auth/verify_email
    Cookie: zendo_access_token=yourAccessToken
    ```

### Task Endpoints (`/tasks`)

All task endpoints require authentication via `zendo_access_token` cookie.

*   **Get All Tasks:**
    ```http
    GET /tasks?page=1&limit=10
    ```
*   **Search Tasks:**
    ```http
    GET /tasks/search?search=keyword
    ```
*   **Create a New Task:**
    ```http
    POST /tasks/new
    Content-Type: application/json

    {
        "title": "Build Zendo Frontend",
        "description": "Develop the client-side application for Zendo.",
        "dueDate": "2024-12-31T23:59:59Z",
        "status": "pending",
        "subTasks": [
            {"_id": "sub1", "title": "Design UI", "completed": false},
            {"_id": "sub2", "title": "Implement API integration", "completed": false}
        ],
        "projectId": "optionalProjectId",
        "team_id": "optionalTeamId",
        "assignees": ["optionalAssigneeId1", "optionalAssigneeId2"]
    }
    ```
*   **Get Task by ID:**
    ```http
    GET /tasks/:id
    ```
*   **Update Task:**
    ```http
    PUT /tasks/:id
    Content-Type: application/json

    {
        "title": "Updated Task Title",
        "status": "completed",
        "dueDate": "2025-01-15T10:00:00Z"
    }
    ```
*   **Update Subtask:**
    ```http
    PUT /tasks/:id/subtask/:subTaskId
    Content-Type: application/json

    {
        "completed": true
    }
    ```
*   **Delete Subtask:**
    ```http
    DELETE /tasks/:id/subtask/:subTaskId
    ```
*   **Delete Task by ID:**
    ```http
    DELETE /tasks/:id
    ```
*   **Delete All Tasks:**
    ```http
    DELETE /tasks/all
    ```

### Project Endpoints (`/projects`)

All project endpoints require authentication via `zendo_access_token` cookie.

*   **Get All Projects:**
    ```http
    GET /projects?page=1&limit=10
    ```
*   **Search Projects:**
    ```http
    GET /projects/search?search=keyword
    ```
*   **Create a New Project:**
    ```http
    POST /projects/new
    Content-Type: application/json

    {
        "name": "Personal Goals",
        "description": "Tracking my personal development objectives."
    }
    ```
*   **Get Project by ID:**
    ```http
    GET /projects/:id
    ```
*   **Get Tasks within a Project:**
    ```http
    GET /projects/:id/tasks
    ```
*   **Update Project:**
    ```http
    PUT /projects/:id
    Content-Type: application/json

    {
        "name": "Updated Project Name",
        "description": "Revised project description."
    }
    ```
*   **Delete Project by ID:**
    ```http
    DELETE /projects/:id
    ```
*   **Delete All Projects:**
    ```http
    DELETE /projects/all
    ```

### Team Endpoints (`/teams`)

All team endpoints require authentication via `zendo_access_token` cookie. Team-specific endpoints often require `teamId` in the path and role-based authorization.

*   **Get All Teams for User:**
    ```http
    GET /teams?page=1&limit=10
    ```
*   **Get User's Overall Team Statistics:**
    ```http
    GET /teams/stats
    ```
*   **Create a New Team:**
    ```http
    POST /teams/new
    Content-Type: application/json

    {
        "name": "Team Alpha",
        "description": "Our primary development team."
    }
    ```
*   **Accept Team Invite:** (After receiving a magic link in email)
    ```http
    GET /teams/members/invite?token=yourTeamInviteToken
    ```

*   **Get Team by ID (Team Member Only):**
    ```http
    GET /teams/:teamId
    ```
*   **Get Team Statistics by ID (Team Member Only):**
    ```http
    GET /teams/:teamId/stats
    ```
*   **Update Team (Owner Only):**
    ```http
    PUT /teams/:teamId
    Content-Type: application/json

    {
        "name": "Updated Team Name",
        "description": "New team description."
    }
    ```
*   **Delete Team (Owner Only):**
    ```http
    DELETE /teams/:teamId
    ```
*   **Get Team Invites (Owner Only):**
    ```http
    GET /teams/:teamId/invites
    ```
*   **Cancel Team Invite (Owner Only):**
    ```http
    DELETE /teams/:teamId/invites/:inviteId
    ```

### Team Members Endpoints (`/teams/:teamId/members`)

*   **Get All Team Members (Team Member Only):**
    ```http
    GET /teams/:teamId/members?page=1&limit=10
    ```
*   **Invite Member to Team (Admin/Owner Only):**
    ```http
    POST /teams/:teamId/members/invite
    Content-Type: application/json

    {
        "email": "new.member@example.com",
        "role": "member" # or "admin"
    }
    ```
*   **Remove Team Member (Owner/Admin Only - can remove members/admins, owner can remove anyone including self):**
    ```http
    DELETE /teams/:teamId/members/:memberId
    ```

### Team Task Endpoints (`/teams/:teamId/tasks`)

*   **Get All Tasks for a Team (Team Member Only):**
    ```http
    GET /teams/:teamId/tasks?page=1&limit=10
    ```
*   **Get a Specific Task for a Team (Team Member Only):**
    ```http
    GET /teams/:teamId/tasks/:id
    ```
*   **Delete Team Task (Owner Only):**
    ```http
    DELETE /teams/:teamId/tasks/:taskId
    ```

### Statistics Endpoint (`/stats`)

*   **Get User Statistics (Authenticated):**
    ```http
    GET /stats
    Cookie: zendo_access_token=yourAccessToken
    ```

## ✨ Features

Zendo is packed with features designed for effective task and project management:

*   🔒 **Secure User Authentication**: Full user lifecycle management including sign-up, sign-in, JWT-based authentication (access and refresh tokens), and cookie-based session handling.
*   📧 **Email Verification**: Ensures account security and validity with a magic link email verification system.
*   🌐 **Google OAuth2 Integration**: Seamless sign-in experience using Google accounts (fully implemented but optional via environment variables).
*   ✅ **Comprehensive Task Management**: Create, retrieve, update, and delete tasks. Each task can include a title, description, due date, status, sub-tasks, assignees, and can be linked to projects or teams.
*   🔗 **Sub-task Functionality**: Break down complex tasks into manageable sub-tasks with their own completion status.
*   📅 **Automated Email Reminders**: Utilizes cron jobs to send timely email notifications for tasks nearing their due date, helping users stay on track.
*   📊 **Project Organization**: Create and manage projects, linking tasks to specific projects for better organization and overview. Projects track their total number of associated tasks.
*   👨‍👩‍👧‍👦 **Team Collaboration**: Create teams, invite members with specific roles (owner, admin, member), and manage team-specific tasks and projects.
*   📈 **User Performance Statistics**: Provides a personal dashboard with key metrics such as total tasks, total projects, task completion rate, and tasks due today.
*   ⚡ **High Performance & Scalability**: Built with Go and Fiber for lightweight, fast API responses and high concurrency.
*   🚀 **Intelligent Caching with Redis**: Implements Redis for caching frequently accessed data, significantly improving API response times and reducing database load.
*   🛡️ **Robust Input Validation**: Ensures data integrity and security with comprehensive server-side input validation using `go-playground/validator`.
*   ☁️ **Cloud-based Asset Management**: Integrates with Cloudinary for secure and efficient storage and delivery of user avatars.
*   🗄️ **Transactional Data Operations**: Leverages MongoDB's transactional capabilities for critical operations like deleting all projects (and their associated tasks) or creating tasks within projects, ensuring data consistency.
*   📉 **Prometheus Metrics**: Exposes metrics for monitoring HTTP requests, database operations, Redis interactions, cron jobs, and user/project/task creation counts, enabling detailed observability.

## 🛠️ Technologies Used

The Zendo backend is built with a modern tech stack, ensuring reliability and performance:

| Category          | Technology                                                             | Description                                                                  | Link                                                           |
| :---------------- | :--------------------------------------------------------------------- | :--------------------------------------------------------------------------- | :------------------------------------------------------------- |
| **Language**      | [Go](https://go.dev/)                                                  | A compiled, statically typed programming language.                           | [go.dev](https://go.dev/)                                      |
| **Web Framework** | [Fiber](https://gofiber.io/)                                           | An Express-inspired web framework for Go, built on Fasthttp.                 | [gofiber.io](https://gofiber.io/)                              |
| **Database**      | [MongoDB](https://www.mongodb.com/)                                    | A NoSQL document database for flexible data storage.                         | [mongodb.com](https://www.mongodb.com/)                        |
| **ORM/Driver**    | [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)        | Official MongoDB driver for Go.                                              | [github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver) |
| **Caching**       | [Redis](https://redis.io/)                                             | An in-memory data structure store, used as a cache and message broker.       | [redis.io](https://redis.io/)                                  |
| **Redis Client**  | [go-redis/v9](https://github.com/redis/go-redis)                       | Popular Redis client for Go.                                                 | [github.com/redis/go-redis](https://github.com/redis/go-redis) |
| **Authentication**| [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)                 | A Go package for JSON Web Tokens (JWT).                                      | [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt) |
| **Password Hashing** | [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)       | Cryptographic functions for Go (bcrypt for password hashing).                | [pkg.go.dev/golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) |
| **Scheduling**    | [go-co-op/gocron](https://github.com/go-co-op/gocron)                  | A modern Go library for scheduling jobs.                                     | [github.com/go-co-op/gocron](https://github.com/go-co-op/gocron) |
| **Cloud Storage** | [Cloudinary Go SDK](https://github.com/cloudinary/cloudinary-go/v2)    | Official Cloudinary SDK for Go for media management.                         | [github.com/cloudinary/cloudinary-go](https://github.com/cloudinary/cloudinary-go/v2) |
| **OAuth**         | [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2)          | Go's official OAuth2 client library.                                         | [pkg.go.dev/golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) |
| **Environment**   | [joho/godotenv](https://github.com/joho/godotenv)                      | A Go port of the Ruby dotenv library to load .env files.                     | [github.com/joho/godotenv](https://github.com/joho/godotenv) |
| **Validation**    | [go-playground/validator/v10](https://github.com/go-playground/validator) | Struct and field validation for Go.                                          | [github.com/go-playground/validator](https://github.com/go-playground/validator) |
| **Emailing**      | [gopkg.in/mail.v2](https://github.com/go-mail/mail)                    | A Go package for sending emails.                                             | [github.com/go-mail/mail](https://github.com/go-mail/mail)     |
| **Monitoring**    | [Prometheus](https://prometheus.io/) / [client_golang](https://github.com/prometheus/client_golang) | Open-source monitoring system for metrics collection and alerting.           | [prometheus.io](https://prometheus.io/)                        |

## 🤝 Contributing

I welcome contributions to the Zendo Task Management Backend! Whether you want to fix a bug, add a new feature, or improve documentation, your input is highly appreciated.

Here’s how you can contribute:

*   ✨ **Fork the Repository**: Start by forking the `zendo` repository to your GitHub account.
*   🌿 **Create a New Branch**: Choose a descriptive name for your branch (e.g., `feature/add-reminders`, `bugfix/auth-issue`).
    ```bash
    git checkout -b feature/your-feature-name
    ```
*   💻 **Make Your Changes**: Implement your features or bug fixes. Write clean, maintainable, and well-commented code.
*   🧪 **Write Tests (If applicable)**: Ensure your changes are robust by adding unit or integration tests where appropriate.
*   ✅ **Lint Your Code**: Make sure your code adheres to Go best practices and styling.
    ```bash
    go fmt ./...
    go vet ./...
    ```
*   ⬆️ **Commit Your Changes**: Write clear and concise commit messages.
*   🚀 **Push to Your Fork**:
    ```bash
    git push origin feature/your-feature-name
    ```
*   📬 **Create a Pull Request**: Open a pull request from your forked repository to the `main` branch of the original `zendo` repository. Describe your changes thoroughly and link to any relevant issues.

I'll review your contribution as soon as possible. Thank you for making Zendo better!

## 📄 License

This project does not yet have an explicit license file. Please contact the author for licensing information.

## ✍️ Author

Hi, I'm Onos Ejoor, and I built the Zendo Task Management Backend! I'm passionate about crafting robust and scalable software solutions.

Feel free to connect with me:

*   Website: [https://onos-ejoor.vercel.app](https://onos-ejoor.vercel.app)
*   Twitter: [@DevText16](https://x.com/DevText16)

---

[![Readme was generated by Readmit](https://img.shields.io/badge/Readme%20was%20generated%20by-Readmit-brightred)](https://readmit.vercel.app)
