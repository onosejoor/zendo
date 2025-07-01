declare global {
  interface ITask {
    _id: string;
    title: string;
    description: string;
    userId: string;
    subTasks?: SubTask[];
    projectId?: string;
    dueDate: string;
    status: string;
    created_at: string;
  }
  interface SubTask {
    title: string;
    completed: boolean;
  }

  interface IProject {
    _id: string;
    name: string;
    description?: string;
    ownerId: string;
    totalTasks: number;
    created_at: string;
  }

  type SignUpFormData = {
    email: string;
    username: string;
    password: string;
    agreeToTerms: boolean;
  };

  type SigninFormData = {
    email: string;
    password: string;
    rememberMe: boolean;
  };

  type APIRes = {
    success: boolean;
    message: string;
  };

  type UserData = {
    success: boolean;
    email: string;
    created_at: Date;
    username: string;
  };

  type UserRes = {
    success: boolean;
    user: UserData;
  };
}

export {};
