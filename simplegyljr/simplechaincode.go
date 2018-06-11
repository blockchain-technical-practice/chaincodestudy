
package main

import (

	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"github.com/hyperledger/fabric/common/util"
)

type gylchaincode struct {

}

func (t *gylchaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("success init "))
}


func (t *gylchaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {


	_,args := stub.GetFunctionAndParameters()

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]



	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


	// 保存交易数据
	if a_parm == "set"{

		stub.PutState(b_parm,[]byte(c_parm))
		return shim.Success( []byte( "success invok " + c_parm  )  )

	}else if a_parm == "get"{   //查询交易数据

		var keyvalue []byte
		var err error
		keyvalue,err = stub.GetState(b_parm)

		if( err != nil  ){

			return shim.Error(" finad error! ")
		}


		return shim.Success( keyvalue )


	}else{


		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))

	}


}


func main() {
	err := shim.Start(new(gylchaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


