# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserWorkspacesGet**](WorkspacesApi.md#UserWorkspacesGet) | **Get** /user/workspaces | List workspaces for the current user
[**UserWorkspacesWorkspacePermissionGet**](WorkspacesApi.md#UserWorkspacesWorkspacePermissionGet) | **Get** /user/workspaces/{workspace}/permission | Get user permission on a workspace
[**WorkspacesWorkspaceCodesearchPost**](WorkspacesApi.md#WorkspacesWorkspaceCodesearchPost) | **Post** /workspaces/{workspace}/codesearch | Search code in a workspace
[**WorkspacesWorkspaceGet**](WorkspacesApi.md#WorkspacesWorkspaceGet) | **Get** /workspaces/{workspace} | Get a workspace
[**WorkspacesWorkspaceHooksGet**](WorkspacesApi.md#WorkspacesWorkspaceHooksGet) | **Get** /workspaces/{workspace}/hooks | List webhooks for a workspace
[**WorkspacesWorkspaceHooksPost**](WorkspacesApi.md#WorkspacesWorkspaceHooksPost) | **Post** /workspaces/{workspace}/hooks | Create a webhook for a workspace
[**WorkspacesWorkspaceHooksUidDelete**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidDelete) | **Delete** /workspaces/{workspace}/hooks/{uid} | Delete a webhook for a workspace
[**WorkspacesWorkspaceHooksUidGet**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidGet) | **Get** /workspaces/{workspace}/hooks/{uid} | Get a webhook for a workspace
[**WorkspacesWorkspaceHooksUidPut**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidPut) | **Put** /workspaces/{workspace}/hooks/{uid} | Update a webhook for a workspace
[**WorkspacesWorkspaceMembersGet**](WorkspacesApi.md#WorkspacesWorkspaceMembersGet) | **Get** /workspaces/{workspace}/members | List users in a workspace
[**WorkspacesWorkspaceMembersMemberGet**](WorkspacesApi.md#WorkspacesWorkspaceMembersMemberGet) | **Get** /workspaces/{workspace}/members/{member} | Get user membership for a workspace
[**WorkspacesWorkspacePermissionsGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsGet) | **Get** /workspaces/{workspace}/permissions | List user permissions in a workspace
[**WorkspacesWorkspacePermissionsRepositoriesGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsRepositoriesGet) | **Get** /workspaces/{workspace}/permissions/repositories | List all repository permissions for a workspace
[**WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet) | **Get** /workspaces/{workspace}/permissions/repositories/{repo_slug} | List a repository permissions for a workspace
[**WorkspacesWorkspaceProjectsGet**](WorkspacesApi.md#WorkspacesWorkspaceProjectsGet) | **Get** /workspaces/{workspace}/projects | List projects in a workspace
[**WorkspacesWorkspaceProjectsProjectKeyGet**](WorkspacesApi.md#WorkspacesWorkspaceProjectsProjectKeyGet) | **Get** /workspaces/{workspace}/projects/{project_key} | Get a project for a workspace
[**WorkspacesWorkspacePullrequestsSelectedUserGet**](WorkspacesApi.md#WorkspacesWorkspacePullrequestsSelectedUserGet) | **Get** /workspaces/{workspace}/pullrequests/{selected_user} | List workspace pull requests for a user
[**WorkspacesWorkspaceSettingsGpgPublicKeyGet**](WorkspacesApi.md#WorkspacesWorkspaceSettingsGpgPublicKeyGet) | **Get** /workspaces/{workspace}/settings/gpg/public-key | Get the workspace system GPG public key(s)

# **UserWorkspacesGet**
> PaginatedWorkspaceAccess UserWorkspacesGet(ctx, optional)
List workspaces for the current user

Returns an object for each workspace accessible to the caller. This object also contains details on whether the caller has admin permissions on the workspace (`\"administrator\" = true`) or not (`\"administrator\" = false`).  Queries support filtering based on administrator permissions, [sorting](/cloud/bitbucket/rest/intro/#sorting-query-results) or [filtering](/cloud/bitbucket/rest/intro/#filtering) by `slug`. Results can be [paginated](/cloud/bitbucket/rest/intro/#pagination).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkspacesApiUserWorkspacesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiUserWorkspacesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| Name of a response property to sort results (only slug is supported). | 
 **administrator** | **optional.Bool**| Filter workspaces based on which ones the caller has admin permissions or not. | 

### Return type

[**PaginatedWorkspaceAccess**](paginated_workspace_access.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserWorkspacesWorkspacePermissionGet**
> WorkspaceMembership UserWorkspacesWorkspacePermissionGet(ctx, workspace)
Get user permission on a workspace

Returns the caller's effective role; as in, the highest level of privilege the caller has for the workspace. If the calling user is a member of multiple groups with distinct roles, only the highest level is returned.  Permissions can be:  * `owner` * `create-project` * `collaborator` (deprecated; see this [deprecation announcement](/cloud/bitbucket/deprecation-notice-collaborator-role/) for more details) * `member`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WorkspaceMembership**](workspace_membership.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceCodesearchPost**
> WorkspacesWorkspaceCodesearchPost(ctx, workspace)
Search code in a workspace

Search code across repositories in a workspace.  Results are limited to the repositories that the calling user can access.  #### Supported authentication  The endpoint requires a credential that identifies the calling user. The following authentication modes are accepted:  - Browser session cookie (Bitbucket web UI) - Atlassian API token (Basic auth using your email and a scoped API   token that includes both `repository:read` and `workspace:read`) - OAuth 2.0 via Forge or an Atlassian app (3LO or 2LO, with both   `repository:read` and `workspace:read` granted to the app) - Bitbucket repository, project, or workspace access tokens  Authentication modes that **do not** identify the user, including Atlassian Connect JWT, legacy Bitbucket OAuth not migrated to identity, app passwords, and API keys, are rejected with a `403`.  #### Request body  Requires `Content-Type: application/json`. Anything else returns `415`.  `query.search_term` is required. These fields are accepted:  - `query.search_term`: non-blank, at most 1024 characters - `query.filters` - `config.max_results`: a JSON integer (a string or float is rejected),   1 to 100, default 50 - `config.search_mode`: `LEXICAL` or omitted - `request_id` and `pagination.next_cursor`: see \"Pagination\"  Any other field is ignored.  ```json {   \"query\": {     \"search_term\": \"logr\",     \"filters\": { \"attribute\": \"language\", \"eq\": \"go\" }   },   \"config\": {     \"max_results\": 10   } } ```  #### Search mode  This endpoint performs lexical code search only. `config.search_mode` is fixed to `LEXICAL` server-side, so callers may omit the field or set it explicitly to `LEXICAL`. Any other value returns a `400`.  #### Filters  Use `query.filters` to narrow a search. For example, to search for `logr` only in Go files whose path matches `cmd/_*_/main.go`:  ```json {   \"query\": {     \"search_term\": \"logr\",     \"filters\": {       \"and\": [         { \"attribute\": \"language\", \"eq\": \"go\" },         { \"attribute\": \"path\", \"wildcard\": \"cmd/_*_/main.go\" }       ]     }   } } ```  A filter names an `attribute` and exactly one operator:  * `eq` matches a value, and works with every attribute. Values are   not case-sensitive. * `wildcard` matches a pattern, and works only with `path` and   `repository`. In a pattern, `*` matches any run of characters and `?`   matches one character. A pattern must match the whole value. On   `repository` it ignores case, and on `path` it is case-sensitive.  Setting both `eq` and `wildcard` on one filter returns `400`.  These attributes are supported:  * `repository`: a repository slug, such as `my-repo`. Use a wildcard   such as `my-*` to match every slug that starts with `my-`. * `path`: with `eq`, a value such as `src/main` matches any path that   contains those folder or file names in that order, such as   `app/src/main/App.java`. Each name must match in full. With   `wildcard`, the pattern must match the whole file path, such as   `*_/main.go`. * `language`: a language name, such as `go`, `python`, or   `javascript`. To filter by file extension, use `extension`. * `extension`: a file extension without the dot, such as `go` or `py`. * `project`: a project key or name. * `scm`: the source control host of the repository. * `branches`: searches run on the default branch only. Set `branches`   to `\"true\"` to search all indexed branches. * `forks`: by default, results include repositories that are forks. Set   `forks` to `\"true\"` to get results from forks only, or to `\"false\"` to   get results from repositories that are not forks.  `branches` only accepts the value `\"true\"`. `branches` and `forks` only work when they are the whole of `query.filters`. You cannot combine them with other filters.  Combine filters with `and`, `or`, and `not`. Filters can nest up to 5 levels deep, and the workspace filter this endpoint adds uses one of them. That leaves 3 levels of `and`, `or`, or `not` around a filter. Deeper filters return `400`. This filter uses all 3 levels:  ```json {   \"and\": [     { \"or\": [       { \"not\": { \"attribute\": \"language\", \"eq\": \"python\" } }     ] },     { \"attribute\": \"extension\", \"eq\": \"go\" }   ] } ```  #### Workspace scoping  The endpoint always applies a `namespace = <workspace_slug>` filter and passes the workspace's wrapped UUID as `scope.workspace_ids`. Callers must **not** include a `namespace` filter in `query.filters` (anywhere in the tree, including nested AND/OR/NOT). Doing so returns `400`. A `scope.workspace_ids` value in the request body is ignored. The workspace in the URL always sets the scope.  #### Pagination  Pagination uses a cursor. To page beyond the first result set, include the `id` from the previous response as `request_id`, and the previous response's `pagination.next_cursor` as `pagination.next_cursor`. `request_id` is required whenever `pagination` is sent. Repeat until `pagination.has_more` is `false`.  On every page, `query.search_term` and `query.filters` must be the same as on the first page, including the order of items inside `and`/`or`. A different value returns `400`. `config.max_results` may change between pages.  ```json {   \"query\": { \"search_term\": \"logr\" },   \"config\": { \"max_results\": 10 },   \"request_id\": \"550e8400-e29b-41d4-a716-446655440000\",   \"pagination\": { \"next_cursor\": \"eyJlbnRyaWVzIjpbeyJsZXhpY2FsS2V5Ijoi...\" } } ```  #### Response  You can rely on these fields:  - `id`, `pagination.next_cursor`, `pagination.has_more` - `repositories[]`: `url`, `host`, `ns`, `name` - `matched_files[]`: `path`, `name`, `language`, `url`, `repositoryId`,   `version` (the commit hash), `branch_or_tag_name`, `is_default_branch` - `matched_files[].matched_chunks[]`: `source_code`, `line_range`,   `matched_line_numbers`, `matched_ranges` - `metadata.total_count`: an approximate size, meaningful on page 1 only  `branch_or_tag_name` and `is_default_branch` are left out when not known, and `language` can be `null`.  Line numbers start at 1, and `line_range.end` is the last line of the chunk. Each `matched_ranges` entry is `[line, start_column, end_column]`, where `line` is the line number in the file, the columns start at 0, and `end_column` is exclusive.  `matched_line_numbers` and `matched_ranges` mark where a search term appears as a whole word in the chunk. They can be empty even when the chunk matched.  Other fields can change or be removed without notice, so ignore any field not listed here. Treat ids and `pagination.next_cursor` as opaque.  Errors from this endpoint, including errors from the search backend, use the standard Bitbucket error envelope: `{\"type\": \"error\", \"error\": {\"message\": ..., \"code\": ...}}`. A backend error keeps its status code, except that a backend 5xx is returned as `502 Bad Gateway`. The backend's own error code, when there is one, is in `error.detail.upstream_code`. A `429` includes the backend's `Retry-After` header when the backend sends one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceGet**
> Workspace WorkspacesWorkspaceGet(ctx, workspace)
Get a workspace

Returns the requested workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Workspace**](workspace.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksGet**
> PaginatedWebhookSubscriptions WorkspacesWorkspaceHooksGet(ctx, workspace, optional)
List webhooks for a workspace

Returns a paginated list of webhooks installed on this workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
 **optional** | ***WorkspacesApiWorkspacesWorkspaceHooksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspaceHooksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedWebhookSubscriptions**](paginated_webhook_subscriptions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksPost**
> WebhookSubscription WorkspacesWorkspaceHooksPost(ctx, workspace)
Create a webhook for a workspace

Creates a new webhook on the specified workspace.  Workspace webhooks are fired for events from all repositories contained by that workspace.  Example: ``` $ curl -X POST -u credentials -H 'Content-Type: application/json'   https://api.bitbucket.org/2.0/workspaces/my-workspace/hooks   -d '     {       \"description\": \"Webhook Description\",       \"url\": \"https://example.com/\",       \"active\": true,       \"secret\": \"this is a really bad secret\",       \"events\": [         \"repo:push\",         \"issue:created\",         \"issue:updated\"       ]     }' ```  When the `secret` is provided it will be used as the key to generate a HMAC digest value sent in the `X-Hub-Signature` header at delivery time. Passing a `null` or empty `secret` or not passing a `secret` will leave the webhook's secret unset. Bitbucket only generates the `X-Hub-Signature` when the webhook's secret is set.  This call requires the webhook scope, as well as any scope that applies to the events that the webhook subscribes to. In the example above that means: `webhook`, `repository` and `issue`.  The `url` must properly resolve and cannot be an internal, non-routed address.  Only workspace owners can install webhooks on workspaces. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksUidDelete**
> WorkspacesWorkspaceHooksUidDelete(ctx, workspace, uid)
Delete a webhook for a workspace

Deletes the specified webhook subscription from the given workspace. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **uid** | **string**| The webhook&#x27;s id. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksUidGet**
> WebhookSubscription WorkspacesWorkspaceHooksUidGet(ctx, workspace, uid)
Get a webhook for a workspace

Returns the webhook with the specified id installed on the specified workspace. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **uid** | **string**| The webhook&#x27;s id. | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksUidPut**
> WebhookSubscription WorkspacesWorkspaceHooksUidPut(ctx, workspace, uid)
Update a webhook for a workspace

Updates the specified webhook subscription.  The following properties can be mutated:  * `description` * `url` * `secret` * `active` * `events`  The hook's secret is used as a key to generate the HMAC hex digest sent in the X-Hub-Signature header at delivery time. This signature is only generated when the hook has a secret.  Set the hook's secret by passing the new value in the secret field . Passing a null value in the secret field will remove the secret from the hook. The hook's secret can be left unchanged by not passing the secret field in the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **uid** | **string**| The webhook&#x27;s id. | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceMembersGet**
> PaginatedWorkspaceMemberships WorkspacesWorkspaceMembersGet(ctx, workspace, optional)
List users in a workspace

Returns all members of the requested workspace.  This endpoint additionally supports [filtering](/cloud/bitbucket/rest/intro/#filtering) by email address, if called by a workspace administrator, integration or workspace access token. This is done by adding the following query string parameter:  * `q=user.email IN (\"user1@org.com\",\"user2@org.com\")`  When filtering by email, you can query up to 90 addresses at a time. Note that the query parameter values need to be URL escaped, so the final query string should be:  * `q=user.email%20IN%20(%22user1@org.com%22,%22user2@org.com%22)`  Email addresses that you filter by (and only these email addresses) can be included in the response using the `fields` query parameter:  * `&fields=+values.user.email` - add the `email` field to the default `user` response object * `&fields=values.user.email,values.user.account_id` - only return user email addresses and account IDs  Once again, all query parameter values must be URL escaped.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspaceMembersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspaceMembersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedWorkspaceMemberships**](paginated_workspace_memberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceMembersMemberGet**
> WorkspaceMembership WorkspacesWorkspaceMembersMemberGet(ctx, member, workspace)
Get user membership for a workspace

Returns the workspace membership, which includes a `User` object for the member and a `Workspace` object for the requested workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **member** | **string**| Member&#x27;s UUID or Atlassian ID. | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WorkspaceMembership**](workspace_membership.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePermissionsGet**
> PaginatedWorkspaceMemberships WorkspacesWorkspacePermissionsGet(ctx, workspace, optional)
List user permissions in a workspace

Returns the list of members in a workspace and their permission levels. Permission can be: * `owner` * `collaborator` * `member`  **The `collaborator` role is being removed from the Bitbucket Cloud API. For more information, see the [deprecation announcement](/cloud/bitbucket/deprecation-notice-collaborator-role/).**  **When you move your administration from Bitbucket Cloud to admin.atlassian.com, the following fields on `workspace_membership` will no longer be present: `last_accessed` and `added_on`. See the [deprecation announcement](/cloud/bitbucket/announcement-breaking-change-workspace-membership/).**  Results may be further [filtered](/cloud/bitbucket/rest/intro/#filtering) by permission by adding the following query string parameters:  * `q=permission=\"owner\"`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePermissionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedWorkspaceMemberships**](paginated_workspace_memberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePermissionsRepositoriesGet**
> PaginatedRepositoryPermissions WorkspacesWorkspacePermissionsRepositoriesGet(ctx, workspace, optional)
List all repository permissions for a workspace

Returns an object for each repository permission for all of a workspace's repositories.  Permissions returned are effective permissions: the highest level of permission the user has. This does not distinguish between direct and indirect (group) privileges.  Only users with admin permission for the team may access this resource.  Permissions can be:  * `admin` * `write` * `read`  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by repository, user, or permission by adding the following query string parameters:  * `q=repository.name=\"geordi\"` or `q=permission>\"read\"` * `sort=user.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 
 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedRepositoryPermissions**](paginated_repository_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet**
> PaginatedRepositoryPermissions WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet(ctx, repoSlug, workspace, optional)
List a repository permissions for a workspace

Returns an object for the repository permission of each user in the requested repository.  Permissions returned are effective permissions: the highest level of permission the user has. This does not distinguish between direct and indirect (group) privileges.  Only users with admin permission for the repository may access this resource.  Permissions can be:  * `admin` * `write` * `read`  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by user, or permission by adding the following query string parameters:  * `q=permission>\"read\"` * `sort=user.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesRepoSlugGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesRepoSlugGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 
 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedRepositoryPermissions**](paginated_repository_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsGet**
> PaginatedProjects WorkspacesWorkspaceProjectsGet(ctx, workspace, optional)
List projects in a workspace

Returns the list of projects in this workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspaceProjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspaceProjectsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedProjects**](paginated_projects.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyGet**
> Project WorkspacesWorkspaceProjectsProjectKeyGet(ctx, projectKey, workspace)
Get a project for a workspace

Returns the requested project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Project**](project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePullrequestsSelectedUserGet**
> PaginatedPullrequests WorkspacesWorkspacePullrequestsSelectedUserGet(ctx, selectedUser, workspace, optional)
List workspace pull requests for a user

Returns all workspace pull requests authored by the specified user.  By default only open pull requests are returned. This can be controlled using the `state` query parameter. To retrieve pull requests that are in one of multiple states, repeat the `state` parameter for each individual state.  This endpoint also supports filtering and sorting of the results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be the username of the pull request author, the author&#x27;s UUID surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, or the author&#x27;s Atlassian ID.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePullrequestsSelectedUserGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePullrequestsSelectedUserGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **optional.String**| Only return pull requests that are in this state. This parameter can be repeated. | 
 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedPullrequests**](paginated_pullrequests.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceSettingsGpgPublicKeyGet**
> WorkspacesWorkspaceSettingsGpgPublicKeyGet(ctx, workspace)
Get the workspace system GPG public key(s)

Returns the system public GPG key(s). In most cases a single key is returned. During a key rotation period, two keys may be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

