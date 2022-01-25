/*
Copyright © 2022 Flavio Rocha

*/
package main

import (
	"ecr/cmd"
	"os"
)

func main() {
	// set region
	os.Setenv("AWS_REGION", "sa-east-1")

	// read aws credentials file
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	
	cmd.Execute()
}
