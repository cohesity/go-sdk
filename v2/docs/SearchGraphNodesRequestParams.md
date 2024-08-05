# SearchGraphNodesRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadParams** | Pointer to [**NullableAadGraphNodeFilterParams**](AadGraphNodeFilterParams.md) | Specifies the aad adapter specific node filter params. | [optional] 
**Name** | Pointer to **NullableString** | Filters the nodes based on provided current node display name. | [optional] 
**RootOnly** | Pointer to **NullableBool** | If set to true only root nodes would be returned. A root node refers to nodes in the graph with no incoming edges. Defaults to false. | [optional] 
**Count** | Pointer to **int32** | Specifies the number of objects to be fetched for the specified pagination cookie. | [optional] 
**IncludeAttributes** | Pointer to **NullableBool** | If set to false the response will only return name, type and is_root fields filled in each node. If set to true all the attributes for the nodes are also returned. Defaults to true. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies a cookie which can be passed in by the user in order to retrieve the next page of results. | [optional] 
**SessionId** | **string** | Specifies the id of a Session. | 

## Methods

### NewSearchGraphNodesRequestParams

`func NewSearchGraphNodesRequestParams(sessionId string, ) *SearchGraphNodesRequestParams`

NewSearchGraphNodesRequestParams instantiates a new SearchGraphNodesRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchGraphNodesRequestParamsWithDefaults

`func NewSearchGraphNodesRequestParamsWithDefaults() *SearchGraphNodesRequestParams`

NewSearchGraphNodesRequestParamsWithDefaults instantiates a new SearchGraphNodesRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadParams

`func (o *SearchGraphNodesRequestParams) GetAadParams() AadGraphNodeFilterParams`

GetAadParams returns the AadParams field if non-nil, zero value otherwise.

### GetAadParamsOk

`func (o *SearchGraphNodesRequestParams) GetAadParamsOk() (*AadGraphNodeFilterParams, bool)`

GetAadParamsOk returns a tuple with the AadParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadParams

`func (o *SearchGraphNodesRequestParams) SetAadParams(v AadGraphNodeFilterParams)`

SetAadParams sets AadParams field to given value.

### HasAadParams

`func (o *SearchGraphNodesRequestParams) HasAadParams() bool`

HasAadParams returns a boolean if a field has been set.

### SetAadParamsNil

`func (o *SearchGraphNodesRequestParams) SetAadParamsNil(b bool)`

 SetAadParamsNil sets the value for AadParams to be an explicit nil

### UnsetAadParams
`func (o *SearchGraphNodesRequestParams) UnsetAadParams()`

UnsetAadParams ensures that no value is present for AadParams, not even an explicit nil
### GetName

`func (o *SearchGraphNodesRequestParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchGraphNodesRequestParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchGraphNodesRequestParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchGraphNodesRequestParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SearchGraphNodesRequestParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SearchGraphNodesRequestParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRootOnly

`func (o *SearchGraphNodesRequestParams) GetRootOnly() bool`

GetRootOnly returns the RootOnly field if non-nil, zero value otherwise.

### GetRootOnlyOk

`func (o *SearchGraphNodesRequestParams) GetRootOnlyOk() (*bool, bool)`

GetRootOnlyOk returns a tuple with the RootOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootOnly

`func (o *SearchGraphNodesRequestParams) SetRootOnly(v bool)`

SetRootOnly sets RootOnly field to given value.

### HasRootOnly

`func (o *SearchGraphNodesRequestParams) HasRootOnly() bool`

HasRootOnly returns a boolean if a field has been set.

### SetRootOnlyNil

`func (o *SearchGraphNodesRequestParams) SetRootOnlyNil(b bool)`

 SetRootOnlyNil sets the value for RootOnly to be an explicit nil

### UnsetRootOnly
`func (o *SearchGraphNodesRequestParams) UnsetRootOnly()`

UnsetRootOnly ensures that no value is present for RootOnly, not even an explicit nil
### GetCount

`func (o *SearchGraphNodesRequestParams) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchGraphNodesRequestParams) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchGraphNodesRequestParams) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchGraphNodesRequestParams) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIncludeAttributes

`func (o *SearchGraphNodesRequestParams) GetIncludeAttributes() bool`

GetIncludeAttributes returns the IncludeAttributes field if non-nil, zero value otherwise.

### GetIncludeAttributesOk

`func (o *SearchGraphNodesRequestParams) GetIncludeAttributesOk() (*bool, bool)`

GetIncludeAttributesOk returns a tuple with the IncludeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttributes

`func (o *SearchGraphNodesRequestParams) SetIncludeAttributes(v bool)`

SetIncludeAttributes sets IncludeAttributes field to given value.

### HasIncludeAttributes

`func (o *SearchGraphNodesRequestParams) HasIncludeAttributes() bool`

HasIncludeAttributes returns a boolean if a field has been set.

### SetIncludeAttributesNil

`func (o *SearchGraphNodesRequestParams) SetIncludeAttributesNil(b bool)`

 SetIncludeAttributesNil sets the value for IncludeAttributes to be an explicit nil

### UnsetIncludeAttributes
`func (o *SearchGraphNodesRequestParams) UnsetIncludeAttributes()`

UnsetIncludeAttributes ensures that no value is present for IncludeAttributes, not even an explicit nil
### GetPaginationCookie

`func (o *SearchGraphNodesRequestParams) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *SearchGraphNodesRequestParams) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *SearchGraphNodesRequestParams) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *SearchGraphNodesRequestParams) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *SearchGraphNodesRequestParams) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *SearchGraphNodesRequestParams) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil
### GetSessionId

`func (o *SearchGraphNodesRequestParams) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SearchGraphNodesRequestParams) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SearchGraphNodesRequestParams) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


