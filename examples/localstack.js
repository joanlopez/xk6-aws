import {AWSConfig} from "k6/x/aws";

export const config = new AWSConfig(
	{
		region: "us-east-1",
		access_key_id: "test",
		secret_access_key: "test",
		endpoint: {url: "http://localhost:4566", signing_region: "us-east-1"},
	},
);