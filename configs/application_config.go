package configs

import "github.com/kumin/AndPadDating/pkg/envx"

var MaxPageLimit = envx.GetInt("MAX_PAGE_LIMIT", 100)
