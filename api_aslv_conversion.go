/*
AVACloud API 2.0.0

AVACloud API specification

API version: 2.0.0
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

// AslvConversionApiService AslvConversionApi service
type AslvConversionApiService service

type ApiAslvConversionConvertToAvaRequest struct {
	ctx                      context.Context
	ApiService               *AslvConversionApiService
	removePlainTextLongTexts *bool
	removeHtmlLongTexts      *bool
	aslvFile                 *os.File
}

// If set to true, plain text long texts will be removed from the output to reduce response sizes
func (r ApiAslvConversionConvertToAvaRequest) RemovePlainTextLongTexts(removePlainTextLongTexts bool) ApiAslvConversionConvertToAvaRequest {
	r.removePlainTextLongTexts = &removePlainTextLongTexts
	return r
}

// If set to true, html long texts will be removed from the output to reduce response sizes
func (r ApiAslvConversionConvertToAvaRequest) RemoveHtmlLongTexts(removeHtmlLongTexts bool) ApiAslvConversionConvertToAvaRequest {
	r.removeHtmlLongTexts = &removeHtmlLongTexts
	return r
}

// The input file
func (r ApiAslvConversionConvertToAvaRequest) AslvFile(aslvFile *os.File) ApiAslvConversionConvertToAvaRequest {
	r.aslvFile = aslvFile
	return r
}

func (r ApiAslvConversionConvertToAvaRequest) Execute() (*ProjectDto, *http.Response, error) {
	return r.ApiService.AslvConversionConvertToAvaExecute(r)
}

/*
AslvConversionConvertToAva Converts Aslv files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAslvConversionConvertToAvaRequest
*/
func (a *AslvConversionApiService) AslvConversionConvertToAva(ctx context.Context) ApiAslvConversionConvertToAvaRequest {
	return ApiAslvConversionConvertToAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectDto
func (a *AslvConversionApiService) AslvConversionConvertToAvaExecute(r ApiAslvConversionConvertToAvaRequest) (*ProjectDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AslvConversionApiService.AslvConversionConvertToAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/aslv/ava"

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
	var aslvFileLocalVarFormFileName string
	var aslvFileLocalVarFileName string
	var aslvFileLocalVarFileBytes []byte

	aslvFileLocalVarFormFileName = "aslvFile"

	aslvFileLocalVarFile := r.aslvFile

	if aslvFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(aslvFileLocalVarFile)

		aslvFileLocalVarFileBytes = fbs
		aslvFileLocalVarFileName = aslvFileLocalVarFile.Name()
		aslvFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: aslvFileLocalVarFileBytes, fileName: aslvFileLocalVarFileName, formFileName: aslvFileLocalVarFormFileName})
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

type ApiAslvConversionConvertToExcelRequest struct {
	ctx                   context.Context
	ApiService            *AslvConversionApiService
	writePrices           *bool
	writeLongTexts        *bool
	conversionCulture     *string
	includeArticleNumbers *bool
	aslvFile              *os.File
}

// Defaults to true
func (r ApiAslvConversionConvertToExcelRequest) WritePrices(writePrices bool) ApiAslvConversionConvertToExcelRequest {
	r.writePrices = &writePrices
	return r
}

// Defaults to true
func (r ApiAslvConversionConvertToExcelRequest) WriteLongTexts(writeLongTexts bool) ApiAslvConversionConvertToExcelRequest {
	r.writeLongTexts = &writeLongTexts
	return r
}

// The culture that should be used for the conversion process, to have localized Excel files
func (r ApiAslvConversionConvertToExcelRequest) ConversionCulture(conversionCulture string) ApiAslvConversionConvertToExcelRequest {
	r.conversionCulture = &conversionCulture
	return r
}

// If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39;
func (r ApiAslvConversionConvertToExcelRequest) IncludeArticleNumbers(includeArticleNumbers bool) ApiAslvConversionConvertToExcelRequest {
	r.includeArticleNumbers = &includeArticleNumbers
	return r
}

// The input file
func (r ApiAslvConversionConvertToExcelRequest) AslvFile(aslvFile *os.File) ApiAslvConversionConvertToExcelRequest {
	r.aslvFile = aslvFile
	return r
}

func (r ApiAslvConversionConvertToExcelRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.AslvConversionConvertToExcelExecute(r)
}

