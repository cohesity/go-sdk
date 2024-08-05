# AzureDiskExclusionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskIds** | Pointer to **[]string** | Array of disk ids to be excluded during backup. This is only used at object level exclusion. | [optional] 
**RawQuery** | Pointer to **NullableString** | Raw boolean query given as input by the user to exclude volume based on tags. In the current version, the query contains only tags. Eg. query 1 - \&quot;K1\&quot; &#x3D; \&quot;V1\&quot; AND \&quot;K2\&quot; IN (\&quot;V2\&quot;, \&quot;V3\&quot;) AND \&quot;K4\&quot; !&#x3D; \&quot;V4\&quot; Eg. query 2 - \&quot;K1\&quot; !&#x3D; \&quot;V1\&quot; OR \&quot;K2\&quot; NOT IN (\&quot;V2\&quot;, \&quot;V3\&quot;) OR \&quot;K4\&quot; &#x3D; \&quot;V4\&quot; All Keys and Values must be wrapped inside double quotes. Comparision Operators supported - IN, NOT IN, &#x3D;, !&#x3D;. Logical Operators supported - AND, OR. We cannot have AND, OR together in the query. Only one of them is allowed. The processed form for this query is stored in the above tagParamsArray. | [optional] 
**TagParamsArray** | Pointer to [**[]AzureDiskTagParams**](AzureDiskTagParams.md) | Array of TagParams objects. Each TagParams object consists of two vectors: for exclusion and inclusion. Each TagPararms object is present as an ORed item. User can only input queries of form: (&lt;&gt; AND &lt;&gt; AND &lt;&gt; ..) OR (&lt;&gt; AND &lt;&gt; AND &lt;&gt; ..) OR (..) OR (..) OR .. There cannot be an OR operator inside the bracket. Example query: (K1 &#x3D; V1 AND K2 &#x3D; V2 AND K3 !&#x3D; V3) OR (K4 &#x3D; V4 AND K6 !&#x3D; V6). This will lead to formation of two items in tagParamsArray. First item: {exclusionTagArray: [(K1, V1),  (K2, V2)], inclusionTagArray: [(K3, V3)]} Second item: {exclusionTagArray: [(K4, V4)], inclusionTagArray: [(K6, V6)]}. | [optional] 

## Methods

### NewAzureDiskExclusionParams

`func NewAzureDiskExclusionParams() *AzureDiskExclusionParams`

NewAzureDiskExclusionParams instantiates a new AzureDiskExclusionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDiskExclusionParamsWithDefaults

`func NewAzureDiskExclusionParamsWithDefaults() *AzureDiskExclusionParams`

NewAzureDiskExclusionParamsWithDefaults instantiates a new AzureDiskExclusionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskIds

`func (o *AzureDiskExclusionParams) GetDiskIds() []string`

GetDiskIds returns the DiskIds field if non-nil, zero value otherwise.

### GetDiskIdsOk

`func (o *AzureDiskExclusionParams) GetDiskIdsOk() (*[]string, bool)`

GetDiskIdsOk returns a tuple with the DiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIds

`func (o *AzureDiskExclusionParams) SetDiskIds(v []string)`

SetDiskIds sets DiskIds field to given value.

### HasDiskIds

`func (o *AzureDiskExclusionParams) HasDiskIds() bool`

HasDiskIds returns a boolean if a field has been set.

### SetDiskIdsNil

`func (o *AzureDiskExclusionParams) SetDiskIdsNil(b bool)`

 SetDiskIdsNil sets the value for DiskIds to be an explicit nil

### UnsetDiskIds
`func (o *AzureDiskExclusionParams) UnsetDiskIds()`

UnsetDiskIds ensures that no value is present for DiskIds, not even an explicit nil
### GetRawQuery

`func (o *AzureDiskExclusionParams) GetRawQuery() string`

GetRawQuery returns the RawQuery field if non-nil, zero value otherwise.

### GetRawQueryOk

`func (o *AzureDiskExclusionParams) GetRawQueryOk() (*string, bool)`

GetRawQueryOk returns a tuple with the RawQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuery

`func (o *AzureDiskExclusionParams) SetRawQuery(v string)`

SetRawQuery sets RawQuery field to given value.

### HasRawQuery

`func (o *AzureDiskExclusionParams) HasRawQuery() bool`

HasRawQuery returns a boolean if a field has been set.

### SetRawQueryNil

`func (o *AzureDiskExclusionParams) SetRawQueryNil(b bool)`

 SetRawQueryNil sets the value for RawQuery to be an explicit nil

### UnsetRawQuery
`func (o *AzureDiskExclusionParams) UnsetRawQuery()`

UnsetRawQuery ensures that no value is present for RawQuery, not even an explicit nil
### GetTagParamsArray

`func (o *AzureDiskExclusionParams) GetTagParamsArray() []AzureDiskTagParams`

GetTagParamsArray returns the TagParamsArray field if non-nil, zero value otherwise.

### GetTagParamsArrayOk

`func (o *AzureDiskExclusionParams) GetTagParamsArrayOk() (*[]AzureDiskTagParams, bool)`

GetTagParamsArrayOk returns a tuple with the TagParamsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagParamsArray

`func (o *AzureDiskExclusionParams) SetTagParamsArray(v []AzureDiskTagParams)`

SetTagParamsArray sets TagParamsArray field to given value.

### HasTagParamsArray

`func (o *AzureDiskExclusionParams) HasTagParamsArray() bool`

HasTagParamsArray returns a boolean if a field has been set.

### SetTagParamsArrayNil

`func (o *AzureDiskExclusionParams) SetTagParamsArrayNil(b bool)`

 SetTagParamsArrayNil sets the value for TagParamsArray to be an explicit nil

### UnsetTagParamsArray
`func (o *AzureDiskExclusionParams) UnsetTagParamsArray()`

UnsetTagParamsArray ensures that no value is present for TagParamsArray, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


