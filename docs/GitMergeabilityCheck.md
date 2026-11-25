# GitMergeabilityCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***interface{}**](interface{}.md) | Present when this check is blocking: &#x60;{&#x27;details&#x27;: {&#x27;href&#x27;: &#x27;...&#x27;}}&#x60;, linking to this pull request&#x27;s &#x60;conflicts&#x60; endpoint. | [optional] [default to null]
**Reason** | **string** | Present when &#x60;status&#x60; is &#x60;PASSED&#x60; or &#x60;FAILED&#x60;: the Git-level outcome for the selected source and destination heads. Omitted when &#x60;status&#x60; is &#x60;UNKNOWN&#x60; (Bitbucket could not evaluate Git mergeability, e.g. due to a timeout); this doesn&#x27;t imply the pull request is conflict-free. | [optional] [default to null]
**Type_** | **string** |  | [default to null]
**Status** | **string** | Whether this check passed. &#x60;UNKNOWN&#x60; means the check could not be evaluated; it does not establish that the pull request is mergeable. | [default to null]
**Required** | **bool** | Whether this check must pass for the pull request to be mergeable. | [default to null]
**Blocking** | **bool** | Whether this check is currently preventing the pull request from being merged. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

