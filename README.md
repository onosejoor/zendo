# **Zendo: Streamlined Task Management**

<!-- --- -->
<!-- 
**A quick note for recruiters**: While this request specified an "Objective-C project," the provided file structure and code context clearly indicate a full-stack application built with **Go (Golang)** for the backend and **Next.js (React/TypeScript)** for the frontend. This README has been generated based on the actual technologies detected in the project files to provide an accurate and detailed overview. -->

---

Zendo is a robust, full-stack task management application designed to boost your productivity and streamline your workflow. Whether you're juggling personal tasks or collaborating on large projects, Zendo provides an intuitive platform to organize, track, and complete your objectives with ease. Built with modern web technologies, Zendo offers a seamless and efficient experience for managing your daily to-dos and ambitious goals.

## 🚀 Installation

Ready to get Zendo up and running on your local machine? Follow these straightforward steps:

### 1. 📂 Clone the Repository

First, grab the project files from GitHub:

```bash
git clone https://github.com/onosejoor/zendo.git
cd zendo
```

### 2. ⚙️ Backend Setup (Go)

Navigate into the `server` directory and prepare your Go environment:

👉 **Dependencies**:
Ensure you have Go installed (version 1.24.2 or higher, as indicated in `go.mod`).

⚙️ **Environment Variables**:
Create a `.env` file in the `server` directory. Populate it with the necessary variables. Here's an example of what your `.env` file might look like, based on the project's code:

```dotenv
PORT=8080
MONGO_URI="mongodb://localhost:27017" # Your MongoDB connection string
DATABASE="zendo_db" # Your MongoDB database name
JWT_SECRET="your_strong_refresh_token_secret" # Secret for refresh tokens
ACCESS_SECRET="your_strong_access_token_secret" # Secret for access tokens
EMAIL_SECRET="your_strong_email_token_secret" # Secret for email verification tokens
REDIS_UPSTASH_URL="redis://default:YOUR_REDIS_PASSWORD@YOUR_REDIS_HOST:PORT" # Your Redis connection URL (e.g., from Upstash)
EMAIL="your_email@gmail.com" # Email address for sending notifications (e.g., Gmail account)
APP_PASSWORD="your_gmail_app_password" # App password for the email account (generate this in Google account settings)
FRONTEND_URL="http://localhost:3000" # The URL where your frontend is hosted
CLIENT_URL="http://localhost:3000" # Used in CORS configuration (often same as FRONTEND_URL)
ORIGIN="http://localhost:8080" # The server's public URL, used for health checks
CLOUDINARY_URL="cloudinary://API_KEY:API_SECRET@CLOUD_NAME" # Your Cloudinary API environment variable
UPLOAD_PRESET="your_cloudinary_upload_preset" # Cloudinary upload preset name
```

📦 **Install Dependencies**:
From the `server` directory, fetch the Go modules:

```bash
cd server
go mod tidy
```

🚀 **Run the Server**:
Start the Go backend application:

```bash
go run main.go
```
The server should now be listening on the port you specified in your `.env` file (defaulting to 8080).

### 3. 🖥️ Frontend Setup (Next.js)

Open a *new* terminal, navigate into the `client` directory, and prepare the frontend:

⚙️ **Environment Variables**:
Create an `.env.local` file in the `client` directory. This project uses `NEXT_PUBLIC_SERVER_URL` for API routing. Make sure it points to your running backend:

```dotenv
NEXT_PUBLIC_SERVER_URL="http://localhost:8080"
```

📦 **Install Dependencies**:
From the `client` directory, install the Node.js packages:

```bash
cd client
npm install # or yarn install / pnpm install
```

🚀 **Run the Frontend**:
Launch the Next.js development server:

```bash
npm run dev # or yarn dev / pnpm dev
```
The frontend application should now be accessible in your web browser at `http://localhost:3000` (or another port if configured).

## 💡 Usage

Once both the backend and frontend are successfully running, Zendo is ready for you to explore and use!

1.  **Access the Application**:
    Open your favorite web browser and navigate to `http://localhost:3000` (or the address where your frontend is serving).

