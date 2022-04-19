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
		fn: '',
		ln: '',
		phone_no: '',
		address: '',
		preference: '',
		birthday: '',
		// membershipID: -1
	};


	firstname: string = "";
	lastname: string = "";
	phone_no: string = "";
	birthday: string = "";
	address: string = "";
	preference: string = "";

	newCustomer():void {
		this.setValue();
		this.tempCustomer.fn = this.firstname;
		this.tempCustomer.ln = this.lastname;
		this.tempCustomer.phone_no = this.phone_no;
		this.tempCustomer.birthday = this.birthday;
		this.tempCustomer.address = this.address;
		this.tempCustomer.preference = this.preference;
	};

  constructor(
  	private fb: FormBuilder,
  	private dataService: DataService) { }

// userForm = this.fb.group({firstname});
public userForm: FormGroup = this.fb.group({
	firstname:'',
	lastname:'',
	birthday:'',
	phone_no: null,
	address:'',
	preference:''
});

 setValue() {
    this.firstname = this.userForm.get('firstname')?.value; // input value retrieved
    this.lastname = this.userForm.get('lastname')?.value;
    this.birthday = this.userForm.get('birthday')?.value;
    this.phone_no = this.userForm.get('phone_no')?.value;
    this.address = this.userForm.get('address')?.value;
    this.preference = this.userForm.get('preference')?.value;
    // alert("FROM FORM " + this.userForm.get('firstname')?.value);

    //TODO add
  }

  ngOnInit(): void {
  	// this.dataService.newCustomer(this.tempCustomer)
  }

  registerNewCustomer(): void {
  	this.newCustomer();
  	this.dataService.postCustomer(this.tempCustomer);
  }


}
