package ratelimiter

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type RateLimiterMiddleware struct {
	opts        RateLimiterMiddlewareOpts
	rateLimiter *RateLimiter
}

type RateLimiterMiddlewareOpts struct {
	IPMaxRequests    int
	TokenMaxRequests int
	TimeWindow       time.Duration
}

func NewRateLimiterMiddleware(
	store RateLimiterStoreInterface,
	opts RateLimiterMiddlewareOpts,
) (*RateLimiterMiddleware, error) {
	if opts.IPMaxRequests <= 0 {
		return nil, fmt.Errorf("invalid IP max requests: %d", opts.IPMaxRequests)
	}
	if opts.TokenMaxRequests < 0 {
		opts.TokenMaxRequests = 0
	}

	rateLimiter, err := NewRateLimiter(store, opts.TimeWindow)
	if err != nil {
		return nil, err
	}
	return &RateLimiterMiddleware{
		rateLimiter: rateLimiter,
		opts:        opts,
	}, nil
}

func (m *RateLimiterMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key, maxRequests := m.getKeyAndMaxRequests(r)
		if !m.rateLimiter.Allow(key, maxRequests) {
			errMsg := "you have reached the maximum number of requests or actions allowed within a certain time frame"
			http.Error(w, errMsg, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *RateLimiterMiddleware) getKeyAndMaxRequests(r *http.Request) (string, int) {
	var key string
	var maxRequests int

	if m.opts.TokenMaxRequests > 0 {
		key = getAPIKEYToken(r)
		maxRequests = m.opts.TokenMaxRequests
	}

	if key == "" {
		key = getIPAddress(r)
		maxRequests = m.opts.IPMaxRequests
	}
	return key, maxRequests
}

func getIPAddress(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")
	if ips != "" {
		ipsArr := strings.Split(ips, ",")
		return strings.TrimSpace(ipsArr[0])
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func getAPIKEYToken(r *http.Request) string {
	return r.Header.Get("API_KEY")
}
