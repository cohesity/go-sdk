# DirQuotaPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirPath** | Pointer to **NullableString** | Specifies the path of the directory in the view. | [optional] 
**DirWalkPending** | Pointer to **NullableBool** | Denotes directory quota walk is pending or not. | [optional] 
**Policy** | Pointer to [**QuotaPolicy**](QuotaPolicy.md) |  | [optional] 
**UsageBytes** | Pointer to **NullableInt64** | Specifies the current usage (in bytes) by the directory in the view. This is set by the response received from bridge when querying directory quota usage. | [optional] 

## Methods

### NewDirQuotaPolicy

`func NewDirQuotaPolicy() *DirQuotaPolicy`

NewDirQuotaPolicy instantiates a new DirQuotaPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirQuotaPolicyWithDefaults

`func NewDirQuotaPolicyWithDefaults() *DirQuotaPolicy`

NewDirQuotaPolicyWithDefaults instantiates a new DirQuotaPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirPath

`func (o *DirQuotaPolicy) GetDirPath() string`

GetDirPath returns the DirPath field if non-nil, zero value otherwise.

### GetDirPathOk

`func (o *DirQuotaPolicy) GetDirPathOk() (*string, bool)`

GetDirPathOk returns a tuple with the DirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirPath

`func (o *DirQuotaPolicy) SetDirPath(v string)`

SetDirPath sets DirPath field to given value.

### HasDirPath

`func (o *DirQuotaPolicy) HasDirPath() bool`

HasDirPath returns a boolean if a field has been set.

### SetDirPathNil

`func (o *DirQuotaPolicy) SetDirPathNil(b bool)`

 SetDirPathNil sets the value for DirPath to be an explicit nil

### UnsetDirPath
`func (o *DirQuotaPolicy) UnsetDirPath()`

UnsetDirPath ensures that no value is present for DirPath, not even an explicit nil
### GetDirWalkPending

`func (o *DirQuotaPolicy) GetDirWalkPending() bool`

GetDirWalkPending returns the DirWalkPending field if non-nil, zero value otherwise.

### GetDirWalkPendingOk

`func (o *DirQuotaPolicy) GetDirWalkPendingOk() (*bool, bool)`

GetDirWalkPendingOk returns a tuple with the DirWalkPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirWalkPending

`func (o *DirQuotaPolicy) SetDirWalkPending(v bool)`

SetDirWalkPending sets DirWalkPending field to given value.

### HasDirWalkPending

`func (o *DirQuotaPolicy) HasDirWalkPending() bool`

HasDirWalkPending returns a boolean if a field has been set.

### SetDirWalkPendingNil

`func (o *DirQuotaPolicy) SetDirWalkPendingNil(b bool)`

 SetDirWalkPendingNil sets the value for DirWalkPending to be an explicit nil

### UnsetDirWalkPending
`func (o *DirQuotaPolicy) UnsetDirWalkPending()`

UnsetDirWalkPending ensures that no value is present for DirWalkPending, not even an explicit nil
### GetPolicy

`func (o *DirQuotaPolicy) GetPolicy() QuotaPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *DirQuotaPolicy) GetPolicyOk() (*QuotaPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *DirQuotaPolicy) SetPolicy(v QuotaPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *DirQuotaPolicy) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetUsageBytes

`func (o *DirQuotaPolicy) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *DirQuotaPolicy) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *DirQuotaPolicy) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *DirQuotaPolicy) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### SetUsageBytesNil

`func (o *DirQuotaPolicy) SetUsageBytesNil(b bool)`

 SetUsageBytesNil sets the value for UsageBytes to be an explicit nil

### UnsetUsageBytes
`func (o *DirQuotaPolicy) UnsetUsageBytes()`

UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


