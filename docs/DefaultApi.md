# DefaultApi

All URIs are relative to *http://appconfig.us-east-1.amazonaws.com*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createApplication**](DefaultApi.md#createApplication) | **POST** /applications |  |
| [**createConfigurationProfile**](DefaultApi.md#createConfigurationProfile) | **POST** /applications/{ApplicationId}/configurationprofiles |  |
| [**createDeploymentStrategy**](DefaultApi.md#createDeploymentStrategy) | **POST** /deploymentstrategies |  |
| [**createEnvironment**](DefaultApi.md#createEnvironment) | **POST** /applications/{ApplicationId}/environments |  |
| [**createExtension**](DefaultApi.md#createExtension) | **POST** /extensions |  |
| [**createExtensionAssociation**](DefaultApi.md#createExtensionAssociation) | **POST** /extensionassociations |  |
| [**createHostedConfigurationVersion**](DefaultApi.md#createHostedConfigurationVersion) | **POST** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/hostedconfigurationversions#Content-Type |  |
| [**deleteApplication**](DefaultApi.md#deleteApplication) | **DELETE** /applications/{ApplicationId} |  |
| [**deleteConfigurationProfile**](DefaultApi.md#deleteConfigurationProfile) | **DELETE** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId} |  |
| [**deleteDeploymentStrategy**](DefaultApi.md#deleteDeploymentStrategy) | **DELETE** /deployementstrategies/{DeploymentStrategyId} |  |
| [**deleteEnvironment**](DefaultApi.md#deleteEnvironment) | **DELETE** /applications/{ApplicationId}/environments/{EnvironmentId} |  |
| [**deleteExtension**](DefaultApi.md#deleteExtension) | **DELETE** /extensions/{ExtensionIdentifier} |  |
| [**deleteExtensionAssociation**](DefaultApi.md#deleteExtensionAssociation) | **DELETE** /extensionassociations/{ExtensionAssociationId} |  |
| [**deleteHostedConfigurationVersion**](DefaultApi.md#deleteHostedConfigurationVersion) | **DELETE** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/hostedconfigurationversions/{VersionNumber} |  |
| [**getApplication**](DefaultApi.md#getApplication) | **GET** /applications/{ApplicationId} |  |
| [**getConfiguration**](DefaultApi.md#getConfiguration) | **GET** /applications/{Application}/environments/{Environment}/configurations/{Configuration}#client_id |  |
| [**getConfigurationProfile**](DefaultApi.md#getConfigurationProfile) | **GET** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId} |  |
| [**getDeployment**](DefaultApi.md#getDeployment) | **GET** /applications/{ApplicationId}/environments/{EnvironmentId}/deployments/{DeploymentNumber} |  |
| [**getDeploymentStrategy**](DefaultApi.md#getDeploymentStrategy) | **GET** /deploymentstrategies/{DeploymentStrategyId} |  |
| [**getEnvironment**](DefaultApi.md#getEnvironment) | **GET** /applications/{ApplicationId}/environments/{EnvironmentId} |  |
| [**getExtension**](DefaultApi.md#getExtension) | **GET** /extensions/{ExtensionIdentifier} |  |
| [**getExtensionAssociation**](DefaultApi.md#getExtensionAssociation) | **GET** /extensionassociations/{ExtensionAssociationId} |  |
| [**getHostedConfigurationVersion**](DefaultApi.md#getHostedConfigurationVersion) | **GET** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/hostedconfigurationversions/{VersionNumber} |  |
| [**listApplications**](DefaultApi.md#listApplications) | **GET** /applications |  |
| [**listConfigurationProfiles**](DefaultApi.md#listConfigurationProfiles) | **GET** /applications/{ApplicationId}/configurationprofiles |  |
| [**listDeploymentStrategies**](DefaultApi.md#listDeploymentStrategies) | **GET** /deploymentstrategies |  |
| [**listDeployments**](DefaultApi.md#listDeployments) | **GET** /applications/{ApplicationId}/environments/{EnvironmentId}/deployments |  |
| [**listEnvironments**](DefaultApi.md#listEnvironments) | **GET** /applications/{ApplicationId}/environments |  |
| [**listExtensionAssociations**](DefaultApi.md#listExtensionAssociations) | **GET** /extensionassociations |  |
| [**listExtensions**](DefaultApi.md#listExtensions) | **GET** /extensions |  |
| [**listHostedConfigurationVersions**](DefaultApi.md#listHostedConfigurationVersions) | **GET** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/hostedconfigurationversions |  |
| [**listTagsForResource**](DefaultApi.md#listTagsForResource) | **GET** /tags/{ResourceArn} |  |
| [**startDeployment**](DefaultApi.md#startDeployment) | **POST** /applications/{ApplicationId}/environments/{EnvironmentId}/deployments |  |
| [**stopDeployment**](DefaultApi.md#stopDeployment) | **DELETE** /applications/{ApplicationId}/environments/{EnvironmentId}/deployments/{DeploymentNumber} |  |
| [**tagResource**](DefaultApi.md#tagResource) | **POST** /tags/{ResourceArn} |  |
| [**untagResource**](DefaultApi.md#untagResource) | **DELETE** /tags/{ResourceArn}#tagKeys |  |
| [**updateApplication**](DefaultApi.md#updateApplication) | **PATCH** /applications/{ApplicationId} |  |
| [**updateConfigurationProfile**](DefaultApi.md#updateConfigurationProfile) | **PATCH** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId} |  |
| [**updateDeploymentStrategy**](DefaultApi.md#updateDeploymentStrategy) | **PATCH** /deploymentstrategies/{DeploymentStrategyId} |  |
| [**updateEnvironment**](DefaultApi.md#updateEnvironment) | **PATCH** /applications/{ApplicationId}/environments/{EnvironmentId} |  |
| [**updateExtension**](DefaultApi.md#updateExtension) | **PATCH** /extensions/{ExtensionIdentifier} |  |
| [**updateExtensionAssociation**](DefaultApi.md#updateExtensionAssociation) | **PATCH** /extensionassociations/{ExtensionAssociationId} |  |
| [**validateConfiguration**](DefaultApi.md#validateConfiguration) | **POST** /applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/validators#configuration_version |  |


<a id="createApplication"></a>
# **createApplication**
> Application createApplication(createApplicationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Creates an application. In AppConfig, an application is simply an organizational construct like a folder. This organizational construct has a relationship with some unit of executable code. For example, you could create an application called MyMobileApp to organize and manage configuration data for a mobile application installed by your users.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    CreateApplicationRequest createApplicationRequest = new CreateApplicationRequest(); // CreateApplicationRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Application result = apiInstance.createApplication(createApplicationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createApplication");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createApplicationRequest** | [**CreateApplicationRequest**](CreateApplicationRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Application**](Application.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | InternalServerException |  -  |

<a id="createConfigurationProfile"></a>
# **createConfigurationProfile**
> ConfigurationProfile createConfigurationProfile(applicationId, createConfigurationProfileRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Creates a configuration profile, which is information that enables AppConfig to access the configuration source. Valid configuration sources include the following:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Configuration data in YAML, JSON, and other formats stored in the AppConfig hosted configuration store&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Configuration data stored as objects in an Amazon Simple Storage Service (Amazon S3) bucket&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Pipelines stored in CodePipeline&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Secrets stored in Secrets Manager&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Standard and secure string parameters stored in Amazon Web Services Systems Manager Parameter Store&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Configuration data in SSM documents stored in the Systems Manager document store&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt;A configuration profile includes the following information:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;The URI location of the configuration data.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;The Identity and Access Management (IAM) role that provides access to the configuration data.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;A validator for the configuration data. Available validators include either a JSON Schema or an Amazon Web Services Lambda function.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt;For more information, see &lt;a href&#x3D;\&quot;http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile.html\&quot;&gt;Create a Configuration and a Configuration Profile&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.&lt;/p&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    CreateConfigurationProfileRequest createConfigurationProfileRequest = new CreateConfigurationProfileRequest(); // CreateConfigurationProfileRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ConfigurationProfile result = apiInstance.createConfigurationProfile(applicationId, createConfigurationProfileRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createConfigurationProfile");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **createConfigurationProfileRequest** | [**CreateConfigurationProfileRequest**](CreateConfigurationProfileRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ConfigurationProfile**](ConfigurationProfile.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="createDeploymentStrategy"></a>
# **createDeploymentStrategy**
> DeploymentStrategy createDeploymentStrategy(createDeploymentStrategyRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Creates a deployment strategy that defines important criteria for rolling out your configuration to the designated targets. A deployment strategy includes the overall duration required, a percentage of targets to receive the deployment during each interval, an algorithm that defines how percentage grows, and bake time.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    CreateDeploymentStrategyRequest createDeploymentStrategyRequest = new CreateDeploymentStrategyRequest(); // CreateDeploymentStrategyRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      DeploymentStrategy result = apiInstance.createDeploymentStrategy(createDeploymentStrategyRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createDeploymentStrategy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createDeploymentStrategyRequest** | [**CreateDeploymentStrategyRequest**](CreateDeploymentStrategyRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**DeploymentStrategy**](DeploymentStrategy.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | InternalServerException |  -  |
| **481** | BadRequestException |  -  |

<a id="createEnvironment"></a>
# **createEnvironment**
> Environment createEnvironment(applicationId, createEnvironmentRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Creates an environment. For each application, you define one or more environments. An environment is a deployment group of AppConfig targets, such as applications in a &lt;code&gt;Beta&lt;/code&gt; or &lt;code&gt;Production&lt;/code&gt; environment. You can also define environments for application subcomponents such as the &lt;code&gt;Web&lt;/code&gt;, &lt;code&gt;Mobile&lt;/code&gt; and &lt;code&gt;Back-end&lt;/code&gt; components for your application. You can configure Amazon CloudWatch alarms for each environment. The system monitors alarms during a configuration deployment. If an alarm is triggered, the system rolls back the configuration.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    CreateEnvironmentRequest createEnvironmentRequest = new CreateEnvironmentRequest(); // CreateEnvironmentRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Environment result = apiInstance.createEnvironment(applicationId, createEnvironmentRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createEnvironment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **createEnvironmentRequest** | [**CreateEnvironmentRequest**](CreateEnvironmentRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Environment**](Environment.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | InternalServerException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | BadRequestException |  -  |

<a id="createExtension"></a>
# **createExtension**
> Extension createExtension(createExtensionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, latestVersionNumber)



&lt;p&gt;Creates an AppConfig extension. An extension augments your ability to inject logic or behavior at different points during the AppConfig workflow of creating or deploying a configuration.&lt;/p&gt; &lt;p&gt;You can create your own extensions or use the Amazon Web Services authored extensions provided by AppConfig. For most use cases, to create your own extension, you must create an Lambda function to perform any computation and processing defined in the extension. For more information about extensions, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.&lt;/p&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    CreateExtensionRequest createExtensionRequest = new CreateExtensionRequest(); // CreateExtensionRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer latestVersionNumber = 56; // Integer | You can omit this field when you create an extension. When you create a new version, specify the most recent current version number. For example, you create version 3, enter 2 for this field.
    try {
      Extension result = apiInstance.createExtension(createExtensionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, latestVersionNumber);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createExtension");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createExtensionRequest** | [**CreateExtensionRequest**](CreateExtensionRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **latestVersionNumber** | **Integer**| You can omit this field when you create an extension. When you create a new version, specify the most recent current version number. For example, you create version 3, enter 2 for this field. | [optional] |

### Return type

[**Extension**](Extension.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ConflictException |  -  |
| **482** | ServiceQuotaExceededException |  -  |
| **483** | InternalServerException |  -  |

<a id="createExtensionAssociation"></a>
# **createExtensionAssociation**
> ExtensionAssociation createExtensionAssociation(createExtensionAssociationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



When you create an extension or configure an Amazon Web Services authored extension, you associate the extension with an AppConfig application, environment, or configuration profile. For example, you can choose to run the &lt;code&gt;AppConfig deployment events to Amazon SNS&lt;/code&gt; Amazon Web Services authored extension and receive notifications on an Amazon SNS topic anytime a configuration deployment is started for a specific application. Defining which extension to associate with an AppConfig resource is called an &lt;i&gt;extension association&lt;/i&gt;. An extension association is a specified relationship between an extension and an AppConfig resource, such as an application or a configuration profile. For more information about extensions and associations, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    CreateExtensionAssociationRequest createExtensionAssociationRequest = new CreateExtensionAssociationRequest(); // CreateExtensionAssociationRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ExtensionAssociation result = apiInstance.createExtensionAssociation(createExtensionAssociationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createExtensionAssociation");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createExtensionAssociationRequest** | [**CreateExtensionAssociationRequest**](CreateExtensionAssociationRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ExtensionAssociation**](ExtensionAssociation.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |
| **483** | ServiceQuotaExceededException |  -  |

<a id="createHostedConfigurationVersion"></a>
# **createHostedConfigurationVersion**
> HostedConfigurationVersion createHostedConfigurationVersion(applicationId, configurationProfileId, contentType, createHostedConfigurationVersionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, description, latestVersionNumber, versionLabel)



Creates a new configuration in the AppConfig hosted configuration store.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String configurationProfileId = "configurationProfileId_example"; // String | The configuration profile ID.
    String contentType = "contentType_example"; // String | A standard MIME type describing the format of the configuration content. For more information, see <a href=\"https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17\">Content-Type</a>.
    CreateHostedConfigurationVersionRequest createHostedConfigurationVersionRequest = new CreateHostedConfigurationVersionRequest(); // CreateHostedConfigurationVersionRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String description = "description_example"; // String | A description of the configuration.
    Integer latestVersionNumber = 56; // Integer | An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version. To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.
    String versionLabel = "versionLabel_example"; // String | An optional, user-defined label for the AppConfig hosted configuration version. This value must contain at least one non-numeric character. For example, \"v2.2.0\".
    try {
      HostedConfigurationVersion result = apiInstance.createHostedConfigurationVersion(applicationId, configurationProfileId, contentType, createHostedConfigurationVersionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, description, latestVersionNumber, versionLabel);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createHostedConfigurationVersion");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **configurationProfileId** | **String**| The configuration profile ID. | |
| **contentType** | **String**| A standard MIME type describing the format of the configuration content. For more information, see &lt;a href&#x3D;\&quot;https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17\&quot;&gt;Content-Type&lt;/a&gt;. | |
| **createHostedConfigurationVersionRequest** | [**CreateHostedConfigurationVersionRequest**](CreateHostedConfigurationVersionRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **description** | **String**| A description of the configuration. | [optional] |
| **latestVersionNumber** | **Integer**| An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version. To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version. | [optional] |
| **versionLabel** | **String**| An optional, user-defined label for the AppConfig hosted configuration version. This value must contain at least one non-numeric character. For example, \&quot;v2.2.0\&quot;. | [optional] |

### Return type

[**HostedConfigurationVersion**](HostedConfigurationVersion.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ServiceQuotaExceededException |  -  |
| **482** | ResourceNotFoundException |  -  |
| **483** | ConflictException |  -  |
| **484** | PayloadTooLargeException |  -  |
| **485** | InternalServerException |  -  |

<a id="deleteApplication"></a>
# **deleteApplication**
> deleteApplication(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes an application. Deleting an application does not delete a configuration from a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The ID of the application to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.deleteApplication(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteApplication");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The ID of the application to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="deleteConfigurationProfile"></a>
# **deleteConfigurationProfile**
> deleteConfigurationProfile(applicationId, configurationProfileId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes a configuration profile. Deleting a configuration profile does not delete a configuration from a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID that includes the configuration profile you want to delete.
    String configurationProfileId = "configurationProfileId_example"; // String | The ID of the configuration profile you want to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.deleteConfigurationProfile(applicationId, configurationProfileId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteConfigurationProfile");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID that includes the configuration profile you want to delete. | |
| **configurationProfileId** | **String**| The ID of the configuration profile you want to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | ConflictException |  -  |
| **482** | InternalServerException |  -  |
| **483** | BadRequestException |  -  |

<a id="deleteDeploymentStrategy"></a>
# **deleteDeploymentStrategy**
> deleteDeploymentStrategy(deploymentStrategyId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes a deployment strategy. Deleting a deployment strategy does not delete a configuration from a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String deploymentStrategyId = "deploymentStrategyId_example"; // String | The ID of the deployment strategy you want to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.deleteDeploymentStrategy(deploymentStrategyId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteDeploymentStrategy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentStrategyId** | **String**| The ID of the deployment strategy you want to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="deleteEnvironment"></a>
# **deleteEnvironment**
> deleteEnvironment(applicationId, environmentId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes an environment. Deleting an environment does not delete a configuration from a host.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID that includes the environment that you want to delete.
    String environmentId = "environmentId_example"; // String | The ID of the environment that you want to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.deleteEnvironment(applicationId, environmentId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteEnvironment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID that includes the environment that you want to delete. | |
| **environmentId** | **String**| The ID of the environment that you want to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | ConflictException |  -  |
| **482** | InternalServerException |  -  |
| **483** | BadRequestException |  -  |

<a id="deleteExtension"></a>
# **deleteExtension**
> deleteExtension(extensionIdentifier, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, version)



Deletes an AppConfig extension. You must delete all associations to an extension before you delete the extension.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String extensionIdentifier = "extensionIdentifier_example"; // String | The name, ID, or Amazon Resource Name (ARN) of the extension you want to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer version = 56; // Integer | A specific version of an extension to delete. If omitted, the highest version is deleted.
    try {
      apiInstance.deleteExtension(extensionIdentifier, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, version);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteExtension");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **extensionIdentifier** | **String**| The name, ID, or Amazon Resource Name (ARN) of the extension you want to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **version** | **Integer**| A specific version of an extension to delete. If omitted, the highest version is deleted. | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="deleteExtensionAssociation"></a>
# **deleteExtensionAssociation**
> deleteExtensionAssociation(extensionAssociationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes an extension association. This action doesn&#39;t delete extensions defined in the association.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String extensionAssociationId = "extensionAssociationId_example"; // String | The ID of the extension association to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.deleteExtensionAssociation(extensionAssociationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteExtensionAssociation");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **extensionAssociationId** | **String**| The ID of the extension association to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="deleteHostedConfigurationVersion"></a>
# **deleteHostedConfigurationVersion**
> deleteHostedConfigurationVersion(applicationId, configurationProfileId, versionNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes a version of a configuration from the AppConfig hosted configuration store.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String configurationProfileId = "configurationProfileId_example"; // String | The configuration profile ID.
    Integer versionNumber = 56; // Integer | The versions number to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.deleteHostedConfigurationVersion(applicationId, configurationProfileId, versionNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteHostedConfigurationVersion");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **configurationProfileId** | **String**| The configuration profile ID. | |
| **versionNumber** | **Integer**| The versions number to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="getApplication"></a>
# **getApplication**
> Application getApplication(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves information about an application.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The ID of the application you want to get.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Application result = apiInstance.getApplication(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getApplication");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The ID of the application you want to get. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Application**](Application.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getConfiguration"></a>
# **getConfiguration**
> ModelConfiguration getConfiguration(application, environment, _configuration, clientId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, clientConfigurationVersion)



&lt;p&gt;(Deprecated) Retrieves the latest deployed configuration.&lt;/p&gt; &lt;important&gt; &lt;p&gt;Note the following important information.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;This API action is deprecated. Calls to receive configuration data should use the &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_StartConfigurationSession.html\&quot;&gt;StartConfigurationSession&lt;/a&gt; and &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html\&quot;&gt;GetLatestConfiguration&lt;/a&gt; APIs instead. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;GetConfiguration&lt;/code&gt; is a priced call. For more information, see &lt;a href&#x3D;\&quot;https://aws.amazon.com/systems-manager/pricing/\&quot;&gt;Pricing&lt;/a&gt;.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;/important&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String application = "application_example"; // String | The application to get. Specify either the application name or the application ID.
    String environment = "environment_example"; // String | The environment to get. Specify either the environment name or the environment ID.
    String _configuration = "_configuration_example"; // String | The configuration to get. Specify either the configuration name or the configuration ID.
    String clientId = "clientId_example"; // String | The clientId parameter in the following command is a unique, user-specified ID to identify the client for the configuration. This ID enables AppConfig to deploy the configuration in intervals, as defined in the deployment strategy. 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String clientConfigurationVersion = "clientConfigurationVersion_example"; // String | <p>The configuration version returned in the most recent <code>GetConfiguration</code> response.</p> <important> <p>AppConfig uses the value of the <code>ClientConfigurationVersion</code> parameter to identify the configuration version on your clients. If you don’t send <code>ClientConfigurationVersion</code> with each call to <code>GetConfiguration</code>, your clients receive the current configuration. You are charged each time your clients receive a configuration.</p> <p>To avoid excess charges, we recommend you use the <a href=\"https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/StartConfigurationSession.html\">StartConfigurationSession</a> and <a href=\"https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/GetLatestConfiguration.html\">GetLatestConfiguration</a> APIs, which track the client configuration version on your behalf. If you choose to continue using <code>GetConfiguration</code>, we recommend that you include the <code>ClientConfigurationVersion</code> value with every call to <code>GetConfiguration</code>. The value to use for <code>ClientConfigurationVersion</code> comes from the <code>ConfigurationVersion</code> attribute returned by <code>GetConfiguration</code> when there is new or updated data, and should be saved for subsequent calls to <code>GetConfiguration</code>.</p> </important> <p>For more information about working with configurations, see <a href=\"http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration.html\">Retrieving the Configuration</a> in the <i>AppConfig User Guide</i>.</p>
    try {
      ModelConfiguration result = apiInstance.getConfiguration(application, environment, _configuration, clientId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, clientConfigurationVersion);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getConfiguration");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **application** | **String**| The application to get. Specify either the application name or the application ID. | |
| **environment** | **String**| The environment to get. Specify either the environment name or the environment ID. | |
| **_configuration** | **String**| The configuration to get. Specify either the configuration name or the configuration ID. | |
| **clientId** | **String**| The clientId parameter in the following command is a unique, user-specified ID to identify the client for the configuration. This ID enables AppConfig to deploy the configuration in intervals, as defined in the deployment strategy.  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **clientConfigurationVersion** | **String**| &lt;p&gt;The configuration version returned in the most recent &lt;code&gt;GetConfiguration&lt;/code&gt; response.&lt;/p&gt; &lt;important&gt; &lt;p&gt;AppConfig uses the value of the &lt;code&gt;ClientConfigurationVersion&lt;/code&gt; parameter to identify the configuration version on your clients. If you don’t send &lt;code&gt;ClientConfigurationVersion&lt;/code&gt; with each call to &lt;code&gt;GetConfiguration&lt;/code&gt;, your clients receive the current configuration. You are charged each time your clients receive a configuration.&lt;/p&gt; &lt;p&gt;To avoid excess charges, we recommend you use the &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/StartConfigurationSession.html\&quot;&gt;StartConfigurationSession&lt;/a&gt; and &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/GetLatestConfiguration.html\&quot;&gt;GetLatestConfiguration&lt;/a&gt; APIs, which track the client configuration version on your behalf. If you choose to continue using &lt;code&gt;GetConfiguration&lt;/code&gt;, we recommend that you include the &lt;code&gt;ClientConfigurationVersion&lt;/code&gt; value with every call to &lt;code&gt;GetConfiguration&lt;/code&gt;. The value to use for &lt;code&gt;ClientConfigurationVersion&lt;/code&gt; comes from the &lt;code&gt;ConfigurationVersion&lt;/code&gt; attribute returned by &lt;code&gt;GetConfiguration&lt;/code&gt; when there is new or updated data, and should be saved for subsequent calls to &lt;code&gt;GetConfiguration&lt;/code&gt;.&lt;/p&gt; &lt;/important&gt; &lt;p&gt;For more information about working with configurations, see &lt;a href&#x3D;\&quot;http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration.html\&quot;&gt;Retrieving the Configuration&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.&lt;/p&gt; | [optional] |

### Return type

[**ModelConfiguration**](ModelConfiguration.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getConfigurationProfile"></a>
# **getConfigurationProfile**
> ConfigurationProfile getConfigurationProfile(applicationId, configurationProfileId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves information about a configuration profile.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The ID of the application that includes the configuration profile you want to get.
    String configurationProfileId = "configurationProfileId_example"; // String | The ID of the configuration profile that you want to get.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ConfigurationProfile result = apiInstance.getConfigurationProfile(applicationId, configurationProfileId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getConfigurationProfile");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The ID of the application that includes the configuration profile you want to get. | |
| **configurationProfileId** | **String**| The ID of the configuration profile that you want to get. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ConfigurationProfile**](ConfigurationProfile.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getDeployment"></a>
# **getDeployment**
> Deployment getDeployment(applicationId, environmentId, deploymentNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves information about a configuration deployment.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The ID of the application that includes the deployment you want to get. 
    String environmentId = "environmentId_example"; // String | The ID of the environment that includes the deployment you want to get. 
    Integer deploymentNumber = 56; // Integer | The sequence number of the deployment.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Deployment result = apiInstance.getDeployment(applicationId, environmentId, deploymentNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getDeployment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The ID of the application that includes the deployment you want to get.  | |
| **environmentId** | **String**| The ID of the environment that includes the deployment you want to get.  | |
| **deploymentNumber** | **Integer**| The sequence number of the deployment. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Deployment**](Deployment.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getDeploymentStrategy"></a>
# **getDeploymentStrategy**
> DeploymentStrategy getDeploymentStrategy(deploymentStrategyId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves information about a deployment strategy. A deployment strategy defines important criteria for rolling out your configuration to the designated targets. A deployment strategy includes the overall duration required, a percentage of targets to receive the deployment during each interval, an algorithm that defines how percentage grows, and bake time.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String deploymentStrategyId = "deploymentStrategyId_example"; // String | The ID of the deployment strategy to get.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      DeploymentStrategy result = apiInstance.getDeploymentStrategy(deploymentStrategyId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getDeploymentStrategy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentStrategyId** | **String**| The ID of the deployment strategy to get. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**DeploymentStrategy**](DeploymentStrategy.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getEnvironment"></a>
# **getEnvironment**
> Environment getEnvironment(applicationId, environmentId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves information about an environment. An environment is a deployment group of AppConfig applications, such as applications in a &lt;code&gt;Production&lt;/code&gt; environment or in an &lt;code&gt;EU_Region&lt;/code&gt; environment. Each configuration deployment targets an environment. You can enable one or more Amazon CloudWatch alarms for an environment. If an alarm is triggered during a deployment, AppConfig roles back the configuration.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The ID of the application that includes the environment you want to get.
    String environmentId = "environmentId_example"; // String | The ID of the environment that you want to get.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Environment result = apiInstance.getEnvironment(applicationId, environmentId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getEnvironment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The ID of the application that includes the environment you want to get. | |
| **environmentId** | **String**| The ID of the environment that you want to get. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Environment**](Environment.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getExtension"></a>
# **getExtension**
> Extension getExtension(extensionIdentifier, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, versionNumber)



Returns information about an AppConfig extension.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String extensionIdentifier = "extensionIdentifier_example"; // String | The name, the ID, or the Amazon Resource Name (ARN) of the extension.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer versionNumber = 56; // Integer | The extension version number. If no version number was defined, AppConfig uses the highest version.
    try {
      Extension result = apiInstance.getExtension(extensionIdentifier, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, versionNumber);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getExtension");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **extensionIdentifier** | **String**| The name, the ID, or the Amazon Resource Name (ARN) of the extension. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **versionNumber** | **Integer**| The extension version number. If no version number was defined, AppConfig uses the highest version. | [optional] |

### Return type

[**Extension**](Extension.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="getExtensionAssociation"></a>
# **getExtensionAssociation**
> ExtensionAssociation getExtensionAssociation(extensionAssociationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Returns information about an AppConfig extension association. For more information about extensions and associations, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String extensionAssociationId = "extensionAssociationId_example"; // String | The extension association ID to get.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ExtensionAssociation result = apiInstance.getExtensionAssociation(extensionAssociationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getExtensionAssociation");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **extensionAssociationId** | **String**| The extension association ID to get. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ExtensionAssociation**](ExtensionAssociation.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="getHostedConfigurationVersion"></a>
# **getHostedConfigurationVersion**
> HostedConfigurationVersion getHostedConfigurationVersion(applicationId, configurationProfileId, versionNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves information about a specific configuration version.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String configurationProfileId = "configurationProfileId_example"; // String | The configuration profile ID.
    Integer versionNumber = 56; // Integer | The version.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      HostedConfigurationVersion result = apiInstance.getHostedConfigurationVersion(applicationId, configurationProfileId, versionNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#getHostedConfigurationVersion");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **configurationProfileId** | **String**| The configuration profile ID. | |
| **versionNumber** | **Integer**| The version. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**HostedConfigurationVersion**](HostedConfigurationVersion.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="listApplications"></a>
# **listApplications**
> Applications listApplications(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2)



Lists all applications in your Amazon Web Services account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Next token is a pagination token generated by AppConfig to describe what page the previous List call ended on. For the first List request, the nextToken should not be set. On subsequent calls, the nextToken parameter should be set to the previous responses nextToken value. Use this token to get the next set of results. 
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      Applications result = apiInstance.listApplications(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listApplications");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Next token is a pagination token generated by AppConfig to describe what page the previous List call ended on. For the first List request, the nextToken should not be set. On subsequent calls, the nextToken parameter should be set to the previous responses nextToken value. Use this token to get the next set of results.  | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**Applications**](Applications.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | InternalServerException |  -  |
| **481** | BadRequestException |  -  |

<a id="listConfigurationProfiles"></a>
# **listConfigurationProfiles**
> ConfigurationProfiles listConfigurationProfiles(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, type, maxResults2, nextToken2)



Lists the configuration profiles for an application.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Use this token to get the next set of results.
    String type = "type_example"; // String | A filter based on the type of configurations that the configuration profile contains. A configuration can be a feature flag or a freeform configuration.
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      ConfigurationProfiles result = apiInstance.listConfigurationProfiles(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, type, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listConfigurationProfiles");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Use this token to get the next set of results. | [optional] |
| **type** | **String**| A filter based on the type of configurations that the configuration profile contains. A configuration can be a feature flag or a freeform configuration. | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**ConfigurationProfiles**](ConfigurationProfiles.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="listDeploymentStrategies"></a>
# **listDeploymentStrategies**
> DeploymentStrategies listDeploymentStrategies(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2)



Lists deployment strategies.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Use this token to get the next set of results.
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      DeploymentStrategies result = apiInstance.listDeploymentStrategies(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listDeploymentStrategies");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Use this token to get the next set of results. | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**DeploymentStrategies**](DeploymentStrategies.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | InternalServerException |  -  |
| **481** | BadRequestException |  -  |

<a id="listDeployments"></a>
# **listDeployments**
> Deployments listDeployments(applicationId, environmentId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2)



Lists the deployments for an environment in descending deployment number order.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String environmentId = "environmentId_example"; // String | The environment ID.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items that may be returned for this call. If there are items that have not yet been returned, the response will include a non-null <code>NextToken</code> that you can provide in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | The token returned by a prior call to this operation indicating the next set of results to be returned. If not specified, the operation will return the first set of results.
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      Deployments result = apiInstance.listDeployments(applicationId, environmentId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listDeployments");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **environmentId** | **String**| The environment ID. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items that may be returned for this call. If there are items that have not yet been returned, the response will include a non-null &lt;code&gt;NextToken&lt;/code&gt; that you can provide in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| The token returned by a prior call to this operation indicating the next set of results to be returned. If not specified, the operation will return the first set of results. | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**Deployments**](Deployments.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="listEnvironments"></a>
# **listEnvironments**
> Environments listEnvironments(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2)



Lists the environments for an application.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Use this token to get the next set of results.
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      Environments result = apiInstance.listEnvironments(applicationId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listEnvironments");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Use this token to get the next set of results. | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**Environments**](Environments.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="listExtensionAssociations"></a>
# **listExtensionAssociations**
> ExtensionAssociations listExtensionAssociations(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, resourceIdentifier, extensionIdentifier, extensionVersionNumber, maxResults, nextToken, maxResults2, nextToken2)



Lists all AppConfig extension associations in the account. For more information about extensions and associations, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String resourceIdentifier = "resourceIdentifier_example"; // String | The ARN of an application, configuration profile, or environment.
    String extensionIdentifier = "extensionIdentifier_example"; // String | The name, the ID, or the Amazon Resource Name (ARN) of the extension.
    Integer extensionVersionNumber = 56; // Integer | The version number for the extension defined in the association.
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Use this token to get the next set of results or pass null to get the first set of results. 
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      ExtensionAssociations result = apiInstance.listExtensionAssociations(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, resourceIdentifier, extensionIdentifier, extensionVersionNumber, maxResults, nextToken, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listExtensionAssociations");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **resourceIdentifier** | **String**| The ARN of an application, configuration profile, or environment. | [optional] |
| **extensionIdentifier** | **String**| The name, the ID, or the Amazon Resource Name (ARN) of the extension. | [optional] |
| **extensionVersionNumber** | **Integer**| The version number for the extension defined in the association. | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Use this token to get the next set of results or pass null to get the first set of results.  | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**ExtensionAssociations**](ExtensionAssociations.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | InternalServerException |  -  |
| **481** | BadRequestException |  -  |

<a id="listExtensions"></a>
# **listExtensions**
> Extensions listExtensions(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, name, maxResults2, nextToken2)



Lists all custom and Amazon Web Services authored AppConfig extensions in the account. For more information about extensions, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Use this token to get the next set of results. 
    String name = "name_example"; // String | The extension name.
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      Extensions result = apiInstance.listExtensions(xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, name, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listExtensions");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Use this token to get the next set of results.  | [optional] |
| **name** | **String**| The extension name. | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**Extensions**](Extensions.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | InternalServerException |  -  |
| **481** | BadRequestException |  -  |

<a id="listHostedConfigurationVersions"></a>
# **listHostedConfigurationVersions**
> HostedConfigurationVersions listHostedConfigurationVersions(applicationId, configurationProfileId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, versionLabel, maxResults2, nextToken2)



Lists configurations stored in the AppConfig hosted configuration store by version.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String configurationProfileId = "configurationProfileId_example"; // String | The configuration profile ID.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    Integer maxResults = 56; // Integer | The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results.
    String nextToken = "nextToken_example"; // String | A token to start the list. Use this token to get the next set of results. 
    String versionLabel = "versionLabel_example"; // String | An optional filter that can be used to specify the version label of an AppConfig hosted configuration version. This parameter supports filtering by prefix using a wildcard, for example \"v2*\". If you don't specify an asterisk at the end of the value, only an exact match is returned.
    String maxResults2 = "maxResults_example"; // String | Pagination limit
    String nextToken2 = "nextToken_example"; // String | Pagination token
    try {
      HostedConfigurationVersions result = apiInstance.listHostedConfigurationVersions(applicationId, configurationProfileId, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken, versionLabel, maxResults2, nextToken2);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listHostedConfigurationVersions");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **configurationProfileId** | **String**| The configuration profile ID. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **Integer**| The maximum number of items to return for this call. The call also returns a token that you can specify in a subsequent call to get the next set of results. | [optional] |
| **nextToken** | **String**| A token to start the list. Use this token to get the next set of results.  | [optional] |
| **versionLabel** | **String**| An optional filter that can be used to specify the version label of an AppConfig hosted configuration version. This parameter supports filtering by prefix using a wildcard, for example \&quot;v2*\&quot;. If you don&#39;t specify an asterisk at the end of the value, only an exact match is returned. | [optional] |
| **maxResults2** | **String**| Pagination limit | [optional] |
| **nextToken2** | **String**| Pagination token | [optional] |

### Return type

[**HostedConfigurationVersions**](HostedConfigurationVersions.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="listTagsForResource"></a>
# **listTagsForResource**
> ResourceTags listTagsForResource(resourceArn, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves the list of key-value tags assigned to the resource.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String resourceArn = "resourceArn_example"; // String | The resource ARN.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ResourceTags result = apiInstance.listTagsForResource(resourceArn, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listTagsForResource");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **resourceArn** | **String**| The resource ARN. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ResourceTags**](ResourceTags.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | BadRequestException |  -  |
| **482** | InternalServerException |  -  |

<a id="startDeployment"></a>
# **startDeployment**
> Deployment startDeployment(applicationId, environmentId, startDeploymentRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Starts a deployment.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String environmentId = "environmentId_example"; // String | The environment ID.
    StartDeploymentRequest startDeploymentRequest = new StartDeploymentRequest(); // StartDeploymentRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Deployment result = apiInstance.startDeployment(applicationId, environmentId, startDeploymentRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#startDeployment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **environmentId** | **String**| The environment ID. | |
| **startDeploymentRequest** | [**StartDeploymentRequest**](StartDeploymentRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Deployment**](Deployment.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | ConflictException |  -  |
| **483** | InternalServerException |  -  |

<a id="stopDeployment"></a>
# **stopDeployment**
> Deployment stopDeployment(applicationId, environmentId, deploymentNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Stops a deployment. This API action works only on deployments that have a status of &lt;code&gt;DEPLOYING&lt;/code&gt;. This action moves the deployment to a status of &lt;code&gt;ROLLED_BACK&lt;/code&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String environmentId = "environmentId_example"; // String | The environment ID.
    Integer deploymentNumber = 56; // Integer | The sequence number of the deployment.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Deployment result = apiInstance.stopDeployment(applicationId, environmentId, deploymentNumber, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#stopDeployment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **environmentId** | **String**| The environment ID. | |
| **deploymentNumber** | **Integer**| The sequence number of the deployment. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Deployment**](Deployment.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **202** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | InternalServerException |  -  |
| **482** | BadRequestException |  -  |

<a id="tagResource"></a>
# **tagResource**
> tagResource(resourceArn, tagResourceRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Assigns metadata to an AppConfig resource. Tags help organize and categorize your AppConfig resources. Each tag consists of a key and an optional value, both of which you define. You can specify a maximum of 50 tags for a resource.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String resourceArn = "resourceArn_example"; // String | The ARN of the resource for which to retrieve tags.
    TagResourceRequest tagResourceRequest = new TagResourceRequest(); // TagResourceRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.tagResource(resourceArn, tagResourceRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#tagResource");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **resourceArn** | **String**| The ARN of the resource for which to retrieve tags. | |
| **tagResourceRequest** | [**TagResourceRequest**](TagResourceRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | BadRequestException |  -  |
| **482** | InternalServerException |  -  |

<a id="untagResource"></a>
# **untagResource**
> untagResource(resourceArn, tagKeys, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Deletes a tag key and value from an AppConfig resource.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String resourceArn = "resourceArn_example"; // String | The ARN of the resource for which to remove tags.
    List<String> tagKeys = Arrays.asList(); // List<String> | The tag keys to delete.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.untagResource(resourceArn, tagKeys, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#untagResource");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **resourceArn** | **String**| The ARN of the resource for which to remove tags. | |
| **tagKeys** | [**List&lt;String&gt;**](String.md)| The tag keys to delete. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | ResourceNotFoundException |  -  |
| **481** | BadRequestException |  -  |
| **482** | InternalServerException |  -  |

<a id="updateApplication"></a>
# **updateApplication**
> Application updateApplication(applicationId, updateApplicationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Updates an application.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    UpdateApplicationRequest updateApplicationRequest = new UpdateApplicationRequest(); // UpdateApplicationRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Application result = apiInstance.updateApplication(applicationId, updateApplicationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#updateApplication");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **updateApplicationRequest** | [**UpdateApplicationRequest**](UpdateApplicationRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Application**](Application.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="updateConfigurationProfile"></a>
# **updateConfigurationProfile**
> ConfigurationProfile updateConfigurationProfile(applicationId, configurationProfileId, updateConfigurationProfileRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Updates a configuration profile.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String configurationProfileId = "configurationProfileId_example"; // String | The ID of the configuration profile.
    UpdateConfigurationProfileRequest updateConfigurationProfileRequest = new UpdateConfigurationProfileRequest(); // UpdateConfigurationProfileRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ConfigurationProfile result = apiInstance.updateConfigurationProfile(applicationId, configurationProfileId, updateConfigurationProfileRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#updateConfigurationProfile");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **configurationProfileId** | **String**| The ID of the configuration profile. | |
| **updateConfigurationProfileRequest** | [**UpdateConfigurationProfileRequest**](UpdateConfigurationProfileRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ConfigurationProfile**](ConfigurationProfile.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="updateDeploymentStrategy"></a>
# **updateDeploymentStrategy**
> DeploymentStrategy updateDeploymentStrategy(deploymentStrategyId, updateDeploymentStrategyRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Updates a deployment strategy.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String deploymentStrategyId = "deploymentStrategyId_example"; // String | The deployment strategy ID.
    UpdateDeploymentStrategyRequest updateDeploymentStrategyRequest = new UpdateDeploymentStrategyRequest(); // UpdateDeploymentStrategyRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      DeploymentStrategy result = apiInstance.updateDeploymentStrategy(deploymentStrategyId, updateDeploymentStrategyRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#updateDeploymentStrategy");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deploymentStrategyId** | **String**| The deployment strategy ID. | |
| **updateDeploymentStrategyRequest** | [**UpdateDeploymentStrategyRequest**](UpdateDeploymentStrategyRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**DeploymentStrategy**](DeploymentStrategy.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="updateEnvironment"></a>
# **updateEnvironment**
> Environment updateEnvironment(applicationId, environmentId, updateEnvironmentRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Updates an environment.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String environmentId = "environmentId_example"; // String | The environment ID.
    UpdateEnvironmentRequest updateEnvironmentRequest = new UpdateEnvironmentRequest(); // UpdateEnvironmentRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Environment result = apiInstance.updateEnvironment(applicationId, environmentId, updateEnvironmentRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#updateEnvironment");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **environmentId** | **String**| The environment ID. | |
| **updateEnvironmentRequest** | [**UpdateEnvironmentRequest**](UpdateEnvironmentRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Environment**](Environment.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="updateExtension"></a>
# **updateExtension**
> Extension updateExtension(extensionIdentifier, updateExtensionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Updates an AppConfig extension. For more information about extensions, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String extensionIdentifier = "extensionIdentifier_example"; // String | The name, the ID, or the Amazon Resource Name (ARN) of the extension.
    UpdateExtensionRequest updateExtensionRequest = new UpdateExtensionRequest(); // UpdateExtensionRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Extension result = apiInstance.updateExtension(extensionIdentifier, updateExtensionRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#updateExtension");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **extensionIdentifier** | **String**| The name, the ID, or the Amazon Resource Name (ARN) of the extension. | |
| **updateExtensionRequest** | [**UpdateExtensionRequest**](UpdateExtensionRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**Extension**](Extension.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | ConflictException |  -  |
| **483** | InternalServerException |  -  |

<a id="updateExtensionAssociation"></a>
# **updateExtensionAssociation**
> ExtensionAssociation updateExtensionAssociation(extensionAssociationId, updateExtensionAssociationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Updates an association. For more information about extensions and associations, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html\&quot;&gt;Working with AppConfig extensions&lt;/a&gt; in the &lt;i&gt;AppConfig User Guide&lt;/i&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String extensionAssociationId = "extensionAssociationId_example"; // String | The system-generated ID for the association.
    UpdateExtensionAssociationRequest updateExtensionAssociationRequest = new UpdateExtensionAssociationRequest(); // UpdateExtensionAssociationRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      ExtensionAssociation result = apiInstance.updateExtensionAssociation(extensionAssociationId, updateExtensionAssociationRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#updateExtensionAssociation");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **extensionAssociationId** | **String**| The system-generated ID for the association. | |
| **updateExtensionAssociationRequest** | [**UpdateExtensionAssociationRequest**](UpdateExtensionAssociationRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**ExtensionAssociation**](ExtensionAssociation.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

<a id="validateConfiguration"></a>
# **validateConfiguration**
> validateConfiguration(applicationId, configurationProfileId, configurationVersion, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Uses the validators in a configuration profile to validate a configuration.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://appconfig.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String applicationId = "applicationId_example"; // String | The application ID.
    String configurationProfileId = "configurationProfileId_example"; // String | The configuration profile ID.
    String configurationVersion = "configurationVersion_example"; // String | The version of the configuration to validate.
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      apiInstance.validateConfiguration(applicationId, configurationProfileId, configurationVersion, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#validateConfiguration");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **applicationId** | **String**| The application ID. | |
| **configurationProfileId** | **String**| The configuration profile ID. | |
| **configurationVersion** | **String**| The version of the configuration to validate. | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

null (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | Success |  -  |
| **480** | BadRequestException |  -  |
| **481** | ResourceNotFoundException |  -  |
| **482** | InternalServerException |  -  |

