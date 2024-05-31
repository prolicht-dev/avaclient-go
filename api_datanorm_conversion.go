/*
AVACloud API 1.51.0

AVACloud API specification

API version: 1.51.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)

// DatanormConversionApiService DatanormConversionApi service
type DatanormConversionApiService service

type ApiDatanormConversionConvertToAvaRequest struct {
	ctx                      context.Context
	ApiService               *DatanormConversionApiService
	removePlainTextLongTexts *bool
	removeHtmlLongTexts      *bool
	datanormFile             *os.File
}

// If set to true, plain text long texts will be removed from the output to reduce response sizes
func (r ApiDatanormConversionConvertToAvaRequest) RemovePlainTextLongTexts(removePlainTextLongTexts bool) ApiDatanormConversionConvertToAvaRequest {
	r.removePlainTextLongTexts = &removePlainTextLongTexts
	return r
}

// If set to true, html long texts will be removed from the output to reduce response sizes
func (r ApiDatanormConversionConvertToAvaRequest) RemoveHtmlLongTexts(removeHtmlLongTexts bool) ApiDatanormConversionConvertToAvaRequest {
	r.removeHtmlLongTexts = &removeHtmlLongTexts
	return r
}

// The input file
func (r ApiDatanormConversionConvertToAvaRequest) DatanormFile(datanormFile *os.File) ApiDatanormConversionConvertToAvaRequest {
	r.datanormFile = datanormFile
	return r
}

func (r ApiDatanormConversionConvertToAvaRequest) Execute() (*ProjectDto, *http.Response, error) {
	return r.ApiService.DatanormConversionConvertToAvaExecute(r)
}

/*
DatanormConversionConvertToAva Converts Datanorm files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDatanormConversionConvertToAvaRequest
*/
func (a *DatanormConversionApiService) DatanormConversionConvertToAva(ctx context.Context) ApiDatanormConversionConvertToAvaRequest {
	return ApiDatanormConversionConvertToAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectDto
func (a *DatanormConversionApiService) DatanormConversionConvertToAvaExecute(r ApiDatanormConversionConvertToAvaRequest) (*ProjectDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatanormConversionApiService.DatanormConversionConvertToAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/datanorm/ava"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.removePlainTextLongTexts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RemovePlainTextLongTexts", r.removePlainTextLongTexts, "")
	}
	if r.removeHtmlLongTexts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RemoveHtmlLongTexts", r.removeHtmlLongTexts, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.com.dangl-it.ProjectDto.v1+json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var datanormFileLocalVarFormFileName string
	var datanormFileLocalVarFileName string
	var datanormFileLocalVarFileBytes []byte

	datanormFileLocalVarFormFileName = "datanormFile"

	datanormFileLocalVarFile := r.datanormFile

	if datanormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(datanormFileLocalVarFile)

		datanormFileLocalVarFileBytes = fbs
		datanormFileLocalVarFileName = datanormFileLocalVarFile.Name()
		datanormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: datanormFileLocalVarFileBytes, fileName: datanormFileLocalVarFileName, formFileName: datanormFileLocalVarFormFileName})
	}
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
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiDatanormConversionConvertToFlatAvaRequest struct {
	ctx          context.Context
	ApiService   *DatanormConversionApiService
	datanormFile *os.File
}

// The input file
func (r ApiDatanormConversionConvertToFlatAvaRequest) DatanormFile(datanormFile *os.File) ApiDatanormConversionConvertToFlatAvaRequest {
	r.datanormFile = datanormFile
	return r
}

func (r ApiDatanormConversionConvertToFlatAvaRequest) Execute() (*FlatAvaProject, *http.Response, error) {
	return r.ApiService.DatanormConversionConvertToFlatAvaExecute(r)
}

/*
DatanormConversionConvertToFlatAva Converts Datanorm files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDatanormConversionConvertToFlatAvaRequest
*/
func (a *DatanormConversionApiService) DatanormConversionConvertToFlatAva(ctx context.Context) ApiDatanormConversionConvertToFlatAvaRequest {
	return ApiDatanormConversionConvertToFlatAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlatAvaProject
func (a *DatanormConversionApiService) DatanormConversionConvertToFlatAvaExecute(r ApiDatanormConversionConvertToFlatAvaRequest) (*FlatAvaProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FlatAvaProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatanormConversionApiService.DatanormConversionConvertToFlatAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/datanorm/flat-ava"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var datanormFileLocalVarFormFileName string
	var datanormFileLocalVarFileName string
	var datanormFileLocalVarFileBytes []byte

	datanormFileLocalVarFormFileName = "datanormFile"

	datanormFileLocalVarFile := r.datanormFile

	if datanormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(datanormFileLocalVarFile)

		datanormFileLocalVarFileBytes = fbs
		datanormFileLocalVarFileName = datanormFileLocalVarFile.Name()
		datanormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: datanormFileLocalVarFileBytes, fileName: datanormFileLocalVarFileName, formFileName: datanormFileLocalVarFormFileName})
	}
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
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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
