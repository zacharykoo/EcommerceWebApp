import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpParams } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

import { Products } from '../_objects/products';

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
	   });
	   this.getOrder().subscribe(data => {
	   	console.log(data);
	   });
	   this.getCart().subscribe(data => {
	   	console.log(data);
	   });
	   this.getShipment().subscribe(data => {
	   	console.log(data);
	   });
	   this.getRewards().subscribe(data => {
	   	console.log(data);
	   });
	   this.getRewardPointNumber().subscribe(data => {
	   	console.log(data);
	   });
	   this.getAdmin().subscribe(data => {
	   	console.log(data);
	   });
	   this.getCoupon().subscribe(data => {
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
	 	return this.http.post(this.customerData, JSON.stringify(data));
	 }



	private productDataTest = "http://localhost:3000/products";

	 postProduct(product:Products):void {
	 	alert(product.price);
	 	// var testData = '{"ItemName":"Test","Description":"desc","Price":12,"Category":"category","ProductImage":"www.com"}';
	 	var jsonData = JSON.stringify(product, function(key, value) {
	 		if (key == "price") {
			    return parseInt(value);
			  } else {
			    return value;
			  }
	 	});

	    this.http.post<Products>(this.productData, jsonData).subscribe(
	        val => {
	            console.log("POST to products call successful value returned in body", 
	                        val);
	        },
	        response => {
	            console.log("POST to products call in error", response);
	        },
	        () => {
	            console.log("The POST to products observable is now completed.");
	        }
	    );
	  }

	  addToCart(product:Products):void {
	  	var params = new HttpParams().set("cartID", 1);

	  	var oldProductList = this.http.get(this.cartData, {params: params}).subscribe(
	        val => {
	            console.log("GET to cart call successful value returned in body", 
	                        val);
	        },
	        response => {
	            console.log("GET to cart call in error", response);
	        },
	        () => {
	            console.log("GET to cart products observable is now completed.");
	        }
	    );

	  	// alert(oldProductList);

	  	// oldProductList.append(product.itemName);
	  	// oldProductList.append(1);

	  // 	var jsonData = JSON.stringify(oldProductList, function(key, value) {
	 	// 	if (key == "quantity") {
			//     return parseInt(value);
			//   } else {
			//     return value;
			//   }
	 	// });

	 	this.http.post<Products>(this.cartData, oldProductList).subscribe(
	        val => {
	            console.log("POST to cart call successful value returned in body", 
	                        val);
	        },
	        response => {
	            console.log("POST to cart call in error", response);
	        },
	        () => {
	            console.log("The POST to cart observable is now completed.");
	        }
	    );
	  }
	  // putProduct(product:Products):void {
	 	// alert(product.Price);
	 	// // var testData = '{"ItemName":"Test","Description":"desc","Price":12,"Category":"category","ProductImage":"www.com"}';
	 	// var jsonData = JSON.stringify(product, function(key, value) {
	 	// 	if (key == "Price") {
			//     return parseInt(value);
			//   } else {
			//     return value;
			//   }
	 	// });

	  //   this.http.put<Products>(this.productData, jsonData).subscribe(
	  //       val => {
	  //           console.log("POST call successful value returned in body", 
	  //                       val);
	  //       },
	  //       response => {
	  //           console.log("POST call in error", response);
	  //       },
	  //       () => {
	  //           console.log("The POST observable is now completed.");
	  //       }
	  //   );
	  // }


	// getShop() {
	//   return this.http.get<Shop>(this.shopUrl);
}
