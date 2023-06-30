// The database package is used to include libraries for storing data in a database
package database

import (
	"errors"
)

var ErrDatabaseConnection = errors.New("error connecting to database")
var ErrDatabaseDisconnect = errors.New("error disconnecting from database")
var ErrDatabaseDrop = errors.New("error dropping database")
var ErrInsertFailed = errors.New("insert failed")
var ErrGettingObjectIDFailed = errors.New("getting object id failed")
var ErrGeneratingObjectIDFailed = errors.New("generating object id failed")
var ErrDeleteFailed = errors.New("delete failed")
var ErrUpdateFailed = errors.New("update failed")
var ErrGetFailed = errors.New("get item failed")
var ErrDeleteCollectionFailed = errors.New("delete collection failed")
var ErrNothingFound = errors.New("nothing found")
