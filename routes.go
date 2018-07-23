package nodeinfo

import (
	"encoding/json"
	"github.com/writeas/go-webfinger"
	"log"
	"net/http"
)

// NodeInfoPath defines the default path of the nodeinfo handler.
const NodeInfoPath = "/.well-known/nodeinfo"

type discoverInfo struct {
	Links []webfinger.Link `json:"links"`
}

func (s *Service) NodeInfoDiscover(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	i := discoverInfo{
		Links: []webfinger.Link{
			{
				Rel:  profile,
				HRef: s.InfoURL,
			},
		},
	}

	body, err := json.Marshal(i)
	if err != nil {
		log.Printf("Unable to marshal nodeinfo discovery: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Printf("Unable to write body: %v", err)
		return
	}
}

func (s *Service) NodeInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; profile="+profile+"#")

	i := s.BuildInfo()

	body, err := json.Marshal(i)
	if err != nil {
		log.Printf("Unable to marshal nodeinfo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Printf("Unable to write body: %v", err)
		return
	}
}
