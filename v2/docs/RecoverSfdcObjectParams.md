# RecoverSfdcObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildObjectIds** | Pointer to **[]string** | Specifies a list of child object IDs to include in the recovery. Specified object IDs will also be recovered as part of this recovery. | [optional] 
**FilterQuery** | Pointer to **NullableString** | Specifies a Query to filter the records. This filtered list of records will be used for recovery. | [optional] 
**IncludeDeletedObjects** | **NullableBool** | Specifies whether to include deleted Objects in the recovery. | 
**MutationTypes** | Pointer to **[]string** | Specifies a list of mutuation types for an object. Mutation type is required in conjunction with &#39;filterQuery&#39;. | [optional] 
**ObjectName** | Pointer to **NullableString** | Specifies the name of the object to be restored. | [optional] 
**ParentObjectIds** | Pointer to **[]string** | Specifies a list of parent object IDs to include in recovery. Specified parent objects will also be recovered as part of this recovery. | [optional] 
**Records** | Pointer to **[]string** | Specifies a list of records IDs to be recovered for the specified Object. | [optional] 

## Methods

### NewRecoverSfdcObjectParams

`func NewRecoverSfdcObjectParams(includeDeletedObjects NullableBool, ) *RecoverSfdcObjectParams`

NewRecoverSfdcObjectParams instantiates a new RecoverSfdcObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSfdcObjectParamsWithDefaults

`func NewRecoverSfdcObjectParamsWithDefaults() *RecoverSfdcObjectParams`

NewRecoverSfdcObjectParamsWithDefaults instantiates a new RecoverSfdcObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildObjectIds

`func (o *RecoverSfdcObjectParams) GetChildObjectIds() []string`

GetChildObjectIds returns the ChildObjectIds field if non-nil, zero value otherwise.

### GetChildObjectIdsOk

`func (o *RecoverSfdcObjectParams) GetChildObjectIdsOk() (*[]string, bool)`

GetChildObjectIdsOk returns a tuple with the ChildObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjectIds

`func (o *RecoverSfdcObjectParams) SetChildObjectIds(v []string)`

SetChildObjectIds sets ChildObjectIds field to given value.

### HasChildObjectIds

`func (o *RecoverSfdcObjectParams) HasChildObjectIds() bool`

HasChildObjectIds returns a boolean if a field has been set.

### GetFilterQuery

`func (o *RecoverSfdcObjectParams) GetFilterQuery() string`

GetFilterQuery returns the FilterQuery field if non-nil, zero value otherwise.

### GetFilterQueryOk

`func (o *RecoverSfdcObjectParams) GetFilterQueryOk() (*string, bool)`

GetFilterQueryOk returns a tuple with the FilterQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterQuery

`func (o *RecoverSfdcObjectParams) SetFilterQuery(v string)`

SetFilterQuery sets FilterQuery field to given value.

### HasFilterQuery

`func (o *RecoverSfdcObjectParams) HasFilterQuery() bool`

HasFilterQuery returns a boolean if a field has been set.

### SetFilterQueryNil

`func (o *RecoverSfdcObjectParams) SetFilterQueryNil(b bool)`

 SetFilterQueryNil sets the value for FilterQuery to be an explicit nil

### UnsetFilterQuery
`func (o *RecoverSfdcObjectParams) UnsetFilterQuery()`

UnsetFilterQuery ensures that no value is present for FilterQuery, not even an explicit nil
### GetIncludeDeletedObjects

`func (o *RecoverSfdcObjectParams) GetIncludeDeletedObjects() bool`

GetIncludeDeletedObjects returns the IncludeDeletedObjects field if non-nil, zero value otherwise.

### GetIncludeDeletedObjectsOk

`func (o *RecoverSfdcObjectParams) GetIncludeDeletedObjectsOk() (*bool, bool)`

GetIncludeDeletedObjectsOk returns a tuple with the IncludeDeletedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeletedObjects

