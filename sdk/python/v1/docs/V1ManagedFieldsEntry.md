# V1ManagedFieldsEntry

ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_version** | **str** | APIVersion defines the version of this resource that this field set applies to. The format is \&quot;group/version\&quot; just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted. | [optional] 
**fields_type** | **str** | FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \&quot;FieldsV1\&quot; | [optional] 
**fields_v1** | **object** | FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.  Each key is either a &#39;.&#39; representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: &#39;f:&lt;name&gt;&#39;, where &lt;name&gt; is the name of a field in a struct, or key in a map &#39;v:&lt;value&gt;&#39;, where &lt;value&gt; is the exact json formatted value of a list item &#39;i:&lt;index&gt;&#39;, where &lt;index&gt; is position of a item in a list &#39;k:&lt;keys&gt;&#39;, where &lt;keys&gt; is a map of  a list item&#39;s key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.  The exact format is defined in sigs.k8s.io/structured-merge-diff | [optional] 
**manager** | **str** | Manager is an identifier of the workflow managing these fields. | [optional] 
**operation** | **str** | Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are &#39;Apply&#39; and &#39;Update&#39;. | [optional] 
**subresource** | **str** | Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource. | [optional] 
**time** | **datetime** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Example

```python
from tatope.models.v1_managed_fields_entry import V1ManagedFieldsEntry

# TODO update the JSON string below
json = "{}"
# create an instance of V1ManagedFieldsEntry from a JSON string
v1_managed_fields_entry_instance = V1ManagedFieldsEntry.from_json(json)
# print the JSON string representation of the object
print(V1ManagedFieldsEntry.to_json())

# convert the object into a dict
v1_managed_fields_entry_dict = v1_managed_fields_entry_instance.to_dict()
# create an instance of V1ManagedFieldsEntry from a dict
v1_managed_fields_entry_from_dict = V1ManagedFieldsEntry.from_dict(v1_managed_fields_entry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


