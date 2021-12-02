package xlconfig

type Provider interface {
	GetString(key ...string) (string, error)
	GetBool(key ...string) (bool, error)
	GetInt(key ...string) (int64, error)
	GetFloat(key ...string) (float64, error)
	Populate(i interface{}, key ...string) error
}
