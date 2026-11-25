# PaginatedWebhookSubscriptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]WebhookSubscription**](webhook_subscription.md) |  | [optional] [default to null]
**Pagelen** | **int32** | Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values. | [optional] [default to null]
**Next** | **string** | Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

