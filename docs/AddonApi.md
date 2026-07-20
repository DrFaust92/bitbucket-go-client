# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonAddonKeyClientKeyGet**](AddonApi.md#AddonAddonKeyClientKeyGet) | **Get** /addon/{addon_key}/client-key | Get the client key of a Connect addon
[**AddonDelete**](AddonApi.md#AddonDelete) | **Delete** /addon | Delete an app
[**AddonPut**](AddonApi.md#AddonPut) | **Put** /addon | Update an installed app

# **AddonAddonKeyClientKeyGet**
> AddonAddonKeyClientKeyGet(ctx, addonKey)
Get the client key of a Connect addon

Get the client key of the Connect addon associated with a Forge app install via forgeAppId linkage.  This endpoint is part of the Connect -> Forge migration tooling. It is intended to be used by a Forge app using `asApp().requestBitbucket()` only. Prerequisite: app developer needs to register the linkage between their Connect and Forge app by setting `forgeAppId` in the Connect addon descriptor to `app.id` from Forge app manifest, then update the installations. If the request came from an installation of a registered Forge app, the client key of the linked Connect addon installed in the same workspace will be returned.  ``` api.asApp().requestBitbucket(route`/2.0/addon/{addon-key}/client-key`) ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addonKey** | **string**| The Connect addon key as defined in an application descriptor.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonDelete**
> AddonDelete(ctx, )
Delete an app

Deletes the application for the user.  This endpoint is intended to be used by Bitbucket Connect apps and only supports JWT authentication -- that is how Bitbucket identifies the particular installation of the app. Developers with applications registered in the \"Develop Apps\" section of Bitbucket Marketplace need not use this endpoint as updates for those applications can be sent out via the UI of that section.  ``` $ curl -X DELETE https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" ```

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonPut**
> AddonPut(ctx, )
Update an installed app

Updates the application installation for the user.  This endpoint is intended to be used by Bitbucket Connect apps and only supports JWT authentication -- that is how Bitbucket identifies the particular installation of the app. Developers with applications registered in the \"Develop Apps\" section of Bitbucket need not use this endpoint as updates for those applications can be sent out via the UI of that section.  Passing an empty body will update the installation using the existing descriptor URL.  ``` $ curl -X PUT https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" \\   --header \"Content-Type: application/json\" \\   --data '{}' ```  The new `descriptor` for the installation can be also provided in the body directly.  ``` $ curl -X PUT https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" \\   --header \"Content-Type: application/json\" \\   --data '{\"descriptor\": $NEW_DESCRIPTOR}' ```  In both these modes the URL of the descriptor cannot be changed. To change the descriptor location and upgrade an installation the request must be made exclusively with a `descriptor_url`.   ``` $ curl -X PUT https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" \\   --header \"Content-Type: application/json\" \\   --data '{\"descriptor_url\": $NEW_URL}' ```  The `descriptor_url` must exactly match the marketplace registration that Atlassian has for the application. Contact your Atlassian developer advocate to update this registration. Once the registration has been updated you may call this resource for each installation.  Note that the scopes of the application cannot be increased in the new descriptor nor reduced to none.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

