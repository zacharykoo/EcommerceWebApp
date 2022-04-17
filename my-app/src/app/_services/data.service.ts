import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

// import { Shop } from '../_objects/shop';

@Injectable({
  providedIn: 'root'
})
export class DataService {
	private dataServer = "http://localhost:3000/products";

	constructor(private http: HttpClient) {
	   this.getJSON().subscribe(data => {
	    console.log(data);
	   });
	 }

	handleError(error: HttpErrorResponse) {
    let errorMessage = 'Unknown error!';
    if (error.error instanceof ErrorEvent) {
      // Client-side errors
      errorMessage = `Error: ${error.error.message}`;
    } else {
      // Server-side errors
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    window.alert(errorMessage);
    return throwError(errorMessage);
  }


	 public getJSON(): Observable<any> {
	   return this.http.get(this.dataServer);
	 }


	// getShop() {
	//   return this.http.get<Shop>(this.shopUrl);
}
