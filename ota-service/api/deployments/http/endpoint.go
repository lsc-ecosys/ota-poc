package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"ota/deployments"
)

func getDeploymentEndpoint(svc deployments.Service) endpoint.