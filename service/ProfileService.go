package service

import (
	"go_database/dao"
	"go_database/entity"
)

type ProfileService struct {
	
}

func NewProfileService()  {

}

var profileDao = new(dao.ProfileDao)

func Login(profile *entity.Profile)  {

}