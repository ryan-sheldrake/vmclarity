// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	. "github.com/openclarity/vmclarity/api/models"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all scan configs.
	// (GET /scanConfigs)
	GetScanConfigs(ctx echo.Context, params GetScanConfigsParams) error
	// Create a scan config
	// (POST /scanConfigs)
	PostScanConfigs(ctx echo.Context) error
	// Delete a scan config.
	// (DELETE /scanConfigs/{scanConfigID})
	DeleteScanConfigsScanConfigID(ctx echo.Context, scanConfigID ScanConfigID) error
	// Get the details for a scan config.
	// (GET /scanConfigs/{scanConfigID})
	GetScanConfigsScanConfigID(ctx echo.Context, scanConfigID ScanConfigID) error
	// Patch a scan config.
	// (PATCH /scanConfigs/{scanConfigID})
	PatchScanConfigsScanConfigID(ctx echo.Context, scanConfigID ScanConfigID) error
	// Update a scan config.
	// (PUT /scanConfigs/{scanConfigID})
	PutScanConfigsScanConfigID(ctx echo.Context, scanConfigID ScanConfigID) error
	// Get scan results according to the given filters
	// (GET /scanResults)
	GetScanResults(ctx echo.Context, params GetScanResultsParams) error
	// Create a scan result for a specific target for a specific scan
	// (POST /scanResults)
	PostScanResults(ctx echo.Context) error
	// Get a scan result.
	// (GET /scanResults/{scanResultID})
	GetScanResultsScanResultID(ctx echo.Context, scanResultID ScanResultID, params GetScanResultsScanResultIDParams) error
	// Patch a scan result
	// (PATCH /scanResults/{scanResultID})
	PatchScanResultsScanResultID(ctx echo.Context, scanResultID ScanResultID) error
	// Update a scan result.
	// (PUT /scanResults/{scanResultID})
	PutScanResultsScanResultID(ctx echo.Context, scanResultID ScanResultID) error
	// Get all scans. Each scan contaians details about a multi-target scheduled scan.
	// (GET /scans)
	GetScans(ctx echo.Context, params GetScansParams) error
	// Create a multi-target scheduled scan
	// (POST /scans)
	PostScans(ctx echo.Context) error
	// Delete a scan.
	// (DELETE /scans/{scanID})
	DeleteScansScanID(ctx echo.Context, scanID ScanID) error
	// Get the details for a given multi-target scheduled scan.
	// (GET /scans/{scanID})
	GetScansScanID(ctx echo.Context, scanID ScanID) error
	// Patch a scan.
	// (PATCH /scans/{scanID})
	PatchScansScanID(ctx echo.Context, scanID ScanID) error
	// Update a scan.
	// (PUT /scans/{scanID})
	PutScansScanID(ctx echo.Context, scanID ScanID) error
	// Get targets
	// (GET /targets)
	GetTargets(ctx echo.Context, params GetTargetsParams) error
	// Create target
	// (POST /targets)
	PostTargets(ctx echo.Context) error
	// Delete target.
	// (DELETE /targets/{targetID})
	DeleteTargetsTargetID(ctx echo.Context, targetID TargetID) error
	// Get target.
	// (GET /targets/{targetID})
	GetTargetsTargetID(ctx echo.Context, targetID TargetID) error
	// Update target.
	// (PUT /targets/{targetID})
	PutTargetsTargetID(ctx echo.Context, targetID TargetID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetScanConfigs converts echo context to params.
func (w *ServerInterfaceWrapper) GetScanConfigs(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScanConfigsParams
	// ------------- Optional query parameter "$filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "$filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter $filter: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageSize: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetScanConfigs(ctx, params)
	return err
}

// PostScanConfigs converts echo context to params.
func (w *ServerInterfaceWrapper) PostScanConfigs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostScanConfigs(ctx)
	return err
}

// DeleteScanConfigsScanConfigID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteScanConfigsScanConfigID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanConfigID" -------------
	var scanConfigID ScanConfigID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanConfigID", runtime.ParamLocationPath, ctx.Param("scanConfigID"), &scanConfigID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanConfigID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteScanConfigsScanConfigID(ctx, scanConfigID)
	return err
}

