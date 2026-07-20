# GpgAccountKey

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [***Account**](account.md) |  | [optional] [default to null]
**Key** | **string** | The GPG key value in X format. | [optional] [default to null]
**KeyId** | **string** | The unique identifier for the GPG key | [optional] [default to null]
**Fingerprint** | **string** | The GPG key fingerprint. | [optional] [default to null]
**ParentFingerprint** | **string** | The fingerprint of the parent key. This value is null unless the current key is a subkey. | [optional] [default to null]
**Name** | **string** | The user-defined label for the GPG key | [optional] [default to null]
**ExpiresOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**AddedOn** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastUsed** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Subkeys** | [**[]GpgAccountKey**](GPG_account_key.md) |  | [optional] [default to null]
**Links** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

