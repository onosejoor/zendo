# Zendo: Seamless Task & Project Management ✅

Zendo is a full-stack task and project management application designed to boost your productivity and keep your work organized. This platform provides an intuitive interface for managing individual tasks, grouping them into projects, and tracking your progress, all while offering a robust, high-performance backend. It's built to simplify your workflow and help you achieve your goals effortlessly.

## 🚀 Getting Started

Follow these steps to get Zendo up and running on your local machine.

### 📋 Prerequisites

Before you begin, ensure you have the following installed:

*   **Node.js** (LTS version recommended)
*   **Go** (version 1.22 or higher)
*   **MongoDB**
*   **Redis**

### 📦 Installation

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/onosejoor/zendo.git
    cd zendo
    ```

2.  **Backend Setup (Go Fiber)**:
    *   Navigate to the server directory:
        ```bash
        cd server
        ```
    *   Install Go dependencies:
        ```bash
        go mod tidy
        ```
    *   Create a `.env` file in the `server` directory based on `env.example`:
        ```
        PORT=8080
        CLIENT_URL=http://localhost:3000
        ORIGIN=http://localhost:8080
        MONGO_URI="your_mongodb_connection_string"
        JWT_SECRET="your_jwt_secret"
        ACCESS_SECRET="your_access_token_secret"
        REDIS_UPSTASH_URL="your_redis_connection_string"
        # Optional Google OAuth settings (uncomment in main.go if using)
        # G_CLIENT_ID=""
        # G_CLIENT_SECRET=""
        # G_REDIRECT="http://localhost:8080/auth/callback"
        # FRONTEND_URL="http://localhost:3000"
        ```
        *Replace placeholder values with your actual credentials.*
    *   Run the Go server:
        ```bash
        go run main.go
        ```
        The server will start on `http://localhost:8080` (or your specified `PORT`).

3.  **Frontend Setup (Next.js)**:
    *   Navigate to the client directory:
        ```bash
        cd ../client
        ```
    *   Install Node.js dependencies:
        ```bash
        npm install
        ```
    *   Create a `.env.local` file in the `client` directory based on `env.example`:
        ```
        NEXT_PUBLIC_SERVER_URL="http://localhost:8080"
        ```
    *   Start the Next.js development server:
        ```bash
        npm run dev
        ```
        The frontend will be available at `http://localhost:3000`.

## 🖥️ Usage

Once both the backend and frontend are running, open your browser and navigate to `http://localhost:3000`.

### Account Management
*   **Sign Up**: Create a new Zendo account with your email, username, and password.
*   **Sign In**: Log in to your existing account.
*   **Session Management**: The application handles secure session management with access and refresh tokens, ensuring your data remains protected and sessions persist.

### Dashboard Overview
Upon logging in, you'll land on a personalized dashboard:
*   **Quick Stats**: View your total tasks, active projects, overall completion rate, and tasks due today at a glance.
*   **Recent Activities**: See your most recent tasks and active projects, providing a quick overview of your current workload.
*   **Quick Actions**: Easily create new tasks or projects directly from the dashboard.

### Task Management
*   **View All Tasks**: Navigate to the "Tasks" section to see a comprehensive list of all your tasks.
*   **Create Task**: Add new tasks with a title, description, due date, status (pending, in-progress, completed), and optionally associate them with a project.
*   **View Task Details**: Click on any task to see its full details.
*   **Edit Task**: Update task details, change status, or modify the due date.
*   **Delete Task**: Remove tasks you no longer need.
*   **Status Indicators**: Tasks are visually cued by their status (completed, in-progress, pending) and display an "Expired" badge if overdue.

### Project Management
*   **View All Projects**: Visit the "Projects" section to see all your organized projects.
*   **Create Project**: Start a new project by giving it a name and an optional description.
*   **View Project Details**: Click on a project to view all associated tasks.
*   **Tasks within Projects**: Add, view, edit, and delete tasks directly from within a project's dedicated page.
*   **Edit Project**: Modify project names or descriptions.
*   **Delete Project**: Remove projects, which will also cascade and delete all associated tasks.

