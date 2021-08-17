package main

import (
	"context"
	"encoding/json"
	"fruit/pb"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

type Request struct {
	Num int `json:"num"`
	Name string `json:"name"`
}

func FruitNum(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	req := Request{}

	_ = json.Unmarshal(body, &req)

	c := pb.NewGRpc_2Client(conn)

	//send market grpc

	send_grpc := pb.SelectFruitNum{
		Num : int32(req.Num),
	}

	rows, _ := c.ChoiceNum(context.Background(), &send_grpc)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rows)
}

func FruitName(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	req := Request{}

	_ = json.Unmarshal(body, &req)

	c := pb.NewGRpc_2Client(conn)

	//send market grpc

	send_grpc := pb.SelectFruitName{
		Name : req.Name,
	}

	rows, _ := c.ChoiceName(context.Background(), &send_grpc)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rows)
}


func main() {

	//market client
	conn, _ = grpc.Dial("localhost:8082", grpc.WithInsecure())

	defer conn.Close()

	r := mux.NewRouter()

	r.HandleFunc("/fruitnum", FruitNum).Methods("POST")

	http.ListenAndServe("localhost:8081", r)
}