`func (o *RecoverSfdcObjectParams) SetIncludeDeletedObjects(v bool)`

SetIncludeDeletedObjects sets IncludeDeletedObjects field to given value.


### SetIncludeDeletedObjectsNil

`func (o *RecoverSfdcObjectParams) SetIncludeDeletedObjectsNil(b bool)`

 SetIncludeDeletedObjectsNil sets the value for IncludeDeletedObjects to be an explicit nil

### UnsetIncludeDeletedObjects
`func (o *RecoverSfdcObjectParams) UnsetIncludeDeletedObjects()`

UnsetIncludeDeletedObjects ensures that no value is present for IncludeDeletedObjects, not even an explicit nil
### GetMutationTypes

`func (o *RecoverSfdcObjectParams) GetMutationTypes() []string`

GetMutationTypes returns the MutationTypes field if non-nil, zero value otherwise.

### GetMutationTypesOk

`func (o *RecoverSfdcObjectParams) GetMutationTypesOk() (*[]string, bool)`

GetMutationTypesOk returns a tuple with the MutationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutationTypes

`func (o *RecoverSfdcObjectParams) SetMutationTypes(v []string)`

SetMutationTypes sets MutationTypes field to given value.

### HasMutationTypes

`func (o *RecoverSfdcObjectParams) HasMutationTypes() bool`

HasMutationTypes returns a boolean if a field has been set.

### SetMutationTypesNil

`func (o *RecoverSfdcObjectParams) SetMutationTypesNil(b bool)`

 SetMutationTypesNil sets the value for MutationTypes to be an explicit nil

### UnsetMutationTypes
`func (o *RecoverSfdcObjectParams) UnsetMutationTypes()`

UnsetMutationTypes ensures that no value is present for MutationTypes, not even an explicit nil
### GetObjectName

`func (o *RecoverSfdcObjectParams) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *RecoverSfdcObjectParams) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *RecoverSfdcObjectParams) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *RecoverSfdcObjectParams) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### SetObjectNameNil

`func (o *RecoverSfdcObjectParams) SetObjectNameNil(b bool)`

 SetObjectNameNil sets the value for ObjectName to be an explicit nil

### UnsetObjectName
`func (o *RecoverSfdcObjectParams) UnsetObjectName()`

UnsetObjectName ensures that no value is present for ObjectName, not even an explicit nil
### GetParentObjectIds

`func (o *RecoverSfdcObjectParams) GetParentObjectIds() []string`

GetParentObjectIds returns the ParentObjectIds field if non-nil, zero value otherwise.

### GetParentObjectIdsOk

`func (o *RecoverSfdcObjectParams) GetParentObjectIdsOk() (*[]string, bool)`

GetParentObjectIdsOk returns a tuple with the ParentObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjectIds

`func (o *RecoverSfdcObjectParams) SetParentObjectIds(v []string)`

SetParentObjectIds sets ParentObjectIds field to given value.

### HasParentObjectIds

`func (o *RecoverSfdcObjectParams) HasParentObjectIds() bool`

HasParentObjectIds returns a boolean if a field has been set.

### SetParentObjectIdsNil

`func (o *RecoverSfdcObjectParams) SetParentObjectIdsNil(b bool)`

 SetParentObjectIdsNil sets the value for ParentObjectIds to be an explicit nil

### UnsetParentObjectIds
`func (o *RecoverSfdcObjectParams) UnsetParentObjectIds()`

UnsetParentObjectIds ensures that no value is present for ParentObjectIds, not even an explicit nil
### GetRecords

`func (o *RecoverSfdcObjectParams) GetRecords() []string`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RecoverSfdcObjectParams) GetRecordsOk() (*[]string, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RecoverSfdcObjectParams) SetRecords(v []string)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *RecoverSfdcObjectParams) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### SetRecordsNil

`func (o *RecoverSfdcObjectParams) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *RecoverSfdcObjectParams) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


