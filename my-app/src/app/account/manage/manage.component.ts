import { Component, OnInit } from '@angular/core';
import { DataService } from '../../_services/data.service';

@Component({
  selector: 'app-manage',
  templateUrl: './manage.component.html',
  styleUrls: ['./manage.component.css']
})
export class ManageComponent implements OnInit {

	customers: any[] =[];
	

  constructor(private dataService: DataService) { }

  ngOnInit(): void {
  	this.dataService.getCustomer().subscribe((data: any[])=>{
  		console.log(data);
  		this.customers = data;
  	})
  }

}