// GetScanConfigsScanConfigID converts echo context to params.
func (w *ServerInterfaceWrapper) GetScanConfigsScanConfigID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanConfigID" -------------
	var scanConfigID ScanConfigID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanConfigID", runtime.ParamLocationPath, ctx.Param("scanConfigID"), &scanConfigID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanConfigID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetScanConfigsScanConfigID(ctx, scanConfigID)
	return err
}

// PatchScanConfigsScanConfigID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchScanConfigsScanConfigID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanConfigID" -------------
	var scanConfigID ScanConfigID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanConfigID", runtime.ParamLocationPath, ctx.Param("scanConfigID"), &scanConfigID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanConfigID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchScanConfigsScanConfigID(ctx, scanConfigID)
	return err
}

// PutScanConfigsScanConfigID converts echo context to params.
func (w *ServerInterfaceWrapper) PutScanConfigsScanConfigID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanConfigID" -------------
	var scanConfigID ScanConfigID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanConfigID", runtime.ParamLocationPath, ctx.Param("scanConfigID"), &scanConfigID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanConfigID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutScanConfigsScanConfigID(ctx, scanConfigID)
	return err
}

// GetScanResults converts echo context to params.
func (w *ServerInterfaceWrapper) GetScanResults(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScanResultsParams
	// ------------- Optional query parameter "$filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "$filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter $filter: %s", err))
	}

	// ------------- Optional query parameter "$select" -------------

	err = runtime.BindQueryParameter("form", true, false, "$select", ctx.QueryParams(), &params.Select)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter $select: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageSize: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetScanResults(ctx, params)
	return err
}

// PostScanResults converts echo context to params.
func (w *ServerInterfaceWrapper) PostScanResults(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostScanResults(ctx)
	return err
}

// GetScanResultsScanResultID converts echo context to params.
func (w *ServerInterfaceWrapper) GetScanResultsScanResultID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanResultID" -------------
	var scanResultID ScanResultID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanResultID", runtime.ParamLocationPath, ctx.Param("scanResultID"), &scanResultID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanResultID: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScanResultsScanResultIDParams
	// ------------- Optional query parameter "$select" -------------

	err = runtime.BindQueryParameter("form", true, false, "$select", ctx.QueryParams(), &params.Select)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter $select: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetScanResultsScanResultID(ctx, scanResultID, params)
	return err
}

// PatchScanResultsScanResultID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchScanResultsScanResultID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanResultID" -------------
	var scanResultID ScanResultID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanResultID", runtime.ParamLocationPath, ctx.Param("scanResultID"), &scanResultID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanResultID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchScanResultsScanResultID(ctx, scanResultID)
	return err
}

// PutScanResultsScanResultID converts echo context to params.
func (w *ServerInterfaceWrapper) PutScanResultsScanResultID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanResultID" -------------
	var scanResultID ScanResultID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanResultID", runtime.ParamLocationPath, ctx.Param("scanResultID"), &scanResultID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanResultID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutScanResultsScanResultID(ctx, scanResultID)
	return err
}

// GetScans converts echo context to params.
func (w *ServerInterfaceWrapper) GetScans(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScansParams
	// ------------- Optional query parameter "$filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "$filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter $filter: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageSize: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetScans(ctx, params)
	return err
}

// PostScans converts echo context to params.
func (w *ServerInterfaceWrapper) PostScans(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostScans(ctx)
	return err
}

// DeleteScansScanID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteScansScanID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanID" -------------
	var scanID ScanID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanID", runtime.ParamLocationPath, ctx.Param("scanID"), &scanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteScansScanID(ctx, scanID)
	return err
}

// GetScansScanID converts echo context to params.
func (w *ServerInterfaceWrapper) GetScansScanID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanID" -------------
	var scanID ScanID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanID", runtime.ParamLocationPath, ctx.Param("scanID"), &scanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetScansScanID(ctx, scanID)
	return err
}

// PatchScansScanID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchScansScanID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanID" -------------
	var scanID ScanID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanID", runtime.ParamLocationPath, ctx.Param("scanID"), &scanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchScansScanID(ctx, scanID)
	return err
}

