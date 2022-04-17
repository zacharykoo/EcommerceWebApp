import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ShopComponent } from './shop/shop.component';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { AccountComponent } from './account/account.component';
import { CartComponent } from './cart/cart.component';
import { LoginSignupComponent } from './account/loginsignup/loginsignup.component';
import { LoginComponent } from './account/loginsignup/login/login.component';
import { SignupComponent } from './account/loginsignup/signup/signup.component';
import { ManageComponent } from './account/manage/manage.component';
import { RewardsComponent } from './account/rewards/rewards.component';


const routes: Routes = [
	{ path: 'shop', component: ShopComponent },
	{ path: 'home', component: HomeComponent },
	{ path: 'about', component: AboutComponent },
	{ path: 'account', component: AccountComponent },
	{ path: 'account/loginsignup', component: LoginSignupComponent },
	{ path: 'account/login', component: LoginComponent },
	{ path: 'account/signup', component: SignupComponent },
	{ path: 'account/manage', component: ManageComponent },
	{ path: 'account/rewards', component: RewardsComponent },
	{ path: 'cart', component: CartComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
