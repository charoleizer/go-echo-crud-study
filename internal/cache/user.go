package cache

import "github.com/charoleizer/go-echo-crud-study/internal/models"

// Temporarily store the data
var (
	Users = map[int]*models.User{}
	Seq   = 1
)
