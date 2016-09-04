package configs

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
	"os"
)

/* singleton */
var conf *Config

var tomlFileName string = "./configs/settings.toml"

type Config struct {
	Environment string
	Server      ServerConfig
	Proxy       ProxyConfig
	Auth        AuthConfig
	MySQL       MySQLConfig
	Redis       RedisConfig
	Mongo       MongoConfig `toml:"mongodb"`
	Aws         AwsConfig
	Develop     DevelopConfig
}

type ServerConfig struct {
	Scheme    string          `toml:"scheme"`
	Host      string          `toml:"host"`
	Port      int             `toml:"port"`
	Num       int             `toml:"num"`
	Docs      DocsConfig      `toml:"docs"`
	Log       LogConfig       `toml:"log"`
	Session   SessionConfig   `toml:"session"`
	BasicAuth BasicAuthConfig `toml:"basic_auth"`
}

type DocsConfig struct {
	Path string `toml:"path"`
}

type LogConfig struct {
	Level uint8  `toml:"level"`
	Path  string `toml:"path"`
}

type SessionConfig struct {
	Name     string `toml:"name"`
	Key      string `toml:"key"`
	MaxAge   int    `toml:"max_age"`
	Secure   bool   `toml:"secure"`
	HttpOnly bool   `toml:"http_only"`
}

type BasicAuthConfig struct {
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

type ProxyConfig struct {
	Mode   uint8             `toml:"mode"` //0:off, 1:go, 2,nginx
	Server ProxyServerConfig `toml:"server"`
}

type ProxyServerConfig struct {
	Scheme string    `toml:"scheme"`
	Host   string    `toml:"host"`
	Port   int       `toml:"port"`
	Log    LogConfig `toml:"log"`
}

type AuthConfig struct {
	Api      ApiConfig      `toml:"api"`
	Jwt      JwtConfig      `toml:"jwt"`
	Google   GoogleConfig   `toml:"google"`
	Facebook FacebookConfig `toml:"facebook"`
}

type ApiConfig struct {
	Header string `toml:"header"`
	Key    string `toml:"key"`
	Ajax   bool   `toml:"only_ajax"`
}

type JwtConfig struct {
	Mode       uint8  `toml:"mode"`
	Secret     string `toml:"secret_code"`
	PrivateKey string `toml:"private_key"`
	PublicKey  string `toml:"public_key"`
}

type GoogleConfig struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	CallbackURL  string `toml:"call_back_url"`
}

type FacebookConfig struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	CallbackURL  string `toml:"call_back_url"`
}

type MySQLConfig struct {
	MySQLContentConfig
	Test MySQLContentConfig `toml:"test"`
}

type MySQLContentConfig struct {
	Host   string `toml:"host"`
	Port   uint16 `toml:"port"`
	DbName string `toml:"dbname"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

type RedisConfig struct {
	Host    string `toml:"host"`
	Port    uint16 `toml:"port"`
	Pass    string `toml:"pass"`
	Session bool   `toml:"session"`
}

type MongoConfig struct {
	Host   string `toml:"host"`
	Port   uint16 `toml:"port"`
	DbName string `toml:"dbname"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

type AwsConfig struct {
	AccessKey string `toml:"access_key"`
	SecretKey string `toml:"secret_key"`
	Region    string `toml:"region"`
}

type DevelopConfig struct {
	ProfileEnable bool `toml:"profile_enable"`
	RecoverEnable bool `toml:"recover_enable"`
}

var checkTomlKeys [][]string = [][]string{
	{"environment"},
	{"server", "scheme"},
	{"server", "host"},
	{"server", "port"},
	{"server", "docs", "path"},
	{"server", "log", "level"},
	{"server", "log", "path"},
	{"server", "session", "name"},
	{"server", "session", "key"},
	{"server", "session", "max_age"},
	{"server", "session", "secure"},
	{"server", "session", "http_only"},
	{"server", "basic_auth", "user"},
	{"server", "basic_auth", "pass"},
	{"proxy", "mode"},
	{"proxy", "server", "scheme"},
	{"proxy", "server", "host"},
	{"proxy", "server", "port"},
	{"proxy", "server", "log", "level"},
	{"proxy", "server", "log", "path"},
	{"auth", "api", "header"},
	{"auth", "api", "key"},
	{"auth", "api", "only_ajax"},
	{"auth", "jwt", "mode"},
	{"auth", "jwt", "secret_code"},
	{"auth", "jwt", "private_key"},
	{"auth", "jwt", "public_key"},
	{"auth", "google", "client_id"},
	{"auth", "google", "client_secret"},
	{"auth", "google", "call_back_url"},
	{"mysql", "host"},
	{"mysql", "port"},
	{"mysql", "dbname"},
	{"mysql", "user"},
	{"mysql", "pass"},
	{"mysql", "test", "host"},
	{"mysql", "test", "port"},
	{"mysql", "test", "dbname"},
	{"mysql", "test", "user"},
	{"mysql", "test", "pass"},
	{"redis", "host"},
	{"redis", "port"},
	{"redis", "pass"},
	{"redis", "session"},
	{"mongodb", "host"},
	{"mongodb", "port"},
	{"mongodb", "dbname"},
	{"mongodb", "user"},
	{"mongodb", "pass"},
	{"aws", "access_key"},
	{"aws", "secret_key"},
	{"aws", "region"},
	{"develop", "profile_enable"},
	{"develop", "recover_enable"},
}

func init() {
	tomlFileName = os.Getenv("GOPATH") + "/src/github.com/hiromaily/go-gin-wrapper/configs/settings.toml"
}

//check validation of config
func validateConfig(conf *Config, md *toml.MetaData) error {
	//for protection when debugging on non production environment
	var errStrings []string

	//Check added new items on toml
	// environment
	//if !md.IsDefined("environment") {
	//	errStrings = append(errStrings, "environment")
	//}

	format := "[%s]"
	inValid := false
	for _, keys := range checkTomlKeys {
		if !md.IsDefined(keys...) {
			switch len(keys) {
			case 1:
				format = "[%s]"
			case 2:
				format = "[%s] %s"
			case 3:
				format = "[%s.%s] %s"
			default:
				//invalid check string
				inValid = true
				break
			}
			keysIfc := u.SliceStrToInterface(keys)
			errStrings = append(errStrings, fmt.Sprintf(format, keysIfc...))
		}
	}

	// Error
	if inValid {
		return errors.New("Error: Check Text has wrong number of parameter")
	}
	if len(errStrings) != 0 {
		return fmt.Errorf("Error: There are lacks of keys : %#v \n", errStrings)
	}

	return nil
}

// load configfile
func loadConfig(fileName string) (*Config, error) {
	if fileName != "" {
		tomlFileName = fileName
	}

	d, err := ioutil.ReadFile(tomlFileName)
	if err != nil {
		return nil, fmt.Errorf(
			"Error reading %s: %s", tomlFileName, err)
	}

	var config Config
	md, err := toml.Decode(string(d), &config)
	if err != nil {
		return nil, fmt.Errorf(
			"Error parsing %s: %s(%v)", tomlFileName, err, md)
	}

	//check validation of config
	err = validateConfig(&config, &md)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// singleton architecture
func New(fileName string) {
	var err error
	if conf == nil {
		conf, err = loadConfig(fileName)
	}
	if err != nil {
		panic(err)
	}
}

// singleton architecture
func GetConf() *Config {
	var err error
	if conf == nil {
		conf, err = loadConfig("")
	}
	if err != nil {
		panic(err)
	}

	return conf
}

func SetTomlPath(path string) {
	tomlFileName = path
}
