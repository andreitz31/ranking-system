package playerstorage

import (
	"ranking/models"
	"sync"
)

//PlayerStorage The storage of the Players in game
type PlayerStorage struct {
	models.PlayersList
}

var mut sync.Mutex
var storage *PlayerStorage

// GetInstance returns the PlayerStorage singleton object
func GetInstance() *PlayerStorage {

	// We add a mutex for this method to be thread safe
	mut.Lock()
	defer mut.Unlock()
	if storage == nil {
		storage = &PlayerStorage{}
	}
	return storage
}
