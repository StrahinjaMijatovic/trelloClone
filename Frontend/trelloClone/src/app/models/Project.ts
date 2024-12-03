export interface Project {
    id: string;
    name: string;
    endDate: string | null;
    minMembers: number;
    maxMembers: number;
    managerId: string;
    members: string[];
    createdAt: Date;
    updatedAt: Date;
  }
  