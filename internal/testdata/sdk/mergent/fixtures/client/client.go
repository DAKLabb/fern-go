// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	fixtures "github.com/fern-api/fern-go/internal/testdata/sdk/mergent/fixtures"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/mergent/fixtures/core"
	io "io"
	http "net/http"
)

type Client interface {
	GetTasks(ctx context.Context) ([]*fixtures.Task, error)
	PostTasks(ctx context.Context, request *fixtures.TaskNew) (*fixtures.Task, error)
	GetTasksTaskId(ctx context.Context, taskId fixtures.Id) (*fixtures.Task, error)
	PatchTasksTaskId(ctx context.Context, taskId fixtures.Id, request *fixtures.Task) (*fixtures.Task, error)
	DeleteTasksTaskId(ctx context.Context, taskId fixtures.Id) error
	PostTasksTaskIdRun(ctx context.Context, taskId fixtures.Id) (*fixtures.Task, error)
	PostTasksBatchCreate(ctx context.Context, request []*fixtures.TaskNew) ([]*fixtures.Task, error)
	PostTasksBatchDelete(ctx context.Context, request []fixtures.Id) error
	GetSchedules(ctx context.Context) ([]*fixtures.Schedule, error)
	PostSchedules(ctx context.Context, request *fixtures.ScheduleNew) (*fixtures.Schedule, error)
	GetSchedulesScheduleId(ctx context.Context, scheduleId fixtures.Id) (*fixtures.Schedule, error)
	PatchSchedulesScheduleId(ctx context.Context, scheduleId fixtures.Id, request *fixtures.Schedule) (*fixtures.Schedule, error)
	DeleteSchedulesScheduleId(ctx context.Context, scheduleId fixtures.Id) error
	GetSchedulesScheduleIdTasks(ctx context.Context, scheduleId fixtures.Id) ([]*fixtures.Task, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (c *client) GetTasks(ctx context.Context) ([]*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "tasks"

	var response []*fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) PostTasks(ctx context.Context, request *fixtures.TaskNew) (*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "tasks"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 409:
			value := new(fixtures.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(fixtures.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Task ID
func (c *client) GetTasksTaskId(ctx context.Context, taskId fixtures.Id) (*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"tasks/%v", taskId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Task ID
func (c *client) PatchTasksTaskId(ctx context.Context, taskId fixtures.Id, request *fixtures.Task) (*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"tasks/%v", taskId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 409:
			value := new(fixtures.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(fixtures.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Task ID
func (c *client) DeleteTasksTaskId(ctx context.Context, taskId fixtures.Id) error {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"tasks/%v", taskId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		nil,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

// Reschedules a queued Task to be run immediately.
//
// Task ID
func (c *client) PostTasksTaskIdRun(ctx context.Context, taskId fixtures.Id) (*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"tasks/%v/run", taskId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 409:
			value := new(fixtures.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// A maximum of 100 Tasks are accepted per request.
// This operation is atomic: it will succeed for all Tasks or fail for all
// Tasks; there is no partial success.
// This endpoint is in beta and may change at any time without notice.
func (c *client) PostTasksBatchCreate(ctx context.Context, request []*fixtures.TaskNew) ([]*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "tasks/batch-create"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 409:
			value := new(fixtures.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 413:
			value := new(fixtures.ContentTooLargeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(fixtures.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response []*fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// A maximum of 100 Task IDs are accepted per request.
// This operation is atomic: it will succeed for all Tasks or fail for all
// Tasks; there is no partial success.
// This endpoint is in beta and may change at any time without notice.
func (c *client) PostTasksBatchDelete(ctx context.Context, request []fixtures.Id) error {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "tasks/batch-delete"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 413:
			value := new(fixtures.ContentTooLargeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		nil,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

func (c *client) GetSchedules(ctx context.Context) ([]*fixtures.Schedule, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "schedules"

	var response []*fixtures.Schedule
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) PostSchedules(ctx context.Context, request *fixtures.ScheduleNew) (*fixtures.Schedule, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "schedules"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 422:
			value := new(fixtures.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Schedule
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Schedule ID
func (c *client) GetSchedulesScheduleId(ctx context.Context, scheduleId fixtures.Id) (*fixtures.Schedule, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"schedules/%v", scheduleId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Schedule
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Schedule ID
func (c *client) PatchSchedulesScheduleId(ctx context.Context, scheduleId fixtures.Id, request *fixtures.Schedule) (*fixtures.Schedule, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"schedules/%v", scheduleId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(fixtures.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Schedule
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPatch,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Schedule ID
func (c *client) DeleteSchedulesScheduleId(ctx context.Context, scheduleId fixtures.Id) error {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"schedules/%v", scheduleId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		nil,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return err
	}
	return nil
}

// Schedule ID
func (c *client) GetSchedulesScheduleIdTasks(ctx context.Context, scheduleId fixtures.Id) ([]*fixtures.Task, error) {
	baseURL := "https://api.mergent.co/v2"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"schedules/%v/tasks", scheduleId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(fixtures.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response []*fixtures.Task
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}