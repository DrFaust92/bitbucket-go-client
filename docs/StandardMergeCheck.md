# StandardMergeCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocking** | **bool** | Whether this check is currently preventing the pull request from being merged. | [default to null]
**Check** | [***interface{}**](interface{}.md) | Identifies which branch-restriction-backed check this is, as &#x60;{&#x27;type&#x27;: &#x27;standard_merge_check_definition&#x27;, &#x27;kind&#x27;: &#x27;...&#x27;}&#x60;. See &#x60;check.kind&#x60; in the endpoint documentation for the full list of supported kinds. | [default to null]
**Observed** | [***interface{}**](interface{}.md) | The current observed value for this check, keyed by a check.kind-specific field name (e.g. &#x60;{&#x27;approval_count&#x27;: 1}&#x60;). The shape varies per &#x60;check.kind&#x60;. | [default to null]
**Required** | **bool** | Whether merge-check enforcement is currently active for the pull request (a repository-wide setting), not a property of the individual check. | [default to null]
**Requirement** | [***interface{}**](interface{}.md) | The configured threshold for this check, keyed by a check.kind-specific field name (e.g. &#x60;{&#x27;minimum_approvals&#x27;: 2}&#x60;). The shape varies per &#x60;check.kind&#x60;. | [default to null]
**Type_** | **string** |  | [default to null]
**Status** | **string** | Whether this check passed. &#x60;UNKNOWN&#x60; means the check could not be evaluated; it does not establish that the pull request is mergeable. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

