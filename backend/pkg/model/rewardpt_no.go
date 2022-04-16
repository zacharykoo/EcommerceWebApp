package model

import "gorm.io/gorm"

type Rewardpt_no struct {
	gorm.Model
	MembershipID uint `gorm:"column:membershipID" json:"membershipID"`
	Reward_no    int  `gorm:"column:reward_no" json:"reward_no"`
}
