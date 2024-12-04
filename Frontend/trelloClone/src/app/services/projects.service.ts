import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Project } from '../models/Project';

@Injectable({
  providedIn: 'root'
})
export class ProjectsService {
  private baseUrl = 'http://localhost:8080/projects'; 

  constructor(private http: HttpClient) {}


  getProjectById(id: string): Observable<Project> {
    return this.http.get<Project>(`${this.baseUrl}/${id}`);
  }

  createProject(project: Project): Observable<Project> {
    return this.http.post<Project>(this.baseUrl, project);
  }

  // addMember(projectId: string, memberId: string): Observable<any> {
  //   return this.http.post(`${this.baseUrl}/${projectId}/add-member`, { memberId });
  // }

  // removeMember(projectId: string, memberId: string): Observable<any> {
  //   return this.http.post(`${this.baseUrl}/${projectId}/remove-member`, { memberId });
  // }

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(this.baseUrl);
  }

  addMember(projectId: string, memberId: string, managerId: string): Observable<any> {
    const payload = { memberId, managerId };
    return this.http.post(`${this.baseUrl}/${projectId}/add-member`, payload);
  }

  removeMember(projectId: string, memberId: string, managerId: string): Observable<any> {
    const payload = { memberId, managerId };
    return this.http.post(`${this.baseUrl}/${projectId}/remove-member`, payload);
  }
}
