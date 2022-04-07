import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { faCartShopping } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Ecommerce';
  faCartShopping = faCartShopping;
}
