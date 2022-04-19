import { Component, OnInit } from '@angular/core';
import { DataService } from '../../_services/data.service';

@Component({
  selector: 'app-rewards',
  templateUrl: './rewards.component.html',
  styleUrls: ['./rewards.component.css']
})
export class RewardsComponent implements OnInit {

 rewardPointMembers: any[] = [];
 rewards: any[] = [];
  constructor(private dataService: DataService) { }

  ngOnInit(): void {
  	this.dataService.getRewardPointNumber().subscribe((data: any[])=>{
  		console.log(data);
  		this.rewardPointMembers = data;
  	})
  	this.dataService.getRewards().subscribe((data: any[])=>{
  		console.log(data);
  		this.rewards = data;
  	})
  }

}