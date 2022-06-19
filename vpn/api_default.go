/*
 * VPN API
 *
 * Description of the VPN API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpn

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
}


func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdIpsecGetExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdIpsecGet Returns a list of vpns.
 * Returns a list of VPNs for the given organization and VPC
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param vpcId The vpc Id
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecGet(ctx _context.Context, orgId string, vpcId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecGetExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecGetRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdIpsecGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/ipsec"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	vpnPostReqBody *VpnPostReqBody
}

func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest) VpnPostReqBody(vpnPostReqBody VpnPostReqBody) DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest {
	r.vpnPostReqBody = &vpnPostReqBody
	return r
}

func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest) Execute() (VpnPost, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdIpsecPostExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdIpsecPost Creates a VPN.
 * Create a VPN connection from a VPC, up to a limit of 10 total VPN Connections per
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param vpcId The vpc Id
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecPost(ctx _context.Context, orgId string, vpcId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
	}
}

/*
 * Execute executes the request
 * @return VpnPost
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecPostExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecPostRequest) (VpnPost, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VpnPost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdIpsecPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/ipsec"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.vpnPostReqBody == nil {
		return localVarReturnValue, nil, reportError("vpnPostReqBody is required and must be specified")
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
	localVarPostBody = r.vpnPostReqBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	vpnId string
}


func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete Delete a VPN connection
 * The VPN connection from a VPC.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param vpcId The vpc Id
 * @param vpnId The vpn Id
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete(ctx _context.Context, orgId string, vpcId string, vpnId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
		vpnId: vpnId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/ipsec/{vpnId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpnId"+"}", _neturl.PathEscape(parameterToString(r.vpnId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	vpcId string
	vpnId string
}


func (r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest) Execute() (VpnGet, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetExecute(r)
}

/*
 * OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet Returns a a specific vpn
 * Returns a specific VPN for the given vpn, organization and VPC
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The organization Id
 * @param vpcId The vpc Id
 * @param vpnId The vpn Id
 * @return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet(ctx _context.Context, orgId string, vpcId string, vpnId string) DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest {
	return DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		vpcId: vpcId,
		vpnId: vpnId,
	}
}

/*
 * Execute executes the request
 * @return VpnGet
 */
func (a *DefaultApiService) OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetExecute(r DefaultApiApiOrganizationsOrgIdVpcsVpcIdIpsecVpnIdGetRequest) (VpnGet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VpnGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdVpcsVpcIdIpsecVpnIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/vpcs/{vpcId}/ipsec/{vpnId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpcId"+"}", _neturl.PathEscape(parameterToString(r.vpcId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"vpnId"+"}", _neturl.PathEscape(parameterToString(r.vpnId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}