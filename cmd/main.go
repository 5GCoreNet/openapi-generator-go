package main

import "github.com/5GCoreNet/openapi-generator-go/cmd/generator"

func main() {
	if generator.Execute() != nil {
		panic("failed to execute generator")
	}
}
