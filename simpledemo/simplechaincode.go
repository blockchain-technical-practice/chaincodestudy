/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 41132168111

*/

package main

import (

	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"github.com/hyperledger/fabric/common/util"
	"encoding/json"
	"strings"
	"time"
)

//定义一个机构体，作为chaincode的主对象，可以是任何符合go语言规范的命名方式
type simplechaincode struct {

}


/**

	系统初始化方法， 在部署chaincode的过程中当执行命令

    peer chaincode instantiate -o orderer.robertfabrictest.com:7050 -C
         roberttestchannel -n r_test_cc6 -v 1.0 -c '{"Args":["init","a","100","b","200"]}'
         -P "OR	('Org1MSP.member','Org2MSP.member')"

    的时候会调用该方法


	https://github.com/hyperledger/fabric/blob/release/core/chaincode/shim/interfaces.go  所有的注释这里

*/
func (t *simplechaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println(" <<  ========  success init it is view in docker  ==========  >> ")
	return shim.Success([]byte("success init "))
}

/**

  主业务逻辑，在执行命令
  peer chaincode invoke -o 192.168.23.212:7050 -C roberttestchannel -n r_test_cc6 -c '{"Args":["invoke","a","b","1"]}'

  的时候系统会调用该方法并传入相关的参数，注意 "invoke" 之后的参数是需要传入的参数


*/
func (t *simplechaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {


	_,args := stub.GetFunctionAndParameters()

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  1、success it is view in docker  ========== ")
	fmt.Println("  ========  2、success it is view in docker  ========== ")
	fmt.Println("  ========  3、success it is view in docker  ========== ")
	fmt.Println("  ========  4、success it is view in docker  ========== ")


	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


	// 设定值
	if a_parm == "set"{

		stub.PutState(b_parm,[]byte(c_parm))
		return shim.Success( []byte( "success invok " + c_parm  )  )

	}else if a_parm == "get"{   //取单个值

		var keyvalue []byte
		var err error
		keyvalue,err = stub.GetState(b_parm)

		if( err != nil  ){

			return shim.Error(" finad error! ")
		}


		return shim.Success( keyvalue )



	}else if a_parm == "CreateCompositeKeyandset" { // 设置一个复合键的值

		//parms := []string{ "c_1" , "d_1" , "e_1","f_1","g_1","h_1" }
		parms := strings.Split(c_parm,",")
		ckey ,_ := stub.CreateCompositeKey(b_parm,parms)

		fmt.Println("  ========  GetStateByPartialCompositeKey   ========== ",ckey)


		err := stub.PutState(ckey , []byte(c_parm) )

		if err !=nil{

			fmt.Println("CreateCompositeKeyandset() : Error inserting Object into State Database %s", err)
		}


		return shim.Success([]byte(ckey))



	}else if a_parm == "GetStateByPartialCompositeKey" { // 设置一个复合键的值


		fmt.Println("  ========  GetStateByPartialCompositeKey   ========== ")

		searchparm := strings.Split(c_parm,",")
		rs, err := stub.GetStateByPartialCompositeKey(b_parm,searchparm)
		if err != nil {
			error_str := fmt.Sprintf("GetListOfInitAucs operation failed. Error marshaling JSON: %s", err)
			return shim.Error(error_str)
		}

		defer rs.Close()

		// Iterate through result set
		var i int
		var tlist []string // Define a list
		for i = 0; rs.HasNext(); i++ {

			// We can process whichever return value is of interest
			responseRange, err := rs.Next()

			if err != nil {
				error_str := fmt.Sprintf("GetListOfInitAucs() operation failed - Unmarshall Error. %s", err)
				fmt.Println(error_str)
				return shim.Error(error_str)
			}

			objectType, compositeKeyParts, _ := stub.SplitCompositeKey(responseRange.Key)

			/*fmt.Println(" objectType value is  "+objectType)
			fmt.Println(compositeKeyParts)*/

			returnedColor := compositeKeyParts[0]
			returnedMarbleName := compositeKeyParts[1]
			fmt.Printf("- found a marble from index:%s color:%s name:%s\n", objectType, returnedColor, returnedMarbleName)

			/*for index, attr := range attributes {

				fmt.Sprintf("  The arrt is : index:  =>  %d  ,  the arrt ->   %s   " ,index,attr)

			}*/


			tlist = append(tlist, responseRange.Key)
		}

		jsonRows, err := json.Marshal(tlist)
		if err != nil {
			error_str := fmt.Sprintf("GetListOfInitAucs() operation failed - Unmarshall Error. %s", err)
			fmt.Println(error_str)
			return shim.Error(error_str)
		}



		//fmt.Println("List of Auctions Requested : ", jsonRows)
		return shim.Success(jsonRows)

		//return shim.Success([]byte("dddd"))



	}else if a_parm == "delete" { //删除某个值


		fmt.Println("  ========  delete   ========== %s ",b_parm)

		err := stub.DelState(b_parm)

		if err != nil {
			return shim.Error(" 删除出现错误！！！！！")
		}

		return shim.Success([]byte(" 删除正确！！！！！  "))


	}else if a_parm == "getStatebyrange" { //查询制定范围内的键值 ，目前没有


		fmt.Println("  ========  getStatebyrange   ========== ")

		startkey := b_parm
		endkey := c_parm

		keysIter, err := stub.GetStateByRange( startkey , endkey )

		if err != nil {
			return shim.Error(fmt.Sprintf("keys operation failed. Error accessing state: %s", err))
		}

		defer keysIter.Close()

		var keys []string

		for keysIter.HasNext(){

			response,iterErr := keysIter.Next()

			if iterErr != nil{
				return shim.Error(fmt.Sprintf("find an error %s",iterErr))
			}

			keys = append(keys, response.Key)

		}

		for key, value := range keys {
			fmt.Printf("key %d contains %s\n", key, value)
		}


		jsonKeys, err := json.Marshal(keys)
		if err != nil {
			return shim.Error(fmt.Sprintf("keys operation failed. Error marshaling JSON: %s", err))
		}

		return shim.Success(jsonKeys)



	}else if a_parm == "GetHistoryForKey" { //取单个值的历史记录

		keysIter, err := stub.GetHistoryForKey(b_parm);


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

			keys = append( keys , txid)

		}


		jsonKeys, err := json.Marshal(keys)
		if err != nil {
			return shim.Error(fmt.Sprintf("query operation failed. Error marshaling JSON: %s", err))
		}

		return shim.Success(jsonKeys)


	}else if a_parm == "GetTxID" { //获取当前交易的编号

		txid := stub.GetTxID();
		fmt.Println("  ========  GetTxID   ==========  %s  ",txid)
		return shim.Success([]byte(txid))


	}else if a_parm == "GetTxTimestamp" { //获取当前时间


		txtime,err:= stub.GetTxTimestamp()
		if err != nil {
			fmt.Printf("Error getting transaction timestamp: %s", err)
			return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err))
		}


		tm := time.Unix(txtime.Seconds, 0)

		fmt.Printf("Transaction Time: %v \n ", tm.Format("2006-01-02 03:04:05 PM"))

		return shim.Success([]byte(fmt.Sprint("  time is :   %s   ",tm.Format("2006-01-02 15:04:05"))))


	}else if a_parm == "GetBinding" { //取单个值


		bindtype , err := stub.GetBinding()

		if err != nil {
			fmt.Printf("Error getting transaction timestamp: %s", err)
			return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err))
		}

		fmt.Printf("GetBinding : %s \n ", string(bindtype[:])  )

		return shim.Success(bindtype)

	}else if a_parm == "GetSignedProposal" { //取单个值


		stub.GetSignedProposal()
		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))


	}else if a_parm == "GetCreator" { //取访问请求者的证书

		createbytes , err := stub.GetCreator()

		if err != nil {
			fmt.Printf("Error getting transaction timestamp: %s", err)
			return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err))
		}

		return shim.Success(createbytes)


	}else if a_parm == "GetTransient" { //获取单个值的方式

		transient,err := stub.GetTransient()

		if err != nil {
			fmt.Printf("Error getting transaction timestamp: %s", err)
			return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err))
		}

		for key := range transient {
			fmt.Printf(" the %s    and value is %s  \n", key, string(transient[key]));
		}

		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))

	}else if a_parm == "setloglevel" { //获取单个值的方式

		logleve , _ := shim.LogLevel("dubug")
		shim.SetLoggingLevel(logleve)
		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))

	}else if a_parm == "InvokeChaincode" { //获取单个值的方式


		//queryArgs := util.ToChaincodeArgs("query", "GetCreator","akeym","11234343")

		//parms1 := []string{ "query", "GetCreator","akeym","11234343" }
		parms1 := []string{"query","a"}
		queryArgs := make([][]byte, len(parms1))
		for i, arg := range parms1 {
			queryArgs[i] = []byte(arg)
		}


		response := stub.InvokeChaincode("cc_endfinlshed",queryArgs,"roberttestchannel12")

		if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}

		result := string(response.Payload)



		fmt.Printf(" invoke chaincode  %s " ,result)

		return shim.Success([]byte("success InvokeChaincode  and Not opter !!!!!!!! " + result))

	}else{

		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))

	}








}


func main() {
	err := shim.Start(new(simplechaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


