import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Project } from '../models/Project';

@Injectable({
  providedIn: 'root'
})
export class ProjectsService {
  private baseUrl = 'http://localhost:8080/projects'; // Backend URL

  constructor(private http: HttpClient) {}


  // Get a single project
  getProjectById(id: string): Observable<Project> {
    return this.http.get<Project>(`${this.baseUrl}/${id}`);
  }

  // Create a new project
  createProject(project: Project): Observable<Project> {
    return this.http.post<Project>(this.baseUrl, project);
  }

  // // Add a member to a project
  // addMember(projectId: string, memberId: string): Observable<any> {
  //   return this.http.post(`${this.baseUrl}/${projectId}/add-member`, { memberId });
  // }

  // // Remove a member from a project
  // removeMember(projectId: string, memberId: string): Observable<any> {
  //   return this.http.post(`${this.baseUrl}/${projectId}/remove-member`, { memberId });
  // }

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(this.baseUrl);
  }

  // Dodavanje člana u projekat
  addMember(projectId: string, memberId: string, managerId: string): Observable<any> {
    const payload = { memberId, managerId };
    return this.http.post(`${this.baseUrl}/${projectId}/add-member`, payload);
  }

  // Uklanjanje člana iz projekta
  removeMember(projectId: string, memberId: string, managerId: string): Observable<any> {
    const payload = { memberId, managerId };
    return this.http.post(`${this.baseUrl}/${projectId}/remove-member`, payload);
  }
}
