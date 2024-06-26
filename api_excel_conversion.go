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

// ExcelConversionApiService ExcelConversionApi service
type ExcelConversionApiService service

type ApiExcelConversionConvertToAvaRequest struct {
	ctx                      context.Context
	ApiService               *ExcelConversionApiService
	readNewElements          *bool
	rebuildItemNumberSchema  *bool
	removePlainTextLongTexts *bool
	removeHtmlLongTexts      *bool
	excelFile                *os.File
}

// Defaults to false
func (r ApiExcelConversionConvertToAvaRequest) ReadNewElements(readNewElements bool) ApiExcelConversionConvertToAvaRequest {
	r.readNewElements = &readNewElements
	return r
}

// When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false.
func (r ApiExcelConversionConvertToAvaRequest) RebuildItemNumberSchema(rebuildItemNumberSchema bool) ApiExcelConversionConvertToAvaRequest {
	r.rebuildItemNumberSchema = &rebuildItemNumberSchema
	return r
}

// If set to true, plain text long texts will be removed from the output to reduce response sizes
func (r ApiExcelConversionConvertToAvaRequest) RemovePlainTextLongTexts(removePlainTextLongTexts bool) ApiExcelConversionConvertToAvaRequest {
	r.removePlainTextLongTexts = &removePlainTextLongTexts
	return r
}

// If set to true, html long texts will be removed from the output to reduce response sizes
func (r ApiExcelConversionConvertToAvaRequest) RemoveHtmlLongTexts(removeHtmlLongTexts bool) ApiExcelConversionConvertToAvaRequest {
	r.removeHtmlLongTexts = &removeHtmlLongTexts
	return r
}

// The input file
func (r ApiExcelConversionConvertToAvaRequest) ExcelFile(excelFile *os.File) ApiExcelConversionConvertToAvaRequest {
	r.excelFile = excelFile
	return r
}

func (r ApiExcelConversionConvertToAvaRequest) Execute() (*ProjectDto, *http.Response, error) {
	return r.ApiService.ExcelConversionConvertToAvaExecute(r)
}

/*
ExcelConversionConvertToAva Converts Excel files to Dangl.AVA projects.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExcelConversionConvertToAvaRequest
*/
func (a *ExcelConversionApiService) ExcelConversionConvertToAva(ctx context.Context) ApiExcelConversionConvertToAvaRequest {
	return ApiExcelConversionConvertToAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectDto
func (a *ExcelConversionApiService) ExcelConversionConvertToAvaExecute(r ApiExcelConversionConvertToAvaRequest) (*ProjectDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExcelConversionApiService.ExcelConversionConvertToAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/excel/ava"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.readNewElements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ReadNewElements", r.readNewElements, "")
	}
	if r.rebuildItemNumberSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RebuildItemNumberSchema", r.rebuildItemNumberSchema, "")
	}
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
	var excelFileLocalVarFormFileName string
	var excelFileLocalVarFileName string
	var excelFileLocalVarFileBytes []byte

	excelFileLocalVarFormFileName = "excelFile"

	excelFileLocalVarFile := r.excelFile

	if excelFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(excelFileLocalVarFile)

		excelFileLocalVarFileBytes = fbs
		excelFileLocalVarFileName = excelFileLocalVarFile.Name()
		excelFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: excelFileLocalVarFileBytes, fileName: excelFileLocalVarFileName, formFileName: excelFileLocalVarFormFileName})
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

type ApiExcelConversionConvertToExcelRequest struct {
	ctx                     context.Context
	ApiService              *ExcelConversionApiService
	readNewElements         *bool
	rebuildItemNumberSchema *bool
	writePrices             *bool
	writeLongTexts          *bool
	conversionCulture       *string
	includeArticleNumbers   *bool
	excelFile               *os.File
}

// Defaults to false
func (r ApiExcelConversionConvertToExcelRequest) ReadNewElements(readNewElements bool) ApiExcelConversionConvertToExcelRequest {
	r.readNewElements = &readNewElements
	return r
}

// When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false.
func (r ApiExcelConversionConvertToExcelRequest) RebuildItemNumberSchema(rebuildItemNumberSchema bool) ApiExcelConversionConvertToExcelRequest {
	r.rebuildItemNumberSchema = &rebuildItemNumberSchema
	return r
}

// Defaults to true
func (r ApiExcelConversionConvertToExcelRequest) WritePrices(writePrices bool) ApiExcelConversionConvertToExcelRequest {
	r.writePrices = &writePrices
	return r
}

