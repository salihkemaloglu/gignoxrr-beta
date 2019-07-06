package helpers

import (
	"github.com/spf13/pflag"
)

var (
	flagAllowAllOrigins = pflag.Bool("allow_all_origins", true, "allow requests from any origin.")
	flagAllowedOrigins  = pflag.StringSlice("allowed_origins", nil, "comma-separated list of origin URLs which are allowed to make cross-origin requests.")
)

//MakeHTTPOriginFunc ...
func MakeHTTPOriginFunc(allowedOrigins *allowedOrigins) func(origin string) bool {
	if *flagAllowAllOrigins {
		return func(origin string) bool {
			return true
		}
	}
	return allowedOrigins.IsAllowed
}

//MakeAllowedOrigins ...
func MakeAllowedOrigins() *allowedOrigins {
	o := map[string]struct{}{}
	for _, allowedOrigin := range *flagAllowedOrigins {
		o[allowedOrigin] = struct{}{}
	}
	return &allowedOrigins{
		origins: o,
	}
}

type allowedOrigins struct {
	origins map[string]struct{}
}

func (a *allowedOrigins) IsAllowed(origin string) bool {
	_, ok := a.origins[origin]
	return ok
}
