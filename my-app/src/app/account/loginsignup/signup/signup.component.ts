import { Component, OnInit } from '@angular/core';
import { DataService } from '../../../_services/data.service';

import { Customer } from '../../../_objects/customer';


@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {

	tempCustomer: Customer = {
		firstname: '',
		lastname: '',
		phone_no: -1,
		address: '',
		preference: '',
		birthday: '',
		membershipID: -1
	};


	newCustomer():void {
		alert("Called newCustomer");
		this.tempCustomer.firstname = (<HTMLInputElement>document.getElementById('firstname')).value;

		alert(this.tempCustomer.firstname);

	};

  constructor() { }

  ngOnInit(): void {

	window.addEventListener('load', this.newCustomer,false);
  }

}
