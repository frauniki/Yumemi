package scope

type MirakurunScoper interface {
	Name() string
	Endpoint() string
	IsDefault() bool

	PatchObject() error
	Close() error
}
