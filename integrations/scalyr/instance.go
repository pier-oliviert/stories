package scalyr

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/convertkit/stories/stories"
	"log"
	"net/http"
	"os"
	"time"
)

type Instance struct {
	Url         string
	SessionInfo map[string]string
	Session     uuid.UUID
	Secret      string

	configured bool
}

func (i *Instance) Configure() error {
	if i.configured == true {
		return errors.New("Instance already configured")
	}

	i.Secret = os.Getenv("SCALYR_WRITE_TOKEN")

	if len(i.Secret) == 0 {
		return errors.New("Missing SCALYR_WRITE_TOKEN required to connect to Scalyr")
	}

	i.SessionInfo = make(map[string]string)
	i.SessionInfo["logfile"] = "stories"
	i.SessionInfo["serverHost"] = "test"

	i.Url = "https://www.scalyr.com/addEvents"

	i.Session = uuid.New()
	log.Print("Session UUID: ", i.Session.String())

	i.configured = true

	return nil
}

func (i *Instance) Send(stories []*stories.Story) (*http.Response, error) {
	log.Printf("Sending %d stories.", len(stories))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	payload := NewPayload(i, stories)

	data, err := json.Marshal(payload)

	log.Print(string(data))

	if err != nil {
		log.Print("Error creating a payload to send ", err)
	}

	req, err := http.NewRequest("POST", i.Url, bytes.NewBuffer(data))

	if err != nil {
		log.Print("Error occured creating a request to Scalyr: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}
