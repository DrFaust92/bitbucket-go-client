# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersSelectedUserGpgKeysFingerprintDelete**](GPGApi.md#UsersSelectedUserGpgKeysFingerprintDelete) | **Delete** /users/{selected_user}/gpg-keys/{fingerprint} | Delete a GPG key
[**UsersSelectedUserGpgKeysFingerprintGet**](GPGApi.md#UsersSelectedUserGpgKeysFingerprintGet) | **Get** /users/{selected_user}/gpg-keys/{fingerprint} | Get a GPG key
[**UsersSelectedUserGpgKeysGet**](GPGApi.md#UsersSelectedUserGpgKeysGet) | **Get** /users/{selected_user}/gpg-keys | List GPG keys
[**UsersSelectedUserGpgKeysPost**](GPGApi.md#UsersSelectedUserGpgKeysPost) | **Post** /users/{selected_user}/gpg-keys | Add a new GPG key

# **UsersSelectedUserGpgKeysFingerprintDelete**
> UsersSelectedUserGpgKeysFingerprintDelete(ctx, fingerprint, selectedUser)
Delete a GPG key

Deletes a specific GPG public key from a user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fingerprint** | **string**| A GPG key fingerprint.  | 
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserGpgKeysFingerprintGet**
> GpgAccountKey UsersSelectedUserGpgKeysFingerprintGet(ctx, fingerprint, selectedUser)
Get a GPG key

Returns a specific GPG public key belonging to a user. The `key` and `subkeys` fields can also be requested from the endpoint. See [Partial Responses](/cloud/bitbucket/rest/intro/#partial-response) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fingerprint** | **string**| A GPG key fingerprint.  | 
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 

### Return type

[**GpgAccountKey**](GPG_account_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserGpgKeysGet**
> PaginatedGpgUserKeys UsersSelectedUserGpgKeysGet(ctx, selectedUser)
List GPG keys

Returns a paginated list of the user's GPG public keys. The `key` and `subkeys` fields can also be requested from the endpoint. See [Partial Responses](/cloud/bitbucket/rest/intro/#partial-response) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 

### Return type

[**PaginatedGpgUserKeys**](paginated_gpg_user_keys.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserGpgKeysPost**
> GpgAccountKey UsersSelectedUserGpgKeysPost(ctx, selectedUser, optional)
Add a new GPG key

Adds a new GPG public key to the specified user account and returns the resulting key.  Example:  ``` $ curl -X POST -H \"Content-Type: application/json\" -d '{\"key\": \"<insert GPG Key>\"}' https://api.bitbucket.org/2.0/users/{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}/gpg-keys ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
 **optional** | ***GPGApiUsersSelectedUserGpgKeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GPGApiUsersSelectedUserGpgKeysPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GpgAccountKey**](GpgAccountKey.md)| The new GPG key object. | 

### Return type

[**GpgAccountKey**](GPG_account_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

