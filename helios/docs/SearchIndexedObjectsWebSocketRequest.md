# SearchIndexedObjectsWebSocketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **NullableInt64** | ID of the message, used to uniqely identify a message. The response will contain this messageId to indicate the request to which the response belongs to. Hint: Use current timestamp in micro seconds. | 
**Body** | [**HeliosSearchIndexedObjectsRequest**](HeliosSearchIndexedObjectsRequest.md) |  | 

## Methods

### NewSearchIndexedObjectsWebSocketRequest

`func NewSearchIndexedObjectsWebSocketRequest(messageId NullableInt64, body HeliosSearchIndexedObjectsRequest, ) *SearchIndexedObjectsWebSocketRequest`

NewSearchIndexedObjectsWebSocketRequest instantiates a new SearchIndexedObjectsWebSocketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexedObjectsWebSocketRequestWithDefaults

`func NewSearchIndexedObjectsWebSocketRequestWithDefaults() *SearchIndexedObjectsWebSocketRequest`

NewSearchIndexedObjectsWebSocketRequestWithDefaults instantiates a new SearchIndexedObjectsWebSocketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *SearchIndexedObjectsWebSocketRequest) GetMessageId() int64`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SearchIndexedObjectsWebSocketRequest) GetMessageIdOk() (*int64, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SearchIndexedObjectsWebSocketRequest) SetMessageId(v int64)`

SetMessageId sets MessageId field to given value.


### SetMessageIdNil

`func (o *SearchIndexedObjectsWebSocketRequest) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *SearchIndexedObjectsWebSocketRequest) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetBody

`func (o *SearchIndexedObjectsWebSocketRequest) GetBody() HeliosSearchIndexedObjectsRequest`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SearchIndexedObjectsWebSocketRequest) GetBodyOk() (*HeliosSearchIndexedObjectsRequest, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SearchIndexedObjectsWebSocketRequest) SetBody(v HeliosSearchIndexedObjectsRequest)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


