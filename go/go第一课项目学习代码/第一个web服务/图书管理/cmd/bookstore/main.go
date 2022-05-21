package main

import (
	"bookserver/internal/store"
	"bookserver/server"
	//"bookserver/store/factory"
	mystore "bookserver/store"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	
	
	/*s := store.Memstore{}&MemStore{
		books: make(map[string]*mystore.Book),
	}*/
	s := &store.MemStore{Books: make(map[string]*mystore.Book),}
	srv := server.NewBookStoreServer(":8080", s)
	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <- errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		return 
	}

}