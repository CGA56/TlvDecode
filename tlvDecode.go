package main

import (
	"fmt"
	"log"
	"strconv"
)

var tlvCollection = make(map[int]map[string]string)

type tlv struct {
	lenght       string
	typeTlvValue string
	secretTlv    string
	rest         string
}

type wrapper struct {
	data   map[int]map[string]string
	status string
}

func getSchema(objTlv *tlv) map[string]string {

	return map[string]string{

		"lenght":    objTlv.lenght,
		"typeValue": string(objTlv.typeTlvValue),
		"secretTlv": string(objTlv.secretTlv),
	}

}

func addTlv(objTlv *tlv) string {
	var schema = getSchema(objTlv)
	tlvCollection[len(tlvCollection)+1] = schema
	return "ok"
}

func searchTlv(objTlv string) *tlv {
	if len(objTlv) >= 5 {
		var tlvObj tlv
		lenghtTlvValue, _ := strconv.ParseInt(objTlv[0:2], 10, 0)
		typeTlvValue := objTlv[2:5]
		secretTlv := objTlv[5 : lenghtTlvValue+5]
		rest := objTlv[5+lenghtTlvValue : len(objTlv)]
		tlvObj.lenght = objTlv[0:2]
		tlvObj.typeTlvValue = typeTlvValue
		tlvObj.secretTlv = secretTlv
		tlvObj.rest = rest
		return &tlvObj
	}
	return nil

}

func createWrapper(data map[int]map[string]string, msje string) *wrapper {
	var wrapper wrapper
	wrapper.data = data
	wrapper.status = msje
	return &wrapper
}

func populate(tlvExample string) *wrapper {

	var dataExample = tlvExample
	var data *tlv

	for range dataExample {
		data = searchTlv(dataExample)
		addTlv(data)
		dataExample = data.rest
		if dataExample == "" {
			break
		}
	}
	return createWrapper(tlvCollection, "status: OK")
}

func main() {
	const tlvExample string = "11A05AB398765UJ102N2300"
	defer func() {
		if err := recover(); err != nil {
			str := "status :" + fmt.Sprint(err)
			log.Println(createWrapper(nil, str))
		}
	}()

	var result = populate(tlvExample)
	fmt.Println(result)

}
