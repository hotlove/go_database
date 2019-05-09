package service

import (
	"go_database/dao"
	"go_database/entity"
)

type ProfileService struct {
}

func NewProfileService() (profileService *ProfileService) {
	return &ProfileService{}
}

var profileDao = new(dao.ProfileDao)

func Login(profile *entity.Profile) {

}
