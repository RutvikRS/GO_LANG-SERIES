package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "mydb"
const colName = "TODO-LIST"

//MOST IMPORTANT

var collection *mongo.Collection

// connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB Connection Suceess")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection Database is ready")

}

//MONGO helpers - file

//insert one record

func insertOnedata(data models.TODOLIST) {
	inserted, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 todo-list in db with id: ", inserted.InsertedID)
}

// update one record
func updateOnedata(TODOID string) {
	id, _ := primitive.ObjectIDFromHex(TODOID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"clicked": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)

}

// DELETE ONE RECORD
func deleteOnedata(TODOID string) {
	id, _ := primitive.ObjectIDFromHex(TODOID)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TODO-LIST got deleted with delete count :", deleteCount)

}

// Delete all records from mongo DB

func deleteAlldata() int64 {

	deleteResults, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Numbers of TODO-LIST are deleted: ", deleteResults.DeletedCount)
	return deleteResults.DeletedCount
}

// Get all movies from the database

func GetAllData() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var TODOLISTS []primitive.M

	for cur.Next(context.Background()) {
		var data bson.M
		err := cur.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		TODOLISTS = append(TODOLISTS, data)
	}

	defer cur.Close(context.Background())
	return TODOLISTS
}

//Actual Controllers - file

func GetMyAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	alldata := GetAllData()
	json.NewEncoder(w).Encode(alldata)

}

func CreateTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var data models.TODOLIST
	_ = json.NewDecoder(r.Body).Decode(&data)
	insertOnedata(data)
	json.NewEncoder(w).Encode(data)
}

func MarkAsClicked(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOnedata(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTODOList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOnedata(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllTODOList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAlldata()
	json.NewEncoder(w).Encode(count)

}
