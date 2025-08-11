# Zendo: Streamlined Task & Project Management

Organize your work, track progress, and boost productivity with Zendo, a powerful and intuitive task management application. Built with a robust Go backend and a modern Next.js frontend, Zendo provides a seamless experience for individuals and teams to manage tasks and projects efficiently.

---

## 🚀 Getting Started

Follow these steps to set up and run Zendo on your local machine.

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
APP_PASSWORD="your_gmail_app_password"
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

## ✨ Features

- **User Authentication**: Secure user registration, login, and session management using JWT.
- **Email Verification**: Ensures account security and authenticity through email verification links.
- **Task Management**:
  - Create, view, update, and delete tasks.
  - Add and manage subtasks within larger tasks.
  - Set due dates and track task status (pending, in-progress, completed).
  - Receive automated email reminders for upcoming task due dates.
- **Project Management**:
  - Organize tasks into projects for better structuring.
  - View all tasks associated with a specific project.
  - Create, update, and delete projects.
- **Dashboard & Analytics**:
  - Overview of total tasks, total projects, completion rate, and tasks due today.
- **User Profile Management**: Update username and avatar.
- **Data Caching**: Utilizes Redis for efficient data retrieval and improved performance.
- **Atomic Operations**: Employs MongoDB transactions for reliable data consistency (e.g., when creating/deleting tasks within projects).

---

## 🛠️ Technologies Used

