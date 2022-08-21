package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	l := &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return &MethodLimiter{
		Limiter: l,
	}
}

func (m MethodLimiter) Key(c *gin.Context) string {
	key := c.Request.RequestURI
	index := strings.Index(key, "?")
	if index == -1 {
		return key
	}
	return key[:index]
}

func (m MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limiterBuckets[key]
	return bucket, ok
}

func (m MethodLimiter) AddBuckets(rules ...BucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := m.GetBucket(rule.Key); !ok {
			bucket := ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
			m.limiterBuckets[rule.Key] = bucket
		}
	}
	return m
}
