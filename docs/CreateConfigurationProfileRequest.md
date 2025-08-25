

# CreateConfigurationProfileRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | A name for the configuration profile. |  |
|**description** | **String** | A description of the configuration profile. |  [optional] |
|**locationUri** | **String** | &lt;p&gt;A URI to locate the configuration. You can specify the following:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;For the AppConfig hosted configuration store and for feature flags, specify &lt;code&gt;hosted&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;For an Amazon Web Services Systems Manager Parameter Store parameter, specify either the parameter name in the format &lt;code&gt;ssm-parameter://&amp;lt;parameter name&amp;gt;&lt;/code&gt; or the ARN.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;For an Secrets Manager secret, specify the URI in the following format: &lt;code&gt;secrets-manager&lt;/code&gt;://&amp;lt;secret name&amp;gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;For an Amazon S3 object, specify the URI in the following format: &lt;code&gt;s3://&amp;lt;bucket&amp;gt;/&amp;lt;objectKey&amp;gt; &lt;/code&gt;. Here is an example: &lt;code&gt;s3://my-bucket/my-app/us-east-1/my-config.json&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;For an SSM document, specify either the document name in the format &lt;code&gt;ssm-document://&amp;lt;document name&amp;gt;&lt;/code&gt; or the Amazon Resource Name (ARN).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  |
|**retrievalRoleArn** | **String** | &lt;p&gt;The ARN of an IAM role with permission to access the configuration at the specified &lt;code&gt;LocationUri&lt;/code&gt;.&lt;/p&gt; &lt;important&gt; &lt;p&gt;A retrieval role ARN is not required for configurations stored in the AppConfig hosted configuration store. It is required for all other sources that store your configuration. &lt;/p&gt; &lt;/important&gt; |  [optional] |
|**validators** | [**List&lt;Validator&gt;**](Validator.md) | A list of methods for validating the configuration. |  [optional] |
|**tags** | **Map&lt;String, String&gt;** | Metadata to assign to the configuration profile. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define. |  [optional] |
|**type** | **String** | &lt;p&gt;The type of configurations contained in the profile. AppConfig supports &lt;code&gt;feature flags&lt;/code&gt; and &lt;code&gt;freeform&lt;/code&gt; configurations. We recommend you create feature flag configurations to enable or disable new features and freeform configurations to distribute configurations to an application. When calling this API, enter one of the following values for &lt;code&gt;Type&lt;/code&gt;:&lt;/p&gt; &lt;p&gt; &lt;code&gt;AWS.AppConfig.FeatureFlags&lt;/code&gt; &lt;/p&gt; &lt;p&gt; &lt;code&gt;AWS.Freeform&lt;/code&gt; &lt;/p&gt; |  [optional] |



