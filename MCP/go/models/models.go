package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// CreateMessageResponse represents the CreateMessageResponse schema from the OpenAPI specification
type CreateMessageResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// DeleteMessageResponse represents the DeleteMessageResponse schema from the OpenAPI specification
type DeleteMessageResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Id string `json:"id"` // Unique identifier for the custom field.
	Name string `json:"name,omitempty"` // Name of the custom field.
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// UpdateMessageResponse represents the UpdateMessageResponse schema from the OpenAPI specification
type UpdateMessageResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Id string `json:"id,omitempty"` // Unique identifier for the email address
	TypeField string `json:"type,omitempty"` // Email type
	Email string `json:"email"` // Email address
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// GetMessageResponse represents the GetMessageResponse schema from the OpenAPI specification
type GetMessageResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Message `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// GetMessagesResponse represents the GetMessagesResponse schema from the OpenAPI specification
type GetMessagesResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Message `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// Message represents the Message schema from the OpenAPI specification
type Message struct {
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Subject string `json:"subject,omitempty"`
	TypeField string `json:"type,omitempty"` // Set to sms for SMS messages and mms for MMS messages.
	Number_of_media_files int `json:"number_of_media_files,omitempty"` // The number of media files associated with the message.
	Messaging_service_id string `json:"messaging_service_id,omitempty"` // The ID of the Messaging Service used with the message. In case of Plivo this links to the Powerpack ID.
	Price map[string]interface{} `json:"price,omitempty"` // Price of the message.
	Body string `json:"body"` // The message text.
	Direction string `json:"direction,omitempty"` // The direction of the message.
	Scheduled_at string `json:"scheduled_at,omitempty"` // The scheduled date and time of the message.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Number_of_units int `json:"number_of_units,omitempty"` // The number of units that make up the complete message. Messages can be split up due to the constraints of the message size.
	Sent_at string `json:"sent_at,omitempty"` // The date and time that the message was sent
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	ErrorField map[string]interface{} `json:"error,omitempty"` // The error returned if your message status is failed or undelivered.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	From string `json:"from"` // The phone number that initiated the message.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	To string `json:"to"` // The phone number that received the message.
	Reference string `json:"reference,omitempty"` // A client reference.
	Status string `json:"status,omitempty"` // Status of the delivery of the message.
	Webhook_url string `json:"webhook_url,omitempty"` // Define a webhook to receive delivery notifications.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
}
