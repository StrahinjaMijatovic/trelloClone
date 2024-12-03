import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpParams } from '@angular/common/http';

interface UserResponse {
  firstName: string;
  lastName: string;
  role: string;
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  isLoggedIn: boolean = false;
  firstName: string = '';
  lastName: string = '';
  role: string = '';

  constructor(private router: Router, private http: HttpClient) {}

  ngOnInit() {
    this.checkUserStatus();
  }

  navigateToLogin() {
    this.router.navigate(['/login']);
  }

  navigateToRegister() {
    this.router.navigate(['/register']);
  }

  navigateToProjects() {
    this.router.navigate(['/projects']);
  }

  navigateToCreateProject() {
    this.router.navigate(['/create-project']);
  }

  logout() {
    localStorage.removeItem('token');
    this.isLoggedIn = false;
    this.firstName = '';
    this.lastName = '';
    this.role = '';
    this.router.navigate(['/home']);
  }

  checkUserStatus() {
    const token = localStorage.getItem('token');
    if (token) {
      this.http.post<UserResponse>('http://localhost:8000/verify-token', { token })
        .subscribe((response: UserResponse) => {
          this.isLoggedIn = true;
          this.firstName = response.firstName;
          this.lastName = response.lastName;
          this.role = response.role;
        }, error => {
          console.error('Provera tokena nije uspela:', error);
          this.isLoggedIn = false;
        });
    }
  }
}
