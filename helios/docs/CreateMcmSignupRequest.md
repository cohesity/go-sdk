# CreateMcmSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInfo** | [**McmSignupUserContactInfo**](McmSignupUserContactInfo.md) |  | 
**Workloads** | Pointer to **[]string** | Specifies the workload requirements of the company to use Cohesity MCM. | [optional] 
**DmaasParams** | Pointer to [**DmaasSignupParams**](DmaasSignupParams.md) |  | [optional] 
**AcceptTerms** | **bool** | Specifies whether the User agrees to Terms and conditions for using Cohesity MCM services which are specified at &lt;url&gt; | 

## Methods

### NewCreateMcmSignupRequest

`func NewCreateMcmSignupRequest(contactInfo McmSignupUserContactInfo, acceptTerms bool, ) *CreateMcmSignupRequest`

NewCreateMcmSignupRequest instantiates a new CreateMcmSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMcmSignupRequestWithDefaults

`func NewCreateMcmSignupRequestWithDefaults() *CreateMcmSignupRequest`

NewCreateMcmSignupRequestWithDefaults instantiates a new CreateMcmSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInfo

`func (o *CreateMcmSignupRequest) GetContactInfo() McmSignupUserContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *CreateMcmSignupRequest) GetContactInfoOk() (*McmSignupUserContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *CreateMcmSignupRequest) SetContactInfo(v McmSignupUserContactInfo)`

SetContactInfo sets ContactInfo field to given value.


### GetWorkloads

`func (o *CreateMcmSignupRequest) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *CreateMcmSignupRequest) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *CreateMcmSignupRequest) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *CreateMcmSignupRequest) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.

### GetDmaasParams

`func (o *CreateMcmSignupRequest) GetDmaasParams() DmaasSignupParams`

GetDmaasParams returns the DmaasParams field if non-nil, zero value otherwise.

### GetDmaasParamsOk

`func (o *CreateMcmSignupRequest) GetDmaasParamsOk() (*DmaasSignupParams, bool)`

GetDmaasParamsOk returns a tuple with the DmaasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaasParams

`func (o *CreateMcmSignupRequest) SetDmaasParams(v DmaasSignupParams)`

SetDmaasParams sets DmaasParams field to given value.

### HasDmaasParams

`func (o *CreateMcmSignupRequest) HasDmaasParams() bool`

HasDmaasParams returns a boolean if a field has been set.

### GetAcceptTerms

`func (o *CreateMcmSignupRequest) GetAcceptTerms() bool`

GetAcceptTerms returns the AcceptTerms field if non-nil, zero value otherwise.

### GetAcceptTermsOk

`func (o *CreateMcmSignupRequest) GetAcceptTermsOk() (*bool, bool)`

GetAcceptTermsOk returns a tuple with the AcceptTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTerms

`func (o *CreateMcmSignupRequest) SetAcceptTerms(v bool)`

SetAcceptTerms sets AcceptTerms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


