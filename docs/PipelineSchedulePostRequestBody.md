# PipelineSchedulePostRequestBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Target** | [***PipelineSchedulePostRequestBodyTarget**](pipeline_schedule_post_request_body_target.md) |  | [default to null]
**Enabled** | **bool** | Whether the schedule is enabled. | [optional] [default to null]
**CronPattern** | **string** | The cron expression with second precision (7 fields) that the schedule applies. For example, for expression: 0 0 12 * * ? *, will execute at 12pm UTC every day. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

