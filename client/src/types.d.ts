declare global {
  interface ITask {
    _id: string;
    title: string;
    description: string;
    subTasks?: ISubTask[];
    projectId?: string;
    dueDate: Date | string;
    status: "in-progress" | "completed" | "pending";
    created_at: Date;
  }
  interface ISubTask {
    _id: string;
    title: string;
    completed: boolean;
  }

  interface IProject {
    _id: string;
    name: string;
    description?: string;
    ownerId: string;
    totalTasks: number;
    created_at: Date;
  }

  type SignUpFormData = {
    email: string;
    username: string;
    password: string;
  };

  type SigninFormData = {
    email: string;
    password: string;
  };

  type APIRes = {
    success: boolean;
    message: string;
  };

  interface IUser {
    _id: string;
    email_verified: boolean,
    username: string;
    email: string;
    avatar?: string;
    created_at: Date;
  }

  interface IStats {
    total_tasks: number;
    total_projects: number;
    completion_rate: number;
    completed_tasks: number;
    dueToday: number;
  }

  type UserRes = {
    success: boolean;
    user: IUser;
  };
}

export {};
