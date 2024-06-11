// Code generated by go generate; DO NOT EDIT.
package aws

import (
	"context"

	"github.com/grafana/sobek"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

type KinesisClient struct {
	*AWS
	sdk *kinesis.Client
}

func (a *AWS) newKinesisClient(call sobek.ConstructorCall) *sobek.Object {
	awsCfg := a.constructorCallToConfig("KinesisClient", call)

	sdk := kinesis.NewFromConfig(awsCfg)

	client := &KinesisClient{
		AWS: a,
		sdk: sdk,
	}

	return a.vu.Runtime().ToValue(client).ToObject(a.vu.Runtime())
}


func (c *KinesisClient) AddTagsToStream(obj *sobek.Object,) sobek.Value {
	in := &kinesis.AddTagsToStreamInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.AddTagsToStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) CreateStream(obj *sobek.Object,) sobek.Value {
	in := &kinesis.CreateStreamInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.CreateStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DecreaseStreamRetentionPeriod(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DecreaseStreamRetentionPeriodInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DecreaseStreamRetentionPeriod(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DeleteResourcePolicy(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DeleteResourcePolicyInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DeleteResourcePolicy(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DeleteStream(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DeleteStreamInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DeleteStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DeregisterStreamConsumer(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DeregisterStreamConsumerInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DeregisterStreamConsumer(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DescribeLimits(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DescribeLimitsInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DescribeLimits(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DescribeStream(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DescribeStreamInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DescribeStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DescribeStreamConsumer(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DescribeStreamConsumerInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DescribeStreamConsumer(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DescribeStreamSummary(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DescribeStreamSummaryInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DescribeStreamSummary(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) DisableEnhancedMonitoring(obj *sobek.Object,) sobek.Value {
	in := &kinesis.DisableEnhancedMonitoringInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.DisableEnhancedMonitoring(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) EnableEnhancedMonitoring(obj *sobek.Object,) sobek.Value {
	in := &kinesis.EnableEnhancedMonitoringInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.EnableEnhancedMonitoring(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) GetRecords(obj *sobek.Object,) sobek.Value {
	in := &kinesis.GetRecordsInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.GetRecords(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) GetResourcePolicy(obj *sobek.Object,) sobek.Value {
	in := &kinesis.GetResourcePolicyInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.GetResourcePolicy(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) GetShardIterator(obj *sobek.Object,) sobek.Value {
	in := &kinesis.GetShardIteratorInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.GetShardIterator(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) IncreaseStreamRetentionPeriod(obj *sobek.Object,) sobek.Value {
	in := &kinesis.IncreaseStreamRetentionPeriodInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.IncreaseStreamRetentionPeriod(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) ListShards(obj *sobek.Object,) sobek.Value {
	in := &kinesis.ListShardsInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.ListShards(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) ListStreamConsumers(obj *sobek.Object,) sobek.Value {
	in := &kinesis.ListStreamConsumersInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.ListStreamConsumers(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) ListStreams(obj *sobek.Object,) sobek.Value {
	in := &kinesis.ListStreamsInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.ListStreams(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) ListTagsForStream(obj *sobek.Object,) sobek.Value {
	in := &kinesis.ListTagsForStreamInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.ListTagsForStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) MergeShards(obj *sobek.Object,) sobek.Value {
	in := &kinesis.MergeShardsInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.MergeShards(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) PutRecord(obj *sobek.Object,) sobek.Value {
	in := &kinesis.PutRecordInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.PutRecord(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) PutRecords(obj *sobek.Object,) sobek.Value {
	in := &kinesis.PutRecordsInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.PutRecords(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) PutResourcePolicy(obj *sobek.Object,) sobek.Value {
	in := &kinesis.PutResourcePolicyInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.PutResourcePolicy(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) RegisterStreamConsumer(obj *sobek.Object,) sobek.Value {
	in := &kinesis.RegisterStreamConsumerInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.RegisterStreamConsumer(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) RemoveTagsFromStream(obj *sobek.Object,) sobek.Value {
	in := &kinesis.RemoveTagsFromStreamInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.RemoveTagsFromStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) SplitShard(obj *sobek.Object,) sobek.Value {
	in := &kinesis.SplitShardInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.SplitShard(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) StartStreamEncryption(obj *sobek.Object,) sobek.Value {
	in := &kinesis.StartStreamEncryptionInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.StartStreamEncryption(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) StopStreamEncryption(obj *sobek.Object,) sobek.Value {
	in := &kinesis.StopStreamEncryptionInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.StopStreamEncryption(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) SubscribeToShard(obj *sobek.Object,) sobek.Value {
	in := &kinesis.SubscribeToShardInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.SubscribeToShard(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) UpdateShardCount(obj *sobek.Object,) sobek.Value {
	in := &kinesis.UpdateShardCountInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.UpdateShardCount(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

func (c *KinesisClient) UpdateStreamMode(obj *sobek.Object,) sobek.Value {
	in := &kinesis.UpdateStreamModeInput{}
	if err := fromSobekObject(c.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := c.sdk.UpdateStreamMode(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return c.vu.Runtime().ToValue(out)
}

