# MSSQLFileProtectionGroupHostParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableSourceSideDeduplication** | Pointer to **NullableBool** | Specifies whether or not to disable source side deduplication on this source. The default behavior is false unless the user has set &#39;performSourceSideDeduplication&#39; to true. | [optional] 
**HostId** | **NullableInt64** | Specifies the id of the host container on which databases are hosted. | 
**HostName** | Pointer to **NullableString** | Specifies the name of the host container on which databases are hosted. | [optional] [readonly] 

## Methods

### NewMSSQLFileProtectionGroupHostParams

`func NewMSSQLFileProtectionGroupHostParams(hostId NullableInt64, ) *MSSQLFileProtectionGroupHostParams`

NewMSSQLFileProtectionGroupHostParams instantiates a new MSSQLFileProtectionGroupHostParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLFileProtectionGroupHostParamsWithDefaults

`func NewMSSQLFileProtectionGroupHostParamsWithDefaults() *MSSQLFileProtectionGroupHostParams`

NewMSSQLFileProtectionGroupHostParamsWithDefaults instantiates a new MSSQLFileProtectionGroupHostParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupHostParams) GetDisableSourceSideDeduplication() bool`

GetDisableSourceSideDeduplication returns the DisableSourceSideDeduplication field if non-nil, zero value otherwise.

### GetDisableSourceSideDeduplicationOk

`func (o *MSSQLFileProtectionGroupHostParams) GetDisableSourceSideDeduplicationOk() (*bool, bool)`

GetDisableSourceSideDeduplicationOk returns a tuple with the DisableSourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupHostParams) SetDisableSourceSideDeduplication(v bool)`

SetDisableSourceSideDeduplication sets DisableSourceSideDeduplication field to given value.

### HasDisableSourceSideDeduplication

`func (o *MSSQLFileProtectionGroupHostParams) HasDisableSourceSideDeduplication() bool`

HasDisableSourceSideDeduplication returns a boolean if a field has been set.

### SetDisableSourceSideDeduplicationNil

`func (o *MSSQLFileProtectionGroupHostParams) SetDisableSourceSideDeduplicationNil(b bool)`

 SetDisableSourceSideDeduplicationNil sets the value for DisableSourceSideDeduplication to be an explicit nil

### UnsetDisableSourceSideDeduplication
`func (o *MSSQLFileProtectionGroupHostParams) UnsetDisableSourceSideDeduplication()`

UnsetDisableSourceSideDeduplication ensures that no value is present for DisableSourceSideDeduplication, not even an explicit nil
### GetHostId

`func (o *MSSQLFileProtectionGroupHostParams) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *MSSQLFileProtectionGroupHostParams) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *MSSQLFileProtectionGroupHostParams) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### SetHostIdNil

`func (o *MSSQLFileProtectionGroupHostParams) SetHostIdNil(b bool)`

 SetHostIdNil sets the value for HostId to be an explicit nil

### UnsetHostId
`func (o *MSSQLFileProtectionGroupHostParams) UnsetHostId()`

UnsetHostId ensures that no value is present for HostId, not even an explicit nil
### GetHostName

`func (o *MSSQLFileProtectionGroupHostParams) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *MSSQLFileProtectionGroupHostParams) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *MSSQLFileProtectionGroupHostParams) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *MSSQLFileProtectionGroupHostParams) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *MSSQLFileProtectionGroupHostParams) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *MSSQLFileProtectionGroupHostParams) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