// PutScansScanID converts echo context to params.
func (w *ServerInterfaceWrapper) PutScansScanID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "scanID" -------------
	var scanID ScanID

	err = runtime.BindStyledParameterWithLocation("simple", false, "scanID", runtime.ParamLocationPath, ctx.Param("scanID"), &scanID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter scanID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutScansScanID(ctx, scanID)
	return err
}

// GetTargets converts echo context to params.
func (w *ServerInterfaceWrapper) GetTargets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTargetsParams
	// ------------- Optional query parameter "$filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "$filter", ctx.QueryParams(), &params.Filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter $filter: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageSize: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTargets(ctx, params)
	return err
}

// PostTargets converts echo context to params.
func (w *ServerInterfaceWrapper) PostTargets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostTargets(ctx)
	return err
}

// DeleteTargetsTargetID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTargetsTargetID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "targetID" -------------
	var targetID TargetID

	err = runtime.BindStyledParameterWithLocation("simple", false, "targetID", runtime.ParamLocationPath, ctx.Param("targetID"), &targetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter targetID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTargetsTargetID(ctx, targetID)
	return err
}

// GetTargetsTargetID converts echo context to params.
func (w *ServerInterfaceWrapper) GetTargetsTargetID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "targetID" -------------
	var targetID TargetID

	err = runtime.BindStyledParameterWithLocation("simple", false, "targetID", runtime.ParamLocationPath, ctx.Param("targetID"), &targetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter targetID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTargetsTargetID(ctx, targetID)
	return err
}

// PutTargetsTargetID converts echo context to params.
func (w *ServerInterfaceWrapper) PutTargetsTargetID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "targetID" -------------
	var targetID TargetID

	err = runtime.BindStyledParameterWithLocation("simple", false, "targetID", runtime.ParamLocationPath, ctx.Param("targetID"), &targetID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter targetID: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutTargetsTargetID(ctx, targetID)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/scanConfigs", wrapper.GetScanConfigs)
	router.POST(baseURL+"/scanConfigs", wrapper.PostScanConfigs)
	router.DELETE(baseURL+"/scanConfigs/:scanConfigID", wrapper.DeleteScanConfigsScanConfigID)
	router.GET(baseURL+"/scanConfigs/:scanConfigID", wrapper.GetScanConfigsScanConfigID)
	router.PATCH(baseURL+"/scanConfigs/:scanConfigID", wrapper.PatchScanConfigsScanConfigID)
	router.PUT(baseURL+"/scanConfigs/:scanConfigID", wrapper.PutScanConfigsScanConfigID)
	router.GET(baseURL+"/scanResults", wrapper.GetScanResults)
	router.POST(baseURL+"/scanResults", wrapper.PostScanResults)
	router.GET(baseURL+"/scanResults/:scanResultID", wrapper.GetScanResultsScanResultID)
	router.PATCH(baseURL+"/scanResults/:scanResultID", wrapper.PatchScanResultsScanResultID)
	router.PUT(baseURL+"/scanResults/:scanResultID", wrapper.PutScanResultsScanResultID)
	router.GET(baseURL+"/scans", wrapper.GetScans)
	router.POST(baseURL+"/scans", wrapper.PostScans)
	router.DELETE(baseURL+"/scans/:scanID", wrapper.DeleteScansScanID)
	router.GET(baseURL+"/scans/:scanID", wrapper.GetScansScanID)
	router.PATCH(baseURL+"/scans/:scanID", wrapper.PatchScansScanID)
	router.PUT(baseURL+"/scans/:scanID", wrapper.PutScansScanID)
	router.GET(baseURL+"/targets", wrapper.GetTargets)
	router.POST(baseURL+"/targets", wrapper.PostTargets)
	router.DELETE(baseURL+"/targets/:targetID", wrapper.DeleteTargetsTargetID)
	router.GET(baseURL+"/targets/:targetID", wrapper.GetTargetsTargetID)
	router.PUT(baseURL+"/targets/:targetID", wrapper.PutTargetsTargetID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
