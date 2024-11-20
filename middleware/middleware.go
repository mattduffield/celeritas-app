package middleware

import (
	"myapp/data"

	"github.com/mattduffield/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
