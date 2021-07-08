# AdObjectsRestoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdObjectsRestoreInfo** | Pointer to [**[]AdObjectRestoreInformation**](AdObjectRestoreInformation.md) | Specifies the status of all the AD Objects which were requested to be restored. | [optional] 
**NumObjectsFailed** | Pointer to **NullableInt32** | Specifies the number of AD Objects whose restore is in progress. | [optional] 
**NumObjectsSucceeded** | Pointer to **NullableInt32** | Specifies the number of AD Objects whose restore is successfull. | [optional] 

## Methods

### NewAdObjectsRestoreStatus

`func NewAdObjectsRestoreStatus() *AdObjectsRestoreStatus`

NewAdObjectsRestoreStatus instantiates a new AdObjectsRestoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdObjectsRestoreStatusWithDefaults

`func NewAdObjectsRestoreStatusWithDefaults() *AdObjectsRestoreStatus`

NewAdObjectsRestoreStatusWithDefaults instantiates a new AdObjectsRestoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdObjectsRestoreInfo

`func (o *AdObjectsRestoreStatus) GetAdObjectsRestoreInfo() []AdObjectRestoreInformation`

GetAdObjectsRestoreInfo returns the AdObjectsRestoreInfo field if non-nil, zero value otherwise.

### GetAdObjectsRestoreInfoOk

`func (o *AdObjectsRestoreStatus) GetAdObjectsRestoreInfoOk() (*[]AdObjectRestoreInformation, bool)`

GetAdObjectsRestoreInfoOk returns a tuple with the AdObjectsRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdObjectsRestoreInfo

`func (o *AdObjectsRestoreStatus) SetAdObjectsRestoreInfo(v []AdObjectRestoreInformation)`

SetAdObjectsRestoreInfo sets AdObjectsRestoreInfo field to given value.

### HasAdObjectsRestoreInfo

`func (o *AdObjectsRestoreStatus) HasAdObjectsRestoreInfo() bool`

HasAdObjectsRestoreInfo returns a boolean if a field has been set.

### SetAdObjectsRestoreInfoNil

`func (o *AdObjectsRestoreStatus) SetAdObjectsRestoreInfoNil(b bool)`

 SetAdObjectsRestoreInfoNil sets the value for AdObjectsRestoreInfo to be an explicit nil

### UnsetAdObjectsRestoreInfo
`func (o *AdObjectsRestoreStatus) UnsetAdObjectsRestoreInfo()`

UnsetAdObjectsRestoreInfo ensures that no value is present for AdObjectsRestoreInfo, not even an explicit nil
### GetNumObjectsFailed

`func (o *AdObjectsRestoreStatus) GetNumObjectsFailed() int32`

GetNumObjectsFailed returns the NumObjectsFailed field if non-nil, zero value otherwise.

### GetNumObjectsFailedOk

`func (o *AdObjectsRestoreStatus) GetNumObjectsFailedOk() (*int32, bool)`

GetNumObjectsFailedOk returns a tuple with the NumObjectsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsFailed

`func (o *AdObjectsRestoreStatus) SetNumObjectsFailed(v int32)`

SetNumObjectsFailed sets NumObjectsFailed field to given value.

### HasNumObjectsFailed

`func (o *AdObjectsRestoreStatus) HasNumObjectsFailed() bool`

HasNumObjectsFailed returns a boolean if a field has been set.

### SetNumObjectsFailedNil

`func (o *AdObjectsRestoreStatus) SetNumObjectsFailedNil(b bool)`

 SetNumObjectsFailedNil sets the value for NumObjectsFailed to be an explicit nil

### UnsetNumObjectsFailed
`func (o *AdObjectsRestoreStatus) UnsetNumObjectsFailed()`

UnsetNumObjectsFailed ensures that no value is present for NumObjectsFailed, not even an explicit nil
### GetNumObjectsSucceeded

`func (o *AdObjectsRestoreStatus) GetNumObjectsSucceeded() int32`

GetNumObjectsSucceeded returns the NumObjectsSucceeded field if non-nil, zero value otherwise.

### GetNumObjectsSucceededOk

`func (o *AdObjectsRestoreStatus) GetNumObjectsSucceededOk() (*int32, bool)`

GetNumObjectsSucceededOk returns a tuple with the NumObjectsSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectsSucceeded

`func (o *AdObjectsRestoreStatus) SetNumObjectsSucceeded(v int32)`

SetNumObjectsSucceeded sets NumObjectsSucceeded field to given value.

### HasNumObjectsSucceeded

`func (o *AdObjectsRestoreStatus) HasNumObjectsSucceeded() bool`

HasNumObjectsSucceeded returns a boolean if a field has been set.

### SetNumObjectsSucceededNil

`func (o *AdObjectsRestoreStatus) SetNumObjectsSucceededNil(b bool)`

 SetNumObjectsSucceededNil sets the value for NumObjectsSucceeded to be an explicit nil

### UnsetNumObjectsSucceeded
`func (o *AdObjectsRestoreStatus) UnsetNumObjectsSucceeded()`

UnsetNumObjectsSucceeded ensures that no value is present for NumObjectsSucceeded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


