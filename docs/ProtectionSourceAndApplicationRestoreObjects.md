# ProtectionSourceAndApplicationRestoreObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationRestoreObjects** | Pointer to [**[]ApplicationRestoreObject**](ApplicationRestoreObject.md) | Specifies the Application Server objects whose data should be restored and the restore parameters for each of them. | [optional] 
**HostingProtectionSource** | [**RestoreObjectDetails**](RestoreObjectDetails.md) |  | 

## Methods

### NewProtectionSourceAndApplicationRestoreObjects

`func NewProtectionSourceAndApplicationRestoreObjects(hostingProtectionSource RestoreObjectDetails, ) *ProtectionSourceAndApplicationRestoreObjects`

NewProtectionSourceAndApplicationRestoreObjects instantiates a new ProtectionSourceAndApplicationRestoreObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceAndApplicationRestoreObjectsWithDefaults

`func NewProtectionSourceAndApplicationRestoreObjectsWithDefaults() *ProtectionSourceAndApplicationRestoreObjects`

NewProtectionSourceAndApplicationRestoreObjectsWithDefaults instantiates a new ProtectionSourceAndApplicationRestoreObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationRestoreObjects

`func (o *ProtectionSourceAndApplicationRestoreObjects) GetApplicationRestoreObjects() []ApplicationRestoreObject`

GetApplicationRestoreObjects returns the ApplicationRestoreObjects field if non-nil, zero value otherwise.

### GetApplicationRestoreObjectsOk

`func (o *ProtectionSourceAndApplicationRestoreObjects) GetApplicationRestoreObjectsOk() (*[]ApplicationRestoreObject, bool)`

GetApplicationRestoreObjectsOk returns a tuple with the ApplicationRestoreObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationRestoreObjects

`func (o *ProtectionSourceAndApplicationRestoreObjects) SetApplicationRestoreObjects(v []ApplicationRestoreObject)`

SetApplicationRestoreObjects sets ApplicationRestoreObjects field to given value.

### HasApplicationRestoreObjects

`func (o *ProtectionSourceAndApplicationRestoreObjects) HasApplicationRestoreObjects() bool`

HasApplicationRestoreObjects returns a boolean if a field has been set.

### SetApplicationRestoreObjectsNil

`func (o *ProtectionSourceAndApplicationRestoreObjects) SetApplicationRestoreObjectsNil(b bool)`

 SetApplicationRestoreObjectsNil sets the value for ApplicationRestoreObjects to be an explicit nil

### UnsetApplicationRestoreObjects
`func (o *ProtectionSourceAndApplicationRestoreObjects) UnsetApplicationRestoreObjects()`

UnsetApplicationRestoreObjects ensures that no value is present for ApplicationRestoreObjects, not even an explicit nil
### GetHostingProtectionSource

`func (o *ProtectionSourceAndApplicationRestoreObjects) GetHostingProtectionSource() RestoreObjectDetails`

GetHostingProtectionSource returns the HostingProtectionSource field if non-nil, zero value otherwise.

### GetHostingProtectionSourceOk

`func (o *ProtectionSourceAndApplicationRestoreObjects) GetHostingProtectionSourceOk() (*RestoreObjectDetails, bool)`

GetHostingProtectionSourceOk returns a tuple with the HostingProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingProtectionSource

`func (o *ProtectionSourceAndApplicationRestoreObjects) SetHostingProtectionSource(v RestoreObjectDetails)`

SetHostingProtectionSource sets HostingProtectionSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


