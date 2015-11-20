package builder

import (
	"encoding/json"
	"testing"
)

func TestgetParentKey(t *testing.T) {
	k1 := ".a.b.c"
	k2 := ".a"
	k3 := ""

	if getParentKey(k1) != ".a.b" {
		t.Error("Failed")
	}

	if getParentKey(k2) != "" {
		t.Error("Failed")
	}

	if getParentKey(k3) != "" {
		t.Error("Failed")
	}
}

func TestFlatXmlContentToMap(t *testing.T) {
	xmlContent := `
<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult>
  <Owner>
    <ID>1087645789419058</ID>
    <DisplayName>1087645789419058</DisplayName>
  </Owner>
  <Buckets>
    <Bucket>
      <Location>oss-cn-beijing</Location>
      <Name>datafiddle1</Name>
      <CreationDate>2015-08-18T08:13:29.000Z</CreationDate>
    </Bucket>
  </Buckets>
</ListAllMyBucketsResult>

`

	FlatXmlContentToMap("", []byte(xmlContent), nil)

	t.Log("-------------")
}

func TestFlatJsonObjToMap(t *testing.T) {
	jsonContent := `
   {
   	"k1":"v1",
   	"k2":{
   		"s1:":"s11",
   		"s2":"s22"
   	}
   }
`

	var unmarshaledObj interface{}
	json.Unmarshal([]byte(jsonContent), &unmarshaledObj)

	reconstructedMap := make(map[string]interface{})
	FlatJsonObjToMap("", unmarshaledObj.(map[string]interface{}), reconstructedMap)

	if len(reconstructedMap) != 3 {
		t.Error("TestFlatJsonObjToMap failed")
	}

	if v, ok := reconstructedMap[".k2.s2"]; !ok || v != "s22" {
		t.Error("TestFlatJsonObjToMap failed for value")
	}
}
