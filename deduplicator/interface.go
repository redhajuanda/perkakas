package deduplicator

type IDeduplicator interface {
	Visit(key string) error
	IsVisited(key string) bool
}
