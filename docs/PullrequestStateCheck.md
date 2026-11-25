# PullrequestStateCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**State** | **string** | The public state of the pull request. Note that a queued pull request is reported as &#x60;OPEN&#x60; here but as &#x60;blocking: true&#x60;, since it cannot be merged again while a merge is already in progress. | [default to null]
**Status** | **string** | Whether this check passed. &#x60;UNKNOWN&#x60; means the check could not be evaluated; it does not establish that the pull request is mergeable. | [default to null]
**Required** | **bool** | Whether this check must pass for the pull request to be mergeable. | [default to null]
**Blocking** | **bool** | Whether this check is currently preventing the pull request from being merged. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

