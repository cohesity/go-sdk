# EbsVolumeExclusionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceNames** | Pointer to **[]string** | Array of device names to exclude. Eg - /dev/sda. | [optional] 
**MaxVolumeSizeBytes** | Pointer to **NullableInt64** | Any volume larger than this size will be excluded. | [optional] 
**RawQuery** | Pointer to **NullableString** | Raw boolean query given as input by the user to exclude volume based on tags. In the current version, the query contains only tags. Eg. query 1 - \&quot;K1\&quot; &#x3D; \&quot;V1\&quot; AND \&quot;K2\&quot; IN (\&quot;V2\&quot;, \&quot;V3\&quot;) AND \&quot;K4\&quot; !&#x3D; \&quot;V4\&quot; Eg. query 2 - \&quot;K1\&quot; !&#x3D; \&quot;V1\&quot; OR \&quot;K2\&quot; NOT IN (\&quot;V2\&quot;, \&quot;V3\&quot;) OR \&quot;K4\&quot; &#x3D; \&quot;V4\&quot; All Keys and Values must be wrapped inside double quotes. Comparision Operators supported - IN, NOT IN, &#x3D;, !&#x3D;. Logical Operators supported - AND, OR. We cannot have AND, OR together in the query. Only one of them is allowed. The processed form for this query is stored in the above tagParamsArray. | [optional] 
**TagParamsArray** | Pointer to [**[]TagParams**](TagParams.md) | Array of TagParams objects. Each TagParams object consists of two vectors: for exclusion and inclusion. Each TagPararms object is present as an ORed item. User can only input queries of form: (&lt;&gt; AND &lt;&gt; AND &lt;&gt; ..) OR (&lt;&gt; AND &lt;&gt; AND &lt;&gt; ..) OR (..) OR (..) OR .. There cannot be an OR operator inside the bracket. Example query: (K1 &#x3D; V1 AND K2 &#x3D; V2 AND K3 !&#x3D; V3) OR (K4 &#x3D; V4 AND K6 !&#x3D; V6). This will lead to formation of two items in tagParamsArray. First item: {exclusionTagArray: [(K1, V1),  (K2, V2)], inclusionTagArray: [(K3, V3)]} Second item: {exclusionTagArray: [(K4, V4)], inclusionTagArray: [(K6, V6)]}. | [optional] 
**VolumeIds** | Pointer to **[]string** | Array of volume IDs that are to be excluded. This is only for object level exclusion. | [optional] 
**VolumeTypes** | Pointer to **[]string** | Array of volume types to exclude. Eg - gp2, gp3. | [optional] 

## Methods

### NewEbsVolumeExclusionParams

`func NewEbsVolumeExclusionParams() *EbsVolumeExclusionParams`

NewEbsVolumeExclusionParams instantiates a new EbsVolumeExclusionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbsVolumeExclusionParamsWithDefaults

`func NewEbsVolumeExclusionParamsWithDefaults() *EbsVolumeExclusionParams`

NewEbsVolumeExclusionParamsWithDefaults instantiates a new EbsVolumeExclusionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceNames

`func (o *EbsVolumeExclusionParams) GetDeviceNames() []string`

GetDeviceNames returns the DeviceNames field if non-nil, zero value otherwise.

### GetDeviceNamesOk

`func (o *EbsVolumeExclusionParams) GetDeviceNamesOk() (*[]string, bool)`

GetDeviceNamesOk returns a tuple with the DeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNames

`func (o *EbsVolumeExclusionParams) SetDeviceNames(v []string)`

SetDeviceNames sets DeviceNames field to given value.

### HasDeviceNames

`func (o *EbsVolumeExclusionParams) HasDeviceNames() bool`

HasDeviceNames returns a boolean if a field has been set.

### SetDeviceNamesNil

`func (o *EbsVolumeExclusionParams) SetDeviceNamesNil(b bool)`

 SetDeviceNamesNil sets the value for DeviceNames to be an explicit nil

### UnsetDeviceNames
`func (o *EbsVolumeExclusionParams) UnsetDeviceNames()`

UnsetDeviceNames ensures that no value is present for DeviceNames, not even an explicit nil
### GetMaxVolumeSizeBytes

`func (o *EbsVolumeExclusionParams) GetMaxVolumeSizeBytes() int64`

GetMaxVolumeSizeBytes returns the MaxVolumeSizeBytes field if non-nil, zero value otherwise.

### GetMaxVolumeSizeBytesOk

`func (o *EbsVolumeExclusionParams) GetMaxVolumeSizeBytesOk() (*int64, bool)`

GetMaxVolumeSizeBytesOk returns a tuple with the MaxVolumeSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumeSizeBytes

`func (o *EbsVolumeExclusionParams) SetMaxVolumeSizeBytes(v int64)`

SetMaxVolumeSizeBytes sets MaxVolumeSizeBytes field to given value.

### HasMaxVolumeSizeBytes

`func (o *EbsVolumeExclusionParams) HasMaxVolumeSizeBytes() bool`

HasMaxVolumeSizeBytes returns a boolean if a field has been set.

### SetMaxVolumeSizeBytesNil

`func (o *EbsVolumeExclusionParams) SetMaxVolumeSizeBytesNil(b bool)`

 SetMaxVolumeSizeBytesNil sets the value for MaxVolumeSizeBytes to be an explicit nil

### UnsetMaxVolumeSizeBytes
`func (o *EbsVolumeExclusionParams) UnsetMaxVolumeSizeBytes()`

UnsetMaxVolumeSizeBytes ensures that no value is present for MaxVolumeSizeBytes, not even an explicit nil
### GetRawQuery

`func (o *EbsVolumeExclusionParams) GetRawQuery() string`

GetRawQuery returns the RawQuery field if non-nil, zero value otherwise.

### GetRawQueryOk

`func (o *EbsVolumeExclusionParams) GetRawQueryOk() (*string, bool)`

GetRawQueryOk returns a tuple with the RawQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuery

`func (o *EbsVolumeExclusionParams) SetRawQuery(v string)`

SetRawQuery sets RawQuery field to given value.

### HasRawQuery

`func (o *EbsVolumeExclusionParams) HasRawQuery() bool`

HasRawQuery returns a boolean if a field has been set.

### SetRawQueryNil

`func (o *EbsVolumeExclusionParams) SetRawQueryNil(b bool)`

 SetRawQueryNil sets the value for RawQuery to be an explicit nil

### UnsetRawQuery
`func (o *EbsVolumeExclusionParams) UnsetRawQuery()`

UnsetRawQuery ensures that no value is present for RawQuery, not even an explicit nil
### GetTagParamsArray

`func (o *EbsVolumeExclusionParams) GetTagParamsArray() []TagParams`

GetTagParamsArray returns the TagParamsArray field if non-nil, zero value otherwise.

### GetTagParamsArrayOk

`func (o *EbsVolumeExclusionParams) GetTagParamsArrayOk() (*[]TagParams, bool)`

GetTagParamsArrayOk returns a tuple with the TagParamsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagParamsArray

`func (o *EbsVolumeExclusionParams) SetTagParamsArray(v []TagParams)`

SetTagParamsArray sets TagParamsArray field to given value.

### HasTagParamsArray

`func (o *EbsVolumeExclusionParams) HasTagParamsArray() bool`

HasTagParamsArray returns a boolean if a field has been set.

### SetTagParamsArrayNil

`func (o *EbsVolumeExclusionParams) SetTagParamsArrayNil(b bool)`

 SetTagParamsArrayNil sets the value for TagParamsArray to be an explicit nil

### UnsetTagParamsArray
`func (o *EbsVolumeExclusionParams) UnsetTagParamsArray()`

UnsetTagParamsArray ensures that no value is present for TagParamsArray, not even an explicit nil
### GetVolumeIds

`func (o *EbsVolumeExclusionParams) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *EbsVolumeExclusionParams) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *EbsVolumeExclusionParams) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.

### HasVolumeIds

`func (o *EbsVolumeExclusionParams) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### SetVolumeIdsNil

`func (o *EbsVolumeExclusionParams) SetVolumeIdsNil(b bool)`

 SetVolumeIdsNil sets the value for VolumeIds to be an explicit nil

### UnsetVolumeIds
`func (o *EbsVolumeExclusionParams) UnsetVolumeIds()`

UnsetVolumeIds ensures that no value is present for VolumeIds, not even an explicit nil
### GetVolumeTypes

`func (o *EbsVolumeExclusionParams) GetVolumeTypes() []string`

GetVolumeTypes returns the VolumeTypes field if non-nil, zero value otherwise.

### GetVolumeTypesOk

`func (o *EbsVolumeExclusionParams) GetVolumeTypesOk() (*[]string, bool)`

GetVolumeTypesOk returns a tuple with the VolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypes

`func (o *EbsVolumeExclusionParams) SetVolumeTypes(v []string)`

SetVolumeTypes sets VolumeTypes field to given value.

### HasVolumeTypes

`func (o *EbsVolumeExclusionParams) HasVolumeTypes() bool`

HasVolumeTypes returns a boolean if a field has been set.

### SetVolumeTypesNil

`func (o *EbsVolumeExclusionParams) SetVolumeTypesNil(b bool)`

 SetVolumeTypesNil sets the value for VolumeTypes to be an explicit nil

### UnsetVolumeTypes
`func (o *EbsVolumeExclusionParams) UnsetVolumeTypes()`

UnsetVolumeTypes ensures that no value is present for VolumeTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


