# FlashbladeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationParams** | Pointer to [**FlashBladeRegistrationInfo**](FlashBladeRegistrationInfo.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies uuid of the pure flashblade server. | [optional] [readonly] 
**AssignedDataVips** | Pointer to **[]string** | Specifies list of data vips that are assigned to cohesity cluster to create nfs share mountpoints. | [optional] 
**AssignedCapacityBytes** | Pointer to **NullableInt64** | Specifies the capacity in bytes assigned on pure flashblade for remote storage usage on cohesity cluster. | [optional] 
**IsDedicatedStorage** | Pointer to **NullableBool** | If true, cohesity cluster uses all available capacity on pure flashblade for remote storage. | [optional] 
**AvailableDataVips** | Pointer to **[]string** | Available data vips configured on pure flashblade. | [optional] [readonly] 
**AvailableCapacity** | Pointer to **NullableInt64** | Available capacity on pure flashblade. | [optional] [readonly] 
**CreatedFileSystemCount** | Pointer to **NullableInt64** | Number of new file systems created on pure flashblade when assignedCapacityBytes is updated. | [optional] [readonly] 
**UpdatedFileSystemCount** | Pointer to **NullableInt64** | Number of file systems that are updated on pure flashblade when assignedCapacityBytes is updated. | [optional] [readonly] 
**SoftwareOsVersion** | Pointer to **NullableString** | Software OS and version running on pure flashblade | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Name of the pure flashblade specified on pure storage. | [optional] [readonly] 
**TotalCapacity** | Pointer to **NullableInt64** | Total capacity of pure flashblade. | [optional] [readonly] 

## Methods

### NewFlashbladeParams

`func NewFlashbladeParams() *FlashbladeParams`

NewFlashbladeParams instantiates a new FlashbladeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashbladeParamsWithDefaults

`func NewFlashbladeParamsWithDefaults() *FlashbladeParams`

NewFlashbladeParamsWithDefaults instantiates a new FlashbladeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationParams

`func (o *FlashbladeParams) GetRegistrationParams() FlashBladeRegistrationInfo`

GetRegistrationParams returns the RegistrationParams field if non-nil, zero value otherwise.

### GetRegistrationParamsOk

`func (o *FlashbladeParams) GetRegistrationParamsOk() (*FlashBladeRegistrationInfo, bool)`

GetRegistrationParamsOk returns a tuple with the RegistrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationParams

`func (o *FlashbladeParams) SetRegistrationParams(v FlashBladeRegistrationInfo)`

SetRegistrationParams sets RegistrationParams field to given value.

### HasRegistrationParams

`func (o *FlashbladeParams) HasRegistrationParams() bool`

HasRegistrationParams returns a boolean if a field has been set.

### GetUuid