2.  **Sign Up / Sign In**:
    On the landing page, you'll find clear options to either "Get Started" (for new users to sign up) or "Login" (for existing users).
    *   **Sign Up**: Create your new Zendo account by providing your email, a unique username, and a secure password. An important step here is the email verification: a link will be sent to your registered email address. Verifying your email is crucial to unlock full application features, including automated task reminders.
    *   **Sign In**: If you're a returning user, simply enter your registered email and password to access your account.

3.  **Dashboard Overview**:
    Upon successful login, you'll be directed to your personalized dashboard. This intuitive interface provides a quick, high-level overview of your productivity metrics: total tasks, active projects, your overall task completion rate, and any tasks that are due today. It's designed to give you an immediate snapshot of your most pressing priorities.

4.  **Effective Task Management**:
    *   **Create New Task**: Whether from the dashboard's "New Task" button or the dedicated "Tasks" page in the sidebar, you can easily create new tasks. Each task allows you to specify a title, detailed description, due date, current status (pending, in-progress, completed), and an optional association with an existing project. For more complex items, you can even add subtasks to break down your work further.
    *   **View All Tasks**: Navigate to the "Tasks" section via the left sidebar to see a comprehensive list of all your tasks. The interface includes robust search and filter capabilities, enabling you to quickly locate specific tasks.
    *   **Edit Task Details**: Click on any task card or row to access its dedicated detail page. From here, you have the flexibility to modify any aspect of the task, including its title, description, status, and due date. You can also manage (add, edit, or remove) its subtasks directly on this page.
    *   **Mark as Complete**: Tasks can be effortlessly marked as "completed" by checking the corresponding checkbox directly from the dashboard overview or the main tasks list, providing immediate visual feedback on your progress.
    *   **Manage Subtasks**: For tasks requiring more granular steps, you can define multiple subtasks. Each subtask can be individually toggled to mark its completion, providing a clear breakdown of progress for complex objectives.

5.  **Organized Project Management**:
    *   **Create New Project**: Initiate new projects using the "New Project" button on your dashboard or by visiting the "Projects" page. Simply provide a project name and an optional description to get started.
    *   **Browse Projects**: The "Projects" section in the sidebar gives you a centralized view of all your organized projects.
    *   **Project-Specific Tasks**: Clicking on any project card will lead you to a dedicated page showcasing all tasks that are associated with that particular project. This view also allows you to add new tasks directly to the selected project, keeping everything neatly categorized.
    *   **Edit Project Details**: On the individual project's detail page, you can update its name or description as your project evolves.

6.  **Account Settings**:
    The "Settings" page offers comprehensive control over your Zendo account:
    *   **Profile Customization**: Update your personal information, including your username and the ability to upload a custom avatar to personalize your profile.
    *   **Danger Zone**: This section contains irreversible actions and should be used with extreme caution. Here, you have the option to permanently delete *all* your tasks or *all* your projects (which also deletes all associated tasks). Please be absolutely certain before proceeding with these actions.

## ✨ Features

Zendo is built with a focus on delivering a powerful yet user-friendly experience:

*   🚀 **Intuitive Task & Project Management**: Effortlessly create, view, update, and delete tasks and projects. Tasks support detailed descriptions, due dates, and customizable statuses.
*   ✅ **Granular Subtask Support**: Break down complex tasks into smaller, actionable subtasks, enhancing clarity and enabling more precise progress tracking.
*   📅 **Smart Due Date Tracking with Reminders**: Tasks are visually highlighted based on their due date status. Automated email reminders ensure you're always aware of upcoming deadlines.
*   📊 **Real-time Dashboard Analytics**: Get an instant overview of your productivity with a dashboard displaying total tasks, active projects, completion rates, and tasks due today.
*   🔐 **Robust User Authentication**: Secure sign-up and sign-in processes using email and password, ensuring your data is protected.
*   📧 **Email Verification**: A crucial step to confirm user identity and activate full account functionalities, including task reminders.
*   👤 **Personalized User Profiles**: Update your username and upload a custom avatar to make your Zendo experience truly yours.
*   ⚡ **Lightning-Fast & Responsive UI**: Powered by Next.js and React, the frontend delivers a smooth and adaptable user experience across all device types.
*   🔄 **Efficient Data Management**: Leverages SWR for intelligent client-side data fetching, caching, and revalidation, ensuring your UI is always up-to-date with minimal network requests.
*   💾 **Backend Redis Caching**: The Go backend utilizes Redis to cache frequently accessed data (like user profiles, task lists, and project data), significantly speeding up API response times.
*   🌐 **Secure Cross-Origin Communication**: Properly configured CORS policies ensure secure and seamless communication between the frontend and backend services.
*   🖼️ **Cloudinary Media Integration**: User avatar uploads are handled efficiently and securely via Cloudinary, providing robust media asset management.
*   💖 **Clean & Modern Design**: A clean, minimalistic UI built with Tailwind CSS and Shadcn UI components for a delightful user experience.

