package limiter

import (
	rstruct "github.com/0RAJA/Road/internal/pkg/struct"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type PrefixLimiter struct {
	*Limier
	*rstruct.PrefixTree
}

func NewPrefixLimiter() *PrefixLimiter {
	return &PrefixLimiter{&Limier{limiterBuckets: map[string]*ratelimit.Bucket{}}, rstruct.NewPrefixTree()}
}

func (p *PrefixLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	prefix := strings.Split(uri, "/")
	if result := p.Get(prefix); result != nil {
		return result.(string)
	}
	return ""
}

func (p *PrefixLimiter) testKey(uri string) string {
	prefix := strings.Split(uri, "/")
	result := p.Get(prefix)
	if result != nil {
		return result.(string)
	}
	return ""
}

func (p *PrefixLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := p.limiterBuckets[key]
	return bucket, ok
}

func (p *PrefixLimiter) AddBuckets(rules ...BucketRule) Iface {
	for _, rule := range rules {
		if _, ok := p.limiterBuckets[rule.Key]; !ok {
			p.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Cap, rule.Quantum)
			p.Put(strings.Split(rule.Key, "/"), rule.Key)
		}
	}
	return p
}
