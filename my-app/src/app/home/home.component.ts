import { Component, OnInit } from '@angular/core';
import { DataService } from '../_services/data.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  products: any[] = [];
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

}