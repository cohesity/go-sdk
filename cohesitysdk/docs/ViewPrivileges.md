# ViewPrivileges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivilegesType** | Pointer to **NullableString** | Specifies if all, none or specific views are allowed to be accessed. Specifies if all, none or specific views are allowed to be accessed. kNone - None of the views have access. kAll - All the views have access. kSpecific - Only specific views have access. | [optional] 
**ViewIds** | Pointer to **[]int64** | Specifies the ids of the views which are allowed to be accessed in case the privilege type is kSpecific. | [optional] 

## Methods

### NewViewPrivileges

`func NewViewPrivileges() *ViewPrivileges`

NewViewPrivileges instantiates a new ViewPrivileges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPrivilegesWithDefaults

`func NewViewPrivilegesWithDefaults() *ViewPrivileges`

NewViewPrivilegesWithDefaults instantiates a new ViewPrivileges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivilegesType

`func (o *ViewPrivileges) GetPrivilegesType() string`

GetPrivilegesType returns the PrivilegesType field if non-nil, zero value otherwise.

### GetPrivilegesTypeOk

`func (o *ViewPrivileges) GetPrivilegesTypeOk() (*string, bool)`

GetPrivilegesTypeOk returns a tuple with the PrivilegesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesType

`func (o *ViewPrivileges) SetPrivilegesType(v string)`

SetPrivilegesType sets PrivilegesType field to given value.

### HasPrivilegesType

`func (o *ViewPrivileges) HasPrivilegesType() bool`

HasPrivilegesType returns a boolean if a field has been set.

### SetPrivilegesTypeNil

`func (o *ViewPrivileges) SetPrivilegesTypeNil(b bool)`

 SetPrivilegesTypeNil sets the value for PrivilegesType to be an explicit nil

### UnsetPrivilegesType
`func (o *ViewPrivileges) UnsetPrivilegesType()`

UnsetPrivilegesType ensures that no value is present for PrivilegesType, not even an explicit nil
### GetViewIds

`func (o *ViewPrivileges) GetViewIds() []int64`

GetViewIds returns the ViewIds field if non-nil, zero value otherwise.

### GetViewIdsOk

`func (o *ViewPrivileges) GetViewIdsOk() (*[]int64, bool)`

GetViewIdsOk returns a tuple with the ViewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewIds

`func (o *ViewPrivileges) SetViewIds(v []int64)`

SetViewIds sets ViewIds field to given value.

### HasViewIds

`func (o *ViewPrivileges) HasViewIds() bool`

HasViewIds returns a boolean if a field has been set.

### SetViewIdsNil

`func (o *ViewPrivileges) SetViewIdsNil(b bool)`

 SetViewIdsNil sets the value for ViewIds to be an explicit nil

### UnsetViewIds
`func (o *ViewPrivileges) UnsetViewIds()`

UnsetViewIds ensures that no value is present for ViewIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


