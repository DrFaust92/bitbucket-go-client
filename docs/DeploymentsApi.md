# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](DeploymentsApi.md#CreateEnvironment) | **Post** /repositories/{workspace}/{repo_slug}/environments | Create an environment
[**DeleteEnvironmentForRepository**](DeploymentsApi.md#DeleteEnvironmentForRepository) | **Delete** /repositories/{workspace}/{repo_slug}/environments/{environment_uuid} | Delete an environment
[**GetDeploymentForRepository**](DeploymentsApi.md#GetDeploymentForRepository) | **Get** /repositories/{workspace}/{repo_slug}/deployments/{deployment_uuid} | Get a deployment
[**GetDeploymentsForRepository**](DeploymentsApi.md#GetDeploymentsForRepository) | **Get** /repositories/{workspace}/{repo_slug}/deployments | List deployments
[**GetEnvironmentForRepository**](DeploymentsApi.md#GetEnvironmentForRepository) | **Get** /repositories/{workspace}/{repo_slug}/environments/{environment_uuid} | Get an environment
[**GetEnvironmentsForRepository**](DeploymentsApi.md#GetEnvironmentsForRepository) | **Get** /repositories/{workspace}/{repo_slug}/environments | List environments
[**RepositoriesWorkspaceRepoSlugDeployKeysGet**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysGet) | **Get** /repositories/{workspace}/{repo_slug}/deploy-keys | List repository deploy keys
[**RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id} | Delete a repository deploy key
[**RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet) | **Get** /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id} | Get a repository deploy key
[**RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut) | **Put** /repositories/{workspace}/{repo_slug}/deploy-keys/{key_id} | Update a repository deploy key
[**RepositoriesWorkspaceRepoSlugDeployKeysPost**](DeploymentsApi.md#RepositoriesWorkspaceRepoSlugDeployKeysPost) | **Post** /repositories/{workspace}/{repo_slug}/deploy-keys | Add a repository deploy key
[**UpdateEnvironmentForRepository**](DeploymentsApi.md#UpdateEnvironmentForRepository) | **Post** /repositories/{workspace}/{repo_slug}/environments/{environment_uuid}/changes | Update an environment
[**WorkspacesWorkspaceProjectsProjectKeyDeployKeysGet**](DeploymentsApi.md#WorkspacesWorkspaceProjectsProjectKeyDeployKeysGet) | **Get** /workspaces/{workspace}/projects/{project_key}/deploy-keys | List project deploy keys
[**WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdDelete**](DeploymentsApi.md#WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdDelete) | **Delete** /workspaces/{workspace}/projects/{project_key}/deploy-keys/{key_id} | Delete a deploy key from a project
[**WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdGet**](DeploymentsApi.md#WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdGet) | **Get** /workspaces/{workspace}/projects/{project_key}/deploy-keys/{key_id} | Get a project deploy key
[**WorkspacesWorkspaceProjectsProjectKeyDeployKeysPost**](DeploymentsApi.md#WorkspacesWorkspaceProjectsProjectKeyDeployKeysPost) | **Post** /workspaces/{workspace}/projects/{project_key}/deploy-keys | Create a project deploy key

# **CreateEnvironment**
> DeploymentEnvironment CreateEnvironment(ctx, body, workspace, repoSlug)
Create an environment

Create an environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeploymentEnvironment**](DeploymentEnvironment.md)| The environment to create. | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**DeploymentEnvironment**](deployment_environment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironmentForRepository**
> DeleteEnvironmentForRepository(ctx, workspace, repoSlug, environmentUuid)
Delete an environment

Delete an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment UUID. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentForRepository**
> Deployment GetDeploymentForRepository(ctx, workspace, repoSlug, deploymentUuid)
Get a deployment

Retrieve a deployment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **deploymentUuid** | **string**| The deployment UUID. | 

### Return type

[**Deployment**](deployment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentsForRepository**
> PaginatedDeployments GetDeploymentsForRepository(ctx, workspace, repoSlug, optional)
List deployments

Find deployments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
 **optional** | ***DeploymentsApiGetDeploymentsForRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiGetDeploymentsForRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedDeployments**](paginated_deployments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentForRepository**
> DeploymentEnvironment GetEnvironmentForRepository(ctx, workspace, repoSlug, environmentUuid)
Get an environment

Retrieve an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment UUID. | 

### Return type

[**DeploymentEnvironment**](deployment_environment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironmentsForRepository**
> PaginatedEnvironments GetEnvironmentsForRepository(ctx, workspace, repoSlug, optional)
List environments

Find environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
 **optional** | ***DeploymentsApiGetEnvironmentsForRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiGetEnvironmentsForRepositoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedEnvironments**](paginated_environments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDeployKeysGet**
> PaginatedDeployKeys RepositoriesWorkspaceRepoSlugDeployKeysGet(ctx, repoSlug, workspace, optional)
List repository deploy keys

Returns all deploy-keys belonging to a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***DeploymentsApiRepositoriesWorkspaceRepoSlugDeployKeysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiRepositoriesWorkspaceRepoSlugDeployKeysGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedDeployKeys**](paginated_deploy_keys.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete**
> RepositoriesWorkspaceRepoSlugDeployKeysKeyIdDelete(ctx, keyId, repoSlug, workspace)
Delete a repository deploy key

This deletes a deploy key from a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The key ID matching the deploy key. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet**
> DeployKey RepositoriesWorkspaceRepoSlugDeployKeysKeyIdGet(ctx, keyId, repoSlug, workspace)
Get a repository deploy key

Returns the deploy key belonging to a specific key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The key ID matching the deploy key. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**DeployKey**](deploy_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut**
> DeployKey RepositoriesWorkspaceRepoSlugDeployKeysKeyIdPut(ctx, keyId, repoSlug, workspace)
Update a repository deploy key

Create a new deploy key in a repository.  The same key needs to be passed in but the comment and label can change.  Example: ``` $ curl -X PUT \\ -H \"Authorization <auth header>\" \\ -H \"Content-type: application/json\" \\ https://api.bitbucket.org/2.0/repositories/mleu/test/deploy-keys/1234 -d \\ '{     \"label\": \"newlabel\",     \"key\": \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAK/b1cHHDr/TEV1JGQl+WjCwStKG6Bhrv0rFpEsYlyTBm1fzN0VOJJYn4ZOPCPJwqse6fGbXntEs+BbXiptR+++HycVgl65TMR0b5ul5AgwrVdZdT7qjCOCgaSV74/9xlHDK8oqgGnfA7ZoBBU+qpVyaloSjBdJfLtPY/xqj4yHnXKYzrtn/uFc4Kp9Tb7PUg9Io3qohSTGJGVHnsVblq/rToJG7L5xIo0OxK0SJSQ5vuId93ZuFZrCNMXj8JDHZeSEtjJzpRCBEXHxpOPhAcbm4MzULgkFHhAVgp4JbkrT99/wpvZ7r9AdkTg7HGqL3rlaDrEcWfL7Lu6TnhBdq5 newcomment\", }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The key ID matching the deploy key. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**DeployKey**](deploy_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDeployKeysPost**
> DeployKey RepositoriesWorkspaceRepoSlugDeployKeysPost(ctx, repoSlug, workspace)
Add a repository deploy key

Create a new deploy key in a repository. Note: If authenticating a deploy key with an OAuth consumer, any changes to the OAuth consumer will subsequently invalidate the deploy key.   Example: ``` $ curl -X POST \\ -H \"Authorization <auth header>\" \\ -H \"Content-type: application/json\" \\ https://api.bitbucket.org/2.0/repositories/mleu/test/deploy-keys -d \\ '{     \"key\": \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAK/b1cHHDr/TEV1JGQl+WjCwStKG6Bhrv0rFpEsYlyTBm1fzN0VOJJYn4ZOPCPJwqse6fGbXntEs+BbXiptR+++HycVgl65TMR0b5ul5AgwrVdZdT7qjCOCgaSV74/9xlHDK8oqgGnfA7ZoBBU+qpVyaloSjBdJfLtPY/xqj4yHnXKYzrtn/uFc4Kp9Tb7PUg9Io3qohSTGJGVHnsVblq/rToJG7L5xIo0OxK0SJSQ5vuId93ZuFZrCNMXj8JDHZeSEtjJzpRCBEXHxpOPhAcbm4MzULgkFHhAVgp4JbkrT99/wpvZ7r9AdkTg7HGqL3rlaDrEcWfL7Lu6TnhBdq5 mleu@C02W454JHTD8\",     \"label\": \"mydeploykey\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**DeployKey**](deploy_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironmentForRepository**
> UpdateEnvironmentForRepository(ctx, workspace, repoSlug, environmentUuid)
Update an environment

Update an environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment UUID. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDeployKeysGet**
> PaginatedProjectDeployKeys WorkspacesWorkspaceProjectsProjectKeyDeployKeysGet(ctx, projectKey, workspace, optional)
List project deploy keys

Returns all deploy keys belonging to a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***DeploymentsApiWorkspacesWorkspaceProjectsProjectKeyDeployKeysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeploymentsApiWorkspacesWorkspaceProjectsProjectKeyDeployKeysGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| page | 

### Return type

[**PaginatedProjectDeployKeys**](paginated_project_deploy_keys.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdDelete**
> WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdDelete(ctx, keyId, projectKey, workspace)
Delete a deploy key from a project

This deletes a deploy key from a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The key ID matching the project deploy key. | 
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdGet**
> ProjectDeployKey WorkspacesWorkspaceProjectsProjectKeyDeployKeysKeyIdGet(ctx, keyId, projectKey, workspace)
Get a project deploy key

Returns the deploy key belonging to a specific key ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | **string**| The key ID matching the project deploy key. | 
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ProjectDeployKey**](project_deploy_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDeployKeysPost**
> ProjectDeployKey WorkspacesWorkspaceProjectsProjectKeyDeployKeysPost(ctx, projectKey, workspace)
Create a project deploy key

Create a new deploy key in a project.  Example: ``` $ curl -X POST \\ -H \"Authorization <auth header>\" \\ -H \"Content-type: application/json\" \\ https://api.bitbucket.org/2.0/workspaces/standard/projects/TEST_PROJECT/deploy-keys/ -d \\ '{     \"key\": \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAK/b1cHHDr/TEV1JGQl+WjCwStKG6Bhrv0rFpEsYlyTBm1fzN0VOJJYn4ZOPCPJwqse6fGbXntEs+BbXiptR+++HycVgl65TMR0b5ul5AgwrVdZdT7qjCOCgaSV74/9xlHDK8oqgGnfA7ZoBBU+qpVyaloSjBdJfLtPY/xqj4yHnXKYzrtn/uFc4Kp9Tb7PUg9Io3qohSTGJGVHnsVblq/rToJG7L5xIo0OxK0SJSQ5vuId93ZuFZrCNMXj8JDHZeSEtjJzpRCBEXHxpOPhAcbm4MzULgkFHhAVgp4JbkrT99/wpvZ7r9AdkTg7HGqL3rlaDrEcWfL7Lu6TnhBdq5 mleu@C02W454JHTD8\",     \"label\": \"mydeploykey\" }' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ProjectDeployKey**](project_deploy_key.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

