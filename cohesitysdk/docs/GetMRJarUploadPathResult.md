# GetMRJarUploadPathResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**JarUploadPath** | Pointer to **NullableString** | Path where Jars can be uploaded by Iris. | [optional] 

## Methods

### NewGetMRJarUploadPathResult

`func NewGetMRJarUploadPathResult() *GetMRJarUploadPathResult`

NewGetMRJarUploadPathResult instantiates a new GetMRJarUploadPathResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMRJarUploadPathResultWithDefaults

`func NewGetMRJarUploadPathResultWithDefaults() *GetMRJarUploadPathResult`

NewGetMRJarUploadPathResultWithDefaults instantiates a new GetMRJarUploadPathResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetMRJarUploadPathResult) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetMRJarUploadPathResult) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetMRJarUploadPathResult) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *GetMRJarUploadPathResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetJarUploadPath

`func (o *GetMRJarUploadPathResult) GetJarUploadPath() string`

GetJarUploadPath returns the JarUploadPath field if non-nil, zero value otherwise.

### GetJarUploadPathOk

`func (o *GetMRJarUploadPathResult) GetJarUploadPathOk() (*string, bool)`

GetJarUploadPathOk returns a tuple with the JarUploadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJarUploadPath

`func (o *GetMRJarUploadPathResult) SetJarUploadPath(v string)`

SetJarUploadPath sets JarUploadPath field to given value.

### HasJarUploadPath

`func (o *GetMRJarUploadPathResult) HasJarUploadPath() bool`

HasJarUploadPath returns a boolean if a field has been set.

### SetJarUploadPathNil

`func (o *GetMRJarUploadPathResult) SetJarUploadPathNil(b bool)`

 SetJarUploadPathNil sets the value for JarUploadPath to be an explicit nil

### UnsetJarUploadPath
`func (o *GetMRJarUploadPathResult) UnsetJarUploadPath()`

UnsetJarUploadPath ensures that no value is present for JarUploadPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


