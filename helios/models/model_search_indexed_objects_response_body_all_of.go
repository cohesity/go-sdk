/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// SearchIndexedObjectsResponseBodyAllOf struct for SearchIndexedObjectsResponseBodyAllOf
type SearchIndexedObjectsResponseBodyAllOf struct {
	// Specifies the indexed emails and email folders.
	Emails *[]Email `json:"emails,omitempty"`
	// Specifies the indexed files.
	Files *[]File `json:"files,omitempty"`
	// Specifies the indexed Cassandra objects.
	CassandraObjects *[]CassandraIndexedObject `json:"cassandraObjects,omitempty"`
	// Specifies the indexed Couchbase objects.
	CouchbaseObjects *[]CouchbaseIndexedObject `json:"couchbaseObjects,omitempty"`
	// Specifies the indexed Hbase objects.
	HbaseObjects *[]HbaseIndexedObject `json:"hbaseObjects,omitempty"`
	// Specifies the indexed Hive objects.
	HiveObjects *[]HiveIndexedObject `json:"hiveObjects,omitempty"`
	// Specifies the indexed Mongo objects.
	MongoObjects *[]MongoIndexedObject `json:"mongoObjects,omitempty"`
	// Specifies the indexed HDFS objects.
	HdfsObjects *[]HDFSIndexedObject `json:"hdfsObjects,omitempty"`
	// Specifies the indexed HDFS objects.
	ExchangeObjects *[]ExchangeIndexedObject `json:"exchangeObjects,omitempty"`
	// Specifies the indexed Public folder items.
	PublicFolderItems *[]PublicFolderItem `json:"publicFolderItems,omitempty"`
	// Specifies the indexed Sharepoint items.
	SharepointItems *[]SharepointItem `json:"sharepointItems,omitempty"`
}

// NewSearchIndexedObjectsResponseBodyAllOf instantiates a new SearchIndexedObjectsResponseBodyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexedObjectsResponseBodyAllOf() *SearchIndexedObjectsResponseBodyAllOf {
	this := SearchIndexedObjectsResponseBodyAllOf{}
	return &this
}

// NewSearchIndexedObjectsResponseBodyAllOfWithDefaults instantiates a new SearchIndexedObjectsResponseBodyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexedObjectsResponseBodyAllOfWithDefaults() *SearchIndexedObjectsResponseBodyAllOf {
	this := SearchIndexedObjectsResponseBodyAllOf{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetEmails() []Email {
	if o == nil || o.Emails == nil {
		var ret []Email
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetEmailsOk() (*[]Email, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []Email and assigns it to the Emails field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetEmails(v []Email) {
	o.Emails = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetFiles() []File {
	if o == nil || o.Files == nil {
		var ret []File
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetFilesOk() (*[]File, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []File and assigns it to the Files field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetFiles(v []File) {
	o.Files = &v
}

// GetCassandraObjects returns the CassandraObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetCassandraObjects() []CassandraIndexedObject {
	if o == nil || o.CassandraObjects == nil {
		var ret []CassandraIndexedObject
		return ret
	}
	return *o.CassandraObjects
}

// GetCassandraObjectsOk returns a tuple with the CassandraObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetCassandraObjectsOk() (*[]CassandraIndexedObject, bool) {
	if o == nil || o.CassandraObjects == nil {
		return nil, false
	}
	return o.CassandraObjects, true
}

// HasCassandraObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasCassandraObjects() bool {
	if o != nil && o.CassandraObjects != nil {
		return true
	}

	return false
}

// SetCassandraObjects gets a reference to the given []CassandraIndexedObject and assigns it to the CassandraObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetCassandraObjects(v []CassandraIndexedObject) {
	o.CassandraObjects = &v
}

// GetCouchbaseObjects returns the CouchbaseObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetCouchbaseObjects() []CouchbaseIndexedObject {
	if o == nil || o.CouchbaseObjects == nil {
		var ret []CouchbaseIndexedObject
		return ret
	}
	return *o.CouchbaseObjects
}

// GetCouchbaseObjectsOk returns a tuple with the CouchbaseObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetCouchbaseObjectsOk() (*[]CouchbaseIndexedObject, bool) {
	if o == nil || o.CouchbaseObjects == nil {
		return nil, false
	}
	return o.CouchbaseObjects, true
}

// HasCouchbaseObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasCouchbaseObjects() bool {
	if o != nil && o.CouchbaseObjects != nil {
		return true
	}

	return false
}

// SetCouchbaseObjects gets a reference to the given []CouchbaseIndexedObject and assigns it to the CouchbaseObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetCouchbaseObjects(v []CouchbaseIndexedObject) {
	o.CouchbaseObjects = &v
}

