package demo1

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/RoaringBitmap/roaring"
)

func Test1(t *testing.T) {
	rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	t.Log(rb1.String())
}

type tests struct {
	ID int `json:"id"`
}

func Test2(t *testing.T) {
	idList := []uint32{}
	idList2 := []uint16{}
	idStruct := []tests{}
	for i := 0; i < 500; i++ {
		idList = append(idList, uint32(i))
		idList2 = append(idList2, uint16(i))
		idStruct = append(idStruct, tests{ID: i})
	}
	t.Log("idList len: ", len(idList))
	rb1 := roaring.BitmapOf(idList...)
	t.Log(rb1.String())
	b, _ := rb1.ToBytes()
	t.Log(string(b))
	rb2 := roaring.New()
	rb2.FromBuffer(b)
	t.Log(rb2.String())
	b64, _ := rb2.ToBase64()

	t.Log("base64: ", len(b64), b64)
	// base64:  1356 OjAAAAEAAAAAAPMBEAAAAAAAAQACAAMABAAFAAYABwAIAAkACgALAAwADQAOAA8AEAARABIAEwAUABUAFgAXABgAGQAaABsAHAAdAB4AHwAgACEAIgAjACQAJQAmACcAKAApACoAKwAsAC0ALgAvADAAMQAyADMANAA1ADYANwA4ADkAOgA7ADwAPQA+AD8AQABBAEIAQwBEAEUARgBHAEgASQBKAEsATABNAE4ATwBQAFEAUgBTAFQAVQBWAFcAWABZAFoAWwBcAF0AXgBfAGAAYQBiAGMAZABlAGYAZwBoAGkAagBrAGwAbQBuAG8AcABxAHIAcwB0AHUAdgB3AHgAeQB6AHsAfAB9AH4AfwCAAIEAggCDAIQAhQCGAIcAiACJAIoAiwCMAI0AjgCPAJAAkQCSAJMAlACVAJYAlwCYAJkAmgCbAJwAnQCeAJ8AoAChAKIAowCkAKUApgCnAKgAqQCqAKsArACtAK4ArwCwALEAsgCzALQAtQC2ALcAuAC5ALoAuwC8AL0AvgC/AMAAwQDCAMMAxADFAMYAxwDIAMkAygDLAMwAzQDOAM8A0ADRANIA0wDUANUA1gDXANgA2QDaANsA3ADdAN4A3wDgAOEA4gDjAOQA5QDmAOcA6ADpAOoA6wDsAO0A7gDvAPAA8QDyAPMA9AD1APYA9wD4APkA+gD7APwA/QD+AP8AAAEBAQIBAwEEAQUBBgEHAQgBCQEKAQsBDAENAQ4BDwEQAREBEgETARQBFQEWARcBGAEZARoBGwEcAR0BHgEfASABIQEiASMBJAElASYBJwEoASkBKgErASwBLQEuAS8BMAExATIBMwE0ATUBNgE3ATgBOQE6ATsBPAE9AT4BPwFAAUEBQgFDAUQBRQFGAUcBSAFJAUoBSwFMAU0BTgFPAVABUQFSAVMBVAFVAVYBVwFYAVkBWgFbAVwBXQFeAV8BYAFhAWIBYwFkAWUBZgFnAWgBaQFqAWsBbAFtAW4BbwFwAXEBcgFzAXQBdQF2AXcBeAF5AXoBewF8AX0BfgF/AYABgQGCAYMBhAGFAYYBhwGIAYkBigGLAYwBjQGOAY8BkAGRAZIBkwGUAZUBlgGXAZgBmQGaAZsBnAGdAZ4BnwGgAaEBogGjAaQBpQGmAacBqAGpAaoBqwGsAa0BrgGvAbABsQGyAbMBtAG1AbYBtwG4AbkBugG7AbwBvQG+Ab8BwAHBAcIBwwHEAcUBxgHHAcgByQHKAcsBzAHNAc4BzwHQAdEB0gHTAdQB1QHWAdcB2AHZAdoB2wHcAd0B3gHfAeAB4QHiAeMB5AHlAeYB5wHoAekB6gHrAewB7QHuAe8B8AHxAfIB8wE=
	rb3 := roaring.New()
	rb3.FromBase64(b64)
	t.Log("rb3: ", rb3.String())

	strBody, _ := json.Marshal(idStruct)
	t.Log("json len: ", len(strBody))
	strBody, _ = json.Marshal(idList)
	t.Log("json len: ", len(strBody))

	strBody, _ = json.Marshal(idList2)
	t.Log("json len: ", len(strBody))
}

func Test21(t *testing.T) {
	idList := []uint32{}
	for i := 0; i < 500; i++ {
		idList = append(idList, uint32(i))
	}
	rb1 := roaring.BitmapOf(idList...)
	t.Log(rb1.String())
	buf := new(bytes.Buffer)
	rb1.WriteTo(buf)
	t.Log(len(buf.Bytes()))
	toB, _ := rb1.ToBytes()
	t.Log("to byte len: ", len(toB))
	to64, _ := rb1.ToBase64()
	t.Log("to b 64 len: ", len(to64))
	strBody, _ := json.Marshal(idList)
	t.Log("json len: ", len(strBody))
}

func Test3(t *testing.T) {
	idList := []uint32{}
	for i := 0; i < 500; i++ {
		idList = append(idList, rand.Uint32())
	}
	rb1 := roaring.BitmapOf(idList...)
	t.Log(rb1.String())
	buf := new(bytes.Buffer)
	rb1.WriteTo(buf)
	t.Log(len(buf.Bytes()))
	toB, _ := rb1.ToBytes()
	t.Log("to byte len: ", len(toB))
}

func Test4(t *testing.T) {
	rb1 := roaring.New()
	rb1.FromBuffer([]byte(`{7, 8, 19, 22, 2206820,6062535,8177164,12210768,19089668,20958506,28462600,32990530,55086860,57060992,63683067,65203089,80372672,85077591,92486895,95052488,99254840, 1, 2 ,3}`))
	t.Log(rb1.String())
}
