# Office365ObjectProtectionObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the ID of the objects to be excluded in the Object Protection. | [optional] 
**Id** | **int64** | Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**ShouldAutoProtectObject** | Pointer to **NullableBool** | Specifies if the object has to be autoprotected. This is applicable only for sharepoint sites. | [optional] 

## Methods

### NewOffice365ObjectProtectionObjectParams

`func NewOffice365ObjectProtectionObjectParams(id int64, ) *Office365ObjectProtectionObjectParams`

NewOffice365ObjectProtectionObjectParams instantiates a new Office365ObjectProtectionObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ObjectProtectionObjectParamsWithDefaults

`func NewOffice365ObjectProtectionObjectParamsWithDefaults() *Office365ObjectProtectionObjectParams`

NewOffice365ObjectProtectionObjectParamsWithDefaults instantiates a new Office365ObjectProtectionObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *Office365ObjectProtectionObjectParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *Office365ObjectProtectionObjectParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *Office365ObjectProtectionObjectParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *Office365ObjectProtectionObjectParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *Office365ObjectProtectionObjectParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *Office365ObjectProtectionObjectParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetId

`func (o *Office365ObjectProtectionObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Office365ObjectProtectionObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Office365ObjectProtectionObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Office365ObjectProtectionObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Office365ObjectProtectionObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Office365ObjectProtectionObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Office365ObjectProtectionObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Office365ObjectProtectionObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Office365ObjectProtectionObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShouldAutoProtectObject

`func (o *Office365ObjectProtectionObjectParams) GetShouldAutoProtectObject() bool`

GetShouldAutoProtectObject returns the ShouldAutoProtectObject field if non-nil, zero value otherwise.

### GetShouldAutoProtectObjectOk

`func (o *Office365ObjectProtectionObjectParams) GetShouldAutoProtectObjectOk() (*bool, bool)`

GetShouldAutoProtectObjectOk returns a tuple with the ShouldAutoProtectObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAutoProtectObject

`func (o *Office365ObjectProtectionObjectParams) SetShouldAutoProtectObject(v bool)`

SetShouldAutoProtectObject sets ShouldAutoProtectObject field to given value.

### HasShouldAutoProtectObject

`func (o *Office365ObjectProtectionObjectParams) HasShouldAutoProtectObject() bool`

HasShouldAutoProtectObject returns a boolean if a field has been set.

### SetShouldAutoProtectObjectNil

`func (o *Office365ObjectProtectionObjectParams) SetShouldAutoProtectObjectNil(b bool)`

 SetShouldAutoProtectObjectNil sets the value for ShouldAutoProtectObject to be an explicit nil

### UnsetShouldAutoProtectObject
`func (o *Office365ObjectProtectionObjectParams) UnsetShouldAutoProtectObject()`

UnsetShouldAutoProtectObject ensures that no value is present for ShouldAutoProtectObject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


