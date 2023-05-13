/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bitbucket

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type PropertiesApiService service

/*
PropertiesApiService Delete a commit application property
Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a commit.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param commit The commit.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) DeleteCommitHostedPropertyValue(ctx context.Context, workspace string, repoSlug string, commit string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", fmt.Sprintf("%v", commit), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Delete a pull request application property
Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a pull request.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param pullrequestId The pull request ID.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) DeletePullRequestHostedPropertyValue(ctx context.Context, workspace string, repoSlug string, pullrequestId string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullrequest_id"+"}", fmt.Sprintf("%v", pullrequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Delete a repository application property
Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a repository.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) DeleteRepositoryHostedPropertyValue(ctx context.Context, workspace string, repoSlug string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Delete a user application property
Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a user.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param selectedUser Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) DeleteUserHostedPropertyValue(ctx context.Context, selectedUser string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{selected_user}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", fmt.Sprintf("%v", selectedUser), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Get a commit application property
Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a commit.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param commit The commit.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.

@return ModelError
*/
func (a *PropertiesApiService) GetCommitHostedPropertyValue(ctx context.Context, workspace string, repoSlug string, commit string, appKey string, propertyName string) (ModelError, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ModelError
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", fmt.Sprintf("%v", commit), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
PropertiesApiService Get a pull request application property
Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a pull request.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param pullrequestId The pull request ID.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.

@return ModelError
*/
func (a *PropertiesApiService) GetPullRequestHostedPropertyValue(ctx context.Context, workspace string, repoSlug string, pullrequestId string, appKey string, propertyName string) (ModelError, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ModelError
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullrequest_id"+"}", fmt.Sprintf("%v", pullrequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
PropertiesApiService Get a repository application property
Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a repository.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.

@return ModelError
*/
func (a *PropertiesApiService) GetRepositoryHostedPropertyValue(ctx context.Context, workspace string, repoSlug string, appKey string, propertyName string) (ModelError, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ModelError
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
PropertiesApiService Get a user application property
Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a user.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param selectedUser Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.

@return ModelError
*/
func (a *PropertiesApiService) RetrieveUserHostedPropertyValue(ctx context.Context, selectedUser string, appKey string, propertyName string) (ModelError, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ModelError
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{selected_user}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", fmt.Sprintf("%v", selectedUser), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
PropertiesApiService Update a commit application property
Update an [application property](/cloud/bitbucket/application-properties/) value stored against a commit.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The application property to create or update.
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param commit The commit.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) UpdateCommitHostedPropertyValue(ctx context.Context, body ModelError, workspace string, repoSlug string, commit string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commit"+"}", fmt.Sprintf("%v", commit), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Update a pull request application property
Update an [application property](/cloud/bitbucket/application-properties/) value stored against a pull request.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The application property to create or update.
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param pullrequestId The pull request ID.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) UpdatePullRequestHostedPropertyValue(ctx context.Context, body ModelError, workspace string, repoSlug string, pullrequestId string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullrequest_id"+"}", fmt.Sprintf("%v", pullrequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Update a repository application property
Update an [application property](/cloud/bitbucket/application-properties/) value stored against a repository.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The application property to create or update.
  - @param workspace The repository container; either the workspace slug or the UUID in curly braces.
  - @param repoSlug The repository.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) UpdateRepositoryHostedPropertyValue(ctx context.Context, body ModelError, workspace string, repoSlug string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", fmt.Sprintf("%v", workspace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
PropertiesApiService Update a user application property
Update an [application property](/cloud/bitbucket/application-properties/) value stored against a user.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body The application property to create or update.
  - @param selectedUser Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.
  - @param appKey The key of the Connect app.
  - @param propertyName The name of the property.
*/
func (a *PropertiesApiService) UpdateUserHostedPropertyValue(ctx context.Context, body ModelError, selectedUser string, appKey string, propertyName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{selected_user}/properties/{app_key}/{property_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"selected_user"+"}", fmt.Sprintf("%v", selectedUser), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_key"+"}", fmt.Sprintf("%v", appKey), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"property_name"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
