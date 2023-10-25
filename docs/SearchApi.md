# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAccount**](SearchApi.md#SearchAccount) | **Get** /users/{selected_user}/search/code | Search for code in a user&#x27;s repositories
[**SearchTeam**](SearchApi.md#SearchTeam) | **Get** /teams/{username}/search/code | Search for code in a team&#x27;s repositories
[**SearchWorkspace**](SearchApi.md#SearchWorkspace) | **Get** /workspaces/{workspace}/search/code | Search for code in a workspace

# **SearchAccount**
> SearchResultPage SearchAccount(ctx, selectedUser, searchQuery, optional)
Search for code in a user's repositories

Search for code in the repositories of the specified user.  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both.  You can use the same syntax for the search query as in the UI. E.g. to search for \"foo\" only within the repository \"demo\", use the query parameter `search_query=foo+repo:demo`.  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files, use the query parameter `search_query=foo&fields=%2Bvalues.file.commit.repository` (the `%2B` is a URL-encoded `+`). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **searchQuery** | **string**| The search query | 
 **optional** | ***SearchApiSearchAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page of the search results to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](search_result_page.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTeam**
> SearchResultPage SearchTeam(ctx, username, searchQuery, optional)
Search for code in a team's repositories

Search for code in the repositories of the specified team.  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both.  You can use the same syntax for the search query as in the UI. E.g. to search for \"foo\" only within the repository \"demo\", use the query parameter `search_query=foo+repo:demo`.  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files, use the query parameter `search_query=foo&fields=%2Bvalues.file.commit.repository` (the `%2B` is a URL-encoded `+`).  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account to search in; either the username or the UUID in curly braces | 
  **searchQuery** | **string**| The search query | 
 **optional** | ***SearchApiSearchTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchTeamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page of the search results to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](search_result_page.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchWorkspace**
> SearchResultPage SearchWorkspace(ctx, workspace, searchQuery, optional)
Search for code in a workspace

Search for code in the repositories of the specified workspace.  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both.  You can use the same syntax for the search query as in the UI. E.g. to search for \"foo\" only within the repository \"demo\", use the query parameter `search_query=foo+repo:demo`.  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files, use the query parameter `search_query=foo&fields=%2Bvalues.file.commit.repository` (the `%2B` is a URL-encoded `+`).  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace to search in; either the slug or the UUID in curly braces | 
  **searchQuery** | **string**| The search query | 
 **optional** | ***SearchApiSearchWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchWorkspaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page of the search results to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](search_result_page.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

