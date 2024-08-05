# CreateGaiaIndexParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | Pointer to **NullableString** | Opaque Id passed to emblem service on where to store the generated vector embeddings for these documents | [optional] 
**EmblemServiceInfo** | Pointer to [**NullableEmblemServiceInfo**](EmblemServiceInfo.md) |  | [optional] 
**JobHandle** | Pointer to **NullableString** | Job handle for this request. | [optional] 
**MaxDocumentSize** | Pointer to **NullableInt64** | Specifies the number of Bytes. | [optional] 
**Objects** | [**[]ObjectInfo**](ObjectInfo.md) | The objects and their snapshots to index. | 

## Methods

### NewCreateGaiaIndexParams

`func NewCreateGaiaIndexParams(objects []ObjectInfo, ) *CreateGaiaIndexParams`

NewCreateGaiaIndexParams instantiates a new CreateGaiaIndexParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGaiaIndexParamsWithDefaults

`func NewCreateGaiaIndexParamsWithDefaults() *CreateGaiaIndexParams`

NewCreateGaiaIndexParamsWithDefaults instantiates a new CreateGaiaIndexParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CreateGaiaIndexParams) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CreateGaiaIndexParams) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CreateGaiaIndexParams) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CreateGaiaIndexParams) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *CreateGaiaIndexParams) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *CreateGaiaIndexParams) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil
### GetEmblemServiceInfo

`func (o *CreateGaiaIndexParams) GetEmblemServiceInfo() EmblemServiceInfo`

GetEmblemServiceInfo returns the EmblemServiceInfo field if non-nil, zero value otherwise.

### GetEmblemServiceInfoOk

`func (o *CreateGaiaIndexParams) GetEmblemServiceInfoOk() (*EmblemServiceInfo, bool)`

GetEmblemServiceInfoOk returns a tuple with the EmblemServiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmblemServiceInfo

`func (o *CreateGaiaIndexParams) SetEmblemServiceInfo(v EmblemServiceInfo)`

SetEmblemServiceInfo sets EmblemServiceInfo field to given value.

### HasEmblemServiceInfo

`func (o *CreateGaiaIndexParams) HasEmblemServiceInfo() bool`

HasEmblemServiceInfo returns a boolean if a field has been set.

### SetEmblemServiceInfoNil

`func (o *CreateGaiaIndexParams) SetEmblemServiceInfoNil(b bool)`

 SetEmblemServiceInfoNil sets the value for EmblemServiceInfo to be an explicit nil

### UnsetEmblemServiceInfo
`func (o *CreateGaiaIndexParams) UnsetEmblemServiceInfo()`

UnsetEmblemServiceInfo ensures that no value is present for EmblemServiceInfo, not even an explicit nil
### GetJobHandle

`func (o *CreateGaiaIndexParams) GetJobHandle() string`

GetJobHandle returns the JobHandle field if non-nil, zero value otherwise.

### GetJobHandleOk

`func (o *CreateGaiaIndexParams) GetJobHandleOk() (*string, bool)`

GetJobHandleOk returns a tuple with the JobHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobHandle

`func (o *CreateGaiaIndexParams) SetJobHandle(v string)`

SetJobHandle sets JobHandle field to given value.

### HasJobHandle

`func (o *CreateGaiaIndexParams) HasJobHandle() bool`

HasJobHandle returns a boolean if a field has been set.

### SetJobHandleNil

`func (o *CreateGaiaIndexParams) SetJobHandleNil(b bool)`

 SetJobHandleNil sets the value for JobHandle to be an explicit nil

### UnsetJobHandle
`func (o *CreateGaiaIndexParams) UnsetJobHandle()`

UnsetJobHandle ensures that no value is present for JobHandle, not even an explicit nil
### GetMaxDocumentSize

`func (o *CreateGaiaIndexParams) GetMaxDocumentSize() int64`

GetMaxDocumentSize returns the MaxDocumentSize field if non-nil, zero value otherwise.

### GetMaxDocumentSizeOk

`func (o *CreateGaiaIndexParams) GetMaxDocumentSizeOk() (*int64, bool)`

GetMaxDocumentSizeOk returns a tuple with the MaxDocumentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDocumentSize

`func (o *CreateGaiaIndexParams) SetMaxDocumentSize(v int64)`

SetMaxDocumentSize sets MaxDocumentSize field to given value.

### HasMaxDocumentSize

`func (o *CreateGaiaIndexParams) HasMaxDocumentSize() bool`

HasMaxDocumentSize returns a boolean if a field has been set.

### SetMaxDocumentSizeNil

`func (o *CreateGaiaIndexParams) SetMaxDocumentSizeNil(b bool)`

 SetMaxDocumentSizeNil sets the value for MaxDocumentSize to be an explicit nil

### UnsetMaxDocumentSize
`func (o *CreateGaiaIndexParams) UnsetMaxDocumentSize()`

UnsetMaxDocumentSize ensures that no value is present for MaxDocumentSize, not even an explicit nil
### GetObjects

`func (o *CreateGaiaIndexParams) GetObjects() []ObjectInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CreateGaiaIndexParams) GetObjectsOk() (*[]ObjectInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CreateGaiaIndexParams) SetObjects(v []ObjectInfo)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *CreateGaiaIndexParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *CreateGaiaIndexParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


