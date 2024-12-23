// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package gen

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Checks the health of API
	// (GET /healthcheck)
	Healthcheck(c *gin.Context)
	// Delete credit card
	// (DELETE /v1/payments/cards)
	DeleteCreditCard(c *gin.Context)
	// Get credit card list
	// (GET /v1/payments/cards)
	GetCreditCards(c *gin.Context)
	// Create new credit card
	// (POST /v1/payments/cards)
	CreateCreditCard(c *gin.Context)
	// Create new reservation
	// (POST /v1/payments/reservations)
	CreateReservation(c *gin.Context)
	// Get list of products
	// (GET /v1/products)
	GetProducts(c *gin.Context, params GetProductsParams)
	// Get product details
	// (GET /v1/products/{productID})
	GetProductByID(c *gin.Context, productID int64)
	// Create new user
	// (POST /v1/users)
	CreateUser(c *gin.Context)
	// Get user information about myself
	// (GET /v1/users/me)
	GetMe(c *gin.Context)
	// Create user profile
	// (POST /v1/users/profiles)
	CreateProfile(c *gin.Context)
	// Get profile information about myself
	// (GET /v1/users/profiles/me)
	GetProfileMe(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// Healthcheck operation middleware
func (siw *ServerInterfaceWrapper) Healthcheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Healthcheck(c)
}

// DeleteCreditCard operation middleware
func (siw *ServerInterfaceWrapper) DeleteCreditCard(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteCreditCard(c)
}

// GetCreditCards operation middleware
func (siw *ServerInterfaceWrapper) GetCreditCards(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCreditCards(c)
}

// CreateCreditCard operation middleware
func (siw *ServerInterfaceWrapper) CreateCreditCard(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateCreditCard(c)
}

// CreateReservation operation middleware
func (siw *ServerInterfaceWrapper) CreateReservation(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateReservation(c)
}

// GetProducts operation middleware
func (siw *ServerInterfaceWrapper) GetProducts(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProductsParams

	// ------------- Optional query parameter "cursor" -------------

	err = runtime.BindQueryParameter("form", true, false, "cursor", c.Request.URL.Query(), &params.Cursor)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cursor: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProducts(c, params)
}

// GetProductByID operation middleware
func (siw *ServerInterfaceWrapper) GetProductByID(c *gin.Context) {

	var err error

	// ------------- Path parameter "productID" -------------
	var productID int64

	err = runtime.BindStyledParameter("simple", false, "productID", c.Param("productID"), &productID)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter productID: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductByID(c, productID)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateUser(c)
}

// GetMe operation middleware
func (siw *ServerInterfaceWrapper) GetMe(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetMe(c)
}

// CreateProfile operation middleware
func (siw *ServerInterfaceWrapper) CreateProfile(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateProfile(c)
}

