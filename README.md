# xk6-aws

xk6-aws is an [extension for k6](https://k6.io/docs/extensions). 
It adds support to interact with [Amazon Web Services](https://aws.amazon.com/es/) from your k6 scripts.

```javascript
import {S3Client} from "k6/x/aws";

export default function () {
  // By default, it uses the default AWS credentials (from environment variables or shared credentials file)
  const s3 = new S3Client();
	
  const {contents} = s3.listObjects({bucket: "my-bucket"})
}
```

Check the [examples](#examples) section below for an extensive and complete implementation.

## Getting started

Using the xk6-aws extension involves building a k6 binary incorporating it. 
A detailed guide on how to do this using a [Docker](https://www.docker.com/) or [Go](https://go.dev/) environment
is available in the [extension's documentation](https://grafana.com/docs/k6/latest/extensions/build-k6-binary-using-go/).

In the current state, building directly from the source code using Go could be helpful. We list below the suggested steps:

### Prepare the local environment

1. Make sure `git` and `go` are available commands.
2. Install [xk6](https://github.com/grafana/xk6#local-installation) as suggested in the [local installation](https://github.com/grafana/xk6#local-installation) documentation's section.
3. Clone the xk6-aws repository and move inside the project's folder

### Build the binary

1. Build a k6 binary incorporating the xk6-aws extension
```bash
xk6 build --with github.com/joanlopez/xk6-aws=.
```

2. Run a test script with the newly built binary
```bash
./k6 run script.js
```

## Usage

Once [built](#getting-started) into a k6 executable using [xk6](https://github.com/grafana/xk6),
the extension can be imported by load test scripts as the `k6/x/aws` JavaScript module.

```javascript
import {AWSConfig} from 'k6/x/aws';
```

### AWSConfig

The module exports a `AWSConfig` type which can be used to configure a service client. 
Construct an `AWSConfig` instance by passing the region, credentials and optionally the endpoint details.

```js
const config = new AWSConfig(
  {
    region: "us-east-1",
    access_key_id: "test",
    secret_access_key: "test",
    endpoint: {url: "http://localhost:4566", signing_region: "us-east-1"},
  },
);
```

The table below details the expected arguments:

| Argument                | Type   | Required | Description                                                                                          |
|-------------------------|--------|----------|------------------------------------------------------------------------------------------------------|
| Region                  | string | No       | The AWS region where the resources will be created or accessed.                                      |
| Access key id           | string | No       | The AWS access key ID for the account.                                                               |
| Secret access key       | string | No       | The AWS secret access key for the account.                                                           |
| Endpoint url            | string | No       | The complete URL to use for the constructed client. Overrides the default endpoint for the service.  |
| Endpoint signing region | string | No       | The region to use for signing the request. Useful when the endpoint URL differs from the AWS region. |

An initialized configuration can then be used to instantiate a service client.

```js
const s3 = new S3Client(config);
```

When no `AWSConfig` is provided, service clients will use the default AWS credentials (_from environment variables or shared credentials file_):
- `AWS_REGION`
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`
- `AWS_ENDPOINT_URL`

```js
const s3 = new S3Client();
```

The returned client (`s3` in this example) can be used to interact with the corresponding service.

## Examples

In [examples](./examples) you can find a [fully working testing environment](./examples/docker-compose.yml),
based on [LocalStack](https://www.localstack.cloud/) and script examples for different services 
([s3.js](./examples/s3.js), [kinesis.js](./examples/kinesis.js), [eventbridge.js](./examples/eventbridge.js)...) 
are available to demonstrate how to use the `k6/x/aws` module to interact with each of the services.

## Support

Please, note that this extension is not officially supported by Grafana Labs/k6 core team.

## Development

This extension is built on top of the [AWS SDK for Go v2 (aws-sdk-go-v2)](https://github.com/aws/aws-sdk-go-v2).