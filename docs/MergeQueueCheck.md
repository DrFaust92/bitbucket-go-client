# MergeQueueCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Required** | **bool** | Whether the destination branch currently requires queueing. | [default to null]
**Blocking** | **bool** | Whether this check is currently preventing the pull request from being merged. | [default to null]
**Queued** | **bool** | Whether the pull request is already queued, according to its current lifecycle state. | [default to null]
**MergeQueue** | [***MergeQueue**](merge_queue.md) |  | [default to null]
**Status** | **string** | Whether this check passed. &#x60;UNKNOWN&#x60; means the check could not be evaluated; it does not establish that the pull request is mergeable. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

