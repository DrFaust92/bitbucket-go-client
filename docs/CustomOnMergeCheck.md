# CustomOnMergeCheck

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Check** | [***interface{}**](interface{}.md) | Identifies which custom check this is, as &#x60;{&#x27;type&#x27;: &#x27;custom_merge_check_definition&#x27;, &#x27;id&#x27;: &#x27;...&#x27;, &#x27;name&#x27;: &#x27;...&#x27;}&#x60;, where &#x60;id&#x60; is the ARI of the custom check extension. | [default to null]
**Message** | **string** | Present when the custom check provides one; omitted when absent. | [optional] [default to null]
**Uuid** | **string** | Present when a persisted result exists: the persisted custom check result UUID. Omitted for on-merge checks known but not yet triggered. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

