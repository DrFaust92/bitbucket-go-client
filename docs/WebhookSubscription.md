# WebhookSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Uuid** | **string** | The webhook&#x27;s id | [optional] [default to null]
**Url** | **string** | The URL events get delivered to. | [optional] [default to null]
**Description** | **string** | A user-defined description of the webhook. | [optional] [default to null]
**SubjectType** | **string** | The type of entity. Set to either &#x60;repository&#x60; or &#x60;workspace&#x60; based on where the subscription is defined. | [optional] [default to null]
**Subject** | [***ModelError**](map.md) |  | [optional] [default to null]
**Active** | **bool** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Events** | **[]string** | The events this webhook is subscribed to. | [optional] [default to null]
**SecretSet** | **bool** | Indicates whether or not the hook has an associated secret. It is not possible to see the hook&#x27;s secret. This field is ignored during updates. | [optional] [default to null]
**Secret** | **string** | The secret to associate with the hook. The secret is never returned via the API. As such, this field is only used during updates. The secret can be set to &#x60;null&#x60; or \&quot;\&quot; to remove the secret (or create a hook with no secret). Leaving out the secret field during updates will leave the secret unchanged. Leaving out the secret during creation will create a hook with no secret. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

