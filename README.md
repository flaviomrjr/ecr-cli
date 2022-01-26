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
This function creates one or more ECR repos

Usage:
  ecr create-repo [flags]

Flags:
  -h, --help                 help for create-repo
      --no-policy            Do not apply repo policy
  -f, --policy-file string   Set Policy Path. e.g.: ~/Documents/policy.json (default "policy.json")

Global Flags:
  -p, --profile string   Set AWS profile (default "shared")
      --region string    Set AWS region (default "sa-east-1")
  -r, --repo strings     Set repo name

Examples:
```shell
$ ./ecr create-repo --repo apps/mynewrepo

$ ./ecr create-repo --repo apps/mynewrepo --profile staging

$ ./ecr create-repo --repo apps/mynewrepo --profile development -f ~/Documents/my-new-policy.json

$ ./ecr create-repo --repo apps/mynewrepo --region us-east-1

// Create ECR repo withot policy
$ ./ecr create-repo --repo apps/mynewrepo --no-policy

// Create more than one ECR repo
$ ./ecr create-repo --repo apps/mynewrepo --repo apps/mynewrepo2
```

**delete-repo:**

This function deletes one or more ECR repos

Usage:
  ecr delete-repo [flags]

Flags:
  -h, --help   help for delete-repo

Global Flags:
  -p, --profile string   Set AWS profile (default "shared")
      --region string    Set AWS region (default "sa-east-1")
  -r, --repo strings     Set repo name

Examples:
```shell
$ ./ecr delete-repo --repo apps/mynewrepo

$ ./ecr delete-repo --repo apps/mynewrepo1 --repo apps/mynewrepo2
```

More functions are comming!

by Flavio Rocha