// GetProfileMe operation middleware
func (siw *ServerInterfaceWrapper) GetProfileMe(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProfileMe(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/healthcheck", wrapper.Healthcheck)
	router.DELETE(options.BaseURL+"/v1/payments/cards", wrapper.DeleteCreditCard)
	router.GET(options.BaseURL+"/v1/payments/cards", wrapper.GetCreditCards)
	router.POST(options.BaseURL+"/v1/payments/cards", wrapper.CreateCreditCard)
	router.POST(options.BaseURL+"/v1/payments/reservations", wrapper.CreateReservation)
	router.GET(options.BaseURL+"/v1/products", wrapper.GetProducts)
	router.GET(options.BaseURL+"/v1/products/:productID", wrapper.GetProductByID)
	router.POST(options.BaseURL+"/v1/users", wrapper.CreateUser)
	router.GET(options.BaseURL+"/v1/users/me", wrapper.GetMe)
	router.POST(options.BaseURL+"/v1/users/profiles", wrapper.CreateProfile)
	router.GET(options.BaseURL+"/v1/users/profiles/me", wrapper.GetProfileMe)
}

type AlreadyExistsResponse struct {
}

type BadRequestResponse struct {
}

type ForbiddenResponse struct {
}

type InternalServerErrorResponse struct {
}

type NotFoundResponse struct {
}

type UnauthorizedResponse struct {
}

type HealthcheckRequestObject struct {
}

type HealthcheckResponseObject interface {
	VisitHealthcheckResponse(w http.ResponseWriter) error
}

type Healthcheck200JSONResponse HealthCheck

func (response Healthcheck200JSONResponse) VisitHealthcheckResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteCreditCardRequestObject struct {
}

type DeleteCreditCardResponseObject interface {
	VisitDeleteCreditCardResponse(w http.ResponseWriter) error
}

type DeleteCreditCard204Response struct {
}

func (response DeleteCreditCard204Response) VisitDeleteCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteCreditCard400Response = BadRequestResponse

func (response DeleteCreditCard400Response) VisitDeleteCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type DeleteCreditCard401Response = UnauthorizedResponse

func (response DeleteCreditCard401Response) VisitDeleteCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type DeleteCreditCard404Response = NotFoundResponse

func (response DeleteCreditCard404Response) VisitDeleteCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type DeleteCreditCard500Response = InternalServerErrorResponse

func (response DeleteCreditCard500Response) VisitDeleteCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetCreditCardsRequestObject struct {
}

type GetCreditCardsResponseObject interface {
	VisitGetCreditCardsResponse(w http.ResponseWriter) error
}

type GetCreditCards200JSONResponse GetCreditCards

func (response GetCreditCards200JSONResponse) VisitGetCreditCardsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetCreditCards400Response = BadRequestResponse

func (response GetCreditCards400Response) VisitGetCreditCardsResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetCreditCards401Response = UnauthorizedResponse

func (response GetCreditCards401Response) VisitGetCreditCardsResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type GetCreditCards500Response = InternalServerErrorResponse

func (response GetCreditCards500Response) VisitGetCreditCardsResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type CreateCreditCardRequestObject struct {
	Body *CreateCreditCardJSONRequestBody
}

type CreateCreditCardResponseObject interface {
	VisitCreateCreditCardResponse(w http.ResponseWriter) error
}

type CreateCreditCard204Response struct {
}

func (response CreateCreditCard204Response) VisitCreateCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type CreateCreditCard400Response = BadRequestResponse

func (response CreateCreditCard400Response) VisitCreateCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreateCreditCard401Response = UnauthorizedResponse

func (response CreateCreditCard401Response) VisitCreateCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type CreateCreditCard500Response = InternalServerErrorResponse

func (response CreateCreditCard500Response) VisitCreateCreditCardResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type CreateReservationRequestObject struct {
	Body *CreateReservationJSONRequestBody
}

type CreateReservationResponseObject interface {
	VisitCreateReservationResponse(w http.ResponseWriter) error
}

type CreateReservation201JSONResponse ReservationResponse

func (response CreateReservation201JSONResponse) VisitCreateReservationResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateReservation400Response = BadRequestResponse

func (response CreateReservation400Response) VisitCreateReservationResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreateReservation401Response = UnauthorizedResponse

func (response CreateReservation401Response) VisitCreateReservationResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type CreateReservation500Response = InternalServerErrorResponse

func (response CreateReservation500Response) VisitCreateReservationResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetProductsRequestObject struct {
	Params GetProductsParams
}

type GetProductsResponseObject interface {
	VisitGetProductsResponse(w http.ResponseWriter) error
}

type GetProducts200JSONResponse Products

func (response GetProducts200JSONResponse) VisitGetProductsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProductByIDRequestObject struct {
	ProductID int64 `json:"productID"`
}

type GetProductByIDResponseObject interface {
	VisitGetProductByIDResponse(w http.ResponseWriter) error
}

type GetProductByID200JSONResponse Product

func (response GetProductByID200JSONResponse) VisitGetProductByIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProductByID400Response = BadRequestResponse

func (response GetProductByID400Response) VisitGetProductByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetProductByID401Response = UnauthorizedResponse

func (response GetProductByID401Response) VisitGetProductByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type GetProductByID404Response = NotFoundResponse

func (response GetProductByID404Response) VisitGetProductByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetProductByID500Response = InternalServerErrorResponse

func (response GetProductByID500Response) VisitGetProductByIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type CreateUserRequestObject struct {
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser201JSONResponse CreateUserResponse

func (response CreateUser201JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser400Response = BadRequestResponse

func (response CreateUser400Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreateUser401Response = UnauthorizedResponse

func (response CreateUser401Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type CreateUser409Response = AlreadyExistsResponse

func (response CreateUser409Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type CreateUser500Response = InternalServerErrorResponse

func (response CreateUser500Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetMeRequestObject struct {
}

type GetMeResponseObject interface {
	VisitGetMeResponse(w http.ResponseWriter) error
}

type GetMe200JSONResponse Me

func (response GetMe200JSONResponse) VisitGetMeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetMe401Response = UnauthorizedResponse

func (response GetMe401Response) VisitGetMeResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type GetMe404Response = NotFoundResponse

func (response GetMe404Response) VisitGetMeResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetMe500Response = InternalServerErrorResponse

func (response GetMe500Response) VisitGetMeResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type CreateProfileRequestObject struct {
	Body *CreateProfileJSONRequestBody
}

type CreateProfileResponseObject interface {
	VisitCreateProfileResponse(w http.ResponseWriter) error
}

type CreateProfile201JSONResponse Profile

func (response CreateProfile201JSONResponse) VisitCreateProfileResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateProfile400Response = BadRequestResponse

func (response CreateProfile400Response) VisitCreateProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreateProfile401Response = UnauthorizedResponse

func (response CreateProfile401Response) VisitCreateProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type CreateProfile403Response = ForbiddenResponse

func (response CreateProfile403Response) VisitCreateProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type CreateProfile409Response = AlreadyExistsResponse

func (response CreateProfile409Response) VisitCreateProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type CreateProfile500Response = InternalServerErrorResponse

func (response CreateProfile500Response) VisitCreateProfileResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type GetProfileMeRequestObject struct {
}

type GetProfileMeResponseObject interface {
	VisitGetProfileMeResponse(w http.ResponseWriter) error
}

type GetProfileMe200JSONResponse Profile

func (response GetProfileMe200JSONResponse) VisitGetProfileMeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProfileMe401Response = UnauthorizedResponse

func (response GetProfileMe401Response) VisitGetProfileMeResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type GetProfileMe404Response = NotFoundResponse

func (response GetProfileMe404Response) VisitGetProfileMeResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetProfileMe500Response = InternalServerErrorResponse

func (response GetProfileMe500Response) VisitGetProfileMeResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Checks the health of API
	// (GET /healthcheck)
	Healthcheck(ctx *gin.Context, request HealthcheckRequestObject) (HealthcheckResponseObject, error)
	// Delete credit card
	// (DELETE /v1/payments/cards)
	DeleteCreditCard(ctx *gin.Context, request DeleteCreditCardRequestObject) (DeleteCreditCardResponseObject, error)
	// Get credit card list
	// (GET /v1/payments/cards)
	GetCreditCards(ctx *gin.Context, request GetCreditCardsRequestObject) (GetCreditCardsResponseObject, error)
	// Create new credit card
	// (POST /v1/payments/cards)
	CreateCreditCard(ctx *gin.Context, request CreateCreditCardRequestObject) (CreateCreditCardResponseObject, error)
	// Create new reservation
	// (POST /v1/payments/reservations)
	CreateReservation(ctx *gin.Context, request CreateReservationRequestObject) (CreateReservationResponseObject, error)
	// Get list of products
	// (GET /v1/products)
	GetProducts(ctx *gin.Context, request GetProductsRequestObject) (GetProductsResponseObject, error)
	// Get product details
	// (GET /v1/products/{productID})
	GetProductByID(ctx *gin.Context, request GetProductByIDRequestObject) (GetProductByIDResponseObject, error)
	// Create new user
	// (POST /v1/users)
	CreateUser(ctx *gin.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)
	// Get user information about myself
	// (GET /v1/users/me)
	GetMe(ctx *gin.Context, request GetMeRequestObject) (GetMeResponseObject, error)
	// Create user profile
	// (POST /v1/users/profiles)
	CreateProfile(ctx *gin.Context, request CreateProfileRequestObject) (CreateProfileResponseObject, error)
	// Get profile information about myself
	// (GET /v1/users/profiles/me)
	GetProfileMe(ctx *gin.Context, request GetProfileMeRequestObject) (GetProfileMeResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// Healthcheck operation middleware
func (sh *strictHandler) Healthcheck(ctx *gin.Context) {
	var request HealthcheckRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Healthcheck(ctx, request.(HealthcheckRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Healthcheck")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(HealthcheckResponseObject); ok {
		if err := validResponse.VisitHealthcheckResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteCreditCard operation middleware
func (sh *strictHandler) DeleteCreditCard(ctx *gin.Context) {
	var request DeleteCreditCardRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteCreditCard(ctx, request.(DeleteCreditCardRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteCreditCard")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(DeleteCreditCardResponseObject); ok {
		if err := validResponse.VisitDeleteCreditCardResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetCreditCards operation middleware
func (sh *strictHandler) GetCreditCards(ctx *gin.Context) {
	var request GetCreditCardsRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetCreditCards(ctx, request.(GetCreditCardsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCreditCards")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetCreditCardsResponseObject); ok {
		if err := validResponse.VisitGetCreditCardsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateCreditCard operation middleware
func (sh *strictHandler) CreateCreditCard(ctx *gin.Context) {
	var request CreateCreditCardRequestObject

	var body CreateCreditCardJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateCreditCard(ctx, request.(CreateCreditCardRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateCreditCard")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreateCreditCardResponseObject); ok {
		if err := validResponse.VisitCreateCreditCardResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateReservation operation middleware
func (sh *strictHandler) CreateReservation(ctx *gin.Context) {
	var request CreateReservationRequestObject

	var body CreateReservationJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateReservation(ctx, request.(CreateReservationRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateReservation")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreateReservationResponseObject); ok {
		if err := validResponse.VisitCreateReservationResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProducts operation middleware
func (sh *strictHandler) GetProducts(ctx *gin.Context, params GetProductsParams) {
	var request GetProductsRequestObject

	request.Params = params

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProducts(ctx, request.(GetProductsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProducts")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetProductsResponseObject); ok {
		if err := validResponse.VisitGetProductsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProductByID operation middleware
func (sh *strictHandler) GetProductByID(ctx *gin.Context, productID int64) {
	var request GetProductByIDRequestObject

	request.ProductID = productID

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProductByID(ctx, request.(GetProductByIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProductByID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetProductByIDResponseObject); ok {
		if err := validResponse.VisitGetProductByIDResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(ctx *gin.Context) {
	var request CreateUserRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx, request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		if err := validResponse.VisitCreateUserResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetMe operation middleware
func (sh *strictHandler) GetMe(ctx *gin.Context) {
	var request GetMeRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMe(ctx, request.(GetMeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMe")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetMeResponseObject); ok {
		if err := validResponse.VisitGetMeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateProfile operation middleware
func (sh *strictHandler) CreateProfile(ctx *gin.Context) {
	var request CreateProfileRequestObject

	var body CreateProfileJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateProfile(ctx, request.(CreateProfileRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateProfile")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreateProfileResponseObject); ok {
		if err := validResponse.VisitCreateProfileResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetProfileMe operation middleware
func (sh *strictHandler) GetProfileMe(ctx *gin.Context) {
	var request GetProfileMeRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProfileMe(ctx, request.(GetProfileMeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProfileMe")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetProfileMeResponseObject); ok {
		if err := validResponse.VisitGetProfileMeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
