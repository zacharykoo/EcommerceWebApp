import { Component, OnInit } from '@angular/core';
import { DataService } from '../_services/data.service';
import { FormGroup, FormControl, Validators, FormBuilder }  from '@angular/forms';
import { FormsModule,ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

	// signup: FormGroup;
	customers: any[] =[];

  constructor(private dataService: DataService) { }

  ngOnInit(): void {
  	this.dataService.getCustomer().subscribe((data: any[])=>{
  		console.log(data);
  		this.customers = data;
  	})
  }

  newCustomer(data:any): void {
  	console.log("Signup work");
  }

}
