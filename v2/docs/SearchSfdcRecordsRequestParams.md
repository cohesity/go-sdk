# SearchSfdcRecordsRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MutationTypes** | **[]string** | Specifies a list of mutuation types for an object. | 
**ObjectName** | **NullableString** | Specifies the name of the object. | 
**QueryString** | Pointer to **NullableString** | Specifies the query string to search records. Query string can be one or multiples clauses joined together by &#39;AND&#39; or &#39;OR&#39; claused. | [optional] 
**SnapshotId** | **NullableString** | Specifies the id of the snapshot for the object. | 

## Methods

### NewSearchSfdcRecordsRequestParams

`func NewSearchSfdcRecordsRequestParams(mutationTypes []string, objectName NullableString, snapshotId NullableString, ) *SearchSfdcRecordsRequestParams`

NewSearchSfdcRecordsRequestParams instantiates a new SearchSfdcRecordsRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSfdcRecordsRequestParamsWithDefaults

`func NewSearchSfdcRecordsRequestParamsWithDefaults() *SearchSfdcRecordsRequestParams`

NewSearchSfdcRecordsRequestParamsWithDefaults instantiates a new SearchSfdcRecordsRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMutationTypes

`func (o *SearchSfdcRecordsRequestParams) GetMutationTypes() []string`

GetMutationTypes returns the MutationTypes field if non-nil, zero value otherwise.

### GetMutationTypesOk

`func (o *SearchSfdcRecordsRequestParams) GetMutationTypesOk() (*[]string, bool)`

GetMutationTypesOk returns a tuple with the MutationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationTypes

`func (o *SearchSfdcRecordsRequestParams) SetMutationTypes(v []string)`

SetMutationTypes sets MutationTypes field to given value.


### SetMutationTypesNil

`func (o *SearchSfdcRecordsRequestParams) SetMutationTypesNil(b bool)`

 SetMutationTypesNil sets the value for MutationTypes to be an explicit nil

### UnsetMutationTypes
`func (o *SearchSfdcRecordsRequestParams) UnsetMutationTypes()`

UnsetMutationTypes ensures that no value is present for MutationTypes, not even an explicit nil
### GetObjectName

`func (o *SearchSfdcRecordsRequestParams) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *SearchSfdcRecordsRequestParams) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *SearchSfdcRecordsRequestParams) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.


### SetObjectNameNil

`func (o *SearchSfdcRecordsRequestParams) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *SearchSfdcRecordsRequestParams) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetQueryString

`func (o *SearchSfdcRecordsRequestParams) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *SearchSfdcRecordsRequestParams) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *SearchSfdcRecordsRequestParams) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *SearchSfdcRecordsRequestParams) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### SetQueryStringNil

`func (o *SearchSfdcRecordsRequestParams) SetQueryStringNil(b bool)`

 SetQueryStringNil sets the value for QueryString to be an explicit nil

### UnsetQueryString
`func (o *SearchSfdcRecordsRequestParams) UnsetQueryString()`

UnsetQueryString ensures that no value is present for QueryString, not even an explicit nil
### GetSnapshotId

`func (o *SearchSfdcRecordsRequestParams) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SearchSfdcRecordsRequestParams) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SearchSfdcRecordsRequestParams) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### SetSnapshotIdNil

`func (o *SearchSfdcRecordsRequestParams) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *SearchSfdcRecordsRequestParams) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


