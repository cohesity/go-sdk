# McmSignup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Specifies the id of the signup request. | [optional] 
**CreatedAtMsecs** | Pointer to **NullableInt64** | Specifies the time at which the signup request is created in Unix timestamp epoch in milliseconds. | [optional] 
**UpdatedAtMsecs** | Pointer to **NullableInt64** | Specifies the time at which the signup request is updated in Unix timestamp epoch in milliseconds. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the MCM signup request. | [optional] 
**AcceptTerms** | Pointer to **NullableBool** | Specifies whether the User agrees to Terms and conditions for using Cohesity MCM services which are specified at &lt;url&gt; | [optional] 
**ContactInfo** | [**McmSignupUserContactInfo**](McmSignupUserContactInfo.md) |  | 
**Workloads** | Pointer to **[]string** | Specifies the workload requirements of the company to use Cohesity MCM. | [optional] 
**DmaasParams** | Pointer to [**DmaasSignupParams**](DmaasSignupParams.md) |  | [optional] 

## Methods

### NewMcmSignup

`func NewMcmSignup(contactInfo McmSignupUserContactInfo, ) *McmSignup`

NewMcmSignup instantiates a new McmSignup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSignupWithDefaults

`func NewMcmSignupWithDefaults() *McmSignup`

NewMcmSignupWithDefaults instantiates a new McmSignup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McmSignup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmSignup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmSignup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *McmSignup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmSignup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmSignup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedAtMsecs

`func (o *McmSignup) GetCreatedAtMsecs() int64`

GetCreatedAtMsecs returns the CreatedAtMsecs field if non-nil, zero value otherwise.

### GetCreatedAtMsecsOk

`func (o *McmSignup) GetCreatedAtMsecsOk() (*int64, bool)`

GetCreatedAtMsecsOk returns a tuple with the CreatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtMsecs

`func (o *McmSignup) SetCreatedAtMsecs(v int64)`

SetCreatedAtMsecs sets CreatedAtMsecs field to given value.

### HasCreatedAtMsecs

`func (o *McmSignup) HasCreatedAtMsecs() bool`

HasCreatedAtMsecs returns a boolean if a field has been set.

### SetCreatedAtMsecsNil

`func (o *McmSignup) SetCreatedAtMsecsNil(b bool)`

 SetCreatedAtMsecsNil sets the value for CreatedAtMsecs to be an explicit nil

### UnsetCreatedAtMsecs
`func (o *McmSignup) UnsetCreatedAtMsecs()`

UnsetCreatedAtMsecs ensures that no value is present for CreatedAtMsecs, not even an explicit nil
### GetUpdatedAtMsecs

`func (o *McmSignup) GetUpdatedAtMsecs() int64`

GetUpdatedAtMsecs returns the UpdatedAtMsecs field if non-nil, zero value otherwise.

### GetUpdatedAtMsecsOk

`func (o *McmSignup) GetUpdatedAtMsecsOk() (*int64, bool)`

GetUpdatedAtMsecsOk returns a tuple with the UpdatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtMsecs

`func (o *McmSignup) SetUpdatedAtMsecs(v int64)`

SetUpdatedAtMsecs sets UpdatedAtMsecs field to given value.

### HasUpdatedAtMsecs

`func (o *McmSignup) HasUpdatedAtMsecs() bool`

HasUpdatedAtMsecs returns a boolean if a field has been set.

### SetUpdatedAtMsecsNil

`func (o *McmSignup) SetUpdatedAtMsecsNil(b bool)`

 SetUpdatedAtMsecsNil sets the value for UpdatedAtMsecs to be an explicit nil

### UnsetUpdatedAtMsecs
`func (o *McmSignup) UnsetUpdatedAtMsecs()`

UnsetUpdatedAtMsecs ensures that no value is present for UpdatedAtMsecs, not even an explicit nil
### GetStatus

`func (o *McmSignup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McmSignup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McmSignup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *McmSignup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *McmSignup) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *McmSignup) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAcceptTerms

`func (o *McmSignup) GetAcceptTerms() bool`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *McmSignup) GetAcceptTermsOk() (*bool, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *McmSignup) SetAcceptTerms(v bool)`

SetAcceptTerms sets AcceptTerms field to given value.

### HasAcceptTerms

`func (o *McmSignup) HasAcceptTerms() bool`

HasAcceptTerms returns a boolean if a field has been set.

### SetAcceptTermsNil

`func (o *McmSignup) SetAcceptTermsNil(b bool)`

 SetAcceptTermsNil sets the value for AcceptTerms to be an explicit nil

### UnsetAcceptTerms
`func (o *McmSignup) UnsetAcceptTerms()`

UnsetAcceptTerms ensures that no value is present for AcceptTerms, not even an explicit nil
### GetContactInfo

`func (o *McmSignup) GetContactInfo() McmSignupUserContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *McmSignup) GetContactInfoOk() (*McmSignupUserContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *McmSignup) SetContactInfo(v McmSignupUserContactInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetWorkloads

`func (o *McmSignup) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *McmSignup) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *McmSignup) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *McmSignup) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetDmaasParams

`func (o *McmSignup) GetDmaasParams() DmaasSignupParams`

GetDmaasParams returns the DmaasParams field if non-nil, zero value otherwise.

### GetDmaasParamsOk

`func (o *McmSignup) GetDmaasParamsOk() (*DmaasSignupParams, bool)`

GetDmaasParamsOk returns a tuple with the DmaasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaasParams

`func (o *McmSignup) SetDmaasParams(v DmaasSignupParams)`

SetDmaasParams sets DmaasParams field to given value.

### HasDmaasParams

`func (o *McmSignup) HasDmaasParams() bool`

HasDmaasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


