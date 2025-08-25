

# CreateExtensionRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | A name for the extension. Each extension name in your account must be unique. Extension versions use the same name. |  |
|**description** | **String** | Information about the extension. |  [optional] |
|**actions** | **Map&lt;String, List&lt;Action&gt;&gt;** | The actions defined in the extension. |  |
|**parameters** | [**Map&lt;String, Parameter&gt;**](Parameter.md) | The parameters accepted by the extension. You specify parameter values when you associate the extension to an AppConfig resource by using the &lt;code&gt;CreateExtensionAssociation&lt;/code&gt; API action. For Lambda extension actions, these parameters are included in the Lambda request object. |  [optional] |
|**tags** | **Map&lt;String, String&gt;** | Adds one or more tags for the specified extension. Tags are metadata that help you categorize resources in different ways, for example, by purpose, owner, or environment. Each tag consists of a key and an optional value, both of which you define.  |  [optional] |



