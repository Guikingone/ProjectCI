package Cache

type CacheInterface interface {
	getItem(itemIdentifier string) ItemInterface
}
