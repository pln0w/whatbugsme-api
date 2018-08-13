package db

import (
	"fmt"
	"whatbugsme/infrastructure/env"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DatabaseConnection struct {
	session *mgo.Session
}

const (
	databaseName = "whatbugsme"
)

var dbC *DatabaseConnection

func init() {
	dbC = NewDatabaseConnection()
}

// NewDatabaseConnection returns pointer to DB conn structure
// having session pointer inside
func NewDatabaseConnection() *DatabaseConnection {

	// MongoDB dial up
	session, err := mgo.Dial(env.Get().Server)
	if err != nil {
		panic("Could not dial MongoDB")
	}

	// Connection handler init and return
	dbConn := &DatabaseConnection{
		session: session,
	}

	return dbConn
}

// GetCollection returns pointer to mgo Collection
// allows to query database on given collection name
func (dbConn *DatabaseConnection) GetCollection(collName string) *mgo.Collection {

	// Create new session
	session := dbConn.session.New()

	// Get collection pointer
	c := session.DB(databaseName).C(collName)

	return c
}

// Insert executes query inserting given structure to its collection
func Insert(collection string, obj interface{}) error {

	err := dbC.GetCollection(collection).Insert(&obj)
	if err != nil {
		return err
	}

	return nil
}

// ExistBy executes query counting number of items in collection
func ExistBy(collection string, params map[string]string) (int, error) {

	fields := bson.M{}

	for k, v := range params {
		fields[k] = v
	}

	n, err := dbC.GetCollection(collection).Find(fields).Count()
	if err != nil {
		return 0, err
	}

	return n, nil
}

// FindOneBy executes query looking for single record in collection
func FindOneBy(collection string, params map[string]string, ids map[string]bson.ObjectId) (interface{}, error) {

	var obj interface{}

	fields := bson.M{}

	for k, v := range params {
		fields[k] = v
	}
	for k, v := range ids {
		fields[k] = v
	}

	err := dbC.GetCollection(collection).Find(fields).One(&obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// FindAllBy executes query looking for multiple record in collection
func FindAllBy(collection string, params map[string]string, ids map[string]bson.ObjectId, s string) ([]interface{}, error) {

	var results []interface{}

	fields := bson.M{}

	for k, v := range params {
		fields[k] = v
	}
	for k, v := range ids {
		fields[k] = v
	}

	sort := "$natural"
	if s != "" {
		sort = s
	}
	err := dbC.GetCollection(collection).Find(fields).Sort(sort).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// IncrementField executes query updating particular integer field
// for all records on given items from collection
func IncrementFieldWhere(collection string, field string, value int, params map[string]string, ids map[string]bson.ObjectId) error {

	fields := bson.M{}

	for k, v := range params {
		fields[k] = v
	}
	for k, v := range ids {
		fields[k] = v
	}

	updErr := dbC.GetCollection(collection).Update(fields, bson.M{"$inc": bson.M{fmt.Sprint(field): value}})
	if updErr != nil {
		return updErr
	}

	return nil
}
