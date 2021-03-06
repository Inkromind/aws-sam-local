package cloudformation

// AWSServerlessFunction_KinesisEvent AWS CloudFormation Resource (AWS::Serverless::Function.KinesisEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#kinesis
type AWSServerlessFunction_KinesisEvent struct {

	// BatchSize AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#kinesis
	BatchSize int `json:"BatchSize,omitempty"`

	// StartingPosition AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#kinesis
	StartingPosition string `json:"StartingPosition,omitempty"`

	// Stream AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#kinesis
	Stream string `json:"Stream,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_KinesisEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.KinesisEvent"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSServerlessFunction_KinesisEvent) AWSCloudFormationSpecificationVersion() string {
	return "2016-10-31"
}
