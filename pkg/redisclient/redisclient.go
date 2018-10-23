package redisclient

import (
	"fmt"

	"github.com/kevinwubert/go-simple-redis/pkg/redis"
)

func Main() error {
	r, err := redis.NewClient()
	if err != nil {
		return err
	}

	err = r.Set("foobar", "asdfasdf")
	if err != nil {
		return err
	}

	v, err := r.Get("foobar")
	if err != nil {
		return err
	}

	fmt.Println(v)

	return nil
}
