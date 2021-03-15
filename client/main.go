package main

import (
	"context"
	"fmt"
	junk "github.com/xtreme-jesse-malone/grpc-streaming-example/proto"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"sync"
)

func main() {
	opts := grpc.WithInsecure()
	addr := os.Args[1]
	threads, _ := strconv.Atoi(os.Args[2])
	// Request forever
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			con, err := grpc.Dial(addr, opts)
			if err != nil {
				fmt.Printf("COnnection eerror %v", err)
			}

			defer con.Close()
			client := junk.NewJunkClient(con)

			stream, err := client.TakeJunk(context.Background())
			if err != nil {
				fmt.Printf("Error starting stream %v", err)
			}
			b := 9
			for {
				if b%1000000 == 0 {
					fmt.Println("Sent %v bytes ...", b)
				}
				stream.Send(&junk.JunkMsg{Junk: "Some Junk"})
				b += 9
			}
		}(&wg)
	}
	wg.Wait()
}
