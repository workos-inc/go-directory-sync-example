package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"github.com/workos-inc/workos-go/pkg/directorysync"
)

func main() {
	var conf struct {
		Addr        string
		APIKey      string
		ProjectID   string
		Directory string
	}

	flag.StringVar(&conf.Addr, "addr", ":3042", "The server addr.")
	flag.StringVar(&conf.APIKey, "api-key", "", "The WorkOS API key.")
	flag.StringVar(&conf.Directory, "directory", "", "The WorkOS directory id.")
	flag.Parse()
	
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	// Configure the WorkOS directory sync SDK:
	directorysync.SetAPIKey(conf.APIKey)

		// Handle users redirect:
		http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
	
			// Retrieving user profile:
			users, err := directorysync.ListUsers(context.Background(), directorysync.ListUsersOpts{
				Directory: conf.Directory,
			})
			if err != nil {
				log.Printf("get list users failed: %s", err)
	
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
	
			// Display Lists:
			b, err := json.MarshalIndent(users, "", "    ")
			if err != nil {
				log.Printf("encoding list users failed: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.Write(b)
	
			log.Printf("user list: %s", b)
		})
	
		if err := http.ListenAndServe(conf.Addr, nil); err != nil {
			log.Panic(err)
		}

	
	
}