// Defaults to true
func (r ApiExcelConversionConvertToExcelRequest) WriteLongTexts(writeLongTexts bool) ApiExcelConversionConvertToExcelRequest {
	r.writeLongTexts = &writeLongTexts
	return r
}

// The culture that should be used for the conversion process, to have localized Excel files
func (r ApiExcelConversionConvertToExcelRequest) ConversionCulture(conversionCulture string) ApiExcelConversionConvertToExcelRequest {
	r.conversionCulture = &conversionCulture
	return r
}

// If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39;
func (r ApiExcelConversionConvertToExcelRequest) IncludeArticleNumbers(includeArticleNumbers bool) ApiExcelConversionConvertToExcelRequest {
	r.includeArticleNumbers = &includeArticleNumbers
	return r
}

// The input file
func (r ApiExcelConversionConvertToExcelRequest) ExcelFile(excelFile *os.File) ApiExcelConversionConvertToExcelRequest {
	r.excelFile = excelFile
	return r
}

func (r ApiExcelConversionConvertToExcelRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ExcelConversionConvertToExcelExecute(r)
}

/*
ExcelConversionConvertToExcel Converts Excel files to Excel files. Used, for example, when elements were added in excel to generate or modify a project. The Excel file can then be shared containing the full project with all formattings, formulas and styles applied.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExcelConversionConvertToExcelRequest
*/
func (a *ExcelConversionApiService) ExcelConversionConvertToExcel(ctx context.Context) ApiExcelConversionConvertToExcelRequest {
	return ApiExcelConversionConvertToExcelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *ExcelConversionApiService) ExcelConversionConvertToExcelExecute(r ApiExcelConversionConvertToExcelRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExcelConversionApiService.ExcelConversionConvertToExcel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/excel/excel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.readNewElements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ReadNewElements", r.readNewElements, "")
	}
	if r.rebuildItemNumberSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RebuildItemNumberSchema", r.rebuildItemNumberSchema, "")
	}
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
	var excelFileLocalVarFormFileName string
	var excelFileLocalVarFileName string
	var excelFileLocalVarFileBytes []byte

	excelFileLocalVarFormFileName = "excelFile"

	excelFileLocalVarFile := r.excelFile

	if excelFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(excelFileLocalVarFile)

		excelFileLocalVarFileBytes = fbs
		excelFileLocalVarFileName = excelFileLocalVarFile.Name()
		excelFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: excelFileLocalVarFileBytes, fileName: excelFileLocalVarFileName, formFileName: excelFileLocalVarFormFileName})
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

type ApiExcelConversionConvertToFlatAvaRequest struct {
	ctx                     context.Context
	ApiService              *ExcelConversionApiService
	readNewElements         *bool
	rebuildItemNumberSchema *bool
	excelFile               *os.File
}

// Defaults to false
func (r ApiExcelConversionConvertToFlatAvaRequest) ReadNewElements(readNewElements bool) ApiExcelConversionConvertToFlatAvaRequest {
	r.readNewElements = &readNewElements
	return r
}

// When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false.
func (r ApiExcelConversionConvertToFlatAvaRequest) RebuildItemNumberSchema(rebuildItemNumberSchema bool) ApiExcelConversionConvertToFlatAvaRequest {
	r.rebuildItemNumberSchema = &rebuildItemNumberSchema
	return r
}

// The input file
func (r ApiExcelConversionConvertToFlatAvaRequest) ExcelFile(excelFile *os.File) ApiExcelConversionConvertToFlatAvaRequest {
	r.excelFile = excelFile
	return r
}

func (r ApiExcelConversionConvertToFlatAvaRequest) Execute() (*FlatAvaProject, *http.Response, error) {
	return r.ApiService.ExcelConversionConvertToFlatAvaExecute(r)
}

