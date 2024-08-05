# AcropolisProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDisks** | Pointer to [**[]AcropolisDiskInfo**](AcropolisDiskInfo.md) | Specifies a list of disks to exclude from being protected. This is only applicable to VM objects. | [optional] 
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**IncludeDisks** | Pointer to [**[]AcropolisDiskInfo**](AcropolisDiskInfo.md) | Specifies a list of disks to include in the protection. This is only applicable to VM objects. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the virtual machine. | [optional] [readonly] 

## Methods

### NewAcropolisProtectionGroupObjectParams

`func NewAcropolisProtectionGroupObjectParams(id NullableInt64, ) *AcropolisProtectionGroupObjectParams`

NewAcropolisProtectionGroupObjectParams instantiates a new AcropolisProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisProtectionGroupObjectParamsWithDefaults

`func NewAcropolisProtectionGroupObjectParamsWithDefaults() *AcropolisProtectionGroupObjectParams`

NewAcropolisProtectionGroupObjectParamsWithDefaults instantiates a new AcropolisProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDisks

`func (o *AcropolisProtectionGroupObjectParams) GetExcludeDisks() []AcropolisDiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *AcropolisProtectionGroupObjectParams) GetExcludeDisksOk() (*[]AcropolisDiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *AcropolisProtectionGroupObjectParams) SetExcludeDisks(v []AcropolisDiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *AcropolisProtectionGroupObjectParams) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### SetExcludeDisksNil

`func (o *AcropolisProtectionGroupObjectParams) SetExcludeDisksNil(b bool)`

 SetExcludeDisksNil sets the value for ExcludeDisks to be an explicit nil

### UnsetExcludeDisks
`func (o *AcropolisProtectionGroupObjectParams) UnsetExcludeDisks()`

UnsetExcludeDisks ensures that no value is present for ExcludeDisks, not even an explicit nil
### GetId

`func (o *AcropolisProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcropolisProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcropolisProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AcropolisProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AcropolisProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncludeDisks

`func (o *AcropolisProtectionGroupObjectParams) GetIncludeDisks() []AcropolisDiskInfo`

GetIncludeDisks returns the IncludeDisks field if non-nil, zero value otherwise.

### GetIncludeDisksOk

`func (o *AcropolisProtectionGroupObjectParams) GetIncludeDisksOk() (*[]AcropolisDiskInfo, bool)`

GetIncludeDisksOk returns a tuple with the IncludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisks

`func (o *AcropolisProtectionGroupObjectParams) SetIncludeDisks(v []AcropolisDiskInfo)`

SetIncludeDisks sets IncludeDisks field to given value.

### HasIncludeDisks

`func (o *AcropolisProtectionGroupObjectParams) HasIncludeDisks() bool`

HasIncludeDisks returns a boolean if a field has been set.

### SetIncludeDisksNil

`func (o *AcropolisProtectionGroupObjectParams) SetIncludeDisksNil(b bool)`

 SetIncludeDisksNil sets the value for IncludeDisks to be an explicit nil

### UnsetIncludeDisks
`func (o *AcropolisProtectionGroupObjectParams) UnsetIncludeDisks()`

UnsetIncludeDisks ensures that no value is present for IncludeDisks, not even an explicit nil
### GetName

`func (o *AcropolisProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcropolisProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcropolisProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AcropolisProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AcropolisProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AcropolisProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


