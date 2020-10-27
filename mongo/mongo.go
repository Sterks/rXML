package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Sterks/rXML/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDb Структура mongo
type MongoDb struct {
	Config  *config.Config
	MongoDb *mongo.Client
}

// MongoDBNew Новая структура для монго
func MongoDBNew(config *config.Config) *MongoDb {
	return &MongoDb{
		Config:  config,
		MongoDb: &mongo.Client{},
	}
}

// MongoDBConnect Соединение с mongo
func (m *MongoDb) MongoDBConnect(config *config.Config) {
	clientOptions := options.Client().ApplyURI(config.MongoDb.ConnectMongoDB)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	m.MongoDb = client
	fmt.Println("Connected to MongoDB!")
}

// Save ...
func (m *MongoDb) Save(databaseMongo string, collectionName string, not interface{}) {
	XMLCollection := m.MongoDb.Database(databaseMongo).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := XMLCollection.InsertOne(ctx, &not)
	if err != nil {
		log.Fatal(err)
	}

	// switch not.(type) {
	// case structs.NotificationEA615:
	// 	var ea615 structs.NotificationEA615
	// 	_, err := XMLCollection.InsertOne(ctx, &ea615)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(ea615.Pprf615NotificationEF.VersionNumber, ea615.Pprf615NotificationEF.CommonInfo.PurchaseNumber, ea615.Pprf615NotificationEF.CommonInfo.PurchaseObjectInfo)
	// case structs.NotificationEA44:
	// 	var ea44 structs.NotificationEA44
	// 	_, err := XMLCollection.InsertOne(ctx, &ea44)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(ea44.FcsNotificationEF.SchemeVersion, ea44.FcsNotificationEF.PurchaseNumber, ea44.FcsNotificationEF.PlacingWay.Name)
	// case structs.NotificationNM111:
	// 	var inm111 structs.NotificationNM111
	// 	_, err := XMLCollection.InsertOne(ctx, &inm111)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(inm111.FcsNotification111.SchemeVersion, inm111.FcsNotification111.PurchaseNumber, inm111.FcsNotification111.PlacingWay.Name)
	// case structs.NotificationOK504:
	// 	var ok504 structs.NotificationOK504
	// 	_, err := XMLCollection.InsertOne(ctx, &ok504)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(ok504.EpNotificationEOK.SchemeVersion, ok504.EpNotificationEOK.CommonInfo.PurchaseNumber, ok504.EpNotificationEOK.CommonInfo.PlacingWay.Name)
	// case structs.NotificationOKU504:
	// 	var oku504 structs.NotificationOKU504
	// 	_, err := XMLCollection.InsertOne(ctx, &oku504)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(oku504.EpNotificationEOKOU.SchemeVersion, oku504.EpNotificationEOKOU.CommonInfo.PurchaseNumber, oku504.EpNotificationEOKOU.CommonInfo.PurchaseNumber)
	// case structs.NotificationZA44:
	// 	var za44 structs.NotificationZA44
	// 	_, err := XMLCollection.InsertOne(ctx, &za44)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(za44.FcsNotificationZakA.SchemeVersion, za44.FcsNotificationZakA.PurchaseNumber, za44.FcsNotificationZakA.PlacingWay.Name)
	// case structs.NotificationZK504:
	// 	var zk504 structs.NotificationZK504
	// 	_, err := XMLCollection.InsertOne(ctx, &zk504)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(zk504.EpNotificationEZK.SchemeVersion, zk504.EpNotificationEZK.CommonInfo.PurchaseNumber, zk504.EpNotificationEZK.CommonInfo.PlacingWay.Name)
	// case structs.NotificationZKKP44:
	// 	var zkk44 structs.NotificationZKKP44
	// 	_, err := XMLCollection.InsertOne(ctx, &zkk44)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(zkk44.FcsNotificationZakK.SchemeVersion, zkk44.FcsNotificationZakK.PurchaseNumber, zkk44.FcsNotificationZakK.PlacingWay.Name)
	// case structs.NotificationZP504:
	// 	var zp504 structs.NotificationZP504
	// 	_, err := XMLCollection.InsertOne(ctx, &zp504)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(zp504.EpNotificationEZP.SchemeVersion, zp504.EpNotificationEZP.CommonInfo.PurchaseNumber, zp504.EpNotificationEZP.CommonInfo.PlacingWay.Name)
	// default:
	// 	fmt.Println("Не могу определить тип")
	// }
}
