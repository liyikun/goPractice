package main

import (
	"log"
	"io"
	"sync/atomic"
	"sync"
	"./pool"
	"time"
	"math/rand"
	"./work"
	"./logger"
)

func main()  {
	//goruntine.RunnerMain()
	//entitles.MutixMain()
	//PoolMain()
	//workMain()
	logger.LogMain()
}


const (
	maxGoroutines = 20
	pooledResources = 2
)

type dbConnection struct {
	id uint32
}

func (c *dbConnection) Close() error {
	log.Println("Close: connection", c.id)
	return nil
}

var intCounter uint32

func createConnection() (io.Closer, error) {
	id := atomic.AddUint32(&intCounter,1)

	return &dbConnection{id}, nil
}

func PoolMain()  {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection,pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0;query < maxGoroutines;query ++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program")

	p.Close()

}

func performQueries(query int,p *pool.Pool)  {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n",query,conn.(*dbConnection).id)
}




var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m  *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func workMain()  {

	p := work.New(2)

	var wg2 sync.WaitGroup
	wg2.Add(100 * len(names))

	for i:=0 ;i<100;i++{
		for _, name := range names{
			np := namePrinter{
				name,
			}

			go func() {
				p.Run(&np)
				wg2.Done()
			}()
		}
	}

	wg2.Wait()

	p.Shutdown()
}

