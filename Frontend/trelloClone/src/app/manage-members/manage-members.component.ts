import { Component, OnInit } from '@angular/core';
import { ProjectsService } from '../services/projects.service';
import { Project } from '../models/Project';
import { Router } from '@angular/router';

@Component({
  selector: 'app-manage-members',
  templateUrl: './manage-members.component.html',
  styleUrls: ['./manage-members.component.css']
})
export class ManageMembersComponent implements OnInit {
  projects: Project[] = [];
  selectedProject: Project | null = null;
  newMemberId: string = '';
  managerId: string = '';

  constructor(private projectsService: ProjectsService, private router: Router) {}

  ngOnInit(): void {
    this.setManagerIdFromToken();
    this.loadProjects();
  }

  private setManagerIdFromToken(): void {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken: any = this.decodeToken(token);
      if (decodedToken && decodedToken.userID) {
        this.managerId = decodedToken.userID;
      } else {
        console.error('Invalid or missing user ID in token.');
        this.router.navigate(['/login']);
      }
    } else {
      console.error('No token found, redirecting to login.');
      this.router.navigate(['/login']);
    }
  }

  private decodeToken(token: string): any {
    try {
      return JSON.parse(atob(token.split('.')[1])); // Dekodiranje JWT tokena
    } catch (error) {
      console.error('Invalid token:', error);
      return null;
    }
  }

  loadProjects(): void {
    this.projectsService.getProjects().subscribe({
      next: (data) => {
        this.projects = data.filter((project) => project.managerId === this.managerId);
      },
      error: (err) => {
        console.error('Error loading projects:', err);
      }
    });
  }

  selectProject(project: Project): void {
    this.selectedProject = project;
  }

  addMember(): void {
    if (!this.selectedProject || !this.newMemberId) {
      alert('Please select a project and enter a member ID.');
      return;
    }
  
    console.log('Adding member with payload:', {
      projectId: this.selectedProject.id,
      memberId: this.newMemberId,
      managerId: this.managerId
    });
  
    this.projectsService.addMember(this.selectedProject.id, this.newMemberId, this.managerId).subscribe({
      next: () => {
        alert('Member added successfully!');
        this.selectedProject?.members.push(this.newMemberId);
        this.newMemberId = '';
      },
      error: (err) => {
        console.error('Error adding member:', err);
        alert('Failed to add member.');
      }
    });
  }
  
  

  removeMember(memberId: string): void {
    if (!this.selectedProject) {
      alert('Please select a project.');
      return;
    }
  
    this.projectsService.removeMember(this.selectedProject.id, memberId, this.managerId).subscribe({
      next: () => {
        alert('Member removed successfully!');
        if (this.selectedProject) {
          this.selectedProject.members = this.selectedProject.members.filter((id) => id !== memberId);
        }
      },
      error: (err) => {
        console.error('Error removing member:', err);
        alert('Failed to remove member.');
      }
    });
  }
  
}