/*
AslvConversionConvertToExcel Converts Aslv files to Excel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAslvConversionConvertToExcelRequest
*/
func (a *AslvConversionApiService) AslvConversionConvertToExcel(ctx context.Context) ApiAslvConversionConvertToExcelRequest {
	return ApiAslvConversionConvertToExcelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *AslvConversionApiService) AslvConversionConvertToExcelExecute(r ApiAslvConversionConvertToExcelRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AslvConversionApiService.AslvConversionConvertToExcel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/aslv/excel"

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
	var aslvFileLocalVarFormFileName string
	var aslvFileLocalVarFileName string
	var aslvFileLocalVarFileBytes []byte

	aslvFileLocalVarFormFileName = "aslvFile"

	aslvFileLocalVarFile := r.aslvFile

	if aslvFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(aslvFileLocalVarFile)

		aslvFileLocalVarFileBytes = fbs
		aslvFileLocalVarFileName = aslvFileLocalVarFile.Name()
		aslvFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: aslvFileLocalVarFileBytes, fileName: aslvFileLocalVarFileName, formFileName: aslvFileLocalVarFormFileName})
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

type ApiAslvConversionConvertToFlatAvaRequest struct {
	ctx        context.Context
	ApiService *AslvConversionApiService
	aslvFile   *os.File
}

// The input file
func (r ApiAslvConversionConvertToFlatAvaRequest) AslvFile(aslvFile *os.File) ApiAslvConversionConvertToFlatAvaRequest {
	r.aslvFile = aslvFile
	return r
}

func (r ApiAslvConversionConvertToFlatAvaRequest) Execute() (*FlatAvaProject, *http.Response, error) {
	return r.ApiService.AslvConversionConvertToFlatAvaExecute(r)
}

/*
AslvConversionConvertToFlatAva Converts Aslv files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAslvConversionConvertToFlatAvaRequest
*/
func (a *AslvConversionApiService) AslvConversionConvertToFlatAva(ctx context.Context) ApiAslvConversionConvertToFlatAvaRequest {
	return ApiAslvConversionConvertToFlatAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlatAvaProject
func (a *AslvConversionApiService) AslvConversionConvertToFlatAvaExecute(r ApiAslvConversionConvertToFlatAvaRequest) (*FlatAvaProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FlatAvaProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AslvConversionApiService.AslvConversionConvertToFlatAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/aslv/flat-ava"

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
	var aslvFileLocalVarFormFileName string
	var aslvFileLocalVarFileName string
	var aslvFileLocalVarFileBytes []byte

	aslvFileLocalVarFormFileName = "aslvFile"

	aslvFileLocalVarFile := r.aslvFile

	if aslvFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(aslvFileLocalVarFile)

		aslvFileLocalVarFileBytes = fbs
		aslvFileLocalVarFileName = aslvFileLocalVarFile.Name()
		aslvFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: aslvFileLocalVarFileBytes, fileName: aslvFileLocalVarFileName, formFileName: aslvFileLocalVarFormFileName})
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

type ApiAslvConversionConvertToGaebRequest struct {
	ctx                                   context.Context
	ApiService                            *AslvConversionApiService
	destinationGaebType                   *string
	targetExchangePhaseTransform          *string
	enforceStrictOfferPhaseLongTextOutput *bool
	exportQuantityDetermination           *bool
	removeUnprintableCharactersFromTexts  *bool
	forceIncludeDescriptions              *bool
	treatNullItemNumberSchemaAsInvalid    *bool
	aslvFile                              *os.File
}

// Defaults to GAEB XML V3.2
func (r ApiAslvConversionConvertToGaebRequest) DestinationGaebType(destinationGaebType string) ApiAslvConversionConvertToGaebRequest {
	r.destinationGaebType = &destinationGaebType
	return r
}

// Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86
func (r ApiAslvConversionConvertToGaebRequest) TargetExchangePhaseTransform(targetExchangePhaseTransform string) ApiAslvConversionConvertToGaebRequest {
	r.targetExchangePhaseTransform = &targetExchangePhaseTransform
	return r
}

// Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export.
func (r ApiAslvConversionConvertToGaebRequest) EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput bool) ApiAslvConversionConvertToGaebRequest {
	r.enforceStrictOfferPhaseLongTextOutput = &enforceStrictOfferPhaseLongTextOutput
	return r
}

// Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property.
func (r ApiAslvConversionConvertToGaebRequest) ExportQuantityDetermination(exportQuantityDetermination bool) ApiAslvConversionConvertToGaebRequest {
	r.exportQuantityDetermination = &exportQuantityDetermination
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiAslvConversionConvertToGaebRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiAslvConversionConvertToGaebRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions.
func (r ApiAslvConversionConvertToGaebRequest) ForceIncludeDescriptions(forceIncludeDescriptions bool) ApiAslvConversionConvertToGaebRequest {
	r.forceIncludeDescriptions = &forceIncludeDescriptions
	return r
}

// When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons.
func (r ApiAslvConversionConvertToGaebRequest) TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid bool) ApiAslvConversionConvertToGaebRequest {
	r.treatNullItemNumberSchemaAsInvalid = &treatNullItemNumberSchemaAsInvalid
	return r
}

