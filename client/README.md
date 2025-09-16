# Zendo Task Manager ✔️

Zendo is a sleek and powerful task management application built with Next.js and TypeScript. It provides an intuitive interface for individuals and teams to organize projects, track tasks, and boost productivity through a clean, responsive dashboard.

## ✨ Features

-   **📝 Task Management**: Create, edit, and delete tasks with detailed descriptions, due dates, and status tracking (Pending, In-Progress, Completed).
-   **📂 Project Organization**: Group related tasks into projects to maintain a clear and organized workflow.
-   **🤝 Team Collaboration**: Create teams, invite members via email, and manage collaborative tasks and projects seamlessly.
-   **🔐 Secure Authentication**: Robust user authentication system supporting both email/password sign-up and Google OAuth for quick access.
-   **📊 Interactive Dashboard**: Get a quick overview of your productivity with stats on total tasks, active projects, and completion rates.
-   **✅ Sub-task Functionality**: Break down complex tasks into smaller, manageable sub-tasks for better tracking.
-   **📱 Fully Responsive**: A seamless experience across all devices, from desktops to mobile phones.

## 🛠️ Technologies Used

| Technology         | Description                                        |
| ------------------ | -------------------------------------------------- |
| **Next.js**        | A React framework for building fast, modern web apps. |
| **TypeScript**     | Superset of JavaScript for type-safe code.         |
| **Tailwind CSS**   | A utility-first CSS framework for rapid UI development. |
| **SWR**            | React Hooks for data fetching and state management. |
| **shadcn/ui**      | Re-usable UI components built on Radix UI.         |
| **Axios**          | Promise-based HTTP client for making API requests. |
| **Lucide React**   | A library of simply beautiful open-source icons.   |
| **NextAuth.js**    | Authentication handling and session management.    |

## 🚀 Getting Started

Follow these instructions to get a local copy up and running for development and testing purposes.

### Prerequisites

Make sure you have the following installed on your machine:
-   [Node.js](https://nodejs.org/en/) (v18 or higher)
-   [npm](https://www.npmjs.com/) or [yarn](https://yarnpkg.com/)

### Installation

1.  **Clone the Repository**
    ```bash
    git clone https://github.com/onosejoor/zendo.git
    ```

2.  **Navigate to the Project Directory**
    ```bash
    cd zendo/client
    ```

3.  **Install Dependencies**
    ```bash
    npm install
    ```

4.  **Set Up Environment Variables**
    Create a `.env.local` file in the `client` directory and add the following variables.

    ```env
    # This is only used on the server-side for middleware, not exposed to the client
    JWT_SECRET="your-jwt-secret-here"
    
    # URL of your backend server (for server-to-server communication if needed)
    SERVER_URL="http://localhost:8080"
    
    # Publicly accessible URL of your backend server (for client-side requests)
    NEXT_PUBLIC_SERVER_URL="http://localhost:8080"
    ```

5.  **Run the Development Server**
    ```bash
    npm run dev
    ```
    Open [http://localhost:3000](http://localhost:3000) in your browser to see the application.

## Usage

Once the application is running, you can:

-   **Sign Up**: Create a new account using your email and password or sign in with Google.
-   **Dashboard**: After logging in, you'll be directed to your personal dashboard which displays key statistics about your tasks and projects.
-   **Create Projects**: Navigate to the "Projects" page and create a new project to group your tasks.
-   **Create Tasks**: Create tasks individually from the "Tasks" page or directly within a project. Assign due dates, descriptions, and sub-tasks.
-   **Manage Teams**: Go to the "Teams" section to create a new team, invite members, and start collaborating on shared tasks.

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  🍴 Fork the Project
2.  🌿 Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  ✨ Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  🚀 Push to the Branch (`git push origin feature/AmazingFeature`)
5.  🎉 Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👨‍💻 Author

**Onos Ejoor**

-   **Website**: [onos-ejoor.vercel.app](https://onos-ejoor.vercel.app)
-   **Twitter**: `[@DevText16]`

---

![Next JS](https://img.shields.io/badge/Next-black?style=for-the-badge&logo=next.js&logoColor=white)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![SWR](https://img.shields.io/badge/SWR-000000?style=for-the-badge&logo=swr&logoColor=white)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)