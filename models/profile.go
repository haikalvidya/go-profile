package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Alamat      string `json:"alamat" gorm:"type:text"`
	Pekerjaan   string `json:"pekerjaan" gorm:"type:text"`
	NamaLengkap string `json:"nama_lengkap" gorm:"type:text"`
	Pendidikan  string `json:"pendidikan" gorm:"type:text"`
	NomorTelpon string `json:"nomer_telpon" gorm:"type:text"`
	UserID      uint
	User        User `gorm:"foreignkey:UserID"`
}
