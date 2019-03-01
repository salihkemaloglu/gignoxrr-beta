package repository

type FileRepository interface {
	Insert()(string ,error)
}