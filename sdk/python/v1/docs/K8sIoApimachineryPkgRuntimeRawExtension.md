# K8sIoApimachineryPkgRuntimeRawExtension

RawExtension is used to hold extensions in external versions.  To use this, make a field which has RawExtension as its type in your external, versioned struct, and Object in your internal struct. You also need to register your various plugin types.  // Internal package:   type MyAPIObject struct {   runtime.TypeMeta `json:\",inline\"`   MyPlugin runtime.Object `json:\"myPlugin\"`  }   type PluginA struct {   AOption string `json:\"aOption\"`  }  // External package:   type MyAPIObject struct {   runtime.TypeMeta `json:\",inline\"`   MyPlugin runtime.RawExtension `json:\"myPlugin\"`  }   type PluginA struct {   AOption string `json:\"aOption\"`  }  // On the wire, the JSON will look something like this:   {   \"kind\":\"MyAPIObject\",   \"apiVersion\":\"v1\",   \"myPlugin\": {    \"kind\":\"PluginA\",    \"aOption\":\"foo\",   },  }  So what happens? Decode first uses json or yaml to unmarshal the serialized data into your external MyAPIObject. That causes the raw JSON to be stored, but not unpacked. The next step is to copy (using pkg/conversion) into the internal struct. The runtime package's DefaultScheme has conversion functions installed which will unpack the JSON stored in RawExtension, turning it into the correct object type, and storing it in the Object. (TODO: In the case where the object is of an unknown type, a runtime.Unknown object will be created and stored.)

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------

## Example

```python
from tatope.models.k8s_io_apimachinery_pkg_runtime_raw_extension import K8sIoApimachineryPkgRuntimeRawExtension

# TODO update the JSON string below
json = "{}"
# create an instance of K8sIoApimachineryPkgRuntimeRawExtension from a JSON string
k8s_io_apimachinery_pkg_runtime_raw_extension_instance = K8sIoApimachineryPkgRuntimeRawExtension.from_json(json)
# print the JSON string representation of the object
print(K8sIoApimachineryPkgRuntimeRawExtension.to_json())

# convert the object into a dict
k8s_io_apimachinery_pkg_runtime_raw_extension_dict = k8s_io_apimachinery_pkg_runtime_raw_extension_instance.to_dict()
# create an instance of K8sIoApimachineryPkgRuntimeRawExtension from a dict
k8s_io_apimachinery_pkg_runtime_raw_extension_from_dict = K8sIoApimachineryPkgRuntimeRawExtension.from_dict(k8s_io_apimachinery_pkg_runtime_raw_extension_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