// GetHbaseObjects returns the HbaseObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetHbaseObjects() []HbaseIndexedObject {
	if o == nil || o.HbaseObjects == nil {
		var ret []HbaseIndexedObject
		return ret
	}
	return *o.HbaseObjects
}

// GetHbaseObjectsOk returns a tuple with the HbaseObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetHbaseObjectsOk() (*[]HbaseIndexedObject, bool) {
	if o == nil || o.HbaseObjects == nil {
		return nil, false
	}
	return o.HbaseObjects, true
}

// HasHbaseObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasHbaseObjects() bool {
	if o != nil && o.HbaseObjects != nil {
		return true
	}

	return false
}

// SetHbaseObjects gets a reference to the given []HbaseIndexedObject and assigns it to the HbaseObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetHbaseObjects(v []HbaseIndexedObject) {
	o.HbaseObjects = &v
}

// GetHiveObjects returns the HiveObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetHiveObjects() []HiveIndexedObject {
	if o == nil || o.HiveObjects == nil {
		var ret []HiveIndexedObject
		return ret
	}
	return *o.HiveObjects
}

// GetHiveObjectsOk returns a tuple with the HiveObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetHiveObjectsOk() (*[]HiveIndexedObject, bool) {
	if o == nil || o.HiveObjects == nil {
		return nil, false
	}
	return o.HiveObjects, true
}

// HasHiveObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasHiveObjects() bool {
	if o != nil && o.HiveObjects != nil {
		return true
	}

	return false
}

// SetHiveObjects gets a reference to the given []HiveIndexedObject and assigns it to the HiveObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetHiveObjects(v []HiveIndexedObject) {
	o.HiveObjects = &v
}

// GetMongoObjects returns the MongoObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetMongoObjects() []MongoIndexedObject {
	if o == nil || o.MongoObjects == nil {
		var ret []MongoIndexedObject
		return ret
	}
	return *o.MongoObjects
}

// GetMongoObjectsOk returns a tuple with the MongoObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetMongoObjectsOk() (*[]MongoIndexedObject, bool) {
	if o == nil || o.MongoObjects == nil {
		return nil, false
	}
	return o.MongoObjects, true
}

// HasMongoObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasMongoObjects() bool {
	if o != nil && o.MongoObjects != nil {
		return true
	}

	return false
}

// SetMongoObjects gets a reference to the given []MongoIndexedObject and assigns it to the MongoObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetMongoObjects(v []MongoIndexedObject) {
	o.MongoObjects = &v
}

// GetHdfsObjects returns the HdfsObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetHdfsObjects() []HDFSIndexedObject {
	if o == nil || o.HdfsObjects == nil {
		var ret []HDFSIndexedObject
		return ret
	}
	return *o.HdfsObjects
}

// GetHdfsObjectsOk returns a tuple with the HdfsObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetHdfsObjectsOk() (*[]HDFSIndexedObject, bool) {
	if o == nil || o.HdfsObjects == nil {
		return nil, false
	}
	return o.HdfsObjects, true
}

// HasHdfsObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasHdfsObjects() bool {
	if o != nil && o.HdfsObjects != nil {
		return true
	}

	return false
}

// SetHdfsObjects gets a reference to the given []HDFSIndexedObject and assigns it to the HdfsObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetHdfsObjects(v []HDFSIndexedObject) {
	o.HdfsObjects = &v
}

// GetExchangeObjects returns the ExchangeObjects field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetExchangeObjects() []ExchangeIndexedObject {
	if o == nil || o.ExchangeObjects == nil {
		var ret []ExchangeIndexedObject
		return ret
	}
	return *o.ExchangeObjects
}

// GetExchangeObjectsOk returns a tuple with the ExchangeObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetExchangeObjectsOk() (*[]ExchangeIndexedObject, bool) {
	if o == nil || o.ExchangeObjects == nil {
		return nil, false
	}
	return o.ExchangeObjects, true
}

