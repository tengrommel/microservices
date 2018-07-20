package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Job struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Company string `json:"company"`
	Salary string `json:"salary"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/jobs", jobsPostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}

func jobsPostHandler(w http.ResponseWriter, r *http.Request)  {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		panic(err)
	}
	var _job Job
	err = json.Unmarshal(b, &_job)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	saveJobToKafka(_job)

	// Convert job struct into json
	jsonString, err := json.Marshal(_job)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)
}

func saveJobToKafka(job Job)  {
	fmt.Println("save to kafka")
	jsonString, err :=json.Marshal(job)
	jobString := string(jsonString)
	fmt.Print(jobString)

	p, err :=kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil{
		panic(err)
	}
	topic := "jobs-topic1"
	for _, word :=range []string{string(jobString)}{
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value: []byte(word),
		}, nil)
	}
}
