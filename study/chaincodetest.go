/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321687

 */

package main

import (


	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	 pb "github.com/hyperledger/fabric/protos/peer"

	"strconv"
)

//定义一个机构体，作为chaincode的主对象，可以是任何符合go语言规范的命名方式

type chainCodeStudy1 struct {

}


var A, B string
var Aval, Bval, X int



/**
	系统初始化方法， 在部署chaincode的过程中当执行命令

    peer chaincode instantiate -o orderer.robertfabrictest.com:7050 -C
         roberttestchannel -n r_test_cc6 -v 1.0 -c '{"Args":["init","a","100","b","200"]}'
         -P "OR	('Org1MSP.member','Org2MSP.member')"

    的时候会调用该方法


 */
func (t *chainCodeStudy1) Init(stub shim.ChaincodeStubInterface) pb.Response {


	var err error


	//获取传入参数值，在端采用数组的方式传入相关的参数
	_, args := stub.GetFunctionAndParameters()



	if len(args) != 4 {

		errinfo := fmt.Sprint(" 参数值错误 . Expecting 4 ， 传入值为  %d",len(args))
		return shim.Error(errinfo)
	}

	// 获取第一个和第二个参数，并验证

	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("参数错误，不是数字")
	}

	//获取第三个和第四个参数并验证
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("参数错误，不是数字")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	/************
			// Write the state to the ledger
			err = stub.PutState(A, []byte(strconv.Itoa(Aval))
			if err != nil {
				return nil, err
			}

			stub.PutState(B, []byte(strconv.Itoa(Bval))
			err = stub.PutState(B, []byte(strconv.Itoa(Bval))
			if err != nil {
				return nil, err
			}
	************/
	return shim.Success(nil)
}

/**

     主业务逻辑，在执行命令
     peer chaincode invoke -o 192.168.23.212:7050 -C roberttestchannel -n r_test_cc6 -c '{"Args":["invoke","a","b","1"]}'

     的时候系统会调用该方法并传入相关的参数，注意 "invoke" 之后的参数是需要传入的参数



 */
func (t *chainCodeStudy1) Invoke(stub shim.ChaincodeStubInterface) pb.Response {


	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\"")


}


/**

    业务方法，以为chaincode对外提供的业务接口只有 Invoke方法，如果业务逻辑比较发展，可以将相关的业务分成多个方法，然后在
    Invoke 方法中调用这些方法，

    本方法是其中的一个业务方法，本方法模拟从账号A，转账X 给账户B

 */
func (t *chainCodeStudy1) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// Transaction makes payment of X units from A to B
	X, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("这是服务器的信息 只能显示在服务器端  ", args[0], err)
		return shim.Error( fmt.Sprintf(" 参数 %s 转换时发生错误  %s ！！！！！ " , args[0] , err ) )
	}

	Aval = Aval - X
	Bval = Bval + X
	ts, err2 := stub.GetTxTimestamp()


	if err2 != nil {
		fmt.Printf("Error getting transaction timestamp: %s", err2)
		return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err2))
	}

	fmt.Printf("Transaction Time: %v,Aval = %d, Bval = %d\n", ts, Aval, Bval)
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(chainCodeStudy1))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}