## 🛠️ Technologies Used

This project is a testament to building a robust full-stack application using a modern and efficient technology stack.

| Category    | Technology                                     | Description                                                                 |
| :---------- | :--------------------------------------------- | :-------------------------------------------------------------------------- |
| **Backend** | [Go (Golang)](https://go.dev/)                 | The core language for high-performance and scalable API development.        |
|             | [Fiber](https://gofiber.io/)                   | An extremely fast and expressive web framework for Go.                      |
|             | [MongoDB](https://www.mongodb.com/)            | A versatile NoSQL database for flexible data storage.                      |
|             | [Redis](https://redis.io/)                     | An in-memory data structure store used for caching and quick data retrieval.|
|             | [JWT (golang-jwt/jwt)](https://github.com/golang-jwt/jwt) | JSON Web Tokens for secure API authentication.                            |
|             | [GoCron](https://github.com/go-co-op/gocron)   | A powerful and simple Go library for scheduling background jobs (e.g., email reminders). |
|             | [Cloudinary (Go SDK)](https://cloudinary.com/) | Cloud-based media management for image uploads.                             |
| **Frontend**| [Next.js](https://nextjs.org/)                 | React framework for robust, production-ready web applications.              |
|             | [React](https://react.dev/)                    | A foundational library for building dynamic user interfaces.                |
|             | [TypeScript](https://www.typescriptlang.org/)  | Enhances JavaScript with static type checking for improved code quality.    |
|             | [Tailwind CSS](https://tailwindcss.com/)       | A utility-first CSS framework for rapid UI development.                   |
|             | [Shadcn UI](https://ui.shadcn.com/)            | Reusable, accessible UI components built on Radix UI and Tailwind CSS.      |
|             | [SWR](https://swr.vercel.app/)                 | React Hooks for data fetching, caching, and revalidation.                   |
|             | [Axios](https://axios-http.com/)               | A popular HTTP client for making API requests.                              |

## 🤝 Contributing

We welcome contributions to Zendo! If you're looking to help out, please consider the following guidelines:

*   ✨ **Fork the Repository**: Start by forking the `zendo` repository to your personal GitHub account.
*   🛠️ **Create a Feature Branch**: Work on new features or bug fixes in a dedicated branch. We recommend descriptive names like `feat/add-dark-mode` or `fix/improve-task-sorting`.
*   🐛 **Report Bugs**: If you encounter any issues, please open a detailed issue on the GitHub repository. Include steps to reproduce, expected behavior, and actual behavior.
*   💡 **Suggest Enhancements**: Have an idea for a new feature or an improvement to an existing one? Feel free to open an issue to discuss your ideas!
*   📝 **Write Clear Commit Messages**: Ensure your commit messages are concise, clear, and accurately describe the changes you've made.
*   🚀 **Submit Pull Requests**: Once your changes are ready, submit a pull request against the `main` branch. Provide a comprehensive description of your changes and the motivation behind them.
*   🤝 **Adhere to Code of Conduct**: Please maintain a respectful and inclusive environment for all contributors.

## 📄 License

This project is open-source. Please refer to the `LICENSE` file in the root of the repository for specific licensing details.

## ✍️ Author Info

Crafted with dedication by:

**Onos Ejoor**
*   🌐 [Personal Website](https://onos-ejoor.vercel.app)
<!-- *   👔 [LinkedIn](your_linkedin_profile_url_here) -->
*   🐦 [Twitter](https://x.com/DevText16)

---

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)