// The input file
func (r ApiAslvConversionConvertToGaebRequest) AslvFile(aslvFile *os.File) ApiAslvConversionConvertToGaebRequest {
	r.aslvFile = aslvFile
	return r
}

func (r ApiAslvConversionConvertToGaebRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.AslvConversionConvertToGaebExecute(r)
}

/*
AslvConversionConvertToGaeb Converts Aslv files to GAEB files

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAslvConversionConvertToGaebRequest
*/
func (a *AslvConversionApiService) AslvConversionConvertToGaeb(ctx context.Context) ApiAslvConversionConvertToGaebRequest {
	return ApiAslvConversionConvertToGaebRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *AslvConversionApiService) AslvConversionConvertToGaebExecute(r ApiAslvConversionConvertToGaebRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AslvConversionApiService.AslvConversionConvertToGaeb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/aslv/gaeb"

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
	var aslvFileLocalVarFormFileName string
	var aslvFileLocalVarFileName string
	var aslvFileLocalVarFileBytes []byte

	aslvFileLocalVarFormFileName = "aslvFile"

	aslvFileLocalVarFile := r.aslvFile

	if aslvFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(aslvFileLocalVarFile)

		aslvFileLocalVarFileBytes = fbs
		aslvFileLocalVarFileName = aslvFileLocalVarFile.Name()
		aslvFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: aslvFileLocalVarFileBytes, fileName: aslvFileLocalVarFileName, formFileName: aslvFileLocalVarFormFileName})
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

type ApiAslvConversionConvertToOenormRequest struct {
	ctx                                    context.Context
	ApiService                             *AslvConversionApiService
	destinationOenormType                  *string
	tryRepairProjectStructure              *bool
	skipTryEnforceSchemaCompliantXmlOutput *bool
	removeUnprintableCharactersFromTexts   *bool
	aslvFile                               *os.File
}

// Defaults to Lv2015
func (r ApiAslvConversionConvertToOenormRequest) DestinationOenormType(destinationOenormType string) ApiAslvConversionConvertToOenormRequest {
	r.destinationOenormType = &destinationOenormType
	return r
}

// Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target
func (r ApiAslvConversionConvertToOenormRequest) TryRepairProjectStructure(tryRepairProjectStructure bool) ApiAslvConversionConvertToOenormRequest {
	r.tryRepairProjectStructure = &tryRepairProjectStructure
	return r
}

// If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option.
func (r ApiAslvConversionConvertToOenormRequest) SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput bool) ApiAslvConversionConvertToOenormRequest {
	r.skipTryEnforceSchemaCompliantXmlOutput = &skipTryEnforceSchemaCompliantXmlOutput
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiAslvConversionConvertToOenormRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiAslvConversionConvertToOenormRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// The input file
func (r ApiAslvConversionConvertToOenormRequest) AslvFile(aslvFile *os.File) ApiAslvConversionConvertToOenormRequest {
	r.aslvFile = aslvFile
	return r
}

func (r ApiAslvConversionConvertToOenormRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.AslvConversionConvertToOenormExecute(r)
}

/*
AslvConversionConvertToOenorm Converts Aslv files to Oenorm files

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAslvConversionConvertToOenormRequest
*/
func (a *AslvConversionApiService) AslvConversionConvertToOenorm(ctx context.Context) ApiAslvConversionConvertToOenormRequest {
	return ApiAslvConversionConvertToOenormRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *AslvConversionApiService) AslvConversionConvertToOenormExecute(r ApiAslvConversionConvertToOenormRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AslvConversionApiService.AslvConversionConvertToOenorm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/aslv/oenorm"

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
	var aslvFileLocalVarFormFileName string
	var aslvFileLocalVarFileName string
	var aslvFileLocalVarFileBytes []byte

	aslvFileLocalVarFormFileName = "aslvFile"

	aslvFileLocalVarFile := r.aslvFile

	if aslvFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(aslvFileLocalVarFile)

		aslvFileLocalVarFileBytes = fbs
		aslvFileLocalVarFileName = aslvFileLocalVarFile.Name()
		aslvFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: aslvFileLocalVarFileBytes, fileName: aslvFileLocalVarFileName, formFileName: aslvFileLocalVarFormFileName})
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
