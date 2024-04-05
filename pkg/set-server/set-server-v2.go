package setserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServerV2(ctx context.Context, h http.Handler, host string, port string) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: h,
	}

	go func() {
		// log.Infof("server running at: %s:%s", host, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// log.Errorf("error servier listen and serve: %v", err)
			log.Printf("error servier listen and serve: %v", err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-gracefulStop

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		//log.Fatal("Server Shutdown: %v", err)
		log.Printf("Server Shutdown: %+v", err)
	}

	select {
	case <-ctx.Done():
		// log.Infof("timeout of 5 seconds.")
		log.Printf("timeout of 5 seconds.")
	default:
	}
	log.Printf("Server exiting")
}
