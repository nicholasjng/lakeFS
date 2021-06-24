/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * GarbageCollectionPrepareResponse
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class GarbageCollectionPrepareResponse {
  public static final String SERIALIZED_NAME_RUN_ID = "run_id";
  @SerializedName(SERIALIZED_NAME_RUN_ID)
  private String runId;


  public GarbageCollectionPrepareResponse runId(String runId) {
    
    this.runId = runId;
    return this;
  }

   /**
   * a unique identifier generated for this GC job
   * @return runId
  **/
  @ApiModelProperty(example = "64eaa103-d726-4a33-bcb8-7c0b4abfe09e", required = true, value = "a unique identifier generated for this GC job")

  public String getRunId() {
    return runId;
  }


  public void setRunId(String runId) {
    this.runId = runId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GarbageCollectionPrepareResponse garbageCollectionPrepareResponse = (GarbageCollectionPrepareResponse) o;
    return Objects.equals(this.runId, garbageCollectionPrepareResponse.runId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(runId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GarbageCollectionPrepareResponse {\n");
    sb.append("    runId: ").append(toIndentedString(runId)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
