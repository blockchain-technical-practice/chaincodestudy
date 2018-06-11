/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

 */

package main

import (
	//"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)


func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte){

	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}


func TestExample02_Init(t *testing.T) {

	scc := new(simplechaincode)
	stub := shim.NewMockStub("ex02", scc)


	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})


}