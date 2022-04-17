import { Component, OnInit } from '@angular/core';
import { DataService } from '../_services/data.service';

import { Products } from '../_objects/products';
import { Cart } from '../_objects/cart';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {

	cart: any[] = [];
  orders: any[] = [];
  shipmentOrders: any[] =[];
  // validCoupons: any[] = [];
  validCoupons = ["test", "ship10"];

  products: Products[] = [];

  organizedCart: Cart[] = [];

  // checkCoupons():void {};
  checkCoupons():void{
    var inputCoupon = (<HTMLInputElement>document.getElementById('inputCoupon'))!.value;
    for (let coupon of this.validCoupons) {
      alert(inputCoupon + " vs " + coupon);
      if(coupon == inputCoupon)
      {
        (<HTMLInputElement>document.getElementById('isCouponValid')!).innerHTML = "Coupon has been applied";
        alert("Works " + inputCoupon);
      }
    }
      (<HTMLInputElement>document.getElementById('isCouponValid')!).innerHTML = "Coupon is invalid";
      alert("Fails " + inputCoupon);
  }

  organizeCart():void {
    for(var i = 0; i < this.cart.length/2; i++)
    {
      this.organizedCart[i].productName = this.cart[i];
      this.organizedCart[i].productQuantity = this.cart[i+1];
      for(var j = 0; j < this.products.length; j++)
      {
        if(this.organizedCart[i].productName === this.products[j].itemName)
        {
          this.organizedCart[i].productPrice = this.products[j].price;
        }
      }
    }
  };

  calculateTotal(): number{
    var total = 0;
    for(var i = 0; i < this.organizedCart.length; i++)
    {
      for(var j = 0; j < this.products.length; j++)
      {
        if(this.organizedCart[i].productName === this.products[j].itemName)
        {
          total = total + (this.products[j].price * this.organizedCart[i].productQuantity);
        }
      }
      
    }
    return total;
   };

  constructor(
  	private dataService: DataService) { }

  ngOnInit(): void {
  	this.dataService.getCart().subscribe((data: any[])=>{
  		console.log(data);
  		this.cart = data;
  	})

    this.dataService.getOrder().subscribe((data: any[])=>{
      console.log(data);
      this.orders = data;
    })

    this.dataService.getCoupon().subscribe((data: any[])=>{
      console.log(data);
      this.validCoupons = data;
    })

    this.dataService.getShipment().subscribe((data: any[])=>{
      console.log(data);
      this.shipmentOrders = data;
    })

    this.dataService.getProduct().subscribe((data: any[])=>{
      console.log(data);
      this.products = data;
    })

    this.organizeCart();
  }

}
