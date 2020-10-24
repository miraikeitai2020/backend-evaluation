package config

import(
	"os"
	"log"
	"fmt"
	"time"
	"math"

	"github.com/kelseyhightower/envconfig"
)

func GetRouterAddr() (addr string) {
	port := os.Getenv("PORT")
	addr = ":" + port
	if port == "" {
		addr = ":8080"
	}
	return
}

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

type dataBaseConfig struct {
	User string `envconfig:"DB_USER" default:"miraiketai2020"`
	Pass string `envconfig:"DB_PASS" default:"miraiketai2020"`
	IP   string `envconfig:"DB_IP" default:"localhost"`
	Port string `envconfig:"DB_PORT" default:"3306"`
	Name string `envconfig:"DB_NAME" default:"michishirube"`
}

func GetConnectionToken() string {
	var c dataBaseConfig
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}

// Exponential Backoff
func BackOff(i int) {
	time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))
}
