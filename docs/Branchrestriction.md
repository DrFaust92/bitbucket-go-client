# Branchrestriction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Id** | **int32** | The branch restriction status&#x27; id. | [optional] [default to null]
**Kind** | **string** | The type of restriction that is being applied. | [default to null]
**BranchMatchKind** | **string** | Indicates how the restriction is matched against a branch. The default is &#x60;glob&#x60;. | [default to null]
**BranchType** | **string** | Apply the restriction to branches of this type. Active when &#x60;branch_match_kind&#x60; is &#x60;branching_model&#x60;. The branch type will be calculated using the branching model configured for the repository. | [optional] [default to null]
**Pattern** | **string** | Apply the restriction to branches that match this pattern. Active when &#x60;branch_match_kind&#x60; is &#x60;glob&#x60;. Will be empty when &#x60;branch_match_kind&#x60; is &#x60;branching_model&#x60;. | [default to null]
**Value** | **int32** | Value with kind-specific semantics:  * &#x60;require_approvals_to_merge&#x60; uses it to require a minimum number of approvals on a PR.  * &#x60;require_default_reviewer_approvals_to_merge&#x60; uses it to require a minimum number of approvals from default reviewers on a PR.  * &#x60;require_passing_builds_to_merge&#x60; uses it to require a minimum number of passing builds.  * &#x60;require_commits_behind&#x60; uses it to require the current branch is up to a maximum number of commits behind it destination. | [optional] [default to null]
**Users** | [**[]Account**](account.md) |  | [optional] [default to null]
**Groups** | [**[]Group**](group.md) |  | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

