# CreateProtectedObjectsRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]EnvSpecificObjectProtectionRequestParams**](EnvSpecificObjectProtectionRequestParams.md) | Specifies the list of objects to be protected. Multiple objects from different adapters can be provided as input. | 
**ActivateRemoteObjectProtection** | Pointer to **NullableBool** | If set to true, it will look for the remote backup of the given user and object, and activates it. Creates a new backup if the remote backup is not found. After activation, this object cannot get snapshots from remote clusters. | [optional] 

## Methods

### NewCreateProtectedObjectsRequestAllOf

`func NewCreateProtectedObjectsRequestAllOf(objects []EnvSpecificObjectProtectionRequestParams, ) *CreateProtectedObjectsRequestAllOf`

NewCreateProtectedObjectsRequestAllOf instantiates a new CreateProtectedObjectsRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProtectedObjectsRequestAllOfWithDefaults

`func NewCreateProtectedObjectsRequestAllOfWithDefaults() *CreateProtectedObjectsRequestAllOf`

NewCreateProtectedObjectsRequestAllOfWithDefaults instantiates a new CreateProtectedObjectsRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *CreateProtectedObjectsRequestAllOf) GetObjects() []EnvSpecificObjectProtectionRequestParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CreateProtectedObjectsRequestAllOf) GetObjectsOk() (*[]EnvSpecificObjectProtectionRequestParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CreateProtectedObjectsRequestAllOf) SetObjects(v []EnvSpecificObjectProtectionRequestParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *CreateProtectedObjectsRequestAllOf) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *CreateProtectedObjectsRequestAllOf) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetActivateRemoteObjectProtection

`func (o *CreateProtectedObjectsRequestAllOf) GetActivateRemoteObjectProtection() bool`

GetActivateRemoteObjectProtection returns the ActivateRemoteObjectProtection field if non-nil, zero value otherwise.

### GetActivateRemoteObjectProtectionOk

`func (o *CreateProtectedObjectsRequestAllOf) GetActivateRemoteObjectProtectionOk() (*bool, bool)`

GetActivateRemoteObjectProtectionOk returns a tuple with the ActivateRemoteObjectProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateRemoteObjectProtection

`func (o *CreateProtectedObjectsRequestAllOf) SetActivateRemoteObjectProtection(v bool)`

SetActivateRemoteObjectProtection sets ActivateRemoteObjectProtection field to given value.

### HasActivateRemoteObjectProtection

`func (o *CreateProtectedObjectsRequestAllOf) HasActivateRemoteObjectProtection() bool`

HasActivateRemoteObjectProtection returns a boolean if a field has been set.

### SetActivateRemoteObjectProtectionNil

`func (o *CreateProtectedObjectsRequestAllOf) SetActivateRemoteObjectProtectionNil(b bool)`

 SetActivateRemoteObjectProtectionNil sets the value for ActivateRemoteObjectProtection to be an explicit nil

### UnsetActivateRemoteObjectProtection
`func (o *CreateProtectedObjectsRequestAllOf) UnsetActivateRemoteObjectProtection()`

UnsetActivateRemoteObjectProtection ensures that no value is present for ActivateRemoteObjectProtection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


