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

// RebConversionApiService RebConversionApi service
type RebConversionApiService service

type ApiRebConversionConvertToAvaRequest struct {
	ctx                      context.Context
	ApiService               *RebConversionApiService
	removePlainTextLongTexts *bool
	removeHtmlLongTexts      *bool
	rebFile                  *os.File
}

// If set to true, plain text long texts will be removed from the output to reduce response sizes
func (r ApiRebConversionConvertToAvaRequest) RemovePlainTextLongTexts(removePlainTextLongTexts bool) ApiRebConversionConvertToAvaRequest {
	r.removePlainTextLongTexts = &removePlainTextLongTexts
	return r
}

// If set to true, html long texts will be removed from the output to reduce response sizes
func (r ApiRebConversionConvertToAvaRequest) RemoveHtmlLongTexts(removeHtmlLongTexts bool) ApiRebConversionConvertToAvaRequest {
	r.removeHtmlLongTexts = &removeHtmlLongTexts
	return r
}

// The input file
func (r ApiRebConversionConvertToAvaRequest) RebFile(rebFile *os.File) ApiRebConversionConvertToAvaRequest {
	r.rebFile = rebFile
	return r
}

func (r ApiRebConversionConvertToAvaRequest) Execute() (*ProjectDto, *http.Response, error) {
	return r.ApiService.RebConversionConvertToAvaExecute(r)
}

/*
RebConversionConvertToAva Converts REB files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRebConversionConvertToAvaRequest
*/
func (a *RebConversionApiService) RebConversionConvertToAva(ctx context.Context) ApiRebConversionConvertToAvaRequest {
	return ApiRebConversionConvertToAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectDto
func (a *RebConversionApiService) RebConversionConvertToAvaExecute(r ApiRebConversionConvertToAvaRequest) (*ProjectDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RebConversionApiService.RebConversionConvertToAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/reb/ava"

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
	var rebFileLocalVarFormFileName string
	var rebFileLocalVarFileName string
	var rebFileLocalVarFileBytes []byte

	rebFileLocalVarFormFileName = "rebFile"

	rebFileLocalVarFile := r.rebFile

	if rebFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(rebFileLocalVarFile)

		rebFileLocalVarFileBytes = fbs
		rebFileLocalVarFileName = rebFileLocalVarFile.Name()
		rebFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rebFileLocalVarFileBytes, fileName: rebFileLocalVarFileName, formFileName: rebFileLocalVarFormFileName})
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

type ApiRebConversionConvertToExcelRequest struct {
	ctx                   context.Context
	ApiService            *RebConversionApiService
	writePrices           *bool
	writeLongTexts        *bool
	conversionCulture     *string
	includeArticleNumbers *bool
	rebFile               *os.File
}

// Defaults to true
func (r ApiRebConversionConvertToExcelRequest) WritePrices(writePrices bool) ApiRebConversionConvertToExcelRequest {
	r.writePrices = &writePrices
	return r
}

// Defaults to true
func (r ApiRebConversionConvertToExcelRequest) WriteLongTexts(writeLongTexts bool) ApiRebConversionConvertToExcelRequest {
	r.writeLongTexts = &writeLongTexts
	return r
}

// The culture that should be used for the conversion process, to have localized Excel files
func (r ApiRebConversionConvertToExcelRequest) ConversionCulture(conversionCulture string) ApiRebConversionConvertToExcelRequest {
	r.conversionCulture = &conversionCulture
	return r
}

// If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39;
func (r ApiRebConversionConvertToExcelRequest) IncludeArticleNumbers(includeArticleNumbers bool) ApiRebConversionConvertToExcelRequest {
	r.includeArticleNumbers = &includeArticleNumbers
	return r
}

// The input file
func (r ApiRebConversionConvertToExcelRequest) RebFile(rebFile *os.File) ApiRebConversionConvertToExcelRequest {
	r.rebFile = rebFile
	return r
}

func (r ApiRebConversionConvertToExcelRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.RebConversionConvertToExcelExecute(r)
}

/*
RebConversionConvertToExcel Converts REB files to Excel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRebConversionConvertToExcelRequest
*/
func (a *RebConversionApiService) RebConversionConvertToExcel(ctx context.Context) ApiRebConversionConvertToExcelRequest {
	return ApiRebConversionConvertToExcelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *RebConversionApiService) RebConversionConvertToExcelExecute(r ApiRebConversionConvertToExcelRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RebConversionApiService.RebConversionConvertToExcel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/reb/excel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.writePrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "WritePrices", r.writePrices, "")
	}
	if r.writeLongTexts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "WriteLongTexts", r.writeLongTexts, "")
	}
	if r.conversionCulture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ConversionCulture", r.conversionCulture, "")
	}
	if r.includeArticleNumbers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IncludeArticleNumbers", r.includeArticleNumbers, "")
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
	var rebFileLocalVarFormFileName string
	var rebFileLocalVarFileName string
	var rebFileLocalVarFileBytes []byte

	rebFileLocalVarFormFileName = "rebFile"

	rebFileLocalVarFile := r.rebFile

	if rebFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(rebFileLocalVarFile)

		rebFileLocalVarFileBytes = fbs
		rebFileLocalVarFileName = rebFileLocalVarFile.Name()
		rebFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rebFileLocalVarFileBytes, fileName: rebFileLocalVarFileName, formFileName: rebFileLocalVarFormFileName})
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

