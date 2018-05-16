package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

type Product struct {
	Name        string
	Lecturer    string
	Course      string
	Id          string
	Tags        []string
	Groups      []string
	ContentURIs []string
}

type Database struct {
	lections *mgo.Collection
}

func NewDatabase() (db *Database) {
	return &Database{}
}

func (db *Database) Connect(user, password, dbName, ip string, port int) {
	session, err := mgo.Dial(fmt.Sprintf("%v:%v@%v:%v/%v",
		user,
		password,
		ip,
		port,
		dbName))
	if err != nil {
		panic(err)
		utils.LogCritical(err, "start MongoDB session")
	}
	session.SetMode(mgo.Monotonic, true)
	db.lections = session.DB(dbName).C("lections")
}

func (db *Database) GetLectionUrls(start, quantity int, sort string) (ids []string, err error) {
	urlStructs := []struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	}{}

	err = db.lections.
		Find(bson.M{}).
		Sort(sort).
		Limit(quantity).
		Skip(start).
		All(&urlStructs)
	if err != nil {
		return ids, err
	}

	for _, item := range urlStructs {
		ids = append(ids, item.Id.Hex())
	}
	return ids, err
}

func (db *Database) GetLectionSingle(id string) (product Product, err error) {
	err = db.lections.
		Find(bson.M{
			"_id": bson.ObjectIdHex(id),
		}).
		One(&product)
	product.Id = id
	return product, err
}

func (db *Database) GetLectionMultiple(ids []string) (products []Product, err error) {
	objIds := []bson.ObjectId{}

	for _, id := range ids {
		objIds = append(objIds, bson.ObjectIdHex(id))
	}

	err = db.lections.
		Find(bson.M{
			"_id": bson.M{
				"$in": objIds,
			},
		}).
		All(&products)

	for index := range products {
		products[index].Id = ids[index]
	}

	return products, err
}

func (db *Database) AddLection(product Product) (err error) {
	return db.lections.Insert(&product)
}

func (db *Database) RemoveLection(id string) (err error) {
	return db.lections.
		Remove(bson.M{
			"_id": bson.ObjectIdHex(id),
		})
}

func (db *Database) ChangeNameLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$set": bson.M{
				"name": name,
			},
		})
}

func (db *Database) ChangeLecturerLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$set": bson.M{
				"lecturer": name,
			},
		})
}

func (db *Database) ChangeCourseLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$set": bson.M{
				"course": name,
			},
		})
}

func (db *Database) AddTagLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$push": bson.M{
				"tags": name,
			},
		})
}

func (db *Database) RemoveTagLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$pull": bson.M{
				"tags": name,
			},
		})
}

func (db *Database) AddGroupLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$push": bson.M{
				"groups": name,
			},
		})
}

func (db *Database) RemoveGroupLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$pull": bson.M{
				"grops": name,
			},
		})
}

func (db *Database) AddContentLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$push": bson.M{
				"contenturis": name,
			},
		})
}

func (db *Database) RemoveContentLection(id, name string) (err error) {
	return db.lections.
		Update(bson.M{
			"_id": bson.ObjectIdHex(id),
		}, bson.M{
			"$pull": bson.M{
				"contenturis": name,
			},
		})
}
