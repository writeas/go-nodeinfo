package nodeinfo

type (
	NodeProtocol string
	NodeService  string
)

const (
	// Protocols that can be supported by this node.
	ProtocolActivityPub NodeProtocol = "activitypub"
	ProtocolBuddyCloud               = "buddycloud"
	ProtocolDFRN                     = "dfrn"
	ProtocolDisaspora                = "diaspora"
	ProtocolLibertree                = "libertree"
	ProtocolOStatus                  = "ostatus"
	ProtocolPumpIO                   = "pumpio"
	ProtocolTent                     = "tent"
	ProtocolXMPP                     = "xmpp"
	ProtocolZot                      = "zot"

	// Services that can be supported (inbound or outbound) by this node's API.
	ServiceTwitter NodeService = "twitter"
	ServiceTumblr              = "tumblr"
)

type Config struct {
	BaseURL   string
	InfoURL   string
	Metadata  Metadata
	Protocols []NodeProtocol
	Services  Services
	Software  SoftwareInfo
}

type (
	// NodeInfo includes all required node info.
	NodeInfo struct {
		Metadata          Metadata       `json:"metadata"`
		OpenRegistrations bool           `json:"openRegistrations"`
		Protocols         []NodeProtocol `json:"protocols"`
		Services          Services       `json:"services"`
		Software          SoftwareInfo   `json:"software"`
		Usage             Usage          `json:"usage"`
		Version           string         `json:"version"`
	}

	// Metadata for nodeinfo. Properties are based on what Pleroma uses.
	//
	// From the spec: Free form key value pairs for software specific values.
	// Clients should not rely on any specific key present.
	Metadata struct {
		NodeName        string `json:"nodeName,omitempty"`
		NodeDescription string `json:"nodeDescription,omitempty"`
		Private         bool   `json:"private,omitempty"`
	}

	Services struct {
		Inbound  []NodeService `json:"inbound"`
		Outbound []NodeService `json:"outbound"`
	}

	SoftwareInfo struct {
		// Name (canonical) of this server software.
		Name string `json:"name"`
		// Version of this server software.
		Version string `json:"version"`
	}

	// Usage is usage statistics for this server.
	Usage struct {
		Users         UsageUsers `json:"users"`
		LocalPosts    int        `json:"localPosts,omitempty"`
		LocalComments int        `json:"localComments,omitempty"`
	}

	UsageUsers struct {
		Total          int `json:"total,omitempty"`
		ActiveHalfYear int `json:"activeHalfyear,omitempty"`
		ActiveMonth    int `json:"activeMonth,omitempty"`
	}
)
