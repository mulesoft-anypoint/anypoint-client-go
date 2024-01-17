/*
Authorization API

Description of the Authentication API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiApiV2Oauth2TokenPostRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	credentials *Credentials
}

// Request body containing credentials
func (r DefaultApiApiV2Oauth2TokenPostRequest) Credentials(credentials Credentials) DefaultApiApiV2Oauth2TokenPostRequest {
	r.credentials = &credentials
	return r
}

func (r DefaultApiApiV2Oauth2TokenPostRequest) Execute() (*ApiV2Oauth2TokenPost200Response, *http.Response, error) {
	return r.ApiService.ApiV2Oauth2TokenPostExecute(r)
}

/*
ApiV2Oauth2TokenPost Returns access token

Authenticates a connected app and returns access token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DefaultApiApiV2Oauth2TokenPostRequest
*/
func (a *DefaultApiService) ApiV2Oauth2TokenPost(ctx context.Context) DefaultApiApiV2Oauth2TokenPostRequest {
	return DefaultApiApiV2Oauth2TokenPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiV2Oauth2TokenPost200Response
func (a *DefaultApiService) ApiV2Oauth2TokenPostExecute(r DefaultApiApiV2Oauth2TokenPostRequest) (*ApiV2Oauth2TokenPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiV2Oauth2TokenPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ApiV2Oauth2TokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.credentials == nil {
		return localVarReturnValue, nil, reportError("credentials is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.credentials
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiLoginPostRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	userPwdCredentials *UserPwdCredentials
}

// Request body containing credentials
func (r DefaultApiLoginPostRequest) UserPwdCredentials(userPwdCredentials UserPwdCredentials) DefaultApiLoginPostRequest {
	r.userPwdCredentials = &userPwdCredentials
	return r
}

func (r DefaultApiLoginPostRequest) Execute() (*LoginPost200Response, *http.Response, error) {
	return r.ApiService.LoginPostExecute(r)
}

/*
LoginPost Returns access token

Authenticates a user and returns access token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DefaultApiLoginPostRequest
*/
func (a *DefaultApiService) LoginPost(ctx context.Context) DefaultApiLoginPostRequest {
	return DefaultApiLoginPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LoginPost200Response
func (a *DefaultApiService) LoginPostExecute(r DefaultApiLoginPostRequest) (*LoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.LoginPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userPwdCredentials == nil {
		return localVarReturnValue, nil, reportError("userPwdCredentials is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userPwdCredentials
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
