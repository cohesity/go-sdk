# CommonVmwareObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from being protected. This is only applicable to VM objects. | [optional] 
**TruncateExchangeLogs** | Pointer to **NullableBool** | Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app. | [optional] 

## Methods

### NewCommonVmwareObjectParams

`func NewCommonVmwareObjectParams() *CommonVmwareObjectParams`

NewCommonVmwareObjectParams instantiates a new CommonVmwareObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonVmwareObjectParamsWithDefaults

`func NewCommonVmwareObjectParamsWithDefaults() *CommonVmwareObjectParams`

NewCommonVmwareObjectParamsWithDefaults instantiates a new CommonVmwareObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDisks

`func (o *CommonVmwareObjectParams) GetExcludeDisks() []DiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *CommonVmwareObjectParams) GetExcludeDisksOk() (*[]DiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *CommonVmwareObjectParams) SetExcludeDisks(v []DiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *CommonVmwareObjectParams) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### GetTruncateExchangeLogs

`func (o *CommonVmwareObjectParams) GetTruncateExchangeLogs() bool`

GetTruncateExchangeLogs returns the TruncateExchangeLogs field if non-nil, zero value otherwise.

### GetTruncateExchangeLogsOk

`func (o *CommonVmwareObjectParams) GetTruncateExchangeLogsOk() (*bool, bool)`

GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLogs

`func (o *CommonVmwareObjectParams) SetTruncateExchangeLogs(v bool)`

SetTruncateExchangeLogs sets TruncateExchangeLogs field to given value.

### HasTruncateExchangeLogs

`func (o *CommonVmwareObjectParams) HasTruncateExchangeLogs() bool`

HasTruncateExchangeLogs returns a boolean if a field has been set.

### SetTruncateExchangeLogsNil

`func (o *CommonVmwareObjectParams) SetTruncateExchangeLogsNil(b bool)`

 SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil

### UnsetTruncateExchangeLogs
`func (o *CommonVmwareObjectParams) UnsetTruncateExchangeLogs()`

UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