| Category     | Technology                                                                                                              | Purpose                                                                       |
| :----------- | :---------------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------- |
| **Backend**  | ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)                               | Primary server-side language for robust API development.                      |
|              | ![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=fiber&logoColor=white)                      | Fast, Express-inspired web framework for Go.                                  |
|              | ![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)                | NoSQL database for flexible and scalable data storage.                        |
|              | ![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)                      | In-memory data store used for caching API responses.                          |
|              | ![GoCron](https://img.shields.io/badge/GoCron-018786?style=for-the-badge&logo=go&logoColor=white)                       | Library for scheduling recurring tasks (e.g., email reminders).               |
|              | ![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white)                | JSON Web Tokens for secure authentication and authorization.                  |
|              | ![GoDotEnv](https://img.shields.io/badge/DotEnv-ECD53F?style=for-the-badge&logo=dot-env&logoColor=white)                | Loads environment variables from `.env` files.                                |
| **Frontend** | ![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=next.js&logoColor=white)                | React framework for server-side rendering and static site generation.         |
|              | ![React](https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black)                      | JavaScript library for building interactive user interfaces.                  |
|              | ![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)       | Superset of JavaScript for type safety and improved developer experience.     |
|              | ![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwind-css&logoColor=white) | Utility-first CSS framework for rapid UI development.                         |
|              | ![SWR](https://img.shields.io/badge/SWR-000000?style=for-the-badge&logo=swr&logoColor=white)                            | React Hooks for data fetching with caching, revalidation, and focus-on-mount. |
|              | ![Shadcn/ui](https://img.shields.io/badge/shadcn%2Fui-000000?style=for-the-badge&logo=data-transfer&logoColor=white)    | Reusable UI components for consistent design.                                 |
| **Cloud**    | ![Cloudinary](https://img.shields.io/badge/Cloudinary-3448C5?style=for-the-badge&logo=cloudinary&logoColor=white)       | Cloud-based image and video management for avatar uploads.                    |

---

# Zendo API

## Overview

The Zendo API is a high-performance backend service built with **Go** and the **Fiber** web framework. It leverages **MongoDB** for persistent data storage and **Redis** for efficient caching, providing a robust and scalable solution for task and project management.

## Features

- **Go**: High-performance, concurrent API operations.
- **Fiber**: Fast and minimal web framework for routing and middleware.
- **MongoDB**: NoSQL database for flexible data models (users, tasks, projects).
- **Redis**: In-memory data store for caching frequently accessed data (user profiles, task lists, project lists, and statistics).
- **JWT**: Secure authentication and authorization for all protected endpoints.
- **Bcrypt**: Password hashing for user security.
- **GoCron**: Scheduled jobs for sending email reminders for tasks approaching their due dates.
- **Cloudinary**: External service for managing user avatar uploads.
- **Transactions**: Ensures data consistency for critical operations (e.g., linking tasks to projects, deleting all projects/tasks).

## Getting Started

### Installation

Navigate to the `server` directory:

```bash
cd server
go mod tidy
```

This command resolves and downloads all necessary Go modules.

### Environment Variables

All required environment variables for the server should be set in a `.env` file in the `server` directory.

| Variable Name       | Example Value                                            | Description                                                                          |
| :------------------ | :------------------------------------------------------- | :----------------------------------------------------------------------------------- |
| `PORT`              | `8080`                                                   | The port on which the Go server will listen.                                         |
| `DATABASE`          | `zendo_db`                                               | The name of the MongoDB database to connect to.                                      |
| `MONGO_URI`         | `mongodb+srv://user:pass@cluster.mongodb.net/dbname?...` | Full MongoDB connection URI.                                                         |
| `REDIS_UPSTASH_URL` | `rediss://user:pass@host:port`                           | Connection URL for the Redis server (e.g., Upstash Redis).                           |
| `ACCESS_SECRET`     | `very_secure_access_token_secret_key`                    | Secret key for signing and verifying JWT access tokens.                              |
| `JWT_SECRET`        | `super_secure_refresh_token_secret_key`                  | Secret key for signing and verifying JWT refresh tokens and email magic links.       |
| `EMAIL`             | `your_email@gmail.com`                                   | The email address used for sending automated emails (e.g., reminders, verification). |
| `APP_PASSWORD`      | `your_gmail_app_password`                                | The app password generated for the `EMAIL` account (for Gmail SMTP).                 |
| `FRONTEND_URL`      | `http://localhost:3000`                                  | The base URL of the frontend application.                                            |
| `CLIENT_URL`        | `http://localhost:3000`                                  | Allowed CORS origin for the frontend application.                                    |
| `ENVIRONMENT`       | `development` or `production`                            | Sets the application environment, affecting secure cookie settings.                  |
| `CLOUDINARY_URL`    | `cloudinary://api_key:api_secret@cloud_name`             | Cloudinary API URL for image uploads.                                                |
| `UPLOAD_PRESET`     | `your_cloudinary_upload_preset_name`                     | Cloudinary upload preset for secure image uploads.                                   |
| `G_CLIENT_ID`       | `your-google-oauth-client-id.apps.googleusercontent.com` | Google OAuth Client ID for authentication (if enabled).                              |
| `G_CLIENT_SECRET`   | `your-google-oauth-client-secret`                        | Google OAuth Client Secret for authentication (if enabled).                          |
| `G_REDIRECT`        | `http://localhost:8080/auth/callback`                    | Google OAuth Redirect URL.                                                           |

## API Documentation

### Base URL

The API base URL is dependent on your environment configuration.

- **Local Development**: `http://localhost:8080`
- **Production**: Refer to your deployed server's base URL.

### Endpoints

#### GET /health

Checks the health status of the API.
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "message": "Healthy"
}
```

**Errors**:

- `500 Internal Server Error`: If there's an issue with the server's internal health check logic.

#### GET /stats

Retrieves statistics for the authenticated user, including total tasks, total projects, completion rate, and tasks due today.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "stats": {
    "total_tasks": 10,
    "total_projects": 3,
    "completion_rate": 75,
    "completed_tasks": 5,
    "dueToday": 2
  }
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error when fetching stats.

#### GET /auth/user

Retrieves the profile information of the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "user": {
    "_id": "65f04a3e74b6e8d1a1d9b0c2",
    "email": "user@example.com",
    "username": "testuser",
    "avatar": "https://res.cloudinary.com/example/image/upload/v123456789/avatar.jpg",
    "email_verified": true,
    "created_at": "2024-03-12T10:00:00Z"
  }
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the user does not exist.
- `500 Internal Server Error`: If there's a database error.

#### PUT /auth/user

Updates the profile information for the authenticated user. Supports updating username and avatar.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
`Content-Type: multipart/form-data`

```
Form Data:
username: "newusername"
avatarFile: <file> (optional)
avatar: "http://example.com/existing_avatar.png" (optional, if no new file)
```

**Response**:

```json
{
  "success": true,
  "message": "user updated succesfully"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's an error updating the user in the database or uploading the avatar.

#### GET /auth/verify_email

Verifies a user's email using a magic link token.
**Request**:
Query Parameters:

- `token`: The verification token received in the email.

```
GET /auth/verify_email?token=your_verification_token
```

**Response**:

```json
{
  "success": true,
  "message": "User verified successfully"
}
```

**Errors**:

- `500 Internal Server Error`: If the token is invalid, expired, or a database error occurs during verification.
- `404 Not Found`: If the user associated with the token is already verified or not found.

#### POST /auth/verify_email

Sends a new email verification link to the authenticated user's email address.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "message": "Verification Link Sent to user@example.com"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `409 Conflict`: If the user is already verified.
- `500 Internal Server Error`: If there's an error generating the token or sending the email.

#### GET /auth/refresh-token

Refreshes the user's access token using the refresh token stored in the `zendo_session_token` cookie.
**Request**:
No payload required. Refresh token is taken from `zendo_session_token` cookie.

**Response**:
(Sets `zendo_access_token` cookie)

```json
{
  "success": true,
  "message": "Access token created"
}
```

**Errors**:

- `401 Unauthorized`: If no refresh token is provided or it's invalid/expired.

#### POST /auth/signup

Registers a new user account.
**Request**:

```json
{
  "email": "newuser@example.com",
  "username": "newuser",
  "password": "StrongPassword123"
}
```

**Response**:

```json
{
  "success": true,
  "message": "Verification Link Sent To Email",
  "email_sent": true
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or validation fails (e.g., missing fields, invalid email format).
- `409 Conflict`: If a user with the provided email already exists.
- `500 Internal Server Error`: If there's an error creating the user or sending the email verification link.

#### POST /auth/signin

Authenticates a user and creates a session (sets access and refresh token cookies).
**Request**:

```json
{
  "email": "user@example.com",
  "password": "StrongPassword123"
}
```

**Response**:
(Sets `zendo_session_token` and `zendo_access_token` cookies)

```json
{
  "success": true,
  "message": "Welcome testuser",
  "email_verified": true
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or validation fails (e.g., missing fields).
- `400 Bad Request`: If credentials are incorrect.
- `404 Not Found`: If the user is not found.
- `500 Internal Server Error`: If there's a database error or an error creating the session.

#### GET /tasks

Retrieves all tasks for the authenticated user, sorted by creation date.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "tasks": [
    {
      "_id": "65f04a3e74b6e8d1a1d9b0c2",
      "title": "Complete Zendo API Documentation",
      "description": "Document all endpoints, request/response schemas, and error handling.",
      "userId": "65e8a9f60d7b3c2e1f4a9b8c",
      "subTasks": [],
      "projectId": "",
      "dueDate": "2024-04-30T23:59:00Z",
      "status": "in-progress",
      "created_at": "2024-03-12T10:00:00Z"
    }
  ]
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error.

#### GET /tasks/:id

Retrieves a specific task by its ID for the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the task.

**Response**:

```json
{
  "success": true,
  "task": {
    "_id": "65f04a3e74b6e8d1a1d9b0c2",
    "title": "Complete Zendo API Documentation",
    "description": "Document all endpoints, request/response schemas, and error handling.",
    "userId": "65e8a9f60d7b3c2e1f4a9b8c",
    "subTasks": [],
    "projectId": "",
    "dueDate": "2024-04-30T23:59:00Z",
    "status": "in-progress",
    "created_at": "2024-03-12T10:00:00Z"
  }
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the task is not found.
- `500 Internal Server Error`: If there's a database error.

#### POST /tasks/new

Creates a new task for the authenticated user. Supports linking to a project and adding subtasks.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:

```json
{
  "title": "New Task Title",
  "description": "Detailed description of the new task.",
  "subTasks": [
    {
      "_id": "random_uuid_string_1",
      "title": "Subtask 1",
      "completed": false
    }
  ],
  "projectId": "65f04a3e74b6e8d1a1d9b0c2", # Optional, can be empty string
  "dueDate": "2024-05-15T17:00:00Z",
  "status": "pending" # or "in-progress"
}
```

**Response**:

```json
{
  "success": true,
  "message": "Task created successfully",
  "taskId": "65f04a3e74b6e8d1a1d9b0c3"
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or validation fails (e.g., missing title, due date).
- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error, or an issue with transaction when creating a project-linked task.

#### PUT /tasks/:id

Updates an existing task for the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the task to update.
  Payload (example, all fields optional for update):

```json
{
  "title": "Updated Task Title",
  "description": "New description for the task.",
  "dueDate": "2024-06-01T09:00:00Z",
  "status": "completed",
  "subTasks": [
    {
      "_id": "existing_subtask_id",
      "title": "Updated Subtask 1",
      "completed": true
    },
    {
      "_id": "new_subtask_id",
      "title": "Newly Added Subtask",
      "completed": false
    }
  ]
}
```

**Response**:

```json
{
  "success": true,
  "message": "Task updated"
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or task ID is malformed.
- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the task is not found.
- `500 Internal Server Error`: If there's a database error.

#### PUT /tasks/:id/subtask/:subTaskId

Updates the completion status of a specific subtask within a task.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the parent task.
- `subTaskId`: The ID of the subtask to update.
  Payload:

```json
{
  "completed": true
}
```

**Response**:

```json
{
  "success": true,
  "message": "SubTask updated"
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or IDs are malformed.
- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error.

#### DELETE /tasks/:id/subtask/:subTaskId

Deletes a specific subtask from a task.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the parent task.
- `subTaskId`: The ID of the subtask to delete.

**Response**:

```json
{
  "success": true,
  "message": "SubTask deleted"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the subtask is not found.
- `500 Internal Server Error`: If there's a database error.

#### DELETE /tasks/all

Deletes all tasks for the authenticated user, including removing task counts from associated projects.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "message": "Task deleted"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error during the transaction.

#### DELETE /tasks/:id

Deletes a specific task by its ID for the authenticated user. If the task is linked to a project, it also decrements the project's task count.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the task to delete.

**Response**:

```json
{
  "success": true,
  "message": "task deleted"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the task is not found.
- `500 Internal Server Error`: If there's a database error or an issue with the transaction.

#### GET /projects

Retrieves all projects for the authenticated user, sorted by creation date.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "projects": [
    {
      "_id": "65f04a3e74b6e8d1a1d9b0c2",
      "name": "Project Alpha",
      "description": "First major project.",
      "ownerId": "65e8a9f60d7b3c2e1f4a9b8c",
      "totalTasks": 5,
      "created_at": "2024-03-01T08:00:00Z"
    }
  ]
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error.

#### GET /projects/:id/tasks

Retrieves all tasks associated with a specific project for the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the project.

**Response**:

```json
{
  "success": true,
  "tasks": [
    {
      "_id": "65f04a3e74b6e8d1a1d9b0c3",
      "title": "Task A in Project Alpha",
      "description": "Details for task A.",
      "userId": "65e8a9f60d7b3c2e1f4a9b8c",
      "subTasks": [],
      "projectId": "65f04a3e74b6e8d1a1d9b0c2",
      "dueDate": "2024-04-15T12:00:00Z",
      "status": "pending",
      "created_at": "2024-03-05T09:00:00Z"
    }
  ]
}
```

**Errors**:

- `400 Bad Request`: If the project ID is malformed.
- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error.

#### GET /projects/:id

Retrieves a specific project by its ID for the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the project.

**Response**:

```json
{
  "success": true,
  "project": {
    "_id": "65f04a3e74b6e8d1a1d9b0c2",
    "name": "Project Alpha",
    "description": "First major project.",
    "ownerId": "65e8a9f60d7b3c2e1f4a9b8c",
    "totalTasks": 5,
    "created_at": "2024-03-01T08:00:00Z"
  }
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the project is not found.
- `500 Internal Server Error`: If there's a database error.

#### POST /projects/new

Creates a new project for the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:

```json
{
  "name": "My New Project",
  "description": "This project is about building something awesome."
}
```

**Response**:

```json
{
  "success": true,
  "message": "project created successfully",
  "projectId": "65f04a3e74b6e8d1a1d9b0c4"
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or validation fails (e.g., missing name).
- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error.

#### PUT /projects/:id

Updates an existing project for the authenticated user.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the project to update.
  Payload (example, both fields optional for update):

```json
{
  "name": "Updated Project Name",
  "description": "A revised description for the project."
}
```

**Response**:

```json
{
  "success": true,
  "message": "Project updated"
}
```

**Errors**:

- `400 Bad Request`: If the request body is invalid or project ID is malformed.
- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the project is not found.
- `500 Internal Server Error`: If there's a database error.

#### DELETE /projects/all

Deletes all projects for the authenticated user, including all tasks within those projects, using a database transaction.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
No payload required.

**Response**:

```json
{
  "success": true,
  "message": "Deleted 10 tasks and 3 project(s)"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `500 Internal Server Error`: If there's a database error during the transaction.

#### DELETE /projects/:id

Deletes a specific project by its ID for the authenticated user, also deleting all associated tasks within it.
**Authentication**: Required (via `zendo_access_token` cookie)
**Request**:
URL Parameters:

- `id`: The ID of the project to delete.

**Response**:

```json
{
  "success": true,
  "message": "Project and related tasks deleted"
}
```

**Errors**:

- `401 Unauthorized`: If no access token is provided or it's invalid.
- `404 Not Found`: If the project is not found.
- `500 Internal Server Error`: If there's a database error.

---

## 🤝 Contributing

Contributions are welcome! If you have suggestions for improvements or want to report an issue, please:

- 🐛 **Report Bugs**: Open an issue on GitHub describing the bug and steps to reproduce it.
- 💡 **Suggest Features**: Open an issue to propose new features or enhancements.
- 💻 **Submit Pull Requests**: Fork the repository, create a new branch, commit your changes, and open a pull request. Please ensure your code adheres to the existing style and passes all tests.

---

## 📝 License

No specific license file found in the repository.

---

## 👨‍💻 Author Info

**Onos Ejoor**
A passionate software developer dedicated to building robust and user-friendly applications.

- 🌐 [Portfolio](https://onos-ejoor.vercel.app/)

---

### Badges

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=fiber&logoColor=white)](https://gofiber.io/)
[![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)](https://www.mongodb.com/)
[![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=next.js&logoColor=white)](https://nextjs.org/)
[![React](https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black)](https://react.dev/)
[![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwind-css&logoColor=white)](https://tailwindcss.com/)
[![Cloudinary](https://img.shields.io/badge/Cloudinary-3448C5?style=for-the-badge&logo=cloudinary&logoColor=white)](https://cloudinary.com/)
[![SWR](https://img.shields.io/badge/SWR-000000?style=for-the-badge&logo=swr&logoColor=white)](https://swr.vercel.app/)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
