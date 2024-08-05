# AzureDiskTagParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExclusionTagArray** | Pointer to [**[]AzureDiskTag**](AzureDiskTag.md) | Array which contains tags for AND exclusion. E.g., exclusionTagArray: [(K1, V1),  (K2, V2)] &#x3D;&gt; This will exclude a particular volume iff it has both these tags. | [optional] 
**InclusionTagArray** | Pointer to [**[]AzureDiskTag**](AzureDiskTag.md) | Array which contains tags for AND inclusion. E.g., inclusionTagArray: [(K3, V3),  (K4, V4)] &#x3D;&gt; This will exclude a particular volume iff it does not have both these tags. | [optional] 

## Methods

### NewAzureDiskTagParams

`func NewAzureDiskTagParams() *AzureDiskTagParams`

NewAzureDiskTagParams instantiates a new AzureDiskTagParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDiskTagParamsWithDefaults

`func NewAzureDiskTagParamsWithDefaults() *AzureDiskTagParams`

NewAzureDiskTagParamsWithDefaults instantiates a new AzureDiskTagParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusionTagArray

`func (o *AzureDiskTagParams) GetExclusionTagArray() []AzureDiskTag`

GetExclusionTagArray returns the ExclusionTagArray field if non-nil, zero value otherwise.

### GetExclusionTagArrayOk

`func (o *AzureDiskTagParams) GetExclusionTagArrayOk() (*[]AzureDiskTag, bool)`

GetExclusionTagArrayOk returns a tuple with the ExclusionTagArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionTagArray

`func (o *AzureDiskTagParams) SetExclusionTagArray(v []AzureDiskTag)`

SetExclusionTagArray sets ExclusionTagArray field to given value.

### HasExclusionTagArray

`func (o *AzureDiskTagParams) HasExclusionTagArray() bool`

HasExclusionTagArray returns a boolean if a field has been set.

### SetExclusionTagArrayNil

`func (o *AzureDiskTagParams) SetExclusionTagArrayNil(b bool)`

 SetExclusionTagArrayNil sets the value for ExclusionTagArray to be an explicit nil

### UnsetExclusionTagArray
`func (o *AzureDiskTagParams) UnsetExclusionTagArray()`

UnsetExclusionTagArray ensures that no value is present for ExclusionTagArray, not even an explicit nil
### GetInclusionTagArray

`func (o *AzureDiskTagParams) GetInclusionTagArray() []AzureDiskTag`

GetInclusionTagArray returns the InclusionTagArray field if non-nil, zero value otherwise.

### GetInclusionTagArrayOk

`func (o *AzureDiskTagParams) GetInclusionTagArrayOk() (*[]AzureDiskTag, bool)`

GetInclusionTagArrayOk returns a tuple with the InclusionTagArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusionTagArray

`func (o *AzureDiskTagParams) SetInclusionTagArray(v []AzureDiskTag)`

SetInclusionTagArray sets InclusionTagArray field to given value.

### HasInclusionTagArray

`func (o *AzureDiskTagParams) HasInclusionTagArray() bool`

HasInclusionTagArray returns a boolean if a field has been set.

### SetInclusionTagArrayNil

`func (o *AzureDiskTagParams) SetInclusionTagArrayNil(b bool)`

 SetInclusionTagArrayNil sets the value for InclusionTagArray to be an explicit nil

### UnsetInclusionTagArray
`func (o *AzureDiskTagParams) UnsetInclusionTagArray()`

UnsetInclusionTagArray ensures that no value is present for InclusionTagArray, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


