import { Component, OnInit } from '@angular/core';
import { ProjectsService } from '../services/projects.service';
import { Project } from '../models/Project';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.css']
})
export class ProjectsComponent implements OnInit {
  projects: Project[] = [];
  selectedProject: Project | null = null;
  newMemberId: string = '';
  managerId: string = ''; // Dodato za čuvanje managerId

  constructor(private projectsService: ProjectsService) {}

  ngOnInit(): void {
    this.setManagerIdFromToken();
    this.loadProjects();
  }

  private setManagerIdFromToken(): void {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = JSON.parse(atob(token.split('.')[1])); // Dekodiranje JWT tokena
      this.managerId = decodedToken.userID;
    } else {
      console.error('No token found');
    }
  }

  loadProjects(): void {
    this.projectsService.getProjects().subscribe((data) => {
      this.projects = data;
    });
  }

  selectProject(project: Project): void {
    this.selectedProject = project;
  }

  addMember(): void {
    if (this.selectedProject && this.newMemberId) {
      this.projectsService
        .addMember(this.selectedProject.id, this.newMemberId, this.managerId) // Prosleđivanje managerId
        .subscribe(() => {
          this.selectedProject?.members.push(this.newMemberId);
          this.newMemberId = '';
        });
    }
  }

  removeMember(memberId: string): void {
    if (this.selectedProject) {
      this.projectsService
        .removeMember(this.selectedProject.id, memberId, this.managerId) // Prosleđivanje managerId
        .subscribe(() => {
          this.selectedProject!.members = this.selectedProject!.members.filter(
            (id) => id !== memberId
          );
        });
    }
  }
}
