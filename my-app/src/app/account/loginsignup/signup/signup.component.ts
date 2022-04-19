import { Component, ElementRef, OnInit } from '@angular/core';
import { DataService } from '../../../_services/data.service';

import { Customer } from '../../../_objects/customer';


import { AbstractControl, FormBuilder, FormControl, FormGroup } from '@angular/forms';

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


	firstname: string = "";
	lastname: string = "";
	phone_no: number = -1;
	birthday: string = "";

	newCustomer():void {
		this.setValue();
		this.tempCustomer.firstname = this.firstname;
		this.tempCustomer.lastname = this.lastname;
		this.tempCustomer.phone_no = this.phone_no;
		this.tempCustomer.birthday = this.birthday;
		this.tempCustomer.membershipID = 10;


	};

  constructor(
  	private fb: FormBuilder,
  	private dataService: DataService) { }

// userForm = this.fb.group({firstname});
public userForm: FormGroup = this.fb.group({
	firstname:'',
	lastname:'',
	phone_no: null,
	email:'',
});

 setValue() {
    this.firstname = this.userForm.get('firstname')?.value; // input value retrieved
    alert("FROM FORM " + this.userForm.get('firstname')?.value);

    //TODO add
  }

  ngOnInit(): void {
  	this.dataService.newCustomer(this.tempCustomer)
  }

}
