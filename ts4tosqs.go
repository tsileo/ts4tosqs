package main

import (
	"flag"
	"fmt"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/sqs"
	"github.com/tsileo/ts4/client"
)

func main() {
	startPtr := flag.String("start", "", "start query")
	endPtr := flag.String("end", "", "end query")
	queuePtr := flag.String("queue", "", "SQS queue")
	flag.Parse()
	fmt.Printf("%v/%v/%v", *startPtr, *endPtr, *queuePtr)
	bs := ts4.New("")
	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}
	SQS := sqs.New(auth, aws.USEast)
	queue, err := SQS.GetQueue(*queuePtr)
	if err != nil {
		panic(err)
	}
	blobs := make(chan []byte)
	go bs.Iter(*startPtr, *endPtr, blobs)
	cnt := 0
	size := 0
	for blob := range blobs {
		if _, err := queue.SendMessage(string(blob)); err != nil {
			panic(err)
		}
		cnt++
		size += len(blob)
	}
	fmt.Printf("%d messages sents (%d bytes) to %v", cnt, size, *queuePtr)
}
