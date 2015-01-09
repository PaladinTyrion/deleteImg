package config

import (
	"time"
)

// MongoDB connection information.
const (
	Mongodbhosts   = "localhost:27017"
	Database       = "express4_bootstrap_starter"
	Collection     = "tricks"
	Username       = ""
	Password       = ""
	Timeout        = time.Second * 5
	DeletePathName = "/Users/paladin/workplace/express4-bootstrap-starter/public/screenshot/"
)
