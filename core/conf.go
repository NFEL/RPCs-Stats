package core

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const (
	chainsDir = "./data/chains.json"
)

var (
	Chains = make(map[int][]string, 0)
	//BenchInstances = make(map[string]Benchmark.Benchmark, 0)
)

func init() {

	jsonFile, err := os.Open(chainsDir)
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}(jsonFile)
	if err != nil {
		log.Fatalf("JSONTokenLoader: %s", err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("JSONTokenLoader: %s", err)
	}
	err = json.Unmarshal(byteValue, &Chains)
	if err != nil {
		log.Fatalf("TokenLoader: %s", err)
	}

	//for _, rpcs := range Chains {
	//	for _, rpc := range rpcs {
	//		b, errBench := Benchmark.New(rpc)
	//		if errBench != nil {
	//			log.Errorf("%s", errBench)
	//		}
	//		_ = b
	//		//BenchInstances[rpc] = *b
	//	}
	//}
}
