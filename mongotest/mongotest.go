package main

import (
	"strconv"
	"time"

	mgo "v1/mgo.v2"
	"v1/mgo.v2/bson"
)

var waitChan = make(chan int, 10)

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		Error(err)
	}
	Debug(session)
	str, err := session.DatabaseNames()
	if err != nil {
		Error(err)
	}
	Debug(str)

	c := session.DB("mongotest").C("mongocol")

	// err = c.Insert(&User{
	// 	Id_:       bson.NewObjectId(),
	// 	Name:      "Jimmy Kuu",
	// 	Age:       33,
	// 	JonedAt:   time.Now(),
	// 	Interests: []string{"Develop", "Movie"},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	err = c.Insert(OldSchool{
		Num:  01,
		Age:  11,
		Akk:  "akk is a note",
		Name: "Daming",
	})
	if err != nil {
		Error(err)
	}
	var users []User
	err = c.Find(nil).All(&users)
	if err != nil {
		Error(err)
	}
	Debug(len(users))

	err = c.Find(bson.M{"_id": bson.ObjectIdHex("598bf384a0bc496150000001")}).All(&users)
	if err != nil {
		Error(err)
	}
	Debug(users)
	Debug(len(users))

	err = c.Update(bson.M{"_id": bson.ObjectIdHex("598bf384a0bc496150000001")}, bson.M{"$set": bson.M{"name": "wwwwwww"}, "$push": bson.M{"interests": "golang"}})
	if err != nil {
		Error(err)
	}
	var u User
	q := c.Find(bson.M{"_id": bson.ObjectIdHex("598bf384a0bc496150000001")})
	q.One(&u)
	Debug(u)

	for i := 0; i < 10; i++ {
		go TestInsert(session)
	}
	for i := 0; i < 10; i++ {
		Debug(i)
		<-waitChan
	}
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	if err = c.EnsureIndex(index); err != nil {
		Error(err)
	}
}

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	JonedAt   time.Time     `bson:"joned_at"`
	Interests []string      `bson:"interests"`
}

type OldSchool struct {
	Num  int `bson:"num`
	Age  int
	Akk  string `bson:"akk"`
	Name string `bson:"name"`
}

func TestInsert(s *mgo.Session) {
	session := s.Copy()
	c := session.DB("mongotest").C("mongocolt")
	for i := 0; i < 100; i++ {
		c.Insert(bson.M{"name": "naem" + strconv.Itoa(i), "age": i})

	}
	waitChan <- 1

}
