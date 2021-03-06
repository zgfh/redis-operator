// This file was automatically generated by informer-gen

package v1

import (
	redis_v1 "github.com/jw-s/redis-operator/pkg/apis/redis/v1"
	clientset "github.com/jw-s/redis-operator/pkg/generated/clientset"
	internalinterfaces "github.com/jw-s/redis-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/jw-s/redis-operator/pkg/generated/listers/redis/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// RedisInformer provides access to a shared informer and lister for
// Redises.
type RedisInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RedisLister
}

type redisInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewRedisInformer constructs a new informer for Redis type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRedisInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.RedisV1().Redises(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.RedisV1().Redises(namespace).Watch(options)
			},
		},
		&redis_v1.Redis{},
		resyncPeriod,
		indexers,
	)
}

func defaultRedisInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewRedisInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *redisInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&redis_v1.Redis{}, defaultRedisInformer)
}

func (f *redisInformer) Lister() v1.RedisLister {
	return v1.NewRedisLister(f.Informer().GetIndexer())
}
