# Zendo - Task Management Made Simple 🚀

Zendo is a modern and intuitive task management application designed to help individuals and teams organize their work, track progress, and boost productivity effortlessly. This project provides a robust frontend experience built with Next.js, offering comprehensive features for managing tasks and projects.

## ✨ Features

*   **Intuitive Dashboard**: Get a clear overview of your total tasks, active projects, completion rates, and tasks due today at a glance.
*   **Comprehensive Task Management**: Create, view, edit, and mark tasks as complete. Tasks can include titles, descriptions, due dates, and status (pending, in-progress, completed).
*   **Robust Project Organization**: Group related tasks into projects for better structure and efficient tracking. Manage project details, view associated tasks, and monitor overall progress.
*   **User Authentication**: Secure sign-up and sign-in processes to ensure personalized and protected user data.
*   **Dynamic Data Fetching**: Leverages SWR for efficient, real-time data fetching and caching, providing a smooth user experience.
*   **Modern UI/UX**: Built with shadcn/ui and Tailwind CSS, ensuring a sleek, responsive, and accessible interface across various devices.
*   **API Interception**: Centralized Axios interceptor handles authentication token refresh and robust error management for seamless backend communication.

## 🚀 Getting Started

Follow these steps to get Zendo up and running on your local machine.

### Prerequisites

Before you begin, ensure you have the following installed:

*   Node.js (v18.x or later)
*   npm or Yarn or pnpm

### Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/onosejoor/zendo.git
    ```
2.  **Navigate to the client directory**:
    ```bash
    cd task-manager/client
    ```
3.  **Install dependencies**:
    ```bash
    npm install
    # or yarn install
    # or pnpm install
    ```
4.  **Set up environment variables**:
    Create a `.env.local` file in the `client` directory based on `env.example`:
    ```
    JWT_SECRET="YOUR_JWT_SECRET"
    SERVER_URL="YOUR_BACKEND_SERVER_URL"
    ```
    Replace `YOUR_JWT_SECRET` with a strong, random secret key and `YOUR_BACKEND_SERVER_URL` with the URL of your backend server (e.g., `http://localhost:8080` if running locally).

5.  **Run the development server**:
    ```bash
    npm run dev
    # or yarn dev
    # or pnpm dev
    ```
    Open [http://localhost:3000](http://localhost:3000) with your browser to see the application.

## 💡 Usage

Once the application is running, you can:

1.  **Sign Up / Sign In**:
    Navigate to the `/auth/signup` or `/auth/signin` pages to create a new account or log in with existing credentials. This secures your personal dashboard and data.

2.  **Explore the Dashboard**:
    Upon successful login, you'll land on the dashboard. Here, you'll find key statistics like total tasks, active projects, and completion rates. Recent tasks and projects are also displayed for quick access.

3.  **Manage Tasks**:
    *   Go to the `/dashboard/tasks` page to view all your tasks.
    *   Click the "New Task" button to create a new task, providing a title, description, due date, and optional project association.
    *   Edit existing tasks by clicking on them or using the dropdown menu on task cards.
    *   Mark tasks as completed directly from the dashboard or the tasks list by checking the checkbox.

4.  **Manage Projects**:
    *   Visit the `/dashboard/projects` page to oversee all your projects.
    *   Use the "New Project" button to start a new project, assigning a name and description.
    *   Click on any project card to view its details and all tasks associated with it. From here, you can also add new tasks directly to that project.
    *   Projects can be edited or deleted from their respective detail pages.

## 🛠 Technologies Used

| Technology    | Description                                       | Link                                                 |
| :------------ | :------------------------------------------------ | :--------------------------------------------------- |
| Next.js       | React framework for production                    | [Next.js](https://nextjs.org/)                       |
| React         | JavaScript library for building user interfaces   | [React](https://react.dev/)                          |
| TypeScript    | Strongly typed JavaScript                         | [TypeScript](https://www.typescriptlang.org/)        |
| Tailwind CSS  | Utility-first CSS framework                       | [Tailwind CSS](https://tailwindcss.com/)             |
| shadcn/ui     | Reusable UI components built with Radix UI & Tailwind | [shadcn/ui](https://ui.shadcn.com/)                  |
| SWR           | React Hooks for data fetching                     | [SWR](https://swr.vercel.app/)                       |
| Axios         | Promise-based HTTP client                         | [Axios](https://axios-http.com/)                     |
| Day.js        | Fast, minimalist `Date` library                   | [Day.js](https://day.js.org/)                        |
| Sonner        | An opinionated toast component for React          | [Sonner](https://sonner.emilkowalski.no/)            |
| Lucide React  | A beautiful & consistent icon toolkit             | [Lucide React](https://lucide.dev/guide/packages/lucide-react) |

## 🤝 Contributing

We welcome contributions to Zendo! If you're interested in improving the project, please follow these guidelines:

1.  **Fork the repository**.
2.  **Create a new branch** for your feature or bug fix: `git checkout -b feature/your-feature-name` or `bugfix/fix-description`.
3.  **Make your changes**, ensuring they adhere to the project's coding style.
4.  **Write clear and concise commit messages**.
5.  **Push your branch** to your forked repository.
6.  **Open a Pull Request** to the `main` branch of this repository. Provide a detailed description of your changes.

## 📄 License

This project does not currently have an explicit license. Please contact the author for licensing information.

## 👤 Author

*   **Onos Ejoor**
    *   GitHub: [@onosejoor](https://github.com/onosejoor)
    *   LinkedIn: [YourLinkedInProfile](https://www.linkedin.com/in/onose-ojoor/) 
    *   X (Twitter): [@DevText15](https://twitter.com/DevText16)

---

[![Built with Next.js](https://img.shields.io/badge/Built%20with-Next.js-000000.svg?style=for-the-badge&logo=next.js)](https://nextjs.org/)
[![Styled with Tailwind CSS](https://img.shields.io/badge/Styled%20with-Tailwind%20CSS-06B6D4.svg?style=for-the-badge&logo=tailwind-css)](https://tailwindcss.com/)
[![Type-Checked with TypeScript](https://img.shields.io/badge/Type--Checked%20with-TypeScript-3178C6.svg?style=for-the-badge&logo=typescript)](https://www.typescriptlang.org/)
[![Powered by SWR](https://img.shields.io/badge/Powered%20by-SWR-FF5733.svg?style=for-the-badge&logo=swr)](https://swr.vercel.app/)
[![Deployed on Vercel](https://img.shields.io/badge/Deployed%20on-Vercel-000000.svg?style=for-the-badge&logo=vercel)](https://vercel.com/)

---

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)