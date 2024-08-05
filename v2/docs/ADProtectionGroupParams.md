# ADProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**Objects** | [**[]ActiveDirectoryProtectionGroupObjectParams**](ActiveDirectoryProtectionGroupObjectParams.md) | Specifies the list of object ids to be protected. | 

## Methods

### NewADProtectionGroupParams

`func NewADProtectionGroupParams(objects []ActiveDirectoryProtectionGroupObjectParams, ) *ADProtectionGroupParams`

NewADProtectionGroupParams instantiates a new ADProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADProtectionGroupParamsWithDefaults

`func NewADProtectionGroupParamsWithDefaults() *ADProtectionGroupParams`

NewADProtectionGroupParamsWithDefaults instantiates a new ADProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexingPolicy

`func (o *ADProtectionGroupParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *ADProtectionGroupParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *ADProtectionGroupParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *ADProtectionGroupParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetObjects

`func (o *ADProtectionGroupParams) GetObjects() []ActiveDirectoryProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ADProtectionGroupParams) GetObjectsOk() (*[]ActiveDirectoryProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ADProtectionGroupParams) SetObjects(v []ActiveDirectoryProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


