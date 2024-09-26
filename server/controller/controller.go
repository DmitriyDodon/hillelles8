package controller

import "les8/db"

type Controller struct {
	dbConnection *db.DBConnection
}

func NewController(dbConnection *db.DBConnection) *Controller {
	return &Controller{
		dbConnection: dbConnection,
	}
}
