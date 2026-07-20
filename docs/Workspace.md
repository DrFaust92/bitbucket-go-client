# Workspace

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Uuid** | **string** | The workspace&#x27;s immutable id. | [optional] [default to null]
**Name** | **string** | The name of the workspace. | [optional] [default to null]
**Slug** | **string** | The short label that identifies this workspace. | [optional] [default to null]
**IsPrivate** | **bool** | Indicates whether the workspace is publicly accessible, or whether it is private to the members and consequently only visible to members. | [optional] [default to null]
**IsPrivacyEnforced** | **bool** | Indicates whether the workspace enforces private content, or whether it allows public content. | [optional] [default to null]
**ForkingMode** | **string** | Controls the rules for forking repositories within this workspace.  * **allow_forks**: unrestricted forking * **internal_only**: prevents forking of private repositories outside the workspace or to public repositories  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

