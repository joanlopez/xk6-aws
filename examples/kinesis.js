import {check, sleep} from 'k6';
import {KinesisClient} from "k6/x/aws";
import {config} from './localstack.js';

// Set your k6 run configuration:
// https://k6.io/docs/using-k6/k6-options
export const options = {
	iterations: 1,

	// Demonstrative k6 thresholds.
	thresholds: {
		checks: [{threshold: 'rate == 1.00', abortOnFail: true}],
	},
};

export default function () {
	const kinesis = new KinesisClient(config);

	// Create a Kinesis stream.
	const streamName = 'test-stream';
	const shardCount = 1;
	kinesis.createStream({stream_name: streamName, shard_count: shardCount});

	// List Kinesis streams to confirm creation. It must return at least the stream we created.
	const {stream_names} = kinesis.listStreams();
	check(stream_names, {
		'stream must be in the list': (stream_names) => stream_names.includes(streamName),
	});

	// Wait for the stream to become ACTIVE.
	let streamStatus = 'CREATING';
	while (streamStatus !== 'ACTIVE') {
		const {stream_description} = kinesis.describeStream({stream_name: streamName});
		// Trim to remove leading/trailing whitespace.
		streamStatus = stream_description.stream_status.trim();
		if (streamStatus !== 'ACTIVE') {
			sleep(1); // Wait for 1 second before checking again.
		}
	}

	// Put records onto the Kinesis stream.
	const partitionKey = 'partitionKey';
	const data = 'test-data';
	const putRecordsResponse = kinesis.putRecords({
		stream_name: streamName,
		records: [
			{
				data: data,
				partition_key: partitionKey,
			},
		],
	});

	// Check that the records were successfully put.
	check(putRecordsResponse, {
		'put records must succeed': (res) => res.failed_record_count == 0,
		'records put must have no errors': (res) => res.records.every(r => r.error_code === null && r.error_message === null),
	});

	// Get shard iterator.
	const {shard_iterator} = kinesis.getShardIterator({
		stream_name: streamName,
		shard_id: 'shardId-000000000000',
		shard_iterator_type: 'TRIM_HORIZON',
	});

	check(shard_iterator, {
		'shard iterator must be obtained': (shard_iterator) => shard_iterator !== null,
	});

	// Get records from the stream.
	const {records} = kinesis.getRecords({
		shard_iterator: shard_iterator,
	});

	check(records, {
		'get records must succeed': (records) => records.length > 0,
		//'retrieved data must be correct': (res) => records[0].data === data,
	});

	// Delete the Kinesis stream.
	kinesis.deleteStream({stream_name: streamName});
}