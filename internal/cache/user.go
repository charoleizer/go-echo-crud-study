package cache

import "github.com/charoleizer/go-echo-crud-study/internal/models"

var (
	Users = map[int]*models.User{}
	Seq   = 1
)
