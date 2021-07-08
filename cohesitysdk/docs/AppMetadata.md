# AppMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **NullableString** | Specifies author of the app. | [optional] 
**CreatedDate** | Pointer to **NullableString** | Specifies date when the first version of the app was created. | [optional] 
**Description** | Pointer to **NullableString** | Specifies description about what app does. | [optional] 
**DevVersion** | Pointer to **NullableString** | Specifies version of the app provided by the developer. | [optional] 
**IconImage** | Pointer to **NullableString** | Specifies application icon. | [optional] 
**LastModifiedDate** | Pointer to **NullableString** | Specifies date when the app was last modified. | [optional] 
**Name** | Pointer to **NullableString** | Specifies name of the app. | [optional] 

## Methods

### NewAppMetadata

`func NewAppMetadata() *AppMetadata`

NewAppMetadata instantiates a new AppMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppMetadataWithDefaults

`func NewAppMetadataWithDefaults() *AppMetadata`

NewAppMetadataWithDefaults instantiates a new AppMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AppMetadata) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AppMetadata) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AppMetadata) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AppMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *AppMetadata) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *AppMetadata) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCreatedDate

`func (o *AppMetadata) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AppMetadata) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AppMetadata) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AppMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *AppMetadata) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *AppMetadata) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetDescription

`func (o *AppMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppMetadata) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppMetadata) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDevVersion

`func (o *AppMetadata) GetDevVersion() string`

GetDevVersion returns the DevVersion field if non-nil, zero value otherwise.

### GetDevVersionOk

`func (o *AppMetadata) GetDevVersionOk() (*string, bool)`

GetDevVersionOk returns a tuple with the DevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevVersion

`func (o *AppMetadata) SetDevVersion(v string)`

SetDevVersion sets DevVersion field to given value.

### HasDevVersion

`func (o *AppMetadata) HasDevVersion() bool`

HasDevVersion returns a boolean if a field has been set.

### SetDevVersionNil

`func (o *AppMetadata) SetDevVersionNil(b bool)`

 SetDevVersionNil sets the value for DevVersion to be an explicit nil

### UnsetDevVersion
`func (o *AppMetadata) UnsetDevVersion()`

UnsetDevVersion ensures that no value is present for DevVersion, not even an explicit nil
### GetIconImage

`func (o *AppMetadata) GetIconImage() string`

GetIconImage returns the IconImage field if non-nil, zero value otherwise.

### GetIconImageOk

`func (o *AppMetadata) GetIconImageOk() (*string, bool)`

GetIconImageOk returns a tuple with the IconImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconImage

`func (o *AppMetadata) SetIconImage(v string)`

SetIconImage sets IconImage field to given value.

### HasIconImage

`func (o *AppMetadata) HasIconImage() bool`

HasIconImage returns a boolean if a field has been set.

### SetIconImageNil

`func (o *AppMetadata) SetIconImageNil(b bool)`

 SetIconImageNil sets the value for IconImage to be an explicit nil

### UnsetIconImage
`func (o *AppMetadata) UnsetIconImage()`

UnsetIconImage ensures that no value is present for IconImage, not even an explicit nil
### GetLastModifiedDate

`func (o *AppMetadata) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *AppMetadata) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *AppMetadata) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *AppMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### SetLastModifiedDateNil

`func (o *AppMetadata) SetLastModifiedDateNil(b bool)`

 SetLastModifiedDateNil sets the value for LastModifiedDate to be an explicit nil

### UnsetLastModifiedDate
`func (o *AppMetadata) UnsetLastModifiedDate()`

UnsetLastModifiedDate ensures that no value is present for LastModifiedDate, not even an explicit nil
### GetName

`func (o *AppMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppMetadata) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppMetadata) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


