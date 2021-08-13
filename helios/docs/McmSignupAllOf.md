# McmSignupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Specifies the id of the signup request. | [optional] 
**CreatedAtMsecs** | Pointer to **NullableInt64** | Specifies the time at which the signup request is created in Unix timestamp epoch in milliseconds. | [optional] 
**UpdatedAtMsecs** | Pointer to **NullableInt64** | Specifies the time at which the signup request is updated in Unix timestamp epoch in milliseconds. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the MCM signup request. | [optional] 
**AcceptTerms** | Pointer to **NullableBool** | Specifies whether the User agrees to Terms and conditions for using Cohesity MCM services which are specified at &lt;url&gt; | [optional] 

## Methods

### NewMcmSignupAllOf

`func NewMcmSignupAllOf() *McmSignupAllOf`

NewMcmSignupAllOf instantiates a new McmSignupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSignupAllOfWithDefaults

`func NewMcmSignupAllOfWithDefaults() *McmSignupAllOf`

NewMcmSignupAllOfWithDefaults instantiates a new McmSignupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McmSignupAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmSignupAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmSignupAllOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *McmSignupAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmSignupAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmSignupAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedAtMsecs

`func (o *McmSignupAllOf) GetCreatedAtMsecs() int64`

GetCreatedAtMsecs returns the CreatedAtMsecs field if non-nil, zero value otherwise.

### GetCreatedAtMsecsOk

`func (o *McmSignupAllOf) GetCreatedAtMsecsOk() (*int64, bool)`

GetCreatedAtMsecsOk returns a tuple with the CreatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtMsecs

`func (o *McmSignupAllOf) SetCreatedAtMsecs(v int64)`

SetCreatedAtMsecs sets CreatedAtMsecs field to given value.

### HasCreatedAtMsecs

`func (o *McmSignupAllOf) HasCreatedAtMsecs() bool`

HasCreatedAtMsecs returns a boolean if a field has been set.

### SetCreatedAtMsecsNil

`func (o *McmSignupAllOf) SetCreatedAtMsecsNil(b bool)`

 SetCreatedAtMsecsNil sets the value for CreatedAtMsecs to be an explicit nil

### UnsetCreatedAtMsecs
`func (o *McmSignupAllOf) UnsetCreatedAtMsecs()`

UnsetCreatedAtMsecs ensures that no value is present for CreatedAtMsecs, not even an explicit nil
### GetUpdatedAtMsecs

`func (o *McmSignupAllOf) GetUpdatedAtMsecs() int64`

GetUpdatedAtMsecs returns the UpdatedAtMsecs field if non-nil, zero value otherwise.

### GetUpdatedAtMsecsOk

`func (o *McmSignupAllOf) GetUpdatedAtMsecsOk() (*int64, bool)`

GetUpdatedAtMsecsOk returns a tuple with the UpdatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtMsecs

`func (o *McmSignupAllOf) SetUpdatedAtMsecs(v int64)`

SetUpdatedAtMsecs sets UpdatedAtMsecs field to given value.

### HasUpdatedAtMsecs

`func (o *McmSignupAllOf) HasUpdatedAtMsecs() bool`

HasUpdatedAtMsecs returns a boolean if a field has been set.

### SetUpdatedAtMsecsNil

`func (o *McmSignupAllOf) SetUpdatedAtMsecsNil(b bool)`

 SetUpdatedAtMsecsNil sets the value for UpdatedAtMsecs to be an explicit nil

### UnsetUpdatedAtMsecs
`func (o *McmSignupAllOf) UnsetUpdatedAtMsecs()`

UnsetUpdatedAtMsecs ensures that no value is present for UpdatedAtMsecs, not even an explicit nil
### GetStatus

`func (o *McmSignupAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McmSignupAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McmSignupAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *McmSignupAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *McmSignupAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *McmSignupAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAcceptTerms

`func (o *McmSignupAllOf) GetAcceptTerms() bool`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *McmSignupAllOf) GetAcceptTermsOk() (*bool, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *McmSignupAllOf) SetAcceptTerms(v bool)`

SetAcceptTerms sets AcceptTerms field to given value.

### HasAcceptTerms

`func (o *McmSignupAllOf) HasAcceptTerms() bool`

HasAcceptTerms returns a boolean if a field has been set.

### SetAcceptTermsNil

`func (o *McmSignupAllOf) SetAcceptTermsNil(b bool)`

 SetAcceptTermsNil sets the value for AcceptTerms to be an explicit nil

### UnsetAcceptTerms
`func (o *McmSignupAllOf) UnsetAcceptTerms()`

UnsetAcceptTerms ensures that no value is present for AcceptTerms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


