package messages

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/peterstirrup/messages/internal/messages/usecases"

	"github.com/peterstirrup/messages/internal/repositories/whatsapp"
	transport "github.com/peterstirrup/messages/internal/transport/http"
)

type environment struct {
	debug bool

	messagesapiServicePort string
}

func main() {
	env := setupEnv()

	log.Println("creating whatsapp client")
	c := http.Client{Timeout: 10 * time.Second}
	whatsapp, err := whatsapp.New(&c, "https://whatsapp.com/api/v1")
	if err != nil {
		log.Fatalf("failed to creat whatsapp client with err %s", err)
	}

	log.Println("creating usecases")
	u := usecases.New(whatsapp)

	log.Println("creating server")
	s := transport.NewServer(u)
	handler := s.NewHandler()

	log.Println(fmt.Sprintf("server started on [::]:%s", env.messagesapiServicePort))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", env.messagesapiServicePort), handler); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func setupEnv() environment {
	env := environment{}

	if debug := os.Getenv("DEBUG"); debug != "" {
		env.debug = true
	}

	if env.messagesapiServicePort = os.Getenv("MESSAGES_API_SERVICE_PORT"); env.messagesapiServicePort == "" {
		panic("MESSAGES_API_SERVICE_PORT not set in environment")
	}

	return env
}
