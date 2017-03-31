package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

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
}

func (s Server) callback(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	content := "Error happened"
	if query.Get("state") == s.Cnf.State {
		content = s.authorize(query.Get("code"))
	}
	fmt.Fprintf(w, content)
}

func (s Server) authorize(code string) string {
	bodyString := ""
	auth := fmt.Sprintf("Basic %s", basicAuth(s.Cnf.Key, s.Cnf.Secret))
	resource := "/v1/oauth/tokens"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Add("redirect_uri", s.Cnf.Redirect)
	data.Add("code", code)

	u, _ := url.ParseRequestURI(s.Cnf.BaseURL)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Authorization", auth)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	if resp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}
		bodyString = string(bodyBytes)
	}
	return bodyString
}

// Start service
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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
