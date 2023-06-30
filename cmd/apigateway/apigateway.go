package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/pkg/logger"
	"github.com/markot99/myinventory-backend/utils/envutils"
)

// proxy is a gin middleware that proxies requests to the given URL
func proxy(proxyURL string, forwardPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(proxyURL)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			if forwardPath != "" {
				req.URL.Path = forwardPath
			}
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// ApiGatewayConfiguration represents the configuration for the API Gateway
type ApiGatewayConfiguration struct {
	Host string // host url of the service
	Path string // path to forward to the service
}

// Extracts the configuration for the API Gateway from the environment variables
// The configuration is stored in the following format:
// FORWARD_1=host;path, FORWARD_2=host;path, ...
func extractApiGatewayConfiguration() []ApiGatewayConfiguration {
	variables := []ApiGatewayConfiguration{}
	index := 1
	for {
		config := os.Getenv(fmt.Sprintf("FORWARD_"+"%d", index))
		if config == "" {
			break
		}

		configSplit := strings.Split(config, ";")
		if len(configSplit) != 2 {
			panic(fmt.Sprintf("Invalid configuration for FORWARD_%d", index))
		}

		hostEnv := configSplit[0]
		pathEnv := configSplit[1]

		variables = append(variables, ApiGatewayConfiguration{Host: hostEnv, Path: pathEnv})
		index++
	}
	return variables
}

// @title myInventory API
// @description Swagger documentation to test the myInventory API
// @version 1.0
// @host localhost:8080
// @BasePath /api
//
// @securityDefinitions.apikey	JWT
// @in							header
// @name						Authorization
// @description					Description for what is this security definition being used
func main() {
	logger.Info("Starting server...")
	logger.InitializeLogger()

	port := envutils.GetEnvVariableOrDefault("PORT", "8080")
	router := gin.Default()

	apiGatewayConfiguration := extractApiGatewayConfiguration()

	for _, apiGatewayConfig := range apiGatewayConfiguration {
		router.Any(apiGatewayConfig.Path, proxy(apiGatewayConfig.Host, ""))
		router.Any(apiGatewayConfig.Path+"/*any", proxy(apiGatewayConfig.Host, ""))
	}

	router.Run(":" + port)
	logger.Info("Bye!")
}
