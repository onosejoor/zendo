# Zendo: A Robust Task Management Platform

Zendo is a full-stack task management application designed to help users efficiently organize their work, track progress, and boost productivity. Built with a powerful Go backend and a responsive Next.js frontend, Zendo offers a seamless experience for managing tasks and projects, from creation to completion. ✨

## 🚀 Installation

To get Zendo up and running on your local machine, follow these steps:

### Prerequisites

Before you start, ensure you have the following installed:

*   **Go**: Version 1.20 or newer
*   **Node.js**: Version 18 or newer
*   **npm** or **Yarn**
*   **MongoDB**: Local instance or access to a cloud-hosted instance (e.g., MongoDB Atlas)
*   **Redis**: Local instance or access to a cloud-hosted instance (e.g., Upstash Redis)

### Clone the Repository

Begin by cloning the Zendo repository:

```bash
git clone https://github.com/onosejoor/zendo.git
cd zendo
```

### Backend Setup (Go)

1.  Navigate to the `server` directory:
    ```bash
    cd server
    ```
2.  Download the Go modules:
    ```bash
    go mod download
    ```
3.  Create a `.env` file in the `server` directory based on the `env.example` provided:
    ```
    PORT=8080
    MONGO_URI="your_mongodb_connection_string"
    REDIS_UPSTASH_URL="your_redis_connection_string"
    JWT_SECRET="your_jwt_secret_key"
    ACCESS_SECRET="your_access_secret_key"
    CLIENT_URL="http://localhost:3000" # Your frontend URL
    ORIGIN="http://localhost:8080" # Your backend URL
    # Optional: Cloudinary for avatar uploads
    CLOUDINARY_URL="cloudinary://api_key:api_secret@cloud_name"
    UPLOAD_PRESET="your_cloudinary_upload_preset"
    # Optional: Google OAuth
    G_CLIENT_ID="your_google_client_id"
    G_CLIENT_SECRET="your_google_client_secret"
    G_REDIRECT="http://localhost:8080/auth/callback" # Your backend redirect URL
    ```
    *Make sure to replace placeholder values with your actual credentials.*
4.  Start the Go server:
    ```bash
    go run main.go
    ```
    The server should now be running on `http://localhost:8080`.

### Frontend Setup (Next.js)

1.  Navigate to the `client` directory (from the project root):
    ```bash
    cd ../client
    ```
2.  Install the Node.js dependencies:
    ```bash
    npm install
    # or
    yarn install
    ```
3.  Create a `.env.local` file in the `client` directory based on the `env.example` provided:
    ```
    NEXT_PUBLIC_SERVER_URL="http://localhost:8080" # Your backend URL
    ```
4.  Start the Next.js development server:
    ```bash
    npm run dev
    # or
    yarn dev
    ```
    The frontend application will be accessible at `http://localhost:3000`.

## 🖥️ Usage

Once the server and client are running, you can start using Zendo:

1.  **Account Creation & Authentication**:
    *   Open your browser and navigate to `http://localhost:3000/auth/signup` to create a new user account.
    *   After successful registration, you'll be redirected to `http://localhost:3000/auth/signin` to log in using your newly created credentials.
2.  **Dashboard Overview**:
    *   Upon successful login, you'll land on your personal dashboard (`/dashboard`), which provides a quick overview of your task statistics, including total tasks, active projects, completion rate, and tasks due today.
3.  **Managing Tasks**:
    *   Navigate to the Tasks section via the sidebar or by visiting `http://localhost:3000/dashboard/tasks`.
    *   Click the "New Task" button to create a new task. You can specify a title, description, due date, status (pending, in-progress, completed), and optionally link it to an existing project.
    *   To modify a task, click on its card in the task list or navigate to its detail page (`/dashboard/tasks/[id]`). Use the "Edit Task" button to make changes.
    *   Easily mark tasks as completed directly from the task list by toggling their checkbox.
    *   Delete individual tasks or clear all your tasks from the settings page.
4.  **Managing Projects**:
    *   Access the Projects section via the sidebar or by visiting `http://localhost:3000/dashboard/projects`.
    *   Create new projects using the "New Project" button. Projects help you group related tasks effectively.
    *   To view all tasks associated with a specific project, click on the project card to navigate to its detail page (`/dashboard/projects/[id]`).
    *   Projects can be edited or deleted (which will also remove all linked tasks) from their respective detail pages.
5.  **User Settings**:
    *   Go to `http://localhost:3000/dashboard/settings` to manage your profile.
    *   Update your username and upload a new avatar to personalize your account.
    *   The "Danger Zone" provides options to permanently delete all your tasks or all your projects. Please use these features with extreme caution as they are irreversible.

## ✨ Features

*   **Secure User Authentication**: Robust user registration and login system with password hashing and JWT-based authentication for secure session management.
*   **Comprehensive Task Management**: Create, view, update, and delete tasks. Each task supports a title, description, due date, and status tracking (pending, in-progress, completed).
*   **Intuitive Project Organization**: Group tasks into distinct projects for streamlined organization and an improved overview of related work.
*   **Dynamic Dashboard Statistics**: A personalized dashboard provides real-time insights into your productivity with key metrics like total tasks, active projects, completion rate, and tasks due today.
*   **Efficient Data Handling**: Leverages MongoDB for persistent data storage and integrates Redis for high-performance data caching, significantly reducing database load and improving response times.
*   **Atomic Operations**: Ensures data integrity and consistency through MongoDB transactions for complex operations, such as deleting a project along with all its associated tasks.
*   **Responsive User Interface**: A clean, modern, and highly responsive UI built with Next.js and Tailwind CSS, ensuring a smooth and consistent experience across various devices and screen sizes.
*   **Cloud-Based Image Uploads**: Seamlessly upload and manage user avatars with integration of Cloudinary for efficient media storage and delivery.

## 🛠️ Technologies Used

| Category      | Technology    | Link                                                 |
| :------------ | :------------ | :--------------------------------------------------- |
| **Backend**   | Go            | [golang.org](https://golang.org/)                    |
|               | Fiber         | [gofiber.io](https://gofiber.io/)                    |
|               | MongoDB       | [mongodb.com](https://www.mongodb.com/)              |
|               | Redis         | [redis.io](https://redis.io/)                        |
|               | JWT           | [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt) |
|               | Cloudinary    | [cloudinary.com](https://cloudinary.com/)            |
|               | go-validator  | [github.com/go-playground/validator](https://github.com/go-playground/validator) |
| **Frontend**  | Next.js       | [nextjs.org](https://nextjs.org/)                    |
|               | React         | [react.dev](https://react.dev/)                      |
|               | TypeScript    | [typescriptlang.org](https://www.typescriptlang.org/) |
|               | Tailwind CSS  | [tailwindcss.com](https://tailwindcss.com/)          |
|               | Shadcn UI     | [ui.shadcn.com](https://ui.shadcn.com/)              |
|               | SWR           | [swr.vercel.app](https://swr.vercel.app/)            |
|               | Axios         | [axios-http.com](https://axios-http.com/)            |
| **Deployment**| Vercel        | [vercel.com](https://vercel.com/)                    |

## 🤝 Contributing

We welcome contributions to Zendo! If you're interested in improving this project, please follow these guidelines:

*   🍴 **Fork the repository**: Start by forking the Zendo repository to your GitHub account.
*   🌲 **Clone locally**: `git clone https://github.com/your-username/zendo.git`
*   🌿 **Create a new branch**: For each new feature or bug fix, create a new branch from `main`: `git checkout -b feature/your-feature-name` or `git checkout -b bugfix/issue-description`.
*   💻 **Make your changes**: Implement your new features, fix bugs, or improve existing code. Ensure your changes align with the project's goals and existing architecture.
*   🧪 **Test your changes**: Before submitting, thoroughly test your changes to ensure they work as expected and do not introduce any regressions.
*   💬 **Commit your changes**: Write clear, concise, and descriptive commit messages.
*   ⬆️ **Push to your fork**: Push your new branch to your forked repository: `git push origin feature/your-feature-name`.
*   ✉️ **Open a Pull Request**: Finally, open a pull request to the `main` branch of the original Zendo repository. Clearly describe the purpose of your changes and any relevant details.

## 📄 License

This project is open source. While a dedicated `LICENSE` file is not included in this repository, the code is made available for review and personal use.

## ✍️ Author Info

**Onos Ejoor**

*   Portfolio: [https://onos-ejoor.vercel.app](https://onos-ejoor.vercel.app)
*   GitHub: [@onosejoor](https://github.com/onosejoor)
*   LinkedIn: [Your LinkedIn Profile](https://linkedin.com/in/your-profile)
*   Twitter: [@your_twitter](https://twitter.com/your_twitter)

---
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=next.js&logoColor=white)
![React](https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black)
![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=fiber&logoColor=white)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)