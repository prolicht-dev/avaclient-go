/*
AVACloud API 1.41.8

AVACloud API specification

API version: 1.41.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// DanglIdentityApiService DanglIdentityApi service
type DanglIdentityApiService service

type ApiDanglIdentityLoginAndReturnTokenRequest struct {
	ctx        context.Context
	ApiService *DanglIdentityApiService
	model      *TokenLoginPost
}

func (r ApiDanglIdentityLoginAndReturnTokenRequest) Model(model TokenLoginPost) ApiDanglIdentityLoginAndReturnTokenRequest {
	r.model = &model
	return r
}

func (r ApiDanglIdentityLoginAndReturnTokenRequest) Execute() (*TokenResponseGet, *http.Response, error) {
	return r.ApiService.DanglIdentityLoginAndReturnTokenExecute(r)
}

/*
DanglIdentityLoginAndReturnToken Method for DanglIdentityLoginAndReturnToken

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDanglIdentityLoginAndReturnTokenRequest
*/
func (a *DanglIdentityApiService) DanglIdentityLoginAndReturnToken(ctx context.Context) ApiDanglIdentityLoginAndReturnTokenRequest {
	return ApiDanglIdentityLoginAndReturnTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenResponseGet
func (a *DanglIdentityApiService) DanglIdentityLoginAndReturnTokenExecute(r ApiDanglIdentityLoginAndReturnTokenRequest) (*TokenResponseGet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenResponseGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DanglIdentityApiService.DanglIdentityLoginAndReturnToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/token-login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model == nil {
		return localVarReturnValue, nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.model
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v TokenResponseGet
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiDanglIdentityLoginWithCookieRequest struct {
	ctx         context.Context
	ApiService  *DanglIdentityApiService
	model       *LoginPost
	redirectUrl *string
}

func (r ApiDanglIdentityLoginWithCookieRequest) Model(model LoginPost) ApiDanglIdentityLoginWithCookieRequest {
	r.model = &model
	return r
}

func (r ApiDanglIdentityLoginWithCookieRequest) RedirectUrl(redirectUrl string) ApiDanglIdentityLoginWithCookieRequest {
	r.redirectUrl = &redirectUrl
	return r
}

func (r ApiDanglIdentityLoginWithCookieRequest) Execute() (*http.Response, error) {
	return r.ApiService.DanglIdentityLoginWithCookieExecute(r)
}

/*
DanglIdentityLoginWithCookie Method for DanglIdentityLoginWithCookie

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDanglIdentityLoginWithCookieRequest
*/
func (a *DanglIdentityApiService) DanglIdentityLoginWithCookie(ctx context.Context) ApiDanglIdentityLoginWithCookieRequest {
	return ApiDanglIdentityLoginWithCookieRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DanglIdentityApiService) DanglIdentityLoginWithCookieExecute(r ApiDanglIdentityLoginWithCookieRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DanglIdentityApiService.DanglIdentityLoginWithCookie")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model == nil {
		return nil, reportError("model is required and must be specified")
	}

	if r.redirectUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "redirectUrl", r.redirectUrl, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

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
	// body params
	localVarPostBody = r.model
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDanglIdentityRefreshTokenRequest struct {
	ctx        context.Context
	ApiService *DanglIdentityApiService
	model      *TokenRefreshPost
}

func (r ApiDanglIdentityRefreshTokenRequest) Model(model TokenRefreshPost) ApiDanglIdentityRefreshTokenRequest {
	r.model = &model
	return r
}

func (r ApiDanglIdentityRefreshTokenRequest) Execute() (*TokenResponseGet, *http.Response, error) {
	return r.ApiService.DanglIdentityRefreshTokenExecute(r)
}

/*
DanglIdentityRefreshToken Method for DanglIdentityRefreshToken

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDanglIdentityRefreshTokenRequest
*/
func (a *DanglIdentityApiService) DanglIdentityRefreshToken(ctx context.Context) ApiDanglIdentityRefreshTokenRequest {
	return ApiDanglIdentityRefreshTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenResponseGet
func (a *DanglIdentityApiService) DanglIdentityRefreshTokenExecute(r ApiDanglIdentityRefreshTokenRequest) (*TokenResponseGet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenResponseGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DanglIdentityApiService.DanglIdentityRefreshToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/token-refresh"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model == nil {
		return localVarReturnValue, nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.model
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v TokenResponseGet
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiDanglIdentityRegisterRequest struct {
	ctx           context.Context
	ApiService    *DanglIdentityApiService
	registerModel *RegisterPost
}

func (r ApiDanglIdentityRegisterRequest) RegisterModel(registerModel RegisterPost) ApiDanglIdentityRegisterRequest {
	r.registerModel = &registerModel
	return r
}

func (r ApiDanglIdentityRegisterRequest) Execute() (*http.Response, error) {
	return r.ApiService.DanglIdentityRegisterExecute(r)
}

/*
DanglIdentityRegister Method for DanglIdentityRegister

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDanglIdentityRegisterRequest
*/
func (a *DanglIdentityApiService) DanglIdentityRegister(ctx context.Context) ApiDanglIdentityRegisterRequest {
	return ApiDanglIdentityRegisterRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DanglIdentityApiService) DanglIdentityRegisterExecute(r ApiDanglIdentityRegisterRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DanglIdentityApiService.DanglIdentityRegister")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.registerModel == nil {
		return nil, reportError("registerModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.registerModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDanglIdentityRequestPasswordResetRequest struct {
	ctx                 context.Context
	ApiService          *DanglIdentityApiService
	forgotPasswordModel *ForgotPasswordPost
}

func (r ApiDanglIdentityRequestPasswordResetRequest) ForgotPasswordModel(forgotPasswordModel ForgotPasswordPost) ApiDanglIdentityRequestPasswordResetRequest {
	r.forgotPasswordModel = &forgotPasswordModel
	return r
}

func (r ApiDanglIdentityRequestPasswordResetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DanglIdentityRequestPasswordResetExecute(r)
}

/*
DanglIdentityRequestPasswordReset Method for DanglIdentityRequestPasswordReset

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDanglIdentityRequestPasswordResetRequest
*/
func (a *DanglIdentityApiService) DanglIdentityRequestPasswordReset(ctx context.Context) ApiDanglIdentityRequestPasswordResetRequest {
	return ApiDanglIdentityRequestPasswordResetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DanglIdentityApiService) DanglIdentityRequestPasswordResetExecute(r ApiDanglIdentityRequestPasswordResetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DanglIdentityApiService.DanglIdentityRequestPasswordReset")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/password-forgotten"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.forgotPasswordModel == nil {
		return nil, reportError("forgotPasswordModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

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
	// body params
	localVarPostBody = r.forgotPasswordModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDanglIdentitySignOutWithSignInManagerRequest struct {
	ctx        context.Context
	ApiService *DanglIdentityApiService
}

func (r ApiDanglIdentitySignOutWithSignInManagerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DanglIdentitySignOutWithSignInManagerExecute(r)
}

/*
DanglIdentitySignOutWithSignInManager Method for DanglIdentitySignOutWithSignInManager

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDanglIdentitySignOutWithSignInManagerRequest
*/
func (a *DanglIdentityApiService) DanglIdentitySignOutWithSignInManager(ctx context.Context) ApiDanglIdentitySignOutWithSignInManagerRequest {
	return ApiDanglIdentitySignOutWithSignInManagerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DanglIdentityApiService) DanglIdentitySignOutWithSignInManagerExecute(r ApiDanglIdentitySignOutWithSignInManagerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DanglIdentityApiService.DanglIdentitySignOutWithSignInManager")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
