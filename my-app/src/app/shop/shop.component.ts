import { Component, OnInit } from '@angular/core';

import { Products } from '../_objects/products';
import { DataService } from '../_services/data.service';

@Component({
  selector: 'app-shop',
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.css']
})
export class ShopComponent implements OnInit {

	products: Products[] = [];
	// products = [];
	// shop: Shop | undefined;
  constructor(
  	private dataService: DataService) { }

  ngOnInit(): void {
  	// this.shop = this.shopService.getJSON();
  	this.dataService.getProduct().subscribe((data: any[])=>{
  		console.log(data);
  		this.products = data;
  	})
  }

  addToCart(product:Products):void {
    alert("Added to cart");
    this.dataService.addToCart(product);
  }

}
