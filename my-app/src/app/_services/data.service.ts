import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

// import { Shop } from '../_objects/shop';

@Injectable({
  providedIn: 'root'
})
export class DataService {
	private productData = "http://localhost:8081/api/product";
	private customerData = "http://localhost:8081/api/customer";
	private orderData = "http://localhost:8081/api/order";
	private cartData = "http://localhost:8081/api/shoppingCart";
	private shipmentData = "http://localhost:8081/api/shipment";
	private rewardsData = "http://localhost:8081/api/rewards";
	private rewardPointNumberData = "http://localhost:8081/api/rewardpt_no";
	private adminData = "http://localhost:8081/api/admin";
	private couponData = "http://localhost:8081/api/coupon";
	// private dataServer = "http://localhost:3000/products";


	constructor(private http: HttpClient) {
	   this.getProduct().subscribe(data => {
	    console.log(data);
	   });
	   this.getCustomer().subscribe(data => {
	   	console.log(data);
	   })
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


	 public getProduct(): Observable<any> {
	   return this.http.get(this.productData);
	 }
	 public getCustomer(): Observable<any> {
	   return this.http.get(this.customerData);
	 }
	 public getOrder(): Observable<any> {
	   return this.http.get(this.orderData);
	 }
	 public getCart(): Observable<any> {
	   return this.http.get(this.cartData);
	 }
	 public getShipment(): Observable<any> {
	   return this.http.get(this.shipmentData);
	 }
	 public getRewards(): Observable<any> {
	   return this.http.get(this.rewardsData);
	 }
	 public getRewardPointNumber(): Observable<any> {
	   return this.http.get(this.rewardPointNumberData);
	 }
	 public getAdmin(): Observable<any> {
	   return this.http.get(this.adminData);
	 }
	 public getCoupon(): Observable<any> {
	   return this.http.get(this.couponData);
	 }

	 public newCustomer(data:any): Observable<any> {
	 	return this.http.post(this.customerData, data);
	 }


	// getShop() {
	//   return this.http.get<Shop>(this.shopUrl);
}
