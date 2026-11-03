package main

import ("fmt"
	"time"
)


func main(){
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	count := 0
	fmt.Println("Starting Loop")

	for range ticker.C {
		count++
		fmt.Printf("Ticker #%d at %s\n", count, time.Now().Format("15:04:04"))

		if count >= 100 {
			break
		}

	}

	fmt.Println("Loop Finished")
}