// HasExchangeObjects returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasExchangeObjects() bool {
	if o != nil && o.ExchangeObjects != nil {
		return true
	}

	return false
}

// SetExchangeObjects gets a reference to the given []ExchangeIndexedObject and assigns it to the ExchangeObjects field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetExchangeObjects(v []ExchangeIndexedObject) {
	o.ExchangeObjects = &v
}

// GetPublicFolderItems returns the PublicFolderItems field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetPublicFolderItems() []PublicFolderItem {
	if o == nil || o.PublicFolderItems == nil {
		var ret []PublicFolderItem
		return ret
	}
	return *o.PublicFolderItems
}

// GetPublicFolderItemsOk returns a tuple with the PublicFolderItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetPublicFolderItemsOk() (*[]PublicFolderItem, bool) {
	if o == nil || o.PublicFolderItems == nil {
		return nil, false
	}
	return o.PublicFolderItems, true
}

// HasPublicFolderItems returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasPublicFolderItems() bool {
	if o != nil && o.PublicFolderItems != nil {
		return true
	}

	return false
}

// SetPublicFolderItems gets a reference to the given []PublicFolderItem and assigns it to the PublicFolderItems field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetPublicFolderItems(v []PublicFolderItem) {
	o.PublicFolderItems = &v
}

// GetSharepointItems returns the SharepointItems field value if set, zero value otherwise.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetSharepointItems() []SharepointItem {
	if o == nil || o.SharepointItems == nil {
		var ret []SharepointItem
		return ret
	}
	return *o.SharepointItems
}

// GetSharepointItemsOk returns a tuple with the SharepointItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) GetSharepointItemsOk() (*[]SharepointItem, bool) {
	if o == nil || o.SharepointItems == nil {
		return nil, false
	}
	return o.SharepointItems, true
}

// HasSharepointItems returns a boolean if a field has been set.
func (o *SearchIndexedObjectsResponseBodyAllOf) HasSharepointItems() bool {
	if o != nil && o.SharepointItems != nil {
		return true
	}

	return false
}

// SetSharepointItems gets a reference to the given []SharepointItem and assigns it to the SharepointItems field.
func (o *SearchIndexedObjectsResponseBodyAllOf) SetSharepointItems(v []SharepointItem) {
	o.SharepointItems = &v
}

func (o SearchIndexedObjectsResponseBodyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.CassandraObjects != nil {
		toSerialize["cassandraObjects"] = o.CassandraObjects
	}
	if o.CouchbaseObjects != nil {
		toSerialize["couchbaseObjects"] = o.CouchbaseObjects
	}
	if o.HbaseObjects != nil {
		toSerialize["hbaseObjects"] = o.HbaseObjects
	}
	if o.HiveObjects != nil {
		toSerialize["hiveObjects"] = o.HiveObjects
	}
	if o.MongoObjects != nil {
		toSerialize["mongoObjects"] = o.MongoObjects
	}
	if o.HdfsObjects != nil {
		toSerialize["hdfsObjects"] = o.HdfsObjects
	}
	if o.ExchangeObjects != nil {
		toSerialize["exchangeObjects"] = o.ExchangeObjects
	}
	if o.PublicFolderItems != nil {
		toSerialize["publicFolderItems"] = o.PublicFolderItems
	}
	if o.SharepointItems != nil {
		toSerialize["sharepointItems"] = o.SharepointItems
	}
	return json.Marshal(toSerialize)
}

type NullableSearchIndexedObjectsResponseBodyAllOf struct {
	value *SearchIndexedObjectsResponseBodyAllOf
	isSet bool
}

func (v NullableSearchIndexedObjectsResponseBodyAllOf) Get() *SearchIndexedObjectsResponseBodyAllOf {
	return v.value
}

func (v *NullableSearchIndexedObjectsResponseBodyAllOf) Set(val *SearchIndexedObjectsResponseBodyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIndexedObjectsResponseBodyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIndexedObjectsResponseBodyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIndexedObjectsResponseBodyAllOf(val *SearchIndexedObjectsResponseBodyAllOf) *NullableSearchIndexedObjectsResponseBodyAllOf {
	return &NullableSearchIndexedObjectsResponseBodyAllOf{value: val, isSet: true}
}

func (v NullableSearchIndexedObjectsResponseBodyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIndexedObjectsResponseBodyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


