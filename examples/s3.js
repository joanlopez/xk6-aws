import {check} from 'k6';
import {S3Client} from "k6/x/aws";
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

export default async function () {
	const s3 = new S3Client(config);

	const {buckets} = s3.listBuckets();
	check(buckets, {
			'it must return two buckets': (buckets) => buckets.length === 2,
			'bucket1 must be in the list': (buckets) => buckets.some(b => b.name.normalize() === "bucket1"),
			'bucket2 must be in the list': (buckets) => buckets.some(b => b.name.normalize() === "bucket1")
		}
	);

	const {contents: b1Objects} = s3.listObjects({bucket: "bucket1"})
	check(b1Objects, {
			'it must return one object': (b1Objects) => b1Objects.length === 1,
			'file1.txt must be in the list': (b1Objects) => b1Objects.some(obj => obj.key.normalize() === "file1.txt")
		}
	);

	const {contents: b2Objects} = s3.listObjects({bucket: "bucket2"})
	check(b2Objects, {
			'it must return one object': (b2Objects) => b2Objects.length === 1,
			'file1.txt must be in the list': (b2Objects) => b2Objects.some(obj => obj.key.normalize() === "file2.txt")
		}
	);

	const {body} = s3.getObject({bucket: "bucket1", key: "file1.txt"})
	const reader = body.getReader();
	const {value} = await reader.read();
	check(value, {
			'file1.txt contents are the expected ones': (value) => value === 'Hello, World',
		}
	);
}