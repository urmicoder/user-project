package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"urmi-arch/common"
	"urmi-arch/common/settings"

	"github.com/go-redis/redis"
)

type RateLimiting interface {
	RateLimiting(next http.Handler) http.Handler
}
type rlimit struct {
	rdb *redis.Client
}

func NewRLimit(rdb *redis.Client) *rlimit {
	return &rlimit{
		rdb: rdb,
	}
}

func (rl *rlimit) RateLimiting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqIp := strings.Split(r.RemoteAddr, ":")[0]
		queryKey := settings.ApiRateLimitKey + reqIp
		fmt.Println("Query Key", queryKey)
		reqCount, err := rl.rdb.Get(queryKey).Int()

		if err == redis.Nil || reqCount < settings.ApiRateLimitMaxRequests {
			err = rl.rdb.Incr(queryKey).Err()

			if err == redis.Nil {
				rl.rdb.Expire(queryKey, settings.ApiRateLimitDuration)
			}
			fmt.Println("Rate Limiting Called: Count ", reqCount)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, common.ErrRateLimiting.Error(), http.StatusUnauthorized)
		}
	})
}
