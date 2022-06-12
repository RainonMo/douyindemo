package entity

type User struct{
	Id            uint64 `gorm:"primayKey" json:"id"`
	Name          string `gorm:"name" json:"name"`
	Password      string `gorm:"password"`
	Phone         string `gorm:"phone"`
	FollowCount   int64  `gorm:"follow_count" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"follower_count" json:"follower_count,omitempty"`
	IsFollow      bool   `gorm:"is_follow" json:"is_follow,omitempty"`
	CreateAt      string `gorm:"create_at"`
	UpdateAt      string `gorm:"update_at"`
}