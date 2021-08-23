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
	flag.StringVar(&conf.ProjectID, "project-id", "", "The WorkOS project id.")
	flag.StringVar(&conf.Directory, "directory", "", "The WorkOS directory id.")
	flag.Parse()

  
	//ENDPOINTS
	//users
	//groups
	
	// Configure the WorkOS SSO SDK:
	directorysync.SetAPIKey(conf.APIKey)

		// Handle login redirect:
		http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			//log.Printf("callback is called with %s", r.URL)
	
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

