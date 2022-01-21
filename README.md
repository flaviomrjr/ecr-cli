ECR CLI
=======

This cli was built to help you manage your ECR image repositories and make it easier!

## Json Policy

You can set a Json policy in your image repositories.

To do so you just need to set your json policy in the config file called `createRepo.go`

You will find a var called jsonData:
```json
// set json policy
var jsonData = `{
	"Version": "2008-10-17",
	"Statement": [
	  {
		"Sid": "eks_access",
		"Effect": "Allow",
		"Principal": {
		  "AWS": [
			"arn:aws:iam::000000000000:root",
			"arn:aws:iam::111111111111:root"
		  ]
		},
		"Action": [
		  "ecr:BatchCheckLayerAvailability",
		  "ecr:BatchGetImage",
		  "ecr:CompleteLayerUpload",
		  "ecr:DescribeImages",
		  "ecr:GetDownloadUrlForLayer",
		  "ecr:GetLifecyclePolicy",
		  "ecr:GetLifecyclePolicyPreview",
		  "ecr:GetRepositoryPolicy",
		  "ecr:ListImages",
		  "ecr:DescribeRepositories",
		  "ecr:InitiateLayerUpload",
		  "ecr:UploadLayerPart",
		  "ecr:PutImage",
		  "ecr:PutLifecyclePolicy",
		  "ecr:SetRepositoryPolicy",
		  "ecr:StartLifecyclePolicyPreview"
		]
	  }
	]
  }`
```

## Build

Clone this repo:
```shell
$ git clone https://github.com/flaviomrjr/ecr-cli.git
```

Building this project:
```shell
$ go build
```

## Available functions

**create-repo:**

This function creates a new image repo in ECR with some configs imported from a json-file

Usage:
  ecr create-repo [flags]

Flags:
  -h, --help             help for create-repo
  -p, --profile string   Set AWS profile (default "shared")
  -r, --repo string      Create a new image repo in ECR

Examples:
```shell
$ ./ecr create-repo --repo apps/mynewrepo

$ ./ecr create-repo --repo apps/mynewrepo2 --profile staging
```

More functions are comming!

by Flavio Rocha

