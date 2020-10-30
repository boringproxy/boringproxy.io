package main

import (
	//"errors"
	"fmt"
        "log"
	//"net/smtp"
        "net/http"
        "time"
        "strings"
        "io"
        "io/ioutil"
        "encoding/json"
        "sync"
        "errors"
        "net/smtp"
	"crypto/rand"
	"math/big"
)

type Auth struct {
	pendingRequests map[string]*LoginRequest
	mutex           *sync.Mutex
}

type LoginRequest struct {
	Email string
}

func NewAuth() *Auth {

	pendingRequests := make(map[string]*LoginRequest)
	mutex := &sync.Mutex{}

	return &Auth{pendingRequests, mutex}
}

type Config struct {
        SignupServer string `json:"signup_server"`
        ProxyServer string `json:"proxy_server"`
        Token string
        Smtp *SmtpConfig
}

type SmtpConfig struct {
        Server string
        Username string
        Password string
        Port int
}

func main() {

        var config *Config
        configBytes, err := ioutil.ReadFile("email_signup_config.json")
        if err != nil {
                log.Fatal(err)
        }

        err = json.Unmarshal(configBytes, &config)
        if err != nil {
                log.Fatal(err)
        }

        auth := NewAuth()


        http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now().Format(time.RFC3339)
		srcIp := strings.Split(r.RemoteAddr, ":")[0]
		fmt.Println(fmt.Sprintf("%s %s %s %s %s", timestamp, srcIp, r.Method, r.Host, r.URL.Path))

                if r.Method != "POST" {
                        w.WriteHeader(405)
                        io.WriteString(w, "Invalid method")
                        return
                }

                r.ParseForm()

                email := r.Form.Get("email")

                if email == "" {
                        w.WriteHeader(400)
                        io.WriteString(w, "Invalid email parameter")
                        return
                }

                _, err := auth.Login(email, config)
                if err != nil {
                        w.WriteHeader(400)
                        io.WriteString(w, err.Error())
                        return
                }

                http.Redirect(w, r, "https://boringproxy.io", 303)
	})

        http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now().Format(time.RFC3339)
		srcIp := strings.Split(r.RemoteAddr, ":")[0]
		fmt.Println(fmt.Sprintf("%s %s %s %s %s", timestamp, srcIp, r.Method, r.Host, r.URL.Path))

                if r.Method != "GET" {
                        w.WriteHeader(405)
                        io.WriteString(w, "Invalid method")
                        return
                }

                r.ParseForm()

                key := r.Form.Get("key")

                if key == "" {
                        w.WriteHeader(400)
                        io.WriteString(w, "Invalid key parameter")
                        return
                }

                email , err := auth.Verify(key)
                if err != nil {
                        w.WriteHeader(400)
                        io.WriteString(w, err.Error())
                        return
                }

                apiUrl := fmt.Sprintf("https://%s/api", config.ProxyServer)
                createUserUrl := fmt.Sprintf("%s/users/?access_token=%s&username=%s", apiUrl, config.Token, email)
                fmt.Println(createUserUrl)

                resp, err := http.Post(createUserUrl, "application/json", nil)
                if err != nil {
                        w.WriteHeader(500)
                        io.WriteString(w, err.Error())
                        return
                }
                defer resp.Body.Close()

                if resp.StatusCode != 200 {
                        w.WriteHeader(500)
                        body, _ := ioutil.ReadAll(resp.Body)
                        w.Write(body)
                        return
                }

                createTokenUrl := fmt.Sprintf("%s/tokens/?access_token=%s&owner=%s", apiUrl, config.Token, email)
                resp, err = http.Post(createTokenUrl, "application/json", nil)
                if err != nil {
                        w.WriteHeader(500)
                        io.WriteString(w, err.Error())
                        return
                }
                defer resp.Body.Close()

                token, _ := ioutil.ReadAll(resp.Body)

                if resp.StatusCode != 200 {
                        w.WriteHeader(500)
                        w.Write(token)
                        return
                }

                redirUrl := fmt.Sprintf("https://%s/login?access_token=%s", config.ProxyServer, string(token))
                fmt.Println(redirUrl)
                http.Redirect(w, r, redirUrl, 303)
	})

        log.Println("Running")

        http.ListenAndServe(":3000", nil)
}

func (a *Auth) Login(email string, config *Config) (string, error) {

	key, err := genRandomKey()
	if err != nil {
		return "", errors.New("Error generating key")
	}

	link := fmt.Sprintf("https://%s/verify?key=%s", config.SignupServer, key)

	bodyTemplate := "From: %s <%s>\r\n" +
		"To: %s\r\n" +
		"Subject: Email Verification\r\n" +
		"\r\n" +
		"This is email verification request from %s. Please click the following link to complete the verification:\r\n" +
		"\r\n" +
		"%s\r\n"

	fromText := "boringproxy email verifier"
	fromEmail := fmt.Sprintf("auth@%s", config.SignupServer)
	emailBody := fmt.Sprintf(bodyTemplate, fromText, fromEmail, email, config.SignupServer, link)

	emailAuth := smtp.PlainAuth("", config.Smtp.Username, config.Smtp.Password, config.Smtp.Server)
	srv := fmt.Sprintf("%s:%d", config.Smtp.Server, config.Smtp.Port)
	msg := []byte(emailBody)
	err = smtp.SendMail(srv, emailAuth, fromEmail, []string{email}, msg)
	if err != nil {
		return "", errors.New("Sending email failed. Probably a bad email address.")
	}

	a.mutex.Lock()
	a.pendingRequests[key] = &LoginRequest{email}
	a.mutex.Unlock()

	return "", nil
}

func (a *Auth) Verify(key string) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	req, ok := a.pendingRequests[key]

	if !ok {
		return "", errors.New("No pending request for that key. It may have expired.")
	}

	delete(a.pendingRequests, key)

	return req.Email, nil
}

const chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genRandomKey() (string, error) {
	id := ""
	for i := 0; i < 32; i++ {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		id += string(chars[randIndex.Int64()])
	}
	return id, nil
}
