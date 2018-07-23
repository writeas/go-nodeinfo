package main

import (
	"github.com/writeas/go-nodeinfo"
	"net/http"
)

func main() {
	cfg := nodeinfo.Config{
		BaseURL: "http://localhost:8080",
		InfoURL: "/api/nodeinfo",

		Metadata: nodeinfo.Metadata{
			NodeName:        "Agora",
			NodeDescription: "A federated something-something.",
			Private:         false,
		},
		Protocols: []nodeinfo.NodeProtocol{
			nodeinfo.ProtocolActivityPub,
		},
		Services: nodeinfo.Services{
			Inbound: []nodeinfo.NodeService{},
			Outbound: []nodeinfo.NodeService{
				nodeinfo.ServiceTwitter,
				nodeinfo.ServiceTumblr,
			},
		},
		Software: nodeinfo.SoftwareInfo{
			Name:    "Agora",
			Version: "1.0",
		},
	}
	ni := nodeinfo.NewService(cfg, nodeInfoResolver{})

	http.Handle(nodeinfo.NodeInfoPath, http.HandlerFunc(ni.NodeInfoDiscover))
	http.Handle(cfg.InfoURL, http.HandlerFunc(ni.NodeInfo))

	http.ListenAndServe(":8080", nil)
}

type nodeInfoResolver struct{}

func (r nodeInfoResolver) IsOpenRegistration() (bool, error) {
	return true, nil
}

func (r nodeInfoResolver) Usage() (nodeinfo.Usage, error) {
	u := nodeinfo.Usage{
		Users: nodeinfo.UsageUsers{
			Total: 1,
		},
		LocalPosts: 1,
	}
	return u, nil
}
