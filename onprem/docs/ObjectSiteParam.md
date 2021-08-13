# ObjectSiteParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerInfo** | [**CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the Site owner info. | 
**DocumentLibraryParams** | Pointer to [**[]OneDriveParam**](OneDriveParam.md) | Specifies the parameters to recover a document library | [optional] 

## Methods

### NewObjectSiteParam

`func NewObjectSiteParam(ownerInfo CommonRecoverObjectSnapshotParams, ) *ObjectSiteParam`

NewObjectSiteParam instantiates a new ObjectSiteParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSiteParamWithDefaults

`func NewObjectSiteParamWithDefaults() *ObjectSiteParam`

NewObjectSiteParamWithDefaults instantiates a new ObjectSiteParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerInfo

`func (o *ObjectSiteParam) GetOwnerInfo() CommonRecoverObjectSnapshotParams`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *ObjectSiteParam) GetOwnerInfoOk() (*CommonRecoverObjectSnapshotParams, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *ObjectSiteParam) SetOwnerInfo(v CommonRecoverObjectSnapshotParams)`

SetOwnerInfo sets OwnerInfo field to given value.


### GetDocumentLibraryParams

`func (o *ObjectSiteParam) GetDocumentLibraryParams() []OneDriveParam`

GetDocumentLibraryParams returns the DocumentLibraryParams field if non-nil, zero value otherwise.

### GetDocumentLibraryParamsOk

`func (o *ObjectSiteParam) GetDocumentLibraryParamsOk() (*[]OneDriveParam, bool)`

GetDocumentLibraryParamsOk returns a tuple with the DocumentLibraryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLibraryParams

`func (o *ObjectSiteParam) SetDocumentLibraryParams(v []OneDriveParam)`

SetDocumentLibraryParams sets DocumentLibraryParams field to given value.

### HasDocumentLibraryParams

`func (o *ObjectSiteParam) HasDocumentLibraryParams() bool`

HasDocumentLibraryParams returns a boolean if a field has been set.

### SetDocumentLibraryParamsNil

`func (o *ObjectSiteParam) SetDocumentLibraryParamsNil(b bool)`

 SetDocumentLibraryParamsNil sets the value for DocumentLibraryParams to be an explicit nil

### UnsetDocumentLibraryParams
`func (o *ObjectSiteParam) UnsetDocumentLibraryParams()`

UnsetDocumentLibraryParams ensures that no value is present for DocumentLibraryParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


