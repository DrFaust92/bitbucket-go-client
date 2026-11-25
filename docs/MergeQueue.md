# MergeQueue

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The merge queue UUID. | [default to null]
**Name** | **string** | The merge queue name. | [default to null]
**State** | **string** | The configuration state returned by the merge queue service. Known values are &#x60;ACTIVE&#x60;, &#x60;PAUSED&#x60;, &#x60;DRAINING&#x60;, &#x60;SUSPENDED&#x60;, &#x60;INACTIVE&#x60;, and &#x60;UNKNOWN&#x60;. &#x60;ACTIVE&#x60; accepts entries; &#x60;PAUSED&#x60; accepts entries but pauses landing; &#x60;DRAINING&#x60; and &#x60;SUSPENDED&#x60; block entry; &#x60;INACTIVE&#x60; does not require queueing. Unknown values are returned unchanged and follow the existing queue behavior: only &#x60;DRAINING&#x60; and &#x60;SUSPENDED&#x60; block entry. | [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

