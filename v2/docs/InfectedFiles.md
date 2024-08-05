# InfectedFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookie** | Pointer to **NullableString** | Specifies the pagination cookie. Cookie is used to  resume the enumeration of infected entities. When the cookie is set the fields viewNameVec, includeQuarantinedFiles and include UnquarantinedFiles are ignored.  | [optional] 
**InfectedFiles** | Pointer to [**[]InfectedFile**](InfectedFile.md) | Specifies the list of infected entities. | [optional] 

## Methods

### NewInfectedFiles

`func NewInfectedFiles() *InfectedFiles`

NewInfectedFiles instantiates a new InfectedFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfectedFilesWithDefaults

`func NewInfectedFilesWithDefaults() *InfectedFiles`

NewInfectedFilesWithDefaults instantiates a new InfectedFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookie

`func (o *InfectedFiles) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *InfectedFiles) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *InfectedFiles) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *InfectedFiles) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### SetCookieNil

`func (o *InfectedFiles) SetCookieNil(b bool)`

 SetCookieNil sets the value for Cookie to be an explicit nil

### UnsetCookie
`func (o *InfectedFiles) UnsetCookie()`

UnsetCookie ensures that no value is present for Cookie, not even an explicit nil
### GetInfectedFiles

`func (o *InfectedFiles) GetInfectedFiles() []InfectedFile`

GetInfectedFiles returns the InfectedFiles field if non-nil, zero value otherwise.

### GetInfectedFilesOk

`func (o *InfectedFiles) GetInfectedFilesOk() (*[]InfectedFile, bool)`

GetInfectedFilesOk returns a tuple with the InfectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedFiles

`func (o *InfectedFiles) SetInfectedFiles(v []InfectedFile)`

SetInfectedFiles sets InfectedFiles field to given value.

### HasInfectedFiles

`func (o *InfectedFiles) HasInfectedFiles() bool`

HasInfectedFiles returns a boolean if a field has been set.

### SetInfectedFilesNil

`func (o *InfectedFiles) SetInfectedFilesNil(b bool)`

 SetInfectedFilesNil sets the value for InfectedFiles to be an explicit nil

### UnsetInfectedFiles
`func (o *InfectedFiles) UnsetInfectedFiles()`

UnsetInfectedFiles ensures that no value is present for InfectedFiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


