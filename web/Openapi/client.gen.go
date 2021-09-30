// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package Openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostUser request with any body
	PostUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostUser(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUsersUserId request
	GetUsersUserId(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchUsersUserId request with any body
	PatchUsersUserIdWithBody(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchUsersUserId(ctx context.Context, userId int, body PatchUsersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostUsersUserIdTodo request with any body
	PostUsersUserIdTodoWithBody(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostUsersUserIdTodo(ctx context.Context, userId string, body PostUsersUserIdTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchUsersUserIdTodoTodoId request with any body
	PatchUsersUserIdTodoTodoIdWithBody(ctx context.Context, userId string, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchUsersUserIdTodoTodoId(ctx context.Context, userId string, todoId string, body PatchUsersUserIdTodoTodoIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUsersUserIdTodos request
	GetUsersUserIdTodos(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUser(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUsersUserId(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersUserIdRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUsersUserIdWithBody(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUsersUserIdRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUsersUserId(ctx context.Context, userId int, body PatchUsersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUsersUserIdRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUsersUserIdTodoWithBody(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUsersUserIdTodoRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUsersUserIdTodo(ctx context.Context, userId string, body PostUsersUserIdTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUsersUserIdTodoRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUsersUserIdTodoTodoIdWithBody(ctx context.Context, userId string, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUsersUserIdTodoTodoIdRequestWithBody(c.Server, userId, todoId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUsersUserIdTodoTodoId(ctx context.Context, userId string, todoId string, body PatchUsersUserIdTodoTodoIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUsersUserIdTodoTodoIdRequest(c.Server, userId, todoId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUsersUserIdTodos(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersUserIdTodosRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostUserRequest calls the generic PostUser builder with application/json body
func NewPostUserRequest(server string, body PostUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostUserRequestWithBody(server, "application/json", bodyReader)
}

// NewPostUserRequestWithBody generates requests for PostUser with any type of body
func NewPostUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUsersUserIdRequest generates requests for GetUsersUserId
func NewGetUsersUserIdRequest(server string, userId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPatchUsersUserIdRequest calls the generic PatchUsersUserId builder with application/json body
func NewPatchUsersUserIdRequest(server string, userId int, body PatchUsersUserIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchUsersUserIdRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewPatchUsersUserIdRequestWithBody generates requests for PatchUsersUserId with any type of body
func NewPatchUsersUserIdRequestWithBody(server string, userId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPostUsersUserIdTodoRequest calls the generic PostUsersUserIdTodo builder with application/json body
func NewPostUsersUserIdTodoRequest(server string, userId string, body PostUsersUserIdTodoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostUsersUserIdTodoRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewPostUsersUserIdTodoRequestWithBody generates requests for PostUsersUserIdTodo with any type of body
func NewPostUsersUserIdTodoRequestWithBody(server string, userId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/todo", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPatchUsersUserIdTodoTodoIdRequest calls the generic PatchUsersUserIdTodoTodoId builder with application/json body
func NewPatchUsersUserIdTodoTodoIdRequest(server string, userId string, todoId string, body PatchUsersUserIdTodoTodoIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchUsersUserIdTodoTodoIdRequestWithBody(server, userId, todoId, "application/json", bodyReader)
}

// NewPatchUsersUserIdTodoTodoIdRequestWithBody generates requests for PatchUsersUserIdTodoTodoId with any type of body
func NewPatchUsersUserIdTodoTodoIdRequestWithBody(server string, userId string, todoId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "todoId", runtime.ParamLocationPath, todoId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/todo/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUsersUserIdTodosRequest generates requests for GetUsersUserIdTodos
func NewGetUsersUserIdTodosRequest(server string, userId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/todos", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostUser request with any body
	PostUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUserResponse, error)

	PostUserWithResponse(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUserResponse, error)

	// GetUsersUserId request
	GetUsersUserIdWithResponse(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*GetUsersUserIdResponse, error)

	// PatchUsersUserId request with any body
	PatchUsersUserIdWithBodyWithResponse(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchUsersUserIdResponse, error)

	PatchUsersUserIdWithResponse(ctx context.Context, userId int, body PatchUsersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchUsersUserIdResponse, error)

	// PostUsersUserIdTodo request with any body
	PostUsersUserIdTodoWithBodyWithResponse(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUsersUserIdTodoResponse, error)

	PostUsersUserIdTodoWithResponse(ctx context.Context, userId string, body PostUsersUserIdTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUsersUserIdTodoResponse, error)

	// PatchUsersUserIdTodoTodoId request with any body
	PatchUsersUserIdTodoTodoIdWithBodyWithResponse(ctx context.Context, userId string, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchUsersUserIdTodoTodoIdResponse, error)

	PatchUsersUserIdTodoTodoIdWithResponse(ctx context.Context, userId string, todoId string, body PatchUsersUserIdTodoTodoIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchUsersUserIdTodoTodoIdResponse, error)

	// GetUsersUserIdTodos request
	GetUsersUserIdTodosWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*GetUsersUserIdTodosResponse, error)
}

type PostUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r PostUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsersUserIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r GetUsersUserIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersUserIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchUsersUserIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r PatchUsersUserIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchUsersUserIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostUsersUserIdTodoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Todo
}

// Status returns HTTPResponse.Status
func (r PostUsersUserIdTodoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostUsersUserIdTodoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchUsersUserIdTodoTodoIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Todo
}

// Status returns HTTPResponse.Status
func (r PatchUsersUserIdTodoTodoIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchUsersUserIdTodoTodoIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsersUserIdTodosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Todo
}

// Status returns HTTPResponse.Status
func (r GetUsersUserIdTodosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersUserIdTodosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostUserWithBodyWithResponse request with arbitrary body returning *PostUserResponse
func (c *ClientWithResponses) PostUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUserResponse, error) {
	rsp, err := c.PostUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUserResponse(rsp)
}

func (c *ClientWithResponses) PostUserWithResponse(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUserResponse, error) {
	rsp, err := c.PostUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUserResponse(rsp)
}

// GetUsersUserIdWithResponse request returning *GetUsersUserIdResponse
func (c *ClientWithResponses) GetUsersUserIdWithResponse(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*GetUsersUserIdResponse, error) {
	rsp, err := c.GetUsersUserId(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersUserIdResponse(rsp)
}

// PatchUsersUserIdWithBodyWithResponse request with arbitrary body returning *PatchUsersUserIdResponse
func (c *ClientWithResponses) PatchUsersUserIdWithBodyWithResponse(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchUsersUserIdResponse, error) {
	rsp, err := c.PatchUsersUserIdWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUsersUserIdResponse(rsp)
}

func (c *ClientWithResponses) PatchUsersUserIdWithResponse(ctx context.Context, userId int, body PatchUsersUserIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchUsersUserIdResponse, error) {
	rsp, err := c.PatchUsersUserId(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUsersUserIdResponse(rsp)
}

// PostUsersUserIdTodoWithBodyWithResponse request with arbitrary body returning *PostUsersUserIdTodoResponse
func (c *ClientWithResponses) PostUsersUserIdTodoWithBodyWithResponse(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUsersUserIdTodoResponse, error) {
	rsp, err := c.PostUsersUserIdTodoWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUsersUserIdTodoResponse(rsp)
}

func (c *ClientWithResponses) PostUsersUserIdTodoWithResponse(ctx context.Context, userId string, body PostUsersUserIdTodoJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUsersUserIdTodoResponse, error) {
	rsp, err := c.PostUsersUserIdTodo(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUsersUserIdTodoResponse(rsp)
}

// PatchUsersUserIdTodoTodoIdWithBodyWithResponse request with arbitrary body returning *PatchUsersUserIdTodoTodoIdResponse
func (c *ClientWithResponses) PatchUsersUserIdTodoTodoIdWithBodyWithResponse(ctx context.Context, userId string, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchUsersUserIdTodoTodoIdResponse, error) {
	rsp, err := c.PatchUsersUserIdTodoTodoIdWithBody(ctx, userId, todoId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUsersUserIdTodoTodoIdResponse(rsp)
}

func (c *ClientWithResponses) PatchUsersUserIdTodoTodoIdWithResponse(ctx context.Context, userId string, todoId string, body PatchUsersUserIdTodoTodoIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchUsersUserIdTodoTodoIdResponse, error) {
	rsp, err := c.PatchUsersUserIdTodoTodoId(ctx, userId, todoId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUsersUserIdTodoTodoIdResponse(rsp)
}

// GetUsersUserIdTodosWithResponse request returning *GetUsersUserIdTodosResponse
func (c *ClientWithResponses) GetUsersUserIdTodosWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*GetUsersUserIdTodosResponse, error) {
	rsp, err := c.GetUsersUserIdTodos(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersUserIdTodosResponse(rsp)
}

// ParsePostUserResponse parses an HTTP response from a PostUserWithResponse call
func ParsePostUserResponse(rsp *http.Response) (*PostUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetUsersUserIdResponse parses an HTTP response from a GetUsersUserIdWithResponse call
func ParseGetUsersUserIdResponse(rsp *http.Response) (*GetUsersUserIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetUsersUserIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePatchUsersUserIdResponse parses an HTTP response from a PatchUsersUserIdWithResponse call
func ParsePatchUsersUserIdResponse(rsp *http.Response) (*PatchUsersUserIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PatchUsersUserIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostUsersUserIdTodoResponse parses an HTTP response from a PostUsersUserIdTodoWithResponse call
func ParsePostUsersUserIdTodoResponse(rsp *http.Response) (*PostUsersUserIdTodoResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostUsersUserIdTodoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Todo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePatchUsersUserIdTodoTodoIdResponse parses an HTTP response from a PatchUsersUserIdTodoTodoIdWithResponse call
func ParsePatchUsersUserIdTodoTodoIdResponse(rsp *http.Response) (*PatchUsersUserIdTodoTodoIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PatchUsersUserIdTodoTodoIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Todo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetUsersUserIdTodosResponse parses an HTTP response from a GetUsersUserIdTodosWithResponse call
func ParseGetUsersUserIdTodosResponse(rsp *http.Response) (*GetUsersUserIdTodosResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetUsersUserIdTodosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Todo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

