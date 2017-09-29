//tutorial url: https://www.kieranajp.uk/articles/build-url-shortener-api-golang/

package main

import (
	"fmt"
	"os"

	redis "gopkg.in/redis.v4"

	"github.com/kataras/iris"
)

// Open a connection to Redis
var redisStorage = redis.NewClient(&redis.Options{
	Addr: fmt.Sprintf(
		"%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	),
})

// Custom `*iris.Context` struct to bind our functions to
type UrlAPI struct {
	*iris.Context
}

// Handler for Post route
func (u UrlAPI) Post() {
	url := u.FormValue("url")
	u.Write(url)
}

func main() {
	defer redisStorage.Close()

	iris.API("/", UrlAPI{})
	iris.Listen(fmt.Sprintf("0.0.0.0:%s", os.Getenv("IRIS_PORT")))
}
