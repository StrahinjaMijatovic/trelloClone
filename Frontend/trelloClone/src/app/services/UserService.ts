import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor() { }

  getToken(): string | null {
    return localStorage.getItem('token');
  }

  getClanId(): number | null {
    const token = this.getToken();

    if (token) {
      const payload = this.decodeToken(token);
      return payload ? payload.sub : null; 
    }
    return null;
  }

  private decodeToken(token: string): any {
    const base64Payload = token.split('.')[1];
    const payload = atob(base64Payload);
    return JSON.parse(payload);
  }
}
