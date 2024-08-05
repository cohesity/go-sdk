# SfdcObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]SfdcObjectProtectionObjectParams**](SfdcObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 

## Methods

### NewSfdcObjectProtectionResponseParams

`func NewSfdcObjectProtectionResponseParams(objects []SfdcObjectProtectionObjectParams, ) *SfdcObjectProtectionResponseParams`

NewSfdcObjectProtectionResponseParams instantiates a new SfdcObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcObjectProtectionResponseParamsWithDefaults

`func NewSfdcObjectProtectionResponseParamsWithDefaults() *SfdcObjectProtectionResponseParams`

NewSfdcObjectProtectionResponseParamsWithDefaults instantiates a new SfdcObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *SfdcObjectProtectionResponseParams) GetObjects() []SfdcObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *SfdcObjectProtectionResponseParams) GetObjectsOk() (*[]SfdcObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *SfdcObjectProtectionResponseParams) SetObjects(v []SfdcObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


