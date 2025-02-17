// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
)

func ExampleEvaluationConfig_outputUsage() {
	var union types.EvaluationConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EvaluationConfigMemberAutomated:
		_ = v.Value // Value is types.AutomatedEvaluationConfig

	case *types.EvaluationConfigMemberHuman:
		_ = v.Value // Value is types.HumanEvaluationConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AutomatedEvaluationConfig
var _ *types.HumanEvaluationConfig

func ExampleEvaluationDatasetLocation_outputUsage() {
	var union types.EvaluationDatasetLocation
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EvaluationDatasetLocationMemberS3Uri:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleEvaluationInferenceConfig_outputUsage() {
	var union types.EvaluationInferenceConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EvaluationInferenceConfigMemberModels:
		_ = v.Value // Value is []types.EvaluationModelConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []types.EvaluationModelConfig

func ExampleEvaluationModelConfig_outputUsage() {
	var union types.EvaluationModelConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EvaluationModelConfigMemberBedrockModel:
		_ = v.Value // Value is types.EvaluationBedrockModel

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EvaluationBedrockModel

func ExampleInferenceProfileModelSource_outputUsage() {
	var union types.InferenceProfileModelSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.InferenceProfileModelSourceMemberCopyFrom:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleModelDataSource_outputUsage() {
	var union types.ModelDataSource
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ModelDataSourceMemberS3DataSource:
		_ = v.Value // Value is types.S3DataSource

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3DataSource

func ExampleModelInvocationJobInputDataConfig_outputUsage() {
	var union types.ModelInvocationJobInputDataConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ModelInvocationJobInputDataConfigMemberS3InputDataConfig:
		_ = v.Value // Value is types.ModelInvocationJobS3InputDataConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ModelInvocationJobS3InputDataConfig

func ExampleModelInvocationJobOutputDataConfig_outputUsage() {
	var union types.ModelInvocationJobOutputDataConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ModelInvocationJobOutputDataConfigMemberS3OutputDataConfig:
		_ = v.Value // Value is types.ModelInvocationJobS3OutputDataConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ModelInvocationJobS3OutputDataConfig
