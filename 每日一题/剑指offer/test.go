package main
import (
	"sync"
	"time"
)
func main() {
	var wg sync.WaitGroup
	//wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Add(1)
		wg.Done()
		//wg.Add(1)
	}()
	wg.Wait()
}