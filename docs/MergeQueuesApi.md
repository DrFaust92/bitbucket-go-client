# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMergeQueue**](MergeQueuesApi.md#GetMergeQueue) | **Get** /repositories/{workspace}/{repo_slug}/merge-queues/{merge_queue_uuid} | Get a merge queue
[**GetMergeQueuePullRequest**](MergeQueuesApi.md#GetMergeQueuePullRequest) | **Get** /repositories/{workspace}/{repo_slug}/merge-queues/{merge_queue_uuid}/queue/{pullrequest_id} | Get a pullrequest in a merge queue
[**ListMergeQueuePullRequests**](MergeQueuesApi.md#ListMergeQueuePullRequests) | **Get** /repositories/{workspace}/{repo_slug}/merge-queues/{merge_queue_uuid}/queue | List pullrequests in a merge queue
[**ListMergeQueues**](MergeQueuesApi.md#ListMergeQueues) | **Get** /repositories/{workspace}/{repo_slug}/merge-queues | List merge queues
[**RemovePullRequestFromMergeQueue**](MergeQueuesApi.md#RemovePullRequestFromMergeQueue) | **Delete** /repositories/{workspace}/{repo_slug}/merge-queues/{merge_queue_uuid}/queue/{pullrequest_id} | Remove a pullrequest from a merge queue

# **GetMergeQueue**
> MergeQueue GetMergeQueue(ctx, workspace, repoSlug, mergeQueueUuid)
Get a merge queue

Gets a merge queue by UUID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace containing the repository. | 
  **repoSlug** | **string**| The repository slug. | 
  **mergeQueueUuid** | [**string**](.md)| The UUID of the merge queue. | 

### Return type

[**MergeQueue**](merge_queue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMergeQueuePullRequest**
> MergeQueuePullrequest GetMergeQueuePullRequest(ctx, workspace, repoSlug, mergeQueueUuid, pullrequestId)
Get a pullrequest in a merge queue

Gets a specific pullrequest in a merge queue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace containing the repository. | 
  **repoSlug** | **string**| The repository slug. | 
  **mergeQueueUuid** | [**string**](.md)| The UUID of the merge queue. | 
  **pullrequestId** | **int32**| The ID of the pullrequest. | 

### Return type

[**MergeQueuePullrequest**](merge_queue_pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMergeQueuePullRequests**
> PaginatedMergeQueuePullrequests ListMergeQueuePullRequests(ctx, workspace, repoSlug, mergeQueueUuid, optional)
List pullrequests in a merge queue

Lists pullrequests currently associated with a merge queue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace containing the repository. | 
  **repoSlug** | **string**| The repository slug. | 
  **mergeQueueUuid** | [**string**](.md)| The UUID of the merge queue. | 
 **optional** | ***MergeQueuesApiListMergeQueuePullRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeQueuesApiListMergeQueuePullRequestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**| The page number of elements to retrieve. | [default to 1]
 **pagelen** | **optional.Int32**| The maximum number of results to return. | [default to 10]

### Return type

[**PaginatedMergeQueuePullrequests**](paginated_merge_queue_pullrequests.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMergeQueues**
> PaginatedMergeQueues ListMergeQueues(ctx, workspace, repoSlug, optional)
List merge queues

Lists merge queues configured for a repository. Optionally filter by branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace containing the repository. | 
  **repoSlug** | **string**| The repository slug. | 
 **optional** | ***MergeQueuesApiListMergeQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeQueuesApiListMergeQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetBranch** | **optional.String**| Optional target branch name to filter merge queues. | 
 **page** | **optional.Int32**| The page number of elements to retrieve. | [default to 1]
 **pagelen** | **optional.Int32**| The maximum number of results to return. | [default to 10]

### Return type

[**PaginatedMergeQueues**](paginated_merge_queues.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePullRequestFromMergeQueue**
> RemovePullRequestFromMergeQueue(ctx, workspace, repoSlug, mergeQueueUuid, pullrequestId)
Remove a pullrequest from a merge queue

Removes a pullrequest from the specified merge queue asynchronously.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace containing the repository. | 
  **repoSlug** | **string**| The repository slug. | 
  **mergeQueueUuid** | [**string**](.md)| The UUID of the merge queue. | 
  **pullrequestId** | **int32**| The ID of the pullrequest to remove. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

