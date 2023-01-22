package utils

const (
	// OpenApiClientGeneratorUrl is the URL of the OpenAPI generator
	OpenApiClientGeneratorUrl = "https://api.openapi-generator.tech/api/gen/clients/go"
	// OpenApiGinServerGeneratorUrl is the URL of the OpenAPI generator
	OpenApiGinServerGeneratorUrl = "https://api.openapi-generator.tech/api/gen/servers/go-gin-server"
	// ServerGinSubFolder is the subfolder where the generated server code is normally located in the zip file returned by the OpenAPI generator.
	ServerGinSubFolder = "go-gin-server-server/"
	// ClientSubFolder is the subfolder where the generated client code is normally located in the zip file returned by the OpenAPI generator.
	ClientSubFolder = "go-client/"
)
