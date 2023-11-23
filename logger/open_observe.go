package logger

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type OpenObserveWriter struct {
	URL      string
	User     string
	Password string
	Client   *http.Client
}

func (o OpenObserveWriter) Write(p []byte) (n int, err error) {
	req, err := http.NewRequest("POST", o.URL, bytes.NewReader(p))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(o.User, o.Password)
	resp, err := o.Client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		closeErr := resp.Body.Close()
		if err != nil {
			if closeErr != nil {
				log.Printf("failed to close response body: %v\n", closeErr)
			}
			return
		}
		err = closeErr
	}()
	io.Copy(io.Discard, resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println("OpenObserve status code:", resp.StatusCode, "\n log", string(p))
	}
	return
}
