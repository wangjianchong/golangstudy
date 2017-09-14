package main

import mgo "v1/mgo.v2"

func init() {
	mongoInit()
}

type mongoConfig struct {
	session *mgo.Session
	db      *mgo.Database
}

func mongoInit() error {
	session, err := mgo.Dial("")
	if err != nil {
		return err
	}
	mongo.session = session
	mongo.db = session.DB("Adata")
	return nil
}