## ✨ Features

*   **User Authentication**: Secure signup and sign-in with robust token-based authentication.
*   **Task Management**: Create, read, update, and delete tasks with due dates and statuses.
*   **Project Organization**: Group tasks into projects for better structuring and overview.
*   **Dashboard Insights**: A dynamic dashboard displaying key metrics like total tasks, projects, completion rate, and tasks due today.
*   **Responsive Design**: A user-friendly interface that adapts seamlessly to various screen sizes.
*   **Data Caching**: Utilizes Redis for efficient data retrieval and reduced database load.
*   **Error Handling**: Comprehensive error displays for a smoother user experience.

## 🛠️ Technologies Used

| Category   | Technology                                                | Description                                                               |
| :--------- | :-------------------------------------------------------- | :------------------------------------------------------------------------ |
| **Backend** | [Go](https://go.dev/)                                     | The efficient and performant programming language for the server.         |
|            | [Fiber](https://gofiber.io/)                              | An expressive and fast HTTP framework for Go.                             |
|            | [MongoDB](https://www.mongodb.com/)                       | A NoSQL database for flexible data storage.                               |
|            | [Redis](https://redis.io/)                                | An in-memory data store for caching and session management.               |
|            | [Go-JWT](https://github.com/golang-jwt/jwt)               | JSON Web Tokens for secure authentication.                                |
|            | [Go-Validator](https://github.com/go-playground/validator) | Struct validation for incoming request bodies.                            |
| **Frontend** | [Next.js](https://nextjs.org/) (v15)                      | A React framework for building server-rendered and static web applications. |
|            | [React](https://react.dev/) (v19)                         | A JavaScript library for building user interfaces.                        |
|            | [TypeScript](https://www.typescriptlang.org/)             | A typed superset of JavaScript that compiles to plain JavaScript.         |
|            | [Tailwind CSS](https://tailwindcss.com/)                  | A utility-first CSS framework for rapid UI development.                   |
|            | [Shadcn UI](https://ui.shadcn.com/)                       | A collection of reusable components built with Radix UI and Tailwind CSS. |
|            | [SWR](https://swr.vercel.app/)                            | React Hooks for Data Fetching.                                            |
|            | [Axios](https://axios-http.com/)                          | Promise-based HTTP client for the browser and Node.js.                    |
| **Dev Tools** | [ESLint](https://eslint.org/)                             | For code linting and maintaining consistent code style.                   |
|            | [Prettier](https://prettier.io/)                          | An opinionated code formatter.                                            |

## 🤝 Contributing

We welcome contributions to Zendo! If you're looking to help out, here’s how you can get started:

*   ⭐ **Fork the repository** on GitHub.
*   🌳 **Clone your forked repository** locally.
*   🌿 **Create a new branch** for your feature or bug fix: `git checkout -b feature/your-feature-name`.
*   ✍️ **Make your changes** and ensure they align with the project's coding style.
*   ✅ **Test your changes thoroughly** to prevent regressions.
*   ➕ **Commit your changes** with clear and concise messages: `git commit -m "feat: Add new awesome feature"`.
*   ⬆️ **Push your branch** to your forked repository: `git push origin feature/your-feature-name`.
*   🔄 **Open a Pull Request** against the `main` branch of the original repository, describing your changes in detail.

## 📄 License

This project is licensed under the MIT License.

## 👤 Author

**Onos Ejoor**

*   Website: [https://onos-ejoor.vercel.app](https://onos-ejoor.vercel.app)
*   GitHub: [@onosejoor](https://github.com/onosejoor)
<!-- *   LinkedIn: [your-linkedin-username](https://www.linkedin.com/in/onosejoor) -->

---

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)