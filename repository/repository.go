package repository

import (
	db "github.com/salihkemaloglu/DemRR-beta-001/mongodb"
)

type UserRepository interface {
	Login() error
	GetUser()(*db.User,error)
	CheckUser() error
	Insert() error
	Update() error
	Delete() error
}

type FolderRepository interface {
	GetFolder()(*db.Folder,error)
	GetAllFolders()([]db.Folder,error)
	Insert() error
	Update() error
	Delete() error
}
type FileRepository interface {
	GetFile()(*db.File,error)
	GetAllFiles()([]db.File,error)
	Insert() error
	Update() error
	Delete() error
}

type FollowRepository interface {
	GetFollower()([]db.Follow,error)
	GetFollowed()([]db.Follow,error)
	Insert() error
	Update() error
	Delete() error
}

type BuriedRepository interface {
	GetBuriedFile()(*db.File,error)
	GetAllBuriedFiles()([]db.File,error)
	Insert() error
	Update() error
	Delete() error
}