/*
ExcelConversionConvertToFlatAva Converts Excel files to Dangl.AVA projects.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExcelConversionConvertToFlatAvaRequest
*/
func (a *ExcelConversionApiService) ExcelConversionConvertToFlatAva(ctx context.Context) ApiExcelConversionConvertToFlatAvaRequest {
	return ApiExcelConversionConvertToFlatAvaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlatAvaProject
func (a *ExcelConversionApiService) ExcelConversionConvertToFlatAvaExecute(r ApiExcelConversionConvertToFlatAvaRequest) (*FlatAvaProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FlatAvaProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExcelConversionApiService.ExcelConversionConvertToFlatAva")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/excel/flat-ava"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.readNewElements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ReadNewElements", r.readNewElements, "")
	}
	if r.rebuildItemNumberSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RebuildItemNumberSchema", r.rebuildItemNumberSchema, "")
	}
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
	var excelFileLocalVarFormFileName string
	var excelFileLocalVarFileName string
	var excelFileLocalVarFileBytes []byte

	excelFileLocalVarFormFileName = "excelFile"

	excelFileLocalVarFile := r.excelFile

	if excelFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(excelFileLocalVarFile)

		excelFileLocalVarFileBytes = fbs
		excelFileLocalVarFileName = excelFileLocalVarFile.Name()
		excelFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: excelFileLocalVarFileBytes, fileName: excelFileLocalVarFileName, formFileName: excelFileLocalVarFormFileName})
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

type ApiExcelConversionConvertToGaebRequest struct {
	ctx                                   context.Context
	ApiService                            *ExcelConversionApiService
	readNewElements                       *bool
	rebuildItemNumberSchema               *bool
	destinationGaebType                   *string
	targetExchangePhaseTransform          *string
	enforceStrictOfferPhaseLongTextOutput *bool
	exportQuantityDetermination           *bool
	removeUnprintableCharactersFromTexts  *bool
	forceIncludeDescriptions              *bool
	treatNullItemNumberSchemaAsInvalid    *bool
	excelFile                             *os.File
}

// Defaults to false
func (r ApiExcelConversionConvertToGaebRequest) ReadNewElements(readNewElements bool) ApiExcelConversionConvertToGaebRequest {
	r.readNewElements = &readNewElements
	return r
}

// When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false.
func (r ApiExcelConversionConvertToGaebRequest) RebuildItemNumberSchema(rebuildItemNumberSchema bool) ApiExcelConversionConvertToGaebRequest {
	r.rebuildItemNumberSchema = &rebuildItemNumberSchema
	return r
}

// Defaults to GAEB XML V3.2
func (r ApiExcelConversionConvertToGaebRequest) DestinationGaebType(destinationGaebType string) ApiExcelConversionConvertToGaebRequest {
	r.destinationGaebType = &destinationGaebType
	return r
}

// Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86
func (r ApiExcelConversionConvertToGaebRequest) TargetExchangePhaseTransform(targetExchangePhaseTransform string) ApiExcelConversionConvertToGaebRequest {
	r.targetExchangePhaseTransform = &targetExchangePhaseTransform
	return r
}

// Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export.
func (r ApiExcelConversionConvertToGaebRequest) EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput bool) ApiExcelConversionConvertToGaebRequest {
	r.enforceStrictOfferPhaseLongTextOutput = &enforceStrictOfferPhaseLongTextOutput
	return r
}

// Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property.
func (r ApiExcelConversionConvertToGaebRequest) ExportQuantityDetermination(exportQuantityDetermination bool) ApiExcelConversionConvertToGaebRequest {
	r.exportQuantityDetermination = &exportQuantityDetermination
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiExcelConversionConvertToGaebRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiExcelConversionConvertToGaebRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions.
func (r ApiExcelConversionConvertToGaebRequest) ForceIncludeDescriptions(forceIncludeDescriptions bool) ApiExcelConversionConvertToGaebRequest {
	r.forceIncludeDescriptions = &forceIncludeDescriptions
	return r
}

// When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons.
func (r ApiExcelConversionConvertToGaebRequest) TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid bool) ApiExcelConversionConvertToGaebRequest {
	r.treatNullItemNumberSchemaAsInvalid = &treatNullItemNumberSchemaAsInvalid
	return r
}

// The input file
func (r ApiExcelConversionConvertToGaebRequest) ExcelFile(excelFile *os.File) ApiExcelConversionConvertToGaebRequest {
	r.excelFile = excelFile
	return r
}

func (r ApiExcelConversionConvertToGaebRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ExcelConversionConvertToGaebExecute(r)
}

