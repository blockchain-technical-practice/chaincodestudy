/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

 */

package main

import (

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"time"
	"encoding/json"
)


type dairyfarm struct {}


/**
 系统初始化
 */
func (t *dairyfarm) Init(stub shim.ChaincodeStubInterface) pb.Response {
	

	return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))


}

/**
  系统invoke方法
 */
func (t *dairyfarm) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	_,args := stub.GetFunctionAndParameters()

	var opttype = args[0]
	var assetname = args[1]
	var optcontont = args[2]

	fmt.Printf(" parm is  %s  %s  %s   \n " , opttype , assetname , optcontont )


	if opttype == "putvalue" {  //设置

		stub.PutState(assetname,[]byte(optcontont))
		return shim.Success( []byte( "success put " + optcontont  )  )


	}else if opttype == "getlastvalue"{  //取值


		var keyvalue []byte
		var err error
		keyvalue,err = stub.GetState(assetname)

		if( err != nil  ){

			return shim.Error(" finad error! ")
		}

		return shim.Success( keyvalue )



	}else if opttype == "gethistory"{   //获取交易记录


		keysIter, err := stub.GetHistoryForKey( assetname );

		if err != nil {
			return shim.Error(fmt.Sprintf("GetHistoryForKey failed. Error accessing state: %s", err))
		}
		defer keysIter.Close()

		var keys []string

		for keysIter.HasNext() {

			response, iterErr := keysIter.Next()
			if iterErr != nil {
				return shim.Error(fmt.Sprintf("GetHistoryForKey operation failed. Error accessing state: %s", err))
			}
			//交易编号
			txid := response.TxId
			//交易的值
			txvalue := response.Value
			//当前交易的状态
			txstatus := response.IsDelete
			//交易发生的时间戳
			txtimesamp :=response.Timestamp

			tm := time.Unix(txtimesamp.Seconds, 0)
			datestr := tm.Format("2006-01-02 03:04:05 PM")

			fmt.Printf(" Tx info -   txid : %s   value :  %s  if delete: %t   datetime : %s \n ", txid , string(txvalue) , txstatus , datestr )
			keys = append( keys , string(txvalue) + ":" + datestr)

		}


		jsonKeys, err := json.Marshal(keys)
		if err != nil {
			return shim.Error(fmt.Sprintf("query operation failed. Error marshaling JSON: %s", err))
		}

		return shim.Success(jsonKeys)


	}else{

		return shim.Success([]byte("success invok and No operation !!!!!!!! "))


	}
}



func main() {
	err := shim.Start( new( dairyfarm ) )
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


