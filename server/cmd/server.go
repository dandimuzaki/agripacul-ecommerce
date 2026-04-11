package cmd

import (
	"context"
	"debian-ecommerce/internal/wire"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func APiserver(app *wire.App) {
	fmt.Printf("Server running on port :%s", app.Config.Port)

	// Run cron jobs
	app.Scheduler.Start()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", app.Config.Port),
		Handler: app.Route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("can't run service")
		}
	}()

	// gracefully shutdown ------------------------------------------------------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cronCtx := app.Scheduler.Stop()
	<-cronCtx.Done()

	close(app.Stop)
	app.WG.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("can't shutdown service")
	}

	ctx = app.Scheduler.Stop()
	<-ctx.Done()

	log.Println("server shutdown cleanly")
}

