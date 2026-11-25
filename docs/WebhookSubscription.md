# WebhookSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The webhook&#x27;s id. | [optional] [default to null]
**Url** | **string** | The URL events get delivered to. | [optional] [default to null]
**Description** | **string** | A user-defined description of the webhook. | [optional] [default to null]
**SubjectType** | **string** | The type of entity. Set to either &#x60;repository&#x60; or &#x60;workspace&#x60; based on where the subscription is defined. | [optional] [default to null]
**Subject** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Active** | **bool** | Whether or not the webhook subscription is enabled. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the webhook subscription was created. | [optional] [default to null]
**Events** | **[]string** | The events this webhook is subscribed to. | [optional] [default to null]
**SecretSet** | **bool** | Indicates whether or not the hook has an associated secret. It is not possible to see the hook&#x27;s secret. This field is ignored during updates. | [optional] [default to null]
**Secret** | **string** | The secret to associate with the hook. The secret is never returned via the API. As such, this field is only used during updates. The secret can be set to &#x60;null&#x60;&#x60; or &#x60;\&quot;\&quot;&#x60;&#x60; to remove the secret (or create a hook with no secret). Leaving out the secret field during updates will leave the secret unchanged. Leaving out the secret during creation will create a hook with no secret. | [optional] [default to null]
**Links** | [***WebhooksSubscriptionLinks**](webhooks_subscription_links.md) |  | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

