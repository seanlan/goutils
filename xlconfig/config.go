package xlconfig

import "os"

var GlobalConfig Provider

func Setup() {
	var err error
	switch os.Getenv("APP_CONF_TYPE") {
	case "apollo":
		appid := os.Getenv("APOLLO_APP_ID")
		cluster := os.Getenv("APOLLO_CLUSTER")
		host := os.Getenv("APOLLO_HOST")
		namespace := os.Getenv("APOLLO_NAMESPACE")
		secret := os.Getenv("APOLLO_SECRET")
		GlobalConfig, err = NewApolloProvider(appid, cluster, host, namespace, secret)
		if err != nil {
			panic("set config error")
		}
		break
	default:
		path := os.Getenv("APP_YAML_PATH")
		GlobalConfig, err = NewYAMLProvider(path)
		if err != nil {
			panic("set config error")
		}
	}
}

func GetString(key ...string) string {
	result, _ := GlobalConfig.GetString(key...)
	return result
}

func GetBool(key ...string) bool {
	result, _ := GlobalConfig.GetBool(key...)
	return result
}

func GetInt(key ...string) int64 {
	result, _ := GlobalConfig.GetInt(key...)
	return result
}

func GetFloat(key ...string) float64 {
	result, _ := GlobalConfig.GetFloat(key...)
	return result
}
