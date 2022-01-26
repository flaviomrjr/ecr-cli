ECR CLI
=======

This cli was built to help you manage your ECR image repositories and make it easier!

## Json Policy

You can set a Json policy in your image repositories.

To do so you just need to create a json file with your repo policy!

You will find a file called `policy.json` as a template:
```json
{
	"Version": "2008-10-17",
	"Statement": [
	  {
		"Sid": "eks_access",
		"Effect": "Allow",
		"Principal": {
		  "AWS": [
			"arn:aws:iam::010101010101:root",
			"arn:aws:iam::020202020202:root"
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
  }
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
This function creates a new image repo in ECR

Usage:
  ecr create-repo [flags]

Flags:
  -h, --help                 help for create-repo
  -f, --policy-file string   Set Policy Path. e.g.: ~/Documents/policy.json (default "policy.json")

Global Flags:
  -p, --profile string   Set AWS profile (default "shared")
      --region string    Set AWS region (default "sa-east-1")
  -r, --repo string      Set repo name

Examples:
```shell
$ ./ecr create-repo --repo apps/mynewrepo

$ ./ecr create-repo --repo apps/mynewrepo2 --profile staging

$ ./ecr create-repo --repo apps/mynewrepo3 --profile development -f ~/Documents/my-new-policy.json

$ ./ecr create-repo --repo apps/mynewrepo3 --region us-east-1
```

More functions are comming!

by Flavio Rocha

