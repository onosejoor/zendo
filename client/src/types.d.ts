declare global {
  type Status = "in-progress" | "completed" | "pending";
  interface ITask {
    _id: string;
    title: string;
    description: string;
    subTasks: ISubTask[];
    projectId?: string;
    assignees: IAssignee[];
    team_id?: string;
    dueDate: Date | string;
    status: Status;
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

  type ITeam = {
    _id: string;
    name: string;
    description: string;
    creatorId: string;
    created_at: Date;
    role: TeamRole;
    joined_at: Date;
    members_count: number;
  };

  type TeamRole = "owner" | "admin" | "member";

  interface IMember extends IUser {
    role: Role;
  }

  interface IAssignee {
    email: string;
    _id: string;
    username: string;
  }
  interface ITeamWithMember {
    member: IMember;
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
    email_verified: boolean;
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

  interface Assignees {
    _id: string;
    username: string;
    email: string;
  }
}

export {};
