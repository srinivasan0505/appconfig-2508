/**
 * MCP Server function for Creates a new configuration in the AppConfig hosted configuration store.
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_Applications_Application_Id_Configurationprofiles_Configuration_Profile_Id_Hostedconfigurationversions_Content_TypeMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Applications_Application_Id_Configurationprofiles_Configuration_Profile_Id_Hostedconfigurationversions_Content_TypeHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("ApplicationId")) {
            queryParams.add("ApplicationId=" + args.get("ApplicationId"));
        }
        if (args.containsKey("ConfigurationProfileId")) {
            queryParams.add("ConfigurationProfileId=" + args.get("ConfigurationProfileId"));
        }
        if (args.containsKey("Description")) {
            queryParams.add("Description=" + args.get("Description"));
        }
        if (args.containsKey("Content-Type")) {
            queryParams.add("Content-Type=" + args.get("Content-Type"));
        }
        if (args.containsKey("VersionLabel")) {
            queryParams.add("VersionLabel=" + args.get("VersionLabel"));
        }
        if (args.containsKey("Content")) {
            queryParams.add("Content=" + args.get("Content"));
        }
        if (args.containsKey("Latest-Version-Number")) {
            queryParams.add("Latest-Version-Number=" + args.get("Latest-Version-Number"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_applications_application_id_configurationprofiles_configuration_profile_id_hostedconfigurationversions_content_type" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_Applications_Application_Id_Configurationprofiles_Configuration_Profile_Id_Hostedconfigurationversions_Content_TypeTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> ApplicationIdProperty = new HashMap<>();
        ApplicationIdProperty.put("type", "string");
        ApplicationIdProperty.put("required", true);
        ApplicationIdProperty.put("description", "The application ID.");
        properties.put("ApplicationId", ApplicationIdProperty);
        Map<String, Object> ConfigurationProfileIdProperty = new HashMap<>();
        ConfigurationProfileIdProperty.put("type", "string");
        ConfigurationProfileIdProperty.put("required", true);
        ConfigurationProfileIdProperty.put("description", "The configuration profile ID.");
        properties.put("ConfigurationProfileId", ConfigurationProfileIdProperty);
        Map<String, Object> DescriptionProperty = new HashMap<>();
        DescriptionProperty.put("type", "string");
        DescriptionProperty.put("required", false);
        DescriptionProperty.put("description", "A description of the configuration.");
        properties.put("Description", DescriptionProperty);
        Map<String, Object> Content-TypeProperty = new HashMap<>();
        Content-TypeProperty.put("type", "string");
        Content-TypeProperty.put("required", true);
        Content-TypeProperty.put("description", "A standard MIME type describing the format of the configuration content. For more information, see <a href=\"https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17\">Content-Type</a>.");
        properties.put("Content-Type", Content-TypeProperty);
        Map<String, Object> VersionLabelProperty = new HashMap<>();
        VersionLabelProperty.put("type", "string");
        VersionLabelProperty.put("required", false);
        VersionLabelProperty.put("description", "An optional, user-defined label for the AppConfig hosted configuration version. This value must contain at least one non-numeric character. For example, \"v2.2.0\".");
        properties.put("VersionLabel", VersionLabelProperty);
        Map<String, Object> ContentProperty = new HashMap<>();
        ContentProperty.put("type", "string");
        ContentProperty.put("required", true);
        ContentProperty.put("description", "Input parameter: The content of the configuration or the configuration data.");
        properties.put("Content", ContentProperty);
        Map<String, Object> Latest-Version-NumberProperty = new HashMap<>();
        Latest-Version-NumberProperty.put("type", "string");
        Latest-Version-NumberProperty.put("required", false);
        Latest-Version-NumberProperty.put("description", "An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version. To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.");
        properties.put("Latest-Version-Number", Latest-Version-NumberProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_applications_application_id_configurationprofiles_configuration_profile_id_hostedconfigurationversions_content_type",
            "Creates a new configuration in the AppConfig hosted configuration store.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Applications_Application_Id_Configurationprofiles_Configuration_Profile_Id_Hostedconfigurationversions_Content_TypeHandler(config));
    }
    
}