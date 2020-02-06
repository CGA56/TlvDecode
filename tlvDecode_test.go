package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createMockSchema() *tlv {

	var element tlv
	element.lenght = "11"
	element.typeTlvValue = "A05"
	element.secretTlv = "AB398765UJ1"
	element.rest = ""
	return &element
}

// 11A05AB398765UJ1
func TestGetSchema(t *testing.T) {
	assert := assert.New(t)
	var mock *tlv = createMockSchema()
	var resp = getSchema(mock)
	assert.EqualValues(mock.lenght, resp["lenght"])
	assert.EqualValues(mock.typeTlvValue, resp["typeValue"])
	assert.EqualValues(mock.secretTlv, resp["secretTlv"])
	assert.EqualValues(mock.rest, resp["rest"])

}

func TestAddTlv(t *testing.T) {
	assert := assert.New(t)
	var mock *tlv = createMockSchema()
	var resp = addTlv(mock)
	assert.Equal(resp, "ok")

}

func TestSearchTlv(t *testing.T) {
	assert := assert.New(t)
	const dataDummy = "11A05AB398765UJ1"
	var mock *tlv = createMockSchema()
	resp := searchTlv(dataDummy)
	assert.EqualValues(resp.lenght, mock.lenght)
}

func TestSearchTlvFail(t *testing.T) {
	assert := assert.New(t)
	const dataDummy = "11A0"
	resp := searchTlv(dataDummy)
	assert.Nil(resp)

}

func TestCreateWrapper(t *testing.T) {
	assert := assert.New(t)
	var tlvCollection = make(map[int]map[string]string)
	var mock = createMockSchema()
	resp := getSchema(mock)
	tlvCollection[1] = resp
	wrapper := createWrapper(tlvCollection, "test")
	assert.Equal(wrapper.status, "test")
	assert.Equal(wrapper.data, tlvCollection)

}

func TestPopulate(t *testing.T) {
	assert := assert.New(t)
	tlvExample := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48}
	var wrapper *wrapper = populate(tlvExample)
	assert.Equal(wrapper.status, "status: OK")

}
