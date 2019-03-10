package repository

import (
	db "github.com/salihkemaloglu/DemRR-beta-001/mongodb"
)

type FileRepository interface {
	GetFile()(db.File,error)
	GetAllFiles()(db.File,error)
	Insert() error
	Update() error
	Delete() error
}
type UserRepository interface {
	Login()(bool,error)
	GetUser()(db.User,error)
	Insert() error
	Update() error
	Delete() error
}