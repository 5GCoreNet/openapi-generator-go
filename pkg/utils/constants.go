package utils

const (
	// DefaultRepository is the default repository for the OpenAPI specification
	DefaultRepository = "5GC_APIs"
	// DefaultRepositoryOwner is the default owner of the OpenAPI specification repository
	DefaultRepositoryOwner = "jdegre"
	// DefaultRef is the default commit hash or branch name of the OpenAPI specification repository
	DefaultRef = "Rel-18"
	// DefaultPath is the default path to the OpenAPI specification folder under the repository
	DefaultPath = "."
)

const (
	// OpenApiClientGeneratorUrl is the URL of the OpenAPI generator
	OpenApiClientGeneratorUrl = "https://api.openapi-generator.tech/api/gen/clients/go"
	// OpenApiGinServerGeneratorUrl is the URL of the OpenAPI generator
	OpenApiGinServerGeneratorUrl = "https://api.openapi-generator.tech/api/gen/servers/go-gin-server"
	// ServerGinSubFolder is the subfolder where the generated server code is normally located in the zip file returned by the OpenAPI generator.
	ServerGinSubFolder = "go-gin-server-server/"
	// GoSubFolder is the subfolder where the generated go code is normally located in the zip file returned by the OpenAPI generator.
	GoSubFolder = "go/"
	// ClientSubFolder is the subfolder where the generated client code is normally located in the zip file returned by the OpenAPI generator.
	ClientSubFolder = "go-client/"
)
