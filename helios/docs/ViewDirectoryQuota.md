# ViewDirectoryQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryPath** | **NullableString** | Specifies the directory path. | 
**QuotaPolicy** | [**ViewDirectoryQuotaPolicy**](ViewDirectoryQuotaPolicy.md) | Specifies the directory quota policy. | 
**UsageBytes** | Pointer to **NullableInt64** | Specifies the directory usage in bytes. | [optional] [readonly] 
**DirectoryWalkPending** | Pointer to **NullableBool** | Specifies whether directory quota walk is pending. | [optional] [readonly] 

## Methods

### NewViewDirectoryQuota

`func NewViewDirectoryQuota(directoryPath NullableString, quotaPolicy ViewDirectoryQuotaPolicy, ) *ViewDirectoryQuota`

NewViewDirectoryQuota instantiates a new ViewDirectoryQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewDirectoryQuotaWithDefaults

`func NewViewDirectoryQuotaWithDefaults() *ViewDirectoryQuota`

NewViewDirectoryQuotaWithDefaults instantiates a new ViewDirectoryQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryPath

`func (o *ViewDirectoryQuota) GetDirectoryPath() string`

GetDirectoryPath returns the DirectoryPath field if non-nil, zero value otherwise.

### GetDirectoryPathOk

`func (o *ViewDirectoryQuota) GetDirectoryPathOk() (*string, bool)`

GetDirectoryPathOk returns a tuple with the DirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPath

`func (o *ViewDirectoryQuota) SetDirectoryPath(v string)`

SetDirectoryPath sets DirectoryPath field to given value.


### SetDirectoryPathNil

`func (o *ViewDirectoryQuota) SetDirectoryPathNil(b bool)`

 SetDirectoryPathNil sets the value for DirectoryPath to be an explicit nil

### UnsetDirectoryPath
`func (o *ViewDirectoryQuota) UnsetDirectoryPath()`

UnsetDirectoryPath ensures that no value is present for DirectoryPath, not even an explicit nil
### GetQuotaPolicy

`func (o *ViewDirectoryQuota) GetQuotaPolicy() ViewDirectoryQuotaPolicy`

GetQuotaPolicy returns the QuotaPolicy field if non-nil, zero value otherwise.

### GetQuotaPolicyOk

`func (o *ViewDirectoryQuota) GetQuotaPolicyOk() (*ViewDirectoryQuotaPolicy, bool)`

GetQuotaPolicyOk returns a tuple with the QuotaPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaPolicy

`func (o *ViewDirectoryQuota) SetQuotaPolicy(v ViewDirectoryQuotaPolicy)`

SetQuotaPolicy sets QuotaPolicy field to given value.


### GetUsageBytes

`func (o *ViewDirectoryQuota) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *ViewDirectoryQuota) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *ViewDirectoryQuota) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *ViewDirectoryQuota) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### SetUsageBytesNil

`func (o *ViewDirectoryQuota) SetUsageBytesNil(b bool)`

 SetUsageBytesNil sets the value for UsageBytes to be an explicit nil

### UnsetUsageBytes
`func (o *ViewDirectoryQuota) UnsetUsageBytes()`

UnsetUsageBytes ensures that no value is present for UsageBytes, not even an explicit nil
### GetDirectoryWalkPending

`func (o *ViewDirectoryQuota) GetDirectoryWalkPending() bool`

GetDirectoryWalkPending returns the DirectoryWalkPending field if non-nil, zero value otherwise.

### GetDirectoryWalkPendingOk

`func (o *ViewDirectoryQuota) GetDirectoryWalkPendingOk() (*bool, bool)`

GetDirectoryWalkPendingOk returns a tuple with the DirectoryWalkPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryWalkPending

`func (o *ViewDirectoryQuota) SetDirectoryWalkPending(v bool)`

SetDirectoryWalkPending sets DirectoryWalkPending field to given value.

### HasDirectoryWalkPending

`func (o *ViewDirectoryQuota) HasDirectoryWalkPending() bool`

HasDirectoryWalkPending returns a boolean if a field has been set.

### SetDirectoryWalkPendingNil

`func (o *ViewDirectoryQuota) SetDirectoryWalkPendingNil(b bool)`

 SetDirectoryWalkPendingNil sets the value for DirectoryWalkPending to be an explicit nil

### UnsetDirectoryWalkPending
`func (o *ViewDirectoryQuota) UnsetDirectoryWalkPending()`

UnsetDirectoryWalkPending ensures that no value is present for DirectoryWalkPending, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


