package core

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"pizer_project/conf"
	"pizer_project/globle"
	"pizer_project/utils"
	"strings"
)

var (
	r      map[string]interface{}
	esConf *utils.CfgFileConfig
)

//es操作方法结构体
type EsExecute struct {
	es *elasticsearch.Client
}

//es操作方法 [add source by index]
func (execute EsExecute) AddSourceByIndex(index string, source string) bool {

	res, err := execute.es.Index(
		index,                     // Index name
		strings.NewReader(source), // Document body `{"title" : "Test2222"}`
		//execute.es.Index.WithDocumentID("1"),            // Document ID
		execute.es.Index.WithRefresh("true"), // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return false
	}
	defer res.Body.Close()

	log.Println(res)
	return true
}

func esInit() EsExecute {
	esConf = conf.GetBaseConf()
	addresses := esConf.Get(globle.CONST_CONFIG_SECTION_ES, "addresses")

	cfg := elasticsearch.Config{
		Addresses: []string{
			addresses,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error new client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Deserialize the response into a map. 解析info返回的是否为一个map
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	} else {
		log.Println(res)
	}

	log.Printf("Success connect elasticsearch version : %s", r["version"].(map[string]interface{})["number"])

	intExe := EsExecute{}
	intExe.es = es
	return intExe
}

// 测试ES
//func main() {
//	esInit().AddSourceByIndex("pizer", `{"age":"19"}`)
//}