type ApiRebConversionConvertToFlatAvaRequest struct {
	ctx        context.Context
	ApiService *RebConversionApiService
	rebFile    *os.File
}

// The input file
func (r ApiRebConversionConvertToFlatAvaRequest) RebFile(rebFile *os.File) ApiRebConversionConvertToFlatAvaRequest {
	r.rebFile = rebFile
	return r
}

func (r ApiRebConversionConvertToFlatAvaRequest) Execute() (*FlatAvaProject, *http.Response, error) {
	return r.ApiService.RebConversionConvertToFlatAvaExecute(r)
}

/*
RebConversionConvertToFlatAva Converts REB files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRebConversionConvertToFlatAvaRequest
*/
func (a *RebConversionApiService) RebConversionConvertToFlatAva(ctx context.Context) ApiRebConversionConvertToFlatAvaRequest {
	return ApiRebConversionConvertToFlatAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlatAvaProject
func (a *RebConversionApiService) RebConversionConvertToFlatAvaExecute(r ApiRebConversionConvertToFlatAvaRequest) (*FlatAvaProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FlatAvaProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RebConversionApiService.RebConversionConvertToFlatAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/reb/flat-ava"

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
	var rebFileLocalVarFormFileName string
	var rebFileLocalVarFileName string
	var rebFileLocalVarFileBytes []byte

	rebFileLocalVarFormFileName = "rebFile"

	rebFileLocalVarFile := r.rebFile

	if rebFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(rebFileLocalVarFile)

		rebFileLocalVarFileBytes = fbs
		rebFileLocalVarFileName = rebFileLocalVarFile.Name()
		rebFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rebFileLocalVarFileBytes, fileName: rebFileLocalVarFileName, formFileName: rebFileLocalVarFormFileName})
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

type ApiRebConversionConvertToGaebRequest struct {
	ctx                                   context.Context
	ApiService                            *RebConversionApiService
	destinationGaebType                   *string
	targetExchangePhaseTransform          *string
	enforceStrictOfferPhaseLongTextOutput *bool
	exportQuantityDetermination           *bool
	removeUnprintableCharactersFromTexts  *bool
	forceIncludeDescriptions              *bool
	treatNullItemNumberSchemaAsInvalid    *bool
	rebFile                               *os.File
}

// Defaults to GAEB XML V3.2
func (r ApiRebConversionConvertToGaebRequest) DestinationGaebType(destinationGaebType string) ApiRebConversionConvertToGaebRequest {
	r.destinationGaebType = &destinationGaebType
	return r
}

// Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86
func (r ApiRebConversionConvertToGaebRequest) TargetExchangePhaseTransform(targetExchangePhaseTransform string) ApiRebConversionConvertToGaebRequest {
	r.targetExchangePhaseTransform = &targetExchangePhaseTransform
	return r
}

// Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export.
func (r ApiRebConversionConvertToGaebRequest) EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput bool) ApiRebConversionConvertToGaebRequest {
	r.enforceStrictOfferPhaseLongTextOutput = &enforceStrictOfferPhaseLongTextOutput
	return r
}

// Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property.
func (r ApiRebConversionConvertToGaebRequest) ExportQuantityDetermination(exportQuantityDetermination bool) ApiRebConversionConvertToGaebRequest {
	r.exportQuantityDetermination = &exportQuantityDetermination
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiRebConversionConvertToGaebRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiRebConversionConvertToGaebRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions.
func (r ApiRebConversionConvertToGaebRequest) ForceIncludeDescriptions(forceIncludeDescriptions bool) ApiRebConversionConvertToGaebRequest {
	r.forceIncludeDescriptions = &forceIncludeDescriptions
	return r
}

// When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons.
func (r ApiRebConversionConvertToGaebRequest) TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid bool) ApiRebConversionConvertToGaebRequest {
	r.treatNullItemNumberSchemaAsInvalid = &treatNullItemNumberSchemaAsInvalid
	return r
}

// The input file
func (r ApiRebConversionConvertToGaebRequest) RebFile(rebFile *os.File) ApiRebConversionConvertToGaebRequest {
	r.rebFile = rebFile
	return r
}

func (r ApiRebConversionConvertToGaebRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.RebConversionConvertToGaebExecute(r)
}