`func (o *FlashbladeParams) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FlashbladeParams) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FlashbladeParams) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FlashbladeParams) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *FlashbladeParams) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *FlashbladeParams) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetAssignedDataVips

`func (o *FlashbladeParams) GetAssignedDataVips() []string`

GetAssignedDataVips returns the AssignedDataVips field if non-nil, zero value otherwise.

### GetAssignedDataVipsOk

`func (o *FlashbladeParams) GetAssignedDataVipsOk() (*[]string, bool)`

GetAssignedDataVipsOk returns a tuple with the AssignedDataVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDataVips

`func (o *FlashbladeParams) SetAssignedDataVips(v []string)`

SetAssignedDataVips sets AssignedDataVips field to given value.

### HasAssignedDataVips

`func (o *FlashbladeParams) HasAssignedDataVips() bool`

HasAssignedDataVips returns a boolean if a field has been set.

### SetAssignedDataVipsNil

`func (o *FlashbladeParams) SetAssignedDataVipsNil(b bool)`

 SetAssignedDataVipsNil sets the value for AssignedDataVips to be an explicit nil

### UnsetAssignedDataVips
`func (o *FlashbladeParams) UnsetAssignedDataVips()`

UnsetAssignedDataVips ensures that no value is present for AssignedDataVips, not even an explicit nil
### GetAssignedCapacityBytes

`func (o *FlashbladeParams) GetAssignedCapacityBytes() int64`

GetAssignedCapacityBytes returns the AssignedCapacityBytes field if non-nil, zero value otherwise.

### GetAssignedCapacityBytesOk

`func (o *FlashbladeParams) GetAssignedCapacityBytesOk() (*int64, bool)`

GetAssignedCapacityBytesOk returns a tuple with the AssignedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCapacityBytes

`func (o *FlashbladeParams) SetAssignedCapacityBytes(v int64)`

SetAssignedCapacityBytes sets AssignedCapacityBytes field to given value.

### HasAssignedCapacityBytes

`func (o *FlashbladeParams) HasAssignedCapacityBytes() bool`

HasAssignedCapacityBytes returns a boolean if a field has been set.

### SetAssignedCapacityBytesNil

`func (o *FlashbladeParams) SetAssignedCapacityBytesNil(b bool)`

 SetAssignedCapacityBytesNil sets the value for AssignedCapacityBytes to be an explicit nil

### UnsetAssignedCapacityBytes
`func (o *FlashbladeParams) UnsetAssignedCapacityBytes()`

UnsetAssignedCapacityBytes ensures that no value is present for AssignedCapacityBytes, not even an explicit nil
### GetIsDedicatedStorage

`func (o *FlashbladeParams) GetIsDedicatedStorage() bool`

GetIsDedicatedStorage returns the IsDedicatedStorage field if non-nil, zero value otherwise.

### GetIsDedicatedStorageOk

`func (o *FlashbladeParams) GetIsDedicatedStorageOk() (*bool, bool)`

GetIsDedicatedStorageOk returns a tuple with the IsDedicatedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDedicatedStorage

`func (o *FlashbladeParams) SetIsDedicatedStorage(v bool)`

SetIsDedicatedStorage sets IsDedicatedStorage field to given value.

### HasIsDedicatedStorage

`func (o *FlashbladeParams) HasIsDedicatedStorage() bool`

HasIsDedicatedStorage returns a boolean if a field has been set.

### SetIsDedicatedStorageNil

`func (o *FlashbladeParams) SetIsDedicatedStorageNil(b bool)`

 SetIsDedicatedStorageNil sets the value for IsDedicatedStorage to be an explicit nil

### UnsetIsDedicatedStorage
`func (o *FlashbladeParams) UnsetIsDedicatedStorage()`

UnsetIsDedicatedStorage ensures that no value is present for IsDedicatedStorage, not even an explicit nil
### GetAvailableDataVips

`func (o *FlashbladeParams) GetAvailableDataVips() []string`

GetAvailableDataVips returns the AvailableDataVips field if non-nil, zero value otherwise.

### GetAvailableDataVipsOk

`func (o *FlashbladeParams) GetAvailableDataVipsOk() (*[]string, bool)`

GetAvailableDataVipsOk returns a tuple with the AvailableDataVips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDataVips

`func (o *FlashbladeParams) SetAvailableDataVips(v []string)`

SetAvailableDataVips sets AvailableDataVips field to given value.

### HasAvailableDataVips

`func (o *FlashbladeParams) HasAvailableDataVips() bool`

HasAvailableDataVips returns a boolean if a field has been set.

### GetAvailableCapacity

`func (o *FlashbladeParams) GetAvailableCapacity() int64`

GetAvailableCapacity returns the AvailableCapacity field if non-nil, zero value otherwise.

### GetAvailableCapacityOk

`func (o *FlashbladeParams) GetAvailableCapacityOk() (*int64, bool)`

GetAvailableCapacityOk returns a tuple with the AvailableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCapacity

`func (o *FlashbladeParams) SetAvailableCapacity(v int64)`

SetAvailableCapacity sets AvailableCapacity field to given value.

### HasAvailableCapacity

`func (o *FlashbladeParams) HasAvailableCapacity() bool`

HasAvailableCapacity returns a boolean if a field has been set.

### SetAvailableCapacityNil

`func (o *FlashbladeParams) SetAvailableCapacityNil(b bool)`

 SetAvailableCapacityNil sets the value for AvailableCapacity to be an explicit nil

### UnsetAvailableCapacity
`func (o *FlashbladeParams) UnsetAvailableCapacity()`

UnsetAvailableCapacity ensures that no value is present for AvailableCapacity, not even an explicit nil
### GetCreatedFileSystemCount

`func (o *FlashbladeParams) GetCreatedFileSystemCount() int64`

GetCreatedFileSystemCount returns the CreatedFileSystemCount field if non-nil, zero value otherwise.

### GetCreatedFileSystemCountOk

`func (o *FlashbladeParams) GetCreatedFileSystemCountOk() (*int64, bool)`

GetCreatedFileSystemCountOk returns a tuple with the CreatedFileSystemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFileSystemCount

`func (o *FlashbladeParams) SetCreatedFileSystemCount(v int64)`

SetCreatedFileSystemCount sets CreatedFileSystemCount field to given value.

### HasCreatedFileSystemCount

`func (o *FlashbladeParams) HasCreatedFileSystemCount() bool`

HasCreatedFileSystemCount returns a boolean if a field has been set.

### SetCreatedFileSystemCountNil

`func (o *FlashbladeParams) SetCreatedFileSystemCountNil(b bool)`

 SetCreatedFileSystemCountNil sets the value for CreatedFileSystemCount to be an explicit nil

### UnsetCreatedFileSystemCount
`func (o *FlashbladeParams) UnsetCreatedFileSystemCount()`

UnsetCreatedFileSystemCount ensures that no value is present for CreatedFileSystemCount, not even an explicit nil
### GetUpdatedFileSystemCount

`func (o *FlashbladeParams) GetUpdatedFileSystemCount() int64`

GetUpdatedFileSystemCount returns the UpdatedFileSystemCount field if non-nil, zero value otherwise.

### GetUpdatedFileSystemCountOk

`func (o *FlashbladeParams) GetUpdatedFileSystemCountOk() (*int64, bool)`

GetUpdatedFileSystemCountOk returns a tuple with the UpdatedFileSystemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedFileSystemCount

`func (o *FlashbladeParams) SetUpdatedFileSystemCount(v int64)`

SetUpdatedFileSystemCount sets UpdatedFileSystemCount field to given value.

### HasUpdatedFileSystemCount

`func (o *FlashbladeParams) HasUpdatedFileSystemCount() bool`

HasUpdatedFileSystemCount returns a boolean if a field has been set.

### SetUpdatedFileSystemCountNil

`func (o *FlashbladeParams) SetUpdatedFileSystemCountNil(b bool)`

 SetUpdatedFileSystemCountNil sets the value for UpdatedFileSystemCount to be an explicit nil

### UnsetUpdatedFileSystemCount
`func (o *FlashbladeParams) UnsetUpdatedFileSystemCount()`

UnsetUpdatedFileSystemCount ensures that no value is present for UpdatedFileSystemCount, not even an explicit nil
### GetSoftwareOsVersion

`func (o *FlashbladeParams) GetSoftwareOsVersion() string`

GetSoftwareOsVersion returns the SoftwareOsVersion field if non-nil, zero value otherwise.

### GetSoftwareOsVersionOk

`func (o *FlashbladeParams) GetSoftwareOsVersionOk() (*string, bool)`

GetSoftwareOsVersionOk returns a tuple with the SoftwareOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareOsVersion

`func (o *FlashbladeParams) SetSoftwareOsVersion(v string)`

SetSoftwareOsVersion sets SoftwareOsVersion field to given value.

### HasSoftwareOsVersion

`func (o *FlashbladeParams) HasSoftwareOsVersion() bool`

HasSoftwareOsVersion returns a boolean if a field has been set.

### SetSoftwareOsVersionNil

`func (o *FlashbladeParams) SetSoftwareOsVersionNil(b bool)`

 SetSoftwareOsVersionNil sets the value for SoftwareOsVersion to be an explicit nil

### UnsetSoftwareOsVersion
`func (o *FlashbladeParams) UnsetSoftwareOsVersion()`

UnsetSoftwareOsVersion ensures that no value is present for SoftwareOsVersion, not even an explicit nil
### GetName

`func (o *FlashbladeParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlashbladeParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlashbladeParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlashbladeParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FlashbladeParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FlashbladeParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTotalCapacity

`func (o *FlashbladeParams) GetTotalCapacity() int64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *FlashbladeParams) GetTotalCapacityOk() (*int64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *FlashbladeParams) SetTotalCapacity(v int64)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *FlashbladeParams) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### SetTotalCapacityNil

`func (o *FlashbladeParams) SetTotalCapacityNil(b bool)`

 SetTotalCapacityNil sets the value for TotalCapacity to be an explicit nil

### UnsetTotalCapacity
`func (o *FlashbladeParams) UnsetTotalCapacity()`

UnsetTotalCapacity ensures that no value is present for TotalCapacity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


