package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"
const ServicesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.RWMutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	err := r.SendRequiredServices(reg)
	r.nofity(patch{
		Added: []patchEntry{
			{
				Name: reg.ServiceName,
				URL:  reg.ServiceURL,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r registry) nofity(fullPatch patch) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, reg := range r.registrations {
		go func(reg Registration) {
			for _, reqService := range reg.RequiredServices {
				p := patch{Added: []patchEntry{}, Removed: []patchEntry{}}
				sendUpdate := false
				for _, added := range fullPatch.Added {
					if added.Name == reqService {
						p.Added = append(p.Added, added)
						sendUpdate = true
					}
				}

				for _, removed := range fullPatch.Removed {
					if removed.Name == reqService {
						p.Removed = append(p.Removed, removed)
						sendUpdate = true
					}
				}

				if sendUpdate {
					err := r.sendPatch(p, reg.ServiceUpdateURL)
					if err != nil {
						log.Println(err)
						return
					}
				}

			}
		}(reg)
	}
}

func (r registry) SendRequiredServices(reg Registration) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var p patch
	for _, serviceReg := range r.registrations {
		for _, reqservice := range serviceReg.RequiredServices {
			if serviceReg.ServiceName == reqservice {
				p.Added = append(p.Added, patchEntry{
					Name: serviceReg.ServiceName,
					URL:  serviceReg.ServiceURL,
				})
			}
		}
	}

	err := r.sendPatch(p, reg.ServiceUpdateURL)
	if err != nil {
		return err
	}
	return nil
}

func (r registry) sendPatch(p patch, url string) error {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return err
	}
	return nil
}

var reg = registry{registrations: make([]Registration, 0),
	mutex: new(sync.RWMutex),
}

type RegistryService struct {
	http.Handler
}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder((r.Body))
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding Service: %v with URL: %v\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		url := string(payload)
		log.Printf("Remove service at URL %v", url)
		err = reg.remove(url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r *registry) remove(url string) error {
	for i := range r.registrations {
		if r.registrations[i].ServiceURL == url {
			r.nofity(patch{
				Removed: []patchEntry{
					{
						Name: r.registrations[i].ServiceName,
						URL:  r.registrations[i].ServiceURL,
					},
				},
			})
			r.mutex.Lock()
			r.registrations = append(r.registrations[:i], reg.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at URL %v not found", url)
}
