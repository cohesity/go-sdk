# SharepointDocumentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | Pointer to **NullableString** | Specifies the type of site document(file/folder). Specifies the Sharepoint document type.  &#39;kFile&#39; specifies a file. &#39;kFolder&#39; specifies a folder. | [optional] 

## Methods

### NewSharepointDocumentMetadata

`func NewSharepointDocumentMetadata() *SharepointDocumentMetadata`

NewSharepointDocumentMetadata instantiates a new SharepointDocumentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharepointDocumentMetadataWithDefaults

`func NewSharepointDocumentMetadataWithDefaults() *SharepointDocumentMetadata`

NewSharepointDocumentMetadataWithDefaults instantiates a new SharepointDocumentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *SharepointDocumentMetadata) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *SharepointDocumentMetadata) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *SharepointDocumentMetadata) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *SharepointDocumentMetadata) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### SetDocumentTypeNil

`func (o *SharepointDocumentMetadata) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *SharepointDocumentMetadata) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