/*
RebConversionConvertToGaeb Converts REB files to GAEB files

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRebConversionConvertToGaebRequest
*/
func (a *RebConversionApiService) RebConversionConvertToGaeb(ctx context.Context) ApiRebConversionConvertToGaebRequest {
	return ApiRebConversionConvertToGaebRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *RebConversionApiService) RebConversionConvertToGaebExecute(r ApiRebConversionConvertToGaebRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RebConversionApiService.RebConversionConvertToGaeb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/reb/gaeb"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.destinationGaebType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "DestinationGaebType", r.destinationGaebType, "")
	}
	if r.targetExchangePhaseTransform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "TargetExchangePhaseTransform", r.targetExchangePhaseTransform, "")
	}
	if r.enforceStrictOfferPhaseLongTextOutput != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EnforceStrictOfferPhaseLongTextOutput", r.enforceStrictOfferPhaseLongTextOutput, "")
	}
	if r.exportQuantityDetermination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ExportQuantityDetermination", r.exportQuantityDetermination, "")
	}
	if r.removeUnprintableCharactersFromTexts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RemoveUnprintableCharactersFromTexts", r.removeUnprintableCharactersFromTexts, "")
	}
	if r.forceIncludeDescriptions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ForceIncludeDescriptions", r.forceIncludeDescriptions, "")
	}
	if r.treatNullItemNumberSchemaAsInvalid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "TreatNullItemNumberSchemaAsInvalid", r.treatNullItemNumberSchemaAsInvalid, "")
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
	var rebFileLocalVarFormFileName string
	var rebFileLocalVarFileName string
	var rebFileLocalVarFileBytes []byte

	rebFileLocalVarFormFileName = "rebFile"

	rebFileLocalVarFile := r.rebFile

	if rebFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(rebFileLocalVarFile)

		rebFileLocalVarFileBytes = fbs
		rebFileLocalVarFileName = rebFileLocalVarFile.Name()
		rebFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rebFileLocalVarFileBytes, fileName: rebFileLocalVarFileName, formFileName: rebFileLocalVarFormFileName})
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

type ApiRebConversionConvertToOenormRequest struct {
	ctx                                    context.Context
	ApiService                             *RebConversionApiService
	destinationOenormType                  *string
	tryRepairProjectStructure              *bool
	skipTryEnforceSchemaCompliantXmlOutput *bool
	removeUnprintableCharactersFromTexts   *bool
	rebFile                                *os.File
}

// Defaults to Lv2015
func (r ApiRebConversionConvertToOenormRequest) DestinationOenormType(destinationOenormType string) ApiRebConversionConvertToOenormRequest {
	r.destinationOenormType = &destinationOenormType
	return r
}

// Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target
func (r ApiRebConversionConvertToOenormRequest) TryRepairProjectStructure(tryRepairProjectStructure bool) ApiRebConversionConvertToOenormRequest {
	r.tryRepairProjectStructure = &tryRepairProjectStructure
	return r
}

// If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option.
func (r ApiRebConversionConvertToOenormRequest) SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput bool) ApiRebConversionConvertToOenormRequest {
	r.skipTryEnforceSchemaCompliantXmlOutput = &skipTryEnforceSchemaCompliantXmlOutput
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiRebConversionConvertToOenormRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiRebConversionConvertToOenormRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// The input file
func (r ApiRebConversionConvertToOenormRequest) RebFile(rebFile *os.File) ApiRebConversionConvertToOenormRequest {
	r.rebFile = rebFile
	return r
}

func (r ApiRebConversionConvertToOenormRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.RebConversionConvertToOenormExecute(r)
}

/*
RebConversionConvertToOenorm Converts REB files to Oenorm

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRebConversionConvertToOenormRequest
*/
func (a *RebConversionApiService) RebConversionConvertToOenorm(ctx context.Context) ApiRebConversionConvertToOenormRequest {
	return ApiRebConversionConvertToOenormRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *RebConversionApiService) RebConversionConvertToOenormExecute(r ApiRebConversionConvertToOenormRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RebConversionApiService.RebConversionConvertToOenorm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/reb/oenorm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.destinationOenormType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "DestinationOenormType", r.destinationOenormType, "")
	}
	if r.tryRepairProjectStructure != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "TryRepairProjectStructure", r.tryRepairProjectStructure, "")
	}
	if r.skipTryEnforceSchemaCompliantXmlOutput != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SkipTryEnforceSchemaCompliantXmlOutput", r.skipTryEnforceSchemaCompliantXmlOutput, "")
	}
	if r.removeUnprintableCharactersFromTexts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RemoveUnprintableCharactersFromTexts", r.removeUnprintableCharactersFromTexts, "")
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
	var rebFileLocalVarFormFileName string
	var rebFileLocalVarFileName string
	var rebFileLocalVarFileBytes []byte

	rebFileLocalVarFormFileName = "rebFile"

	rebFileLocalVarFile := r.rebFile

	if rebFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(rebFileLocalVarFile)

		rebFileLocalVarFileBytes = fbs
		rebFileLocalVarFileName = rebFileLocalVarFile.Name()
		rebFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rebFileLocalVarFileBytes, fileName: rebFileLocalVarFileName, formFileName: rebFileLocalVarFormFileName})
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
