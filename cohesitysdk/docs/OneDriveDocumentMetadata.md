# OneDriveDocumentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | Pointer to **NullableString** | Specifies the type of OneDrive document(file/folder). Specifies the OneDrive document type.  &#39;kFile&#39; specifies a file. &#39;kFolder&#39; specifies a folder. | [optional] 

## Methods

### NewOneDriveDocumentMetadata

`func NewOneDriveDocumentMetadata() *OneDriveDocumentMetadata`

NewOneDriveDocumentMetadata instantiates a new OneDriveDocumentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneDriveDocumentMetadataWithDefaults

`func NewOneDriveDocumentMetadataWithDefaults() *OneDriveDocumentMetadata`

NewOneDriveDocumentMetadataWithDefaults instantiates a new OneDriveDocumentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *OneDriveDocumentMetadata) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *OneDriveDocumentMetadata) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *OneDriveDocumentMetadata) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *OneDriveDocumentMetadata) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *OneDriveDocumentMetadata) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *OneDriveDocumentMetadata) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


