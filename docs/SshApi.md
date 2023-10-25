# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersSelectedUserSshKeysGet**](SshApi.md#UsersSelectedUserSshKeysGet) | **Get** /users/{selected_user}/ssh-keys | List SSH keys
[**UsersSelectedUserSshKeysKeyIdDelete**](SshApi.md#UsersSelectedUserSshKeysKeyIdDelete) | **Delete** /users/{selected_user}/ssh-keys/{key_id} | Delete a SSH key
[**UsersSelectedUserSshKeysKeyIdGet**](SshApi.md#UsersSelectedUserSshKeysKeyIdGet) | **Get** /users/{selected_user}/ssh-keys/{key_id} | Get a SSH key
[**UsersSelectedUserSshKeysKeyIdPut**](SshApi.md#UsersSelectedUserSshKeysKeyIdPut) | **Put** /users/{selected_user}/ssh-keys/{key_id} | Update a SSH key
[**UsersSelectedUserSshKeysPost**](SshApi.md#UsersSelectedUserSshKeysPost) | **Post** /users/{selected_user}/ssh-keys | Add a new SSH key

# **UsersSelectedUserSshKeysGet**
> PaginatedSshUserKeys UsersSelectedUserSshKeysGet(ctx, selectedUser, optional)
List SSH keys

Returns a paginated list of the user's SSH public keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
 **optional** | ***SshApiUsersSelectedUserSshKeysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SshApiUsersSelectedUserSshKeysGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedSshUserKeys**](paginated_ssh_user_keys.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserSshKeysKeyIdDelete**
> UsersSelectedUserSshKeysKeyIdDelete(ctx, keyId, selectedUser)
Delete a SSH key

Deletes a specific SSH public key from a user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The SSH key&#x27;s UUID value. | 
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserSshKeysKeyIdGet**
> SshAccountKey UsersSelectedUserSshKeysKeyIdGet(ctx, keyId, selectedUser)
Get a SSH key

Returns a specific SSH public key belonging to a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The SSH key&#x27;s UUID value. | 
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 

### Return type

[**SshAccountKey**](ssh_account_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserSshKeysKeyIdPut**
> SshAccountKey UsersSelectedUserSshKeysKeyIdPut(ctx, keyId, selectedUser, optional)
Update a SSH key

Updates a specific SSH public key on a user's account  Note: Only the 'comment' field can be updated using this API. To modify the key or comment values, you must delete and add the key again.  Example:  ``` $ curl -X PUT -H \"Content-Type: application/json\" -d '{\"label\": \"Work key\"}' https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a} ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The SSH key&#x27;s UUID value. | 
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
 **optional** | ***SshApiUsersSelectedUserSshKeysKeyIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SshApiUsersSelectedUserSshKeysKeyIdPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SshAccountKey**](SshAccountKey.md)| The updated SSH key object | 

### Return type

[**SshAccountKey**](ssh_account_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserSshKeysPost**
> SshAccountKey UsersSelectedUserSshKeysPost(ctx, selectedUser, optional)
Add a new SSH key

Adds a new SSH public key to the specified user account and returns the resulting key.  Example:  ``` $ curl -X POST -H \"Content-Type: application/json\" -d '{\"key\": \"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY user@myhost\"}' https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/ssh-keys ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be an Atlassian Account ID OR the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
 **optional** | ***SshApiUsersSelectedUserSshKeysPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SshApiUsersSelectedUserSshKeysPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SshAccountKey**](SshAccountKey.md)| The new SSH key object. Note that the username property has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis). | 

### Return type

[**SshAccountKey**](ssh_account_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

