import { Component, ElementRef, OnInit } from '@angular/core';
import { DataService } from '../_services/data.service';

import { Products } from '../_objects/products';

import { AbstractControl, FormBuilder, FormControl, FormGroup } from '@angular/forms';


import { HttpClient, HttpErrorResponse } from '@angular/common/http';

import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {


  // private productData = "http://localhost:8081/api/product";
  private productData = "http://localhost:3000/products";

	admin: any[] = [];
  products: Products[] = [];

	productName: string = "";
	productDescription: string = "";
	productPrice: number = -1;
	productImage: string = "";
  productCategory: string = "";


	newProduct: Products = {
		// item_no: -1,
		itemName: "",
		description: "",
		price: -1,
		category: "",
		product_image: ""
	};

  selectedProduct:string ="";

    //TODO add
  constructor(private fb: FormBuilder, private dataService: DataService, private http: HttpClient) { }

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

  public addProduct: FormGroup = this.fb.group({
  	productName:'',
  	productPrice: null,
  	productDescription:'',
  	productImage:'',
    productCategory:''
	});

  ngOnInit(): void {
  	this.dataService.getAdmin().subscribe((data: any[])=>{
  		console.log(data);
  		this.admin = data;
  	})
    this.dataService.getProduct().subscribe((data: any[])=>{
      console.log(data);
      this.products = data;
    })
  	// this.http.post<Products>(this.productData, JSON.stringify(this.newProduct)).subscribe(data => {
	 	// 	this.productName = data.itemName;
	 	// });
  }


	addNewProduct():void {
		this.setValue();
		this.newProduct.itemName = this.productName;
		this.newProduct.description = this.productDescription;
		this.newProduct.price = this.productPrice;
    this.newProduct.category = this.productCategory;
		this.newProduct.product_image = this.productImage;
	}

	setValue() {
	    this.productName = this.addProduct.get('productName')?.value; // input value retrieved
	    this.productDescription = this.addProduct.get('productDescription')?.value;
	    this.productPrice = this.addProduct.get('productPrice')?.value;
      this.productCategory = this.addProduct.get('productCategory')?.value;
	    this.productImage = this.addProduct.get('productImage')?.value;
	}

  bindValue(prod:any):void {
    this.selectedProduct = prod;
    // alert(prod);
  }

	postProduct(): void {
  	// alert("PUSHED");
    this.addNewProduct();
	 	// return this.http.post<Products>(this.productData, JSON.stringify(this.newProduct));
     this.dataService.postProduct(this.newProduct);
  }

  putProduct(): void {
    var itemID = -1;
    var i = 0;
    for(let prod of this.products)
    {
      // alert("Checking " + prod.itemName + " vs " + this.selectedProduct);
      if(prod.itemName === this.selectedProduct)
      {
        itemID = i; 
      }
      i++;

    }
    // var itemID = this.products.indexOf(this.selectedProduct).itemName;
    // alert(itemID);
    this.addNewProduct();
    this.dataService.putProduct(this.newProduct, itemID);

  }


  testButton(): void {
    this.dataService.testFunc();
  }

}
