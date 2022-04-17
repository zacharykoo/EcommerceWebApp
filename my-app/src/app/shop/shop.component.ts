import { Component, OnInit } from '@angular/core';

// import { Shop } from '../_objects/shop';
import { DataService } from '../_services/data.service';

@Component({
  selector: 'app-shop',
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.css']
})
export class ShopComponent implements OnInit {

	products: any[] = [];
	// products = [];
	// shop: Shop | undefined;
  constructor(
  	private dataService: DataService) { }

  ngOnInit(): void {
  	// this.shop = this.shopService.getJSON();
  	this.dataService.getJSON().subscribe((data: any[])=>{
  		console.log(data);
  		this.products = data;
  	})
  }

}
