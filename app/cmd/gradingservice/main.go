package gradingservice

import (
	"app/log"
	"app/registry"
	"app/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	host, port := "gradingservice", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		"Grading Service",
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
