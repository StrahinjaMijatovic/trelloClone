import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor() { }

  // Funkcija za dobijanje tokena iz localStorage
  getToken(): string | null {
    return localStorage.getItem('token');
  }

  // Funkcija za ekstraktovanje ID-a korisnika iz tokena
  getClanId(): number | null {
    const token = this.getToken(); // Pozovi funkciju za dobijanje tokena

    if (token) {
      const payload = this.decodeToken(token);
      return payload ? payload.sub : null; // Vraća ID korisnika ako postoji
    }
    return null;
  }

  // Dekodiranje JWT tokena
  private decodeToken(token: string): any {
    const base64Payload = token.split('.')[1];
    const payload = atob(base64Payload);
    return JSON.parse(payload);
  }
}
