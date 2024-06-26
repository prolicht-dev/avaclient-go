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

// OenormConversionApiService OenormConversionApi service
type OenormConversionApiService service

type ApiOenormConversionConvertToAvaRequest struct {
	ctx                      context.Context
	ApiService               *OenormConversionApiService
	removePlainTextLongTexts *bool
	removeHtmlLongTexts      *bool
	oenormFile               *os.File
}

// If set to true, plain text long texts will be removed from the output to reduce response sizes
func (r ApiOenormConversionConvertToAvaRequest) RemovePlainTextLongTexts(removePlainTextLongTexts bool) ApiOenormConversionConvertToAvaRequest {
	r.removePlainTextLongTexts = &removePlainTextLongTexts
	return r
}

// If set to true, html long texts will be removed from the output to reduce response sizes
func (r ApiOenormConversionConvertToAvaRequest) RemoveHtmlLongTexts(removeHtmlLongTexts bool) ApiOenormConversionConvertToAvaRequest {
	r.removeHtmlLongTexts = &removeHtmlLongTexts
	return r
}

// The input file
func (r ApiOenormConversionConvertToAvaRequest) OenormFile(oenormFile *os.File) ApiOenormConversionConvertToAvaRequest {
	r.oenormFile = oenormFile
	return r
}

func (r ApiOenormConversionConvertToAvaRequest) Execute() (*ProjectDto, *http.Response, error) {
	return r.ApiService.OenormConversionConvertToAvaExecute(r)
}

/*
OenormConversionConvertToAva Converts ÖNorm files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOenormConversionConvertToAvaRequest
*/
func (a *OenormConversionApiService) OenormConversionConvertToAva(ctx context.Context) ApiOenormConversionConvertToAvaRequest {
	return ApiOenormConversionConvertToAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectDto
func (a *OenormConversionApiService) OenormConversionConvertToAvaExecute(r ApiOenormConversionConvertToAvaRequest) (*ProjectDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OenormConversionApiService.OenormConversionConvertToAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/oenorm/ava"

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
	var oenormFileLocalVarFormFileName string
	var oenormFileLocalVarFileName string
	var oenormFileLocalVarFileBytes []byte

	oenormFileLocalVarFormFileName = "oenormFile"

	oenormFileLocalVarFile := r.oenormFile

	if oenormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(oenormFileLocalVarFile)

		oenormFileLocalVarFileBytes = fbs
		oenormFileLocalVarFileName = oenormFileLocalVarFile.Name()
		oenormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: oenormFileLocalVarFileBytes, fileName: oenormFileLocalVarFileName, formFileName: oenormFileLocalVarFormFileName})
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

type ApiOenormConversionConvertToExcelRequest struct {
	ctx                   context.Context
	ApiService            *OenormConversionApiService
	writePrices           *bool
	writeLongTexts        *bool
	conversionCulture     *string
	includeArticleNumbers *bool
	oenormFile            *os.File
}

// Defaults to true
func (r ApiOenormConversionConvertToExcelRequest) WritePrices(writePrices bool) ApiOenormConversionConvertToExcelRequest {
	r.writePrices = &writePrices
	return r
}

// Defaults to true
func (r ApiOenormConversionConvertToExcelRequest) WriteLongTexts(writeLongTexts bool) ApiOenormConversionConvertToExcelRequest {
	r.writeLongTexts = &writeLongTexts
	return r
}

// The culture that should be used for the conversion process, to have localized Excel files
func (r ApiOenormConversionConvertToExcelRequest) ConversionCulture(conversionCulture string) ApiOenormConversionConvertToExcelRequest {
	r.conversionCulture = &conversionCulture
	return r
}

// If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39;
func (r ApiOenormConversionConvertToExcelRequest) IncludeArticleNumbers(includeArticleNumbers bool) ApiOenormConversionConvertToExcelRequest {
	r.includeArticleNumbers = &includeArticleNumbers
	return r
}

// The input file
func (r ApiOenormConversionConvertToExcelRequest) OenormFile(oenormFile *os.File) ApiOenormConversionConvertToExcelRequest {
	r.oenormFile = oenormFile
	return r
}

func (r ApiOenormConversionConvertToExcelRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.OenormConversionConvertToExcelExecute(r)
}

/*
OenormConversionConvertToExcel Converts ÖNorm files to Excel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOenormConversionConvertToExcelRequest
*/
func (a *OenormConversionApiService) OenormConversionConvertToExcel(ctx context.Context) ApiOenormConversionConvertToExcelRequest {
	return ApiOenormConversionConvertToExcelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *OenormConversionApiService) OenormConversionConvertToExcelExecute(r ApiOenormConversionConvertToExcelRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OenormConversionApiService.OenormConversionConvertToExcel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/oenorm/excel"

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
	var oenormFileLocalVarFormFileName string
	var oenormFileLocalVarFileName string
	var oenormFileLocalVarFileBytes []byte

	oenormFileLocalVarFormFileName = "oenormFile"

	oenormFileLocalVarFile := r.oenormFile

	if oenormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(oenormFileLocalVarFile)

		oenormFileLocalVarFileBytes = fbs
		oenormFileLocalVarFileName = oenormFileLocalVarFile.Name()
		oenormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: oenormFileLocalVarFileBytes, fileName: oenormFileLocalVarFileName, formFileName: oenormFileLocalVarFormFileName})
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

type ApiOenormConversionConvertToFlatAvaRequest struct {
	ctx        context.Context
	ApiService *OenormConversionApiService
	oenormFile *os.File
}

// The input file
func (r ApiOenormConversionConvertToFlatAvaRequest) OenormFile(oenormFile *os.File) ApiOenormConversionConvertToFlatAvaRequest {
	r.oenormFile = oenormFile
	return r
}

func (r ApiOenormConversionConvertToFlatAvaRequest) Execute() (*FlatAvaProject, *http.Response, error) {
	return r.ApiService.OenormConversionConvertToFlatAvaExecute(r)
}

/*
OenormConversionConvertToFlatAva Converts ÖNorm files to Dangl.AVA projects

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOenormConversionConvertToFlatAvaRequest
*/
func (a *OenormConversionApiService) OenormConversionConvertToFlatAva(ctx context.Context) ApiOenormConversionConvertToFlatAvaRequest {
	return ApiOenormConversionConvertToFlatAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlatAvaProject
func (a *OenormConversionApiService) OenormConversionConvertToFlatAvaExecute(r ApiOenormConversionConvertToFlatAvaRequest) (*FlatAvaProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FlatAvaProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OenormConversionApiService.OenormConversionConvertToFlatAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/oenorm/flat-ava"

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
	var oenormFileLocalVarFormFileName string
	var oenormFileLocalVarFileName string
	var oenormFileLocalVarFileBytes []byte

	oenormFileLocalVarFormFileName = "oenormFile"

	oenormFileLocalVarFile := r.oenormFile

	if oenormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(oenormFileLocalVarFile)

		oenormFileLocalVarFileBytes = fbs
		oenormFileLocalVarFileName = oenormFileLocalVarFile.Name()
		oenormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: oenormFileLocalVarFileBytes, fileName: oenormFileLocalVarFileName, formFileName: oenormFileLocalVarFormFileName})
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

type ApiOenormConversionConvertToGaebRequest struct {
	ctx                                   context.Context
	ApiService                            *OenormConversionApiService
	destinationGaebType                   *string
	targetExchangePhaseTransform          *string
	enforceStrictOfferPhaseLongTextOutput *bool
	exportQuantityDetermination           *bool
	removeUnprintableCharactersFromTexts  *bool
	forceIncludeDescriptions              *bool
	treatNullItemNumberSchemaAsInvalid    *bool
	oenormFile                            *os.File
}

// Defaults to GAEB XML V3.2
func (r ApiOenormConversionConvertToGaebRequest) DestinationGaebType(destinationGaebType string) ApiOenormConversionConvertToGaebRequest {
	r.destinationGaebType = &destinationGaebType
	return r
}

// Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86
func (r ApiOenormConversionConvertToGaebRequest) TargetExchangePhaseTransform(targetExchangePhaseTransform string) ApiOenormConversionConvertToGaebRequest {
	r.targetExchangePhaseTransform = &targetExchangePhaseTransform
	return r
}

// Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export.
func (r ApiOenormConversionConvertToGaebRequest) EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput bool) ApiOenormConversionConvertToGaebRequest {
	r.enforceStrictOfferPhaseLongTextOutput = &enforceStrictOfferPhaseLongTextOutput
	return r
}

// Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property.
func (r ApiOenormConversionConvertToGaebRequest) ExportQuantityDetermination(exportQuantityDetermination bool) ApiOenormConversionConvertToGaebRequest {
	r.exportQuantityDetermination = &exportQuantityDetermination
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiOenormConversionConvertToGaebRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiOenormConversionConvertToGaebRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions.
func (r ApiOenormConversionConvertToGaebRequest) ForceIncludeDescriptions(forceIncludeDescriptions bool) ApiOenormConversionConvertToGaebRequest {
	r.forceIncludeDescriptions = &forceIncludeDescriptions
	return r
}

// When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons.
func (r ApiOenormConversionConvertToGaebRequest) TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid bool) ApiOenormConversionConvertToGaebRequest {
	r.treatNullItemNumberSchemaAsInvalid = &treatNullItemNumberSchemaAsInvalid
	return r
}

// The input file
func (r ApiOenormConversionConvertToGaebRequest) OenormFile(oenormFile *os.File) ApiOenormConversionConvertToGaebRequest {
	r.oenormFile = oenormFile
	return r
}

func (r ApiOenormConversionConvertToGaebRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.OenormConversionConvertToGaebExecute(r)
}

/*
OenormConversionConvertToGaeb Converts ÖNorm files to GAEB files.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOenormConversionConvertToGaebRequest
*/
func (a *OenormConversionApiService) OenormConversionConvertToGaeb(ctx context.Context) ApiOenormConversionConvertToGaebRequest {
	return ApiOenormConversionConvertToGaebRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *OenormConversionApiService) OenormConversionConvertToGaebExecute(r ApiOenormConversionConvertToGaebRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OenormConversionApiService.OenormConversionConvertToGaeb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/oenorm/gaeb"

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
	var oenormFileLocalVarFormFileName string
	var oenormFileLocalVarFileName string
	var oenormFileLocalVarFileBytes []byte

	oenormFileLocalVarFormFileName = "oenormFile"

	oenormFileLocalVarFile := r.oenormFile

	if oenormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(oenormFileLocalVarFile)

		oenormFileLocalVarFileBytes = fbs
		oenormFileLocalVarFileName = oenormFileLocalVarFile.Name()
		oenormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: oenormFileLocalVarFileBytes, fileName: oenormFileLocalVarFileName, formFileName: oenormFileLocalVarFormFileName})
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

type ApiOenormConversionConvertToOenormRequest struct {
	ctx                                    context.Context
	ApiService                             *OenormConversionApiService
	destinationOenormType                  *string
	tryRepairProjectStructure              *bool
	skipTryEnforceSchemaCompliantXmlOutput *bool
	removeUnprintableCharactersFromTexts   *bool
	oenormFile                             *os.File
}

// Defaults to Lv2015
func (r ApiOenormConversionConvertToOenormRequest) DestinationOenormType(destinationOenormType string) ApiOenormConversionConvertToOenormRequest {
	r.destinationOenormType = &destinationOenormType
	return r
}

// Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target
func (r ApiOenormConversionConvertToOenormRequest) TryRepairProjectStructure(tryRepairProjectStructure bool) ApiOenormConversionConvertToOenormRequest {
	r.tryRepairProjectStructure = &tryRepairProjectStructure
	return r
}

// If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option.
func (r ApiOenormConversionConvertToOenormRequest) SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput bool) ApiOenormConversionConvertToOenormRequest {
	r.skipTryEnforceSchemaCompliantXmlOutput = &skipTryEnforceSchemaCompliantXmlOutput
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiOenormConversionConvertToOenormRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiOenormConversionConvertToOenormRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// The input file
func (r ApiOenormConversionConvertToOenormRequest) OenormFile(oenormFile *os.File) ApiOenormConversionConvertToOenormRequest {
	r.oenormFile = oenormFile
	return r
}

func (r ApiOenormConversionConvertToOenormRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.OenormConversionConvertToOenormExecute(r)
}

/*
OenormConversionConvertToOenorm Converts ÖNorm files to Oenorm files.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOenormConversionConvertToOenormRequest
*/
func (a *OenormConversionApiService) OenormConversionConvertToOenorm(ctx context.Context) ApiOenormConversionConvertToOenormRequest {
	return ApiOenormConversionConvertToOenormRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *OenormConversionApiService) OenormConversionConvertToOenormExecute(r ApiOenormConversionConvertToOenormRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OenormConversionApiService.OenormConversionConvertToOenorm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/oenorm/oenorm"

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
	var oenormFileLocalVarFormFileName string
	var oenormFileLocalVarFileName string
	var oenormFileLocalVarFileBytes []byte

	oenormFileLocalVarFormFileName = "oenormFile"

	oenormFileLocalVarFile := r.oenormFile

	if oenormFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(oenormFileLocalVarFile)

		oenormFileLocalVarFileBytes = fbs
		oenormFileLocalVarFileName = oenormFileLocalVarFile.Name()
		oenormFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: oenormFileLocalVarFileBytes, fileName: oenormFileLocalVarFileName, formFileName: oenormFileLocalVarFormFileName})
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
