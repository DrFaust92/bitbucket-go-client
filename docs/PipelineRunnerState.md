# PipelineRunnerState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The current status of the runner. | [optional] [default to null]
**Version** | [***PipelineRunnerVersion**](pipeline_runner_version.md) |  | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | The timestamp when the runner state was last updated. | [optional] [default to null]
**Cordoned** | **bool** | Whether the runner is cordoned (prevented from accepting new steps). | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

