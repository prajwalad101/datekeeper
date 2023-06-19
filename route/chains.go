package route

import (
	"github.com/Prajwalad101/datekeeper/middleware"
	"github.com/justinas/alice"
)

var Protected alice.Chain

// Initializes new chains of middlewares to be used with handlers
func InitChain() {
	Protected = alice.New(middleware.Auth)
}
