package resources

// Resource interface
type Resource interface {
	Marshal() ([]byte, error)
}
