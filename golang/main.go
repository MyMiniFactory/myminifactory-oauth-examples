package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	BaseURL  string
	Redirect string
	Key      string
	Secret   string
	State    string
}

type Server struct {
	Port string
	Cnf  Config
}

func (c Config) get() Config {
	return Config{
		BaseURL:  "https://auth.myminifactory.com",
		Redirect: "http://localhost:9090/callback",
		Key:      "YOUR_CLIENT_KEY",
		Secret:   "YOUR_CLIENT_SECRET",
		State:    "random_state_1994",
	}
}

func (c *Config) oauthURLBuilder() string {
	return fmt.Sprintf(
		"%s/web/login?client_id=%s&login_redirect_uri=/web/authorize&redirect_uri=%s&response_type=code&state=%s",
		c.BaseURL,
		c.Key,
		c.Redirect,
		c.State,
	)
}

func (s Server) index(w http.ResponseWriter, r *http.Request) {
	url := s.Cnf.oauthURLBuilder()
	content := fmt.Sprintf("<a href='%s'>Login</a>", url)
	fmt.Fprintf(w, content)
	log.Info("Route / -> handled")
}

func (s Server) callback(w http.ResponseWriter, r *http.Request) {
	url := s.Cnf.oauthURLBuilder()
	log.Info(r.URL.Query().Get("code"))
	content := fmt.Sprintf("<a href='%s'>Logins</a>", url)
	fmt.Fprintf(w, content)
	log.Info("Route /callback -> handled")
}

func (s Server) Start() {
	log.Infof("Starting server... (port: '%s')", s.Port)
	http.HandleFunc("/", s.index)
	http.HandleFunc("/callback", s.callback)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		log.Error("ListenAndServe: ", err)
	}
}

func main() {
	config := Config{}
	server := Server{
		Port: ":9090",
		Cnf:  config.get(),
	}
	server.Start()
}
