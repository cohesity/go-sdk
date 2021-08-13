# UpdateMcmSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInfo** | [**McmSignupUserContactInfo**](McmSignupUserContactInfo.md) |  | 
**Workloads** | Pointer to **[]string** | Specifies the workload requirements of the company to use Cohesity MCM. | [optional] 
**DmaasParams** | Pointer to [**DmaasSignupParams**](DmaasSignupParams.md) |  | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the MCM signup request. | [optional] 

## Methods

### NewUpdateMcmSignupRequest

`func NewUpdateMcmSignupRequest(contactInfo McmSignupUserContactInfo, ) *UpdateMcmSignupRequest`

NewUpdateMcmSignupRequest instantiates a new UpdateMcmSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMcmSignupRequestWithDefaults

`func NewUpdateMcmSignupRequestWithDefaults() *UpdateMcmSignupRequest`

NewUpdateMcmSignupRequestWithDefaults instantiates a new UpdateMcmSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInfo

`func (o *UpdateMcmSignupRequest) GetContactInfo() McmSignupUserContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *UpdateMcmSignupRequest) GetContactInfoOk() (*McmSignupUserContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *UpdateMcmSignupRequest) SetContactInfo(v McmSignupUserContactInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetWorkloads

`func (o *UpdateMcmSignupRequest) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *UpdateMcmSignupRequest) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *UpdateMcmSignupRequest) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *UpdateMcmSignupRequest) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetDmaasParams

`func (o *UpdateMcmSignupRequest) GetDmaasParams() DmaasSignupParams`

GetDmaasParams returns the DmaasParams field if non-nil, zero value otherwise.

### GetDmaasParamsOk

`func (o *UpdateMcmSignupRequest) GetDmaasParamsOk() (*DmaasSignupParams, bool)`

GetDmaasParamsOk returns a tuple with the DmaasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaasParams

`func (o *UpdateMcmSignupRequest) SetDmaasParams(v DmaasSignupParams)`

SetDmaasParams sets DmaasParams field to given value.

### HasDmaasParams

`func (o *UpdateMcmSignupRequest) HasDmaasParams() bool`

HasDmaasParams returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateMcmSignupRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateMcmSignupRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateMcmSignupRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateMcmSignupRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateMcmSignupRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateMcmSignupRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


