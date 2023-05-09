package db

import (
	"Server/application/api"
	"Server/zlog"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoManager struct {
	mongoClient   *mongo.Client
	mongoUserInfo *mongo.Collection
	mongoCounters *mongo.Collection
	logger        *zlog.Logger
}

const (
	mongoURI      = "mongodb://localhost:27017"
	dbName        = "hasdb"
	usersCollName = "user_info"
	countersName  = "counters"
)

var Mongomanager *MongoManager

func init() {
	Mongomanager = &MongoManager{
		logger: zlog.NewLogger("[Mongodb]"),
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	collection := client.Database(dbName).Collection(usersCollName)
	counters := client.Database(dbName).Collection(countersName)
	if err != nil {
		panic(err)
	}
	Mongomanager.mongoClient = client
	Mongomanager.mongoUserInfo = collection
	Mongomanager.mongoCounters = counters
	filter := bson.M{"_id": "pid"}
	err = Mongomanager.mongoUserInfo.FindOne(context.Background(), filter).Decode(api.UserInfo{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		doc := bson.D{
			{"_id", "pid"},
			{"counter", 10000},
		}
		Mongomanager.mongoCounters.InsertOne(context.Background(), doc)
	}
}

func (mgdb *MongoManager) FindUserInfoByUsername(username string) (*api.UserInfo, error) {
	userInfo := &api.UserInfo{}
	filter := bson.M{"username": username}
	err := Mongomanager.mongoUserInfo.FindOne(context.Background(), filter).Decode(userInfo)
	if err != nil {
		mgdb.logger.Error("Find Username error:" + err.Error())
		return nil, err
	}
	return userInfo, nil
}

func (mgdb *MongoManager) CreateCounter(counterId string, counter int32) {
	doc := bson.D{
		{"_id", counterId},
		{"counter", counter},
	}
	_, err := mgdb.mongoCounters.InsertOne(context.Background(), doc)
	if err != nil {
		mgdb.logger.Error("Create counter error:" + err.Error())
	}

}
func (mgdb *MongoManager) GetNextPID(counterID string) (*api.Counters, error) {
	filter := bson.M{"_id": counterID}
	update := bson.M{"$inc": bson.M{"counter": 1}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	counter := &api.Counters{}
	err := mgdb.mongoCounters.FindOneAndUpdate(context.Background(), filter, update, options).Decode(counter)
	if err != nil {
		mgdb.logger.Error("Get next pid error:" + err.Error())
		return nil, err
	}
	return counter, nil
}

func (mgdb *MongoManager) UpdatePassword(username string, password string, newpassword string) error {
	filter := bson.M{"username": username}
	tempinfo := &api.UserRegInfo{}
	err := Mongomanager.mongoUserInfo.FindOne(context.Background(), filter).Decode(tempinfo)
	if err != nil {
		mgdb.logger.Error("Update password error:" + err.Error())
		return err
	}
	if password == tempinfo.Password {
		update := bson.M{"password": newpassword}
		opts := options.Update().SetUpsert(false)
		_, err := mgdb.mongoUserInfo.UpdateOne(context.Background(), filter, update, opts)
		if err != nil {
			mgdb.logger.Error("Update password error:" + err.Error())
			return err
		}
	} else {
		mgdb.logger.Error("Update password error:" + err.Error())
		return errors.New("old password incorrect")
	}
	return err
}

func (mgdb *MongoManager) Register(username string, password string) error {
	filter := bson.M{"username": username}
	err := Mongomanager.mongoUserInfo.FindOne(context.Background(), filter).Decode(api.UserInfo{})
	if errors.Is(err, mongo.ErrNoDocuments) {
		pid, err := mgdb.GetNextPID("pid")
		if err != nil {
			mgdb.logger.Error("Register error:" + err.Error())
			return err
		}
		doc := bson.D{
			{"pid", pid.Counter},
			{"username", username},
			{"password", password},
		}
		_, err = mgdb.mongoUserInfo.InsertOne(context.Background(), doc)
		if err != nil {
			mgdb.logger.Error("Register error:" + err.Error())
			return err
		}
	} else {
		return api.ErrUserExists
	}
	return nil
}
