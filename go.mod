module kangkyu.com/rds-prox-demo

go 1.17

require (
	github.com/aws/aws-lambda-go v1.32.1
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.15.13
	github.com/aws/aws-sdk-go-v2/feature/rds/auth v1.1.24
	github.com/aws/aws-sdk-go-v2/service/rds v1.22.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.15.13
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.6
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
