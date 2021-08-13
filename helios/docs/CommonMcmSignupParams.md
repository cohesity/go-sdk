# CommonMcmSignupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInfo** | [**McmSignupUserContactInfo**](McmSignupUserContactInfo.md) |  | 
**Workloads** | Pointer to **[]string** | Specifies the workload requirements of the company to use Cohesity MCM. | [optional] 
**DmaasParams** | Pointer to [**DmaasSignupParams**](DmaasSignupParams.md) |  | [optional] 

## Methods

### NewCommonMcmSignupParams

`func NewCommonMcmSignupParams(contactInfo McmSignupUserContactInfo, ) *CommonMcmSignupParams`

NewCommonMcmSignupParams instantiates a new CommonMcmSignupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMcmSignupParamsWithDefaults

`func NewCommonMcmSignupParamsWithDefaults() *CommonMcmSignupParams`

NewCommonMcmSignupParamsWithDefaults instantiates a new CommonMcmSignupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInfo

`func (o *CommonMcmSignupParams) GetContactInfo() McmSignupUserContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *CommonMcmSignupParams) GetContactInfoOk() (*McmSignupUserContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *CommonMcmSignupParams) SetContactInfo(v McmSignupUserContactInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetWorkloads

`func (o *CommonMcmSignupParams) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *CommonMcmSignupParams) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *CommonMcmSignupParams) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *CommonMcmSignupParams) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetDmaasParams

`func (o *CommonMcmSignupParams) GetDmaasParams() DmaasSignupParams`

GetDmaasParams returns the DmaasParams field if non-nil, zero value otherwise.

### GetDmaasParamsOk

`func (o *CommonMcmSignupParams) GetDmaasParamsOk() (*DmaasSignupParams, bool)`

GetDmaasParamsOk returns a tuple with the DmaasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaasParams

`func (o *CommonMcmSignupParams) SetDmaasParams(v DmaasSignupParams)`

SetDmaasParams sets DmaasParams field to given value.

### HasDmaasParams

`func (o *CommonMcmSignupParams) HasDmaasParams() bool`

HasDmaasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


