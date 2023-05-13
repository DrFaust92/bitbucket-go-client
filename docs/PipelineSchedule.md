# PipelineSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Uuid** | **string** | The UUID identifying the schedule. | [optional] [default to null]
**Enabled** | **bool** | Whether the schedule is enabled. | [optional] [default to null]
**Target** | [***PipelineRefTarget**](pipeline_ref_target.md) |  | [optional] [default to null]
**CronPattern** | **string** | The cron expression with second precision (7 fields) that the schedule applies. For example, for expression: 0 0 12 * * ? *, will execute at 12pm UTC every day. | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | The timestamp when the schedule was created. | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | The timestamp when the schedule was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

