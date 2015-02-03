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
	debugPtr := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if *debugPtr {
		fmt.Println("debug mode: no message will actually be sent.")
	}
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
	batch := []string{}
	for blob := range blobs {
		fmt.Printf("send %v\n", string(blob))
		batch = append(batch, string(blob))
		if !*debugPtr {
			if len(batch) == 10 {
				if _, err := queue.SendMessageBatchString(batch); err != nil {
					panic(err)
				}
				batch = []string{}
			}
		}
		cnt++
		size += len(blob)
	}
	if len(batch) > 0 {
		if _, err := queue.SendMessageBatchString(batch); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%d messages sents (%d bytes) to %v\n", cnt, size, *queuePtr)
}
