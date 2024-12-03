import { Component, OnInit } from '@angular/core';
import { ProjectsService } from '../services/projects.service';
import { Project } from '../models/Project';
import { Router } from '@angular/router';
import { jwtDecode } from 'jwt-decode';

@Component({
  selector: 'app-create-project',
  templateUrl: './create-project.component.html',
  styleUrls: ['./create-project.component.css']
})
export class CreateProjectComponent implements OnInit {
  project: Partial<Project> = {
    name: '',
    endDate: null,
    minMembers: 1,
    maxMembers: 1,
    managerId: ''
  };


  constructor(private projectsService: ProjectsService, private router: Router) {}

  ngOnInit(): void {
    this.setManagerIdFromToken();
  }

  private setManagerIdFromToken(): void {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = this.decodeToken(token);
      if (decodedToken && decodedToken.userID) {
        this.project.managerId = decodedToken.userID;
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
      return jwtDecode(token);
    } catch (error) {
      console.error('Invalid token:', error);
      return null;
    }
  }

  onSubmit(): void {
    if (this.project.minMembers! > this.project.maxMembers!) {
      alert('Minimum members cannot be greater than maximum members.');
      return;
    }
  
    // Ensure endDate is in ISO 8601 format
    if (this.project.endDate) {
      this.project.endDate = new Date(this.project.endDate).toISOString();
    }
  
    this.projectsService.createProject(this.project as Project).subscribe({
      next: (data) => {
        alert('Project created successfully!');
        this.project = {
          name: '',
          endDate: null,
          minMembers: 1,
          maxMembers: 1,
          managerId: ''
        };
      },
      error: (err) => {
        console.error('Error creating project:', err);
        alert('Failed to create project. Please try again.');
      }
    });
  }
}
