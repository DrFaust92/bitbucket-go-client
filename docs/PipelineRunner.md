# PipelineRunner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The UUID identifying the runner. | [optional] [default to null]
**Name** | **string** | The name of the runner. | [optional] [default to null]
**Labels** | **[]string** | Labels assigned to the runner for identification and routing. | [optional] [default to null]
**State** | [***PipelineRunnerState**](pipeline_runner_state.md) |  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | The timestamp when the runner was created. | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | The timestamp when the runner was last updated. | [optional] [default to null]
**OauthClient** | [***PipelineRunnerOauthClient**](pipeline_runner_oauth_client.md) |  | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

