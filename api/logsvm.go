package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Details Memory
type Memory struct {
	Total      float64 `json:"total" bson:"total"`
	Inuse      float64 `json:"inuse" bson:"inuse"`
	Percentage float64 `json:"percentage" bson:"percentage"`
	Free       float64 `json:"free" bson:"free"`
}

// Log from Virtual Machine
type Log struct {
	VmName   string    `json:"vmname" bson:"vmname"`
	Endpoint string    `json:"endpoint" bson:"enpoint"`
	Data     Memory    `json:"data" bson:"data"`
	Date     time.Time `json:"date" bson:"date"`
}

var dbLogs *mongo.Collection
var ctx = context.TODO()

// LogsHandleFunc to be used as http.HandleFunc for User API
func LogsHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		logsobj := FromJSON(body)
		_, err = dbLogs.InsertOne(ctx, logsobj)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("New data inserted")
		}
		writeJSON(w, logsobj)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func Connectdb() bool {

	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:1234@cluster0.fac6w.mongodb.net/User?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	dbLogs = client.Database("User").Collection("user")

	if dbLogs != nil {
		fmt.Println("DB CONNECTED")
		return true
	} else {
		fmt.Println("DB NONE CONNECTED")
		return false
	}
}

// ToJSON to be used for marshalling of User type
func (u Log) ToJSON() []byte {
	ToJSON, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Log {
	user := Log{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	return user
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
