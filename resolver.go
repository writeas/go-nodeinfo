package nodeinfo

type Resolver interface {
	// IsOpenRegistration returns whether or not registration is open on this node.
	IsOpenRegistration() (bool, error)
	// Usage returns usage stats for this node.
	Usage() (Usage, error)
}
