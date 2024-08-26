package main

import (
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)

	//<-time.NewTimer(duration).C

	//ctx, cancel := context.WithTimeout(context.Background(), duration)
	//defer cancel()
	//<-ctx.Done()
}

func main() {
	sleep(time.Second)
}
