package xlconfig

import (
	"errors"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

var TypeError = errors.New("type error")

func handleConvertErr(err *error) {
	if r := recover(); r != nil {
		zap.S().Error(r)
		*err = TypeError
	}
}

type ApolloProvider struct {
	ConfigCache agcache.CacheInterface
}

func NewApolloProvider(appID, cluster, host, namespace, secret string) (provider Provider, err error) {
	c := &config.AppConfig{
		AppID:          appID,
		Cluster:        cluster,
		IP:             host,
		NamespaceName:  namespace,
		IsBackupConfig: true,
		Secret:         secret,
	}
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		return nil, err
	}
	cache := client.GetConfigCache(c.NamespaceName)
	if cache.EntryCount() == 0 {
		err = errors.New("config error")
	}
	provider = ApolloProvider{
		ConfigCache: cache,
	}
	return provider, err
}

func (a ApolloProvider) GetString(key ...string) (result string, err error) {
	defer handleConvertErr(&err)
	value, err := a.ConfigCache.Get(strings.Join(key, "."))
	result = value.(string)
	return
}

func (a ApolloProvider) GetBool(key ...string) (result bool, err error) {
	defer handleConvertErr(&err)
	value, err := a.ConfigCache.Get(strings.Join(key, "."))
	result = value.(bool)
	return
}

func (a ApolloProvider) GetInt(key ...string) (result int64, err error) {
	defer handleConvertErr(&err)
	value, err := a.ConfigCache.Get(strings.Join(key, "."))
	result = int64(value.(int))
	return
}

func (a ApolloProvider) GetFloat(key ...string) (result float64, err error) {
	defer handleConvertErr(&err)
	value, err := a.ConfigCache.Get(strings.Join(key, "."))
	result = value.(float64)
	return
}

func (a ApolloProvider) Populate(i interface{}, key ...string) (err error) {
	defer handleConvertErr(&err)
	value, err := a.ConfigCache.Get(strings.Join(key, "."))
	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(value))
	return
}
