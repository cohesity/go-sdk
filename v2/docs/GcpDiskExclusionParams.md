# GcpDiskExclusionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeVmWithNoDisk** | Pointer to **NullableBool** | Specifies the paramaters to exclude VM without disks. | [optional] 
**RawQuery** | Pointer to **NullableString** | Raw boolean query given as input by the user to exclude disk. User can input params in raw query form: (&lt;&gt; AND &lt;&gt; AND &lt;&gt; ..) OR (&lt;&gt; AND &lt;&gt; AND &lt;&gt; ..) OR (..) OR (..) OR .. There cannot be an OR operator inside the bracket. Example query: (K1 &#x3D; V1 AND K2 &#x3D; V2 AND K3 !&#x3D; V3) OR (K4 &#x3D; V4 AND K6 !&#x3D; V6). | [optional] 

## Methods

### NewGcpDiskExclusionParams

`func NewGcpDiskExclusionParams() *GcpDiskExclusionParams`

NewGcpDiskExclusionParams instantiates a new GcpDiskExclusionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDiskExclusionParamsWithDefaults

`func NewGcpDiskExclusionParamsWithDefaults() *GcpDiskExclusionParams`

NewGcpDiskExclusionParamsWithDefaults instantiates a new GcpDiskExclusionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeVmWithNoDisk

`func (o *GcpDiskExclusionParams) GetExcludeVmWithNoDisk() bool`

GetExcludeVmWithNoDisk returns the ExcludeVmWithNoDisk field if non-nil, zero value otherwise.

### GetExcludeVmWithNoDiskOk

`func (o *GcpDiskExclusionParams) GetExcludeVmWithNoDiskOk() (*bool, bool)`

GetExcludeVmWithNoDiskOk returns a tuple with the ExcludeVmWithNoDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeVmWithNoDisk

`func (o *GcpDiskExclusionParams) SetExcludeVmWithNoDisk(v bool)`

SetExcludeVmWithNoDisk sets ExcludeVmWithNoDisk field to given value.

### HasExcludeVmWithNoDisk

`func (o *GcpDiskExclusionParams) HasExcludeVmWithNoDisk() bool`

HasExcludeVmWithNoDisk returns a boolean if a field has been set.

### SetExcludeVmWithNoDiskNil

`func (o *GcpDiskExclusionParams) SetExcludeVmWithNoDiskNil(b bool)`

 SetExcludeVmWithNoDiskNil sets the value for ExcludeVmWithNoDisk to be an explicit nil

### UnsetExcludeVmWithNoDisk
`func (o *GcpDiskExclusionParams) UnsetExcludeVmWithNoDisk()`

UnsetExcludeVmWithNoDisk ensures that no value is present for ExcludeVmWithNoDisk, not even an explicit nil
### GetRawQuery

`func (o *GcpDiskExclusionParams) GetRawQuery() string`

GetRawQuery returns the RawQuery field if non-nil, zero value otherwise.

### GetRawQueryOk

`func (o *GcpDiskExclusionParams) GetRawQueryOk() (*string, bool)`

GetRawQueryOk returns a tuple with the RawQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuery

`func (o *GcpDiskExclusionParams) SetRawQuery(v string)`

SetRawQuery sets RawQuery field to given value.

### HasRawQuery

`func (o *GcpDiskExclusionParams) HasRawQuery() bool`

HasRawQuery returns a boolean if a field has been set.

### SetRawQueryNil

`func (o *GcpDiskExclusionParams) SetRawQueryNil(b bool)`

 SetRawQueryNil sets the value for RawQuery to be an explicit nil

### UnsetRawQuery
`func (o *GcpDiskExclusionParams) UnsetRawQuery()`

UnsetRawQuery ensures that no value is present for RawQuery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


