# SnmpSysInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **NullableString** | Contact is the system contact. | [optional] 
**Description** | Pointer to **NullableString** | Description is the system description. | [optional] 
**EngineIdType** | Pointer to **NullableInt32** | EngineIdType is the system egineId typeß. | [optional] 
**Location** | Pointer to **NullableString** | Location is the system location. | [optional] 
**Name** | Pointer to **NullableString** | Name is the system name. | [optional] 
**ObjectId** | Pointer to **NullableString** | Object id is the system object id. | [optional] 

## Methods

### NewSnmpSysInfo

`func NewSnmpSysInfo() *SnmpSysInfo`

NewSnmpSysInfo instantiates a new SnmpSysInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpSysInfoWithDefaults

`func NewSnmpSysInfoWithDefaults() *SnmpSysInfo`

NewSnmpSysInfoWithDefaults instantiates a new SnmpSysInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *SnmpSysInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SnmpSysInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SnmpSysInfo) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SnmpSysInfo) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *SnmpSysInfo) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *SnmpSysInfo) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDescription

`func (o *SnmpSysInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnmpSysInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnmpSysInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnmpSysInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SnmpSysInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SnmpSysInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEngineIdType

`func (o *SnmpSysInfo) GetEngineIdType() int32`

GetEngineIdType returns the EngineIdType field if non-nil, zero value otherwise.

### GetEngineIdTypeOk

`func (o *SnmpSysInfo) GetEngineIdTypeOk() (*int32, bool)`

GetEngineIdTypeOk returns a tuple with the EngineIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineIdType

`func (o *SnmpSysInfo) SetEngineIdType(v int32)`

SetEngineIdType sets EngineIdType field to given value.

### HasEngineIdType

`func (o *SnmpSysInfo) HasEngineIdType() bool`

HasEngineIdType returns a boolean if a field has been set.

### SetEngineIdTypeNil

`func (o *SnmpSysInfo) SetEngineIdTypeNil(b bool)`

 SetEngineIdTypeNil sets the value for EngineIdType to be an explicit nil

### UnsetEngineIdType
`func (o *SnmpSysInfo) UnsetEngineIdType()`

UnsetEngineIdType ensures that no value is present for EngineIdType, not even an explicit nil
### GetLocation

`func (o *SnmpSysInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SnmpSysInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SnmpSysInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SnmpSysInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SnmpSysInfo) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SnmpSysInfo) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetName

`func (o *SnmpSysInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnmpSysInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnmpSysInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnmpSysInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SnmpSysInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SnmpSysInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectId

`func (o *SnmpSysInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *SnmpSysInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *SnmpSysInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *SnmpSysInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *SnmpSysInfo) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *SnmpSysInfo) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


