
# 根据文档配置sdk并初始化相关服务
[配置aws-sdk-config](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/)

Like dynamodb config:
```
var _cfg aws.Config

// 初始化使用到的aws-service
func InitAwsService() {
	// Load the Shared AWS Configuration (~/.aws/config)
	var err error
	_cfg, err = config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-2"),
		// local dynamodb
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				}),
		),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("97rwl", "oqzqpi", "")),
	)

	if err != nil {
		log.Fatal(err)
	}

	// init dynamodb
	mdynamodb.Init(_cfg)
}
```

## 使用 mdynamodb
