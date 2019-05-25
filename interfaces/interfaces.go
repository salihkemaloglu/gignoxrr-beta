package interfaces

import (
	repo "github.com/salihkemaloglu/gignox-rr-beta-001/repositories"
)

type IUserRepository interface {
	Login() (*repo.User,error)
	GetUser()(*repo.User,error)
	GetUserByEmail()(*repo.User,error)
	GetUserByUsername()(*repo.User,error)
	CheckUser() error
	UpdateUserPassword() error
	Insert() error
	Update() error
	Delete() error
}

type IFolderRepository interface {
	GetFolder()(*repo.Folder,error)
	GetAllFolders()([]repo.Folder,error)
	Insert() error
	Update() error
	Delete() error
}
type IFileRepository interface {
	GetFile()(*repo.File,error)
	GetAllFiles()([]repo.File,error)
	Insert() error
	Update() error
	Delete() error
}

type IFollowRepository interface {
	IGetFollower()([]repo.Follow,error)
	GetFollowed()([]repo.Follow,error)
	Insert() error
	Update() error
	Delete() error
}

type IBuriedRepository interface {
	IGetBuriedFile()(*repo.File,error)
	GetAllBuriedFiles()([]repo.File,error)
	Insert() error
	Update() error
	Delete() error
}
type IUserTemporaryInformationRepository interface {
	CheckRegisterVerificationToken()(*repo.UserTemporaryInformation,error)
	CheckForgotPasswordVerificationToken()(*repo.UserTemporaryInformation,error)
	CheckVerificationTokenResentEmail()(*repo.UserTemporaryInformation,error)
	GetAllUserTemporaryInformation()([]repo.UserTemporaryInformation,error)
	Insert() error
	Update() error
	UpdateByEmail() error
	Delete() error
}