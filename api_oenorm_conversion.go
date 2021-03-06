/*
 * AVACloud API 1.16.0
 *
 * AVACloud API specification
 *
 * API version: 1.16.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avaclient

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"os"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// OenormConversionApiService OenormConversionApi service
type OenormConversionApiService service

// OenormConversionConvertToAvaOpts Optional parameters for the method 'OenormConversionConvertToAva'
type OenormConversionConvertToAvaOpts struct {
	RemovePlainTextLongTexts optional.Bool
	RemoveHtmlLongTexts      optional.Bool
	OenormFile               optional.Interface
}

/*
OenormConversionConvertToAva Converts ÖNorm files to Dangl.AVA projects
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OenormConversionConvertToAvaOpts - Optional Parameters:
 * @param "RemovePlainTextLongTexts" (optional.Bool) -  If set to true, plain text long texts will be removed from the output to reduce response sizes
 * @param "RemoveHtmlLongTexts" (optional.Bool) -  If set to true, html long texts will be removed from the output to reduce response sizes
 * @param "OenormFile" (optional.Interface of *os.File) -  The input file
@return ProjectDto
*/
func (a *OenormConversionApiService) OenormConversionConvertToAva(ctx _context.Context, localVarOptionals *OenormConversionConvertToAvaOpts) (ProjectDto, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProjectDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversion/oenorm/ava"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.RemovePlainTextLongTexts.IsSet() {
		localVarQueryParams.Add("RemovePlainTextLongTexts", parameterToString(localVarOptionals.RemovePlainTextLongTexts.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RemoveHtmlLongTexts.IsSet() {
		localVarQueryParams.Add("RemoveHtmlLongTexts", parameterToString(localVarOptionals.RemoveHtmlLongTexts.Value(), ""))
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
	localVarFormFileName = "oenormFile"
	var localVarFile *os.File
	if localVarOptionals != nil && localVarOptionals.OenormFile.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.OenormFile.Value().(*os.File)
		if !localVarFileOk {
			return localVarReturnValue, nil, reportError("oenormFile should be *os.File")
		}
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
			newErr.model = v
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

// OenormConversionConvertToExcelOpts Optional parameters for the method 'OenormConversionConvertToExcel'
type OenormConversionConvertToExcelOpts struct {
	WritePrices       optional.Bool
	WriteLongTexts    optional.Bool
	ConversionCulture optional.String
	OenormFile        optional.Interface
}

/*
OenormConversionConvertToExcel Converts ÖNorm files to Excel
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OenormConversionConvertToExcelOpts - Optional Parameters:
 * @param "WritePrices" (optional.Bool) -  Defaults to true
 * @param "WriteLongTexts" (optional.Bool) -  Defaults to true
 * @param "ConversionCulture" (optional.String) -  The culture that should be used for the conversion process, to have localized Excel files
 * @param "OenormFile" (optional.Interface of *os.File) -  The input file
@return *os.File
*/
func (a *OenormConversionApiService) OenormConversionConvertToExcel(ctx _context.Context, localVarOptionals *OenormConversionConvertToExcelOpts) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversion/oenorm/excel"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.WritePrices.IsSet() {
		localVarQueryParams.Add("WritePrices", parameterToString(localVarOptionals.WritePrices.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WriteLongTexts.IsSet() {
		localVarQueryParams.Add("WriteLongTexts", parameterToString(localVarOptionals.WriteLongTexts.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConversionCulture.IsSet() {
		localVarQueryParams.Add("ConversionCulture", parameterToString(localVarOptionals.ConversionCulture.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "oenormFile"
	var localVarFile *os.File
	if localVarOptionals != nil && localVarOptionals.OenormFile.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.OenormFile.Value().(*os.File)
		if !localVarFileOk {
			return localVarReturnValue, nil, reportError("oenormFile should be *os.File")
		}
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
			newErr.model = v
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

// OenormConversionConvertToGaebOpts Optional parameters for the method 'OenormConversionConvertToGaeb'
type OenormConversionConvertToGaebOpts struct {
	DestinationGaebType          optional.String
	TargetExchangePhaseTransform optional.String
	OenormFile                   optional.Interface
}

/*
OenormConversionConvertToGaeb Converts ÖNorm files to GAEB files.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OenormConversionConvertToGaebOpts - Optional Parameters:
 * @param "DestinationGaebType" (optional.String) -  Defaults to GAEB XML V3.2
 * @param "TargetExchangePhaseTransform" (optional.String) -  Defaults to none, meaning no transformation will be done
 * @param "OenormFile" (optional.Interface of *os.File) -  The input file
@return *os.File
*/
func (a *OenormConversionApiService) OenormConversionConvertToGaeb(ctx _context.Context, localVarOptionals *OenormConversionConvertToGaebOpts) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversion/oenorm/gaeb"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.DestinationGaebType.IsSet() {
		localVarQueryParams.Add("DestinationGaebType", parameterToString(localVarOptionals.DestinationGaebType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TargetExchangePhaseTransform.IsSet() {
		localVarQueryParams.Add("TargetExchangePhaseTransform", parameterToString(localVarOptionals.TargetExchangePhaseTransform.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "oenormFile"
	var localVarFile *os.File
	if localVarOptionals != nil && localVarOptionals.OenormFile.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.OenormFile.Value().(*os.File)
		if !localVarFileOk {
			return localVarReturnValue, nil, reportError("oenormFile should be *os.File")
		}
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
			newErr.model = v
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

// OenormConversionConvertToOenormOpts Optional parameters for the method 'OenormConversionConvertToOenorm'
type OenormConversionConvertToOenormOpts struct {
	DestinationOenormType     optional.String
	TryRepairProjectStructure optional.Bool
	OenormFile                optional.Interface
}

/*
OenormConversionConvertToOenorm Converts ÖNorm files to Oenorm files.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OenormConversionConvertToOenormOpts - Optional Parameters:
 * @param "DestinationOenormType" (optional.String) -  Defaults to Lv2015
 * @param "TryRepairProjectStructure" (optional.Bool) -  Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target
 * @param "OenormFile" (optional.Interface of *os.File) -  The input file
@return *os.File
*/
func (a *OenormConversionApiService) OenormConversionConvertToOenorm(ctx _context.Context, localVarOptionals *OenormConversionConvertToOenormOpts) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/conversion/oenorm/oenorm"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.DestinationOenormType.IsSet() {
		localVarQueryParams.Add("DestinationOenormType", parameterToString(localVarOptionals.DestinationOenormType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TryRepairProjectStructure.IsSet() {
		localVarQueryParams.Add("TryRepairProjectStructure", parameterToString(localVarOptionals.TryRepairProjectStructure.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "oenormFile"
	var localVarFile *os.File
	if localVarOptionals != nil && localVarOptionals.OenormFile.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.OenormFile.Value().(*os.File)
		if !localVarFileOk {
			return localVarReturnValue, nil, reportError("oenormFile should be *os.File")
		}
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
			newErr.model = v
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
