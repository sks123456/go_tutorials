package main

import (
	"fmt"
	"sync"
	"time"
)
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{}
func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		// dbCall(i)	//running without concurrency
		wg.Add(1)		//adding increment into counter waitgroup
		go dbCall(i) 	//concurrency happen
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
}

func dbCall(i int){
	//simulate db call delay 
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	wg.Done()		//flagging completion in a task
}

func save(result string){
	m.Lock()
	results = append(results,result)
	m.Unlock()
}

func log(){
	m.RLock() //read only
	fmt.Printf("\nThe Current results are: %v",results)
	m.RUnlock()
}