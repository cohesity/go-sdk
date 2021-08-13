/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// SearchIndexedObjectsWebSocketRequest The request parameters required for MCM Search Indexed Objects Requests, in websocket mode.
type SearchIndexedObjectsWebSocketRequest struct {
	// ID of the message, used to uniqely identify a message. The response will contain this messageId to indicate the request to which the response belongs to. Hint: Use current timestamp in micro seconds.
	MessageId NullableInt64 `json:"messageId"`
	Body HeliosSearchIndexedObjectsRequest `json:"body"`
}

// NewSearchIndexedObjectsWebSocketRequest instantiates a new SearchIndexedObjectsWebSocketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexedObjectsWebSocketRequest(messageId NullableInt64, body HeliosSearchIndexedObjectsRequest) *SearchIndexedObjectsWebSocketRequest {
	this := SearchIndexedObjectsWebSocketRequest{}
	this.MessageId = messageId
	this.Body = body
	return &this
}

// NewSearchIndexedObjectsWebSocketRequestWithDefaults instantiates a new SearchIndexedObjectsWebSocketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexedObjectsWebSocketRequestWithDefaults() *SearchIndexedObjectsWebSocketRequest {
	this := SearchIndexedObjectsWebSocketRequest{}
	return &this
}

// GetMessageId returns the MessageId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *SearchIndexedObjectsWebSocketRequest) GetMessageId() int64 {
	if o == nil || o.MessageId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.MessageId.Get()
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchIndexedObjectsWebSocketRequest) GetMessageIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MessageId.Get(), o.MessageId.IsSet()
}

// SetMessageId sets field value
func (o *SearchIndexedObjectsWebSocketRequest) SetMessageId(v int64) {
	o.MessageId.Set(&v)
}

// GetBody returns the Body field value
func (o *SearchIndexedObjectsWebSocketRequest) GetBody() HeliosSearchIndexedObjectsRequest {
	if o == nil {
		var ret HeliosSearchIndexedObjectsRequest
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsWebSocketRequest) GetBodyOk() (*HeliosSearchIndexedObjectsRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *SearchIndexedObjectsWebSocketRequest) SetBody(v HeliosSearchIndexedObjectsRequest) {
	o.Body = v
}

func (o SearchIndexedObjectsWebSocketRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["messageId"] = o.MessageId.Get()
	}
	if true {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableSearchIndexedObjectsWebSocketRequest struct {
	value *SearchIndexedObjectsWebSocketRequest
	isSet bool
}

func (v NullableSearchIndexedObjectsWebSocketRequest) Get() *SearchIndexedObjectsWebSocketRequest {
	return v.value
}

func (v *NullableSearchIndexedObjectsWebSocketRequest) Set(val *SearchIndexedObjectsWebSocketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIndexedObjectsWebSocketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIndexedObjectsWebSocketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIndexedObjectsWebSocketRequest(val *SearchIndexedObjectsWebSocketRequest) *NullableSearchIndexedObjectsWebSocketRequest {
	return &NullableSearchIndexedObjectsWebSocketRequest{value: val, isSet: true}
}

func (v NullableSearchIndexedObjectsWebSocketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIndexedObjectsWebSocketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


