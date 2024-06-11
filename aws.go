package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/grafana/sobek"
	"go.k6.io/k6/js/common"
)

type Config struct {
	Region string

	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string

	Endpoint Endpoint
}

type Endpoint struct {
	URL           string
	SigningRegion string
}

func (a *AWS) newConfig(call sobek.ConstructorCall) *sobek.Object {
	if len(call.Arguments) == 0 {
		return a.vu.Runtime().ToValue(&Config{}).ToObject(a.vu.Runtime())
	}

	if len(call.Arguments) != 1 || !isObject(call.Arguments[0]) {
		panic(a.vu.Runtime().NewTypeError("AWSConfig constructor expects exactly one object argument"))
	}

	var (
		cfg Config
		obj = call.Arguments[0].ToObject(a.vu.Runtime())
	)
	if err := a.vu.Runtime().ExportTo(obj, &cfg); err != nil {
		panic(a.vu.Runtime().NewTypeError(err.Error()))
	}

	return a.vu.Runtime().ToValue(cfg).ToObject(a.vu.Runtime())
}

func (a *AWS) constructorCallToConfig(id string, call sobek.ConstructorCall) aws.Config {
	rt := a.vu.Runtime()

	// If no arguments are passed, we'll use the default config.
	if len(call.Arguments) == 0 {
		awsCfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			common.Throw(rt, err)
		}
		return awsCfg
	}

	// If more than one argument is passed, we'll throw an error.
	if len(call.Arguments) != 1 {
		panic(a.vu.Runtime().NewTypeError(fmt.Sprintf("%s constructor expects exactly one argument", id)))
	}

	// If the first argument is not a [Config], we'll throw an error.
	cfg, ok := call.Arguments[0].Export().(Config)
	if !ok {
		panic(a.vu.Runtime().NewTypeError(fmt.Sprintf("%s constructor first argument must be an AWSConfig", id)))
	}

	optFns := make([]func(*config.LoadOptions) error, 0)

	// If region is set, we'll use it.
	if len(cfg.Region) > 0 {
		optFns = append(optFns, config.WithRegion(cfg.Region))
	}

	// If any credential field is set, we'll use it.
	if len(cfg.AccessKeyID) > 0 || len(cfg.SecretAccessKey) > 0 || len(cfg.SessionToken) > 0 {
		optFns = append(optFns, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, cfg.SessionToken),
		))
	}

	// If endpoint is set, we'll use it.
	if len(cfg.Endpoint.URL) > 0 {
		optFns = append(optFns, config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           cfg.Endpoint.URL,
					SigningRegion: cfg.Endpoint.SigningRegion,
				}, nil
			},
		)))
	}

	awsCfg, err := config.LoadDefaultConfig(context.TODO(), optFns...)
	if err != nil {
		common.Throw(rt, err)
	}

	return awsCfg
}
