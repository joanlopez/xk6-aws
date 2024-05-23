// Code generated by go generate; DO NOT EDIT.
package aws

import (
	"context"

	"github.com/dop251/goja"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)


func (a *AWS) AddTagsToStream(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.AddTagsToStreamInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).AddTagsToStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) CreateStream(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.CreateStreamInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).CreateStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DecreaseStreamRetentionPeriod(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DecreaseStreamRetentionPeriodInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DecreaseStreamRetentionPeriod(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DeleteResourcePolicy(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DeleteResourcePolicyInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DeleteResourcePolicy(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DeleteStream(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DeleteStreamInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DeleteStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DeregisterStreamConsumer(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DeregisterStreamConsumerInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DeregisterStreamConsumer(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DescribeLimits(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DescribeLimitsInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DescribeLimits(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DescribeStream(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DescribeStreamInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DescribeStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DescribeStreamConsumer(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DescribeStreamConsumerInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DescribeStreamConsumer(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DescribeStreamSummary(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DescribeStreamSummaryInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DescribeStreamSummary(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) DisableEnhancedMonitoring(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.DisableEnhancedMonitoringInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).DisableEnhancedMonitoring(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) EnableEnhancedMonitoring(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.EnableEnhancedMonitoringInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).EnableEnhancedMonitoring(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) GetRecords(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.GetRecordsInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).GetRecords(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) GetResourcePolicy(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.GetResourcePolicyInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).GetResourcePolicy(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) GetShardIterator(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.GetShardIteratorInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).GetShardIterator(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) IncreaseStreamRetentionPeriod(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.IncreaseStreamRetentionPeriodInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).IncreaseStreamRetentionPeriod(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) ListShards(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.ListShardsInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).ListShards(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) ListStreamConsumers(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.ListStreamConsumersInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).ListStreamConsumers(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) ListStreams(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.ListStreamsInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).ListStreams(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) ListTagsForStream(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.ListTagsForStreamInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).ListTagsForStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) MergeShards(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.MergeShardsInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).MergeShards(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) PutRecord(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.PutRecordInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).PutRecord(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) PutRecords(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.PutRecordsInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).PutRecords(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) PutResourcePolicy(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.PutResourcePolicyInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).PutResourcePolicy(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) RegisterStreamConsumer(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.RegisterStreamConsumerInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).RegisterStreamConsumer(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) RemoveTagsFromStream(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.RemoveTagsFromStreamInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).RemoveTagsFromStream(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) SplitShard(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.SplitShardInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).SplitShard(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) StartStreamEncryption(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.StartStreamEncryptionInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).StartStreamEncryption(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) StopStreamEncryption(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.StopStreamEncryptionInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).StopStreamEncryption(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) SubscribeToShard(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.SubscribeToShardInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).SubscribeToShard(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) UpdateShardCount(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.UpdateShardCountInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).UpdateShardCount(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

func (a *AWS) UpdateStreamMode(obj *goja.Object,) goja.Value {
	cfg, err := defaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	in := &kinesis.UpdateStreamModeInput{}
	if err := fromGojaObject(a.vu.Runtime(), obj, in); err != nil {
		panic(err)
	}

	out, err := kinesis.NewFromConfig(cfg).UpdateStreamMode(context.Background(), in, )
    if err != nil {
		panic(err)
	}

	return a.vu.Runtime().ToValue(out)
}