/*
ExcelConversionConvertToGaeb Converts Excel files to GAEB files.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExcelConversionConvertToGaebRequest
*/
func (a *ExcelConversionApiService) ExcelConversionConvertToGaeb(ctx context.Context) ApiExcelConversionConvertToGaebRequest {
	return ApiExcelConversionConvertToGaebRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *ExcelConversionApiService) ExcelConversionConvertToGaebExecute(r ApiExcelConversionConvertToGaebRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExcelConversionApiService.ExcelConversionConvertToGaeb")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/excel/gaeb"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.readNewElements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ReadNewElements", r.readNewElements, "")
	}
	if r.rebuildItemNumberSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RebuildItemNumberSchema", r.rebuildItemNumberSchema, "")
	}
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
	var excelFileLocalVarFormFileName string
	var excelFileLocalVarFileName string
	var excelFileLocalVarFileBytes []byte

	excelFileLocalVarFormFileName = "excelFile"

	excelFileLocalVarFile := r.excelFile

	if excelFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(excelFileLocalVarFile)

		excelFileLocalVarFileBytes = fbs
		excelFileLocalVarFileName = excelFileLocalVarFile.Name()
		excelFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: excelFileLocalVarFileBytes, fileName: excelFileLocalVarFileName, formFileName: excelFileLocalVarFormFileName})
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

type ApiExcelConversionConvertToOenormRequest struct {
	ctx                                    context.Context
	ApiService                             *ExcelConversionApiService
	readNewElements                        *bool
	rebuildItemNumberSchema                *bool
	destinationOenormType                  *string
	tryRepairProjectStructure              *bool
	skipTryEnforceSchemaCompliantXmlOutput *bool
	removeUnprintableCharactersFromTexts   *bool
	excelFile                              *os.File
}

// Defaults to false
func (r ApiExcelConversionConvertToOenormRequest) ReadNewElements(readNewElements bool) ApiExcelConversionConvertToOenormRequest {
	r.readNewElements = &readNewElements
	return r
}

// When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false.
func (r ApiExcelConversionConvertToOenormRequest) RebuildItemNumberSchema(rebuildItemNumberSchema bool) ApiExcelConversionConvertToOenormRequest {
	r.rebuildItemNumberSchema = &rebuildItemNumberSchema
	return r
}

// Defaults to Lv2015
func (r ApiExcelConversionConvertToOenormRequest) DestinationOenormType(destinationOenormType string) ApiExcelConversionConvertToOenormRequest {
	r.destinationOenormType = &destinationOenormType
	return r
}

// Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target
func (r ApiExcelConversionConvertToOenormRequest) TryRepairProjectStructure(tryRepairProjectStructure bool) ApiExcelConversionConvertToOenormRequest {
	r.tryRepairProjectStructure = &tryRepairProjectStructure
	return r
}

// If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option.
func (r ApiExcelConversionConvertToOenormRequest) SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput bool) ApiExcelConversionConvertToOenormRequest {
	r.skipTryEnforceSchemaCompliantXmlOutput = &skipTryEnforceSchemaCompliantXmlOutput
	return r
}

// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
func (r ApiExcelConversionConvertToOenormRequest) RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts bool) ApiExcelConversionConvertToOenormRequest {
	r.removeUnprintableCharactersFromTexts = &removeUnprintableCharactersFromTexts
	return r
}

// The input file
func (r ApiExcelConversionConvertToOenormRequest) ExcelFile(excelFile *os.File) ApiExcelConversionConvertToOenormRequest {
	r.excelFile = excelFile
	return r
}

func (r ApiExcelConversionConvertToOenormRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ExcelConversionConvertToOenormExecute(r)
}

/*
ExcelConversionConvertToOenorm Converts Excel files to Oenorm files.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExcelConversionConvertToOenormRequest
*/
func (a *ExcelConversionApiService) ExcelConversionConvertToOenorm(ctx context.Context) ApiExcelConversionConvertToOenormRequest {
	return ApiExcelConversionConvertToOenormRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return *os.File
func (a *ExcelConversionApiService) ExcelConversionConvertToOenormExecute(r ApiExcelConversionConvertToOenormRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExcelConversionApiService.ExcelConversionConvertToOenorm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conversion/excel/oenorm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.readNewElements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ReadNewElements", r.readNewElements, "")
	}
	if r.rebuildItemNumberSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "RebuildItemNumberSchema", r.rebuildItemNumberSchema, "")
	}
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
	var excelFileLocalVarFormFileName string
	var excelFileLocalVarFileName string
	var excelFileLocalVarFileBytes []byte

	excelFileLocalVarFormFileName = "excelFile"

	excelFileLocalVarFile := r.excelFile

	if excelFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(excelFileLocalVarFile)

		excelFileLocalVarFileBytes = fbs
		excelFileLocalVarFileName = excelFileLocalVarFile.Name()
		excelFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: excelFileLocalVarFileBytes, fileName: excelFileLocalVarFileName, formFileName: excelFileLocalVarFormFileName})
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
