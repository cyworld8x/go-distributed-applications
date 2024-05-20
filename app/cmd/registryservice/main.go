package registryservice

import (
	"app/registry"
	"context"
	"log"
	"net/http"
)

func main() {
	registry.SetupRegistryService()

	http.Handle("/services",
		&registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println((srv.ListenAndServe()))
		cancel()
	}()
	<-ctx.Done()
}
