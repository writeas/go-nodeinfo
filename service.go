package nodeinfo

const (
	profileVer = "2.0"
	profile    = "http://nodeinfo.diaspora.software/ns/schema/" + profileVer
)

type Service struct {
	InfoURL string
	Info    NodeInfo

	resolver Resolver
}

func NewService(cfg Config, r Resolver) *Service {
	return &Service{
		InfoURL: cfg.InfoURL,
		Info: NodeInfo{
			Metadata:  cfg.Metadata,
			Protocols: cfg.Protocols,
			Services:  cfg.Services,
			Software:  cfg.Software,
		},
		resolver: r,
	}
}
