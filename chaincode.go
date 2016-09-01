/*
Copyright Mojing Inc. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"math/rand"
	"strconv"
)

type randomChaincode struct {
}

func (c *randomChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, stub.PutState("random", []byte(strconv.Itoa(rand.Intn(100))))
}

func (c *randomChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	r, err := stub.GetState("random")
	if err != nil {
		return nil, err
	}
	fmt.Printf("get r: %v\n", r)

	return nil, stub.PutState("random", []byte("100"))
}

func (c *randomChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return stub.GetState("random")
}

func main() {
	err := shim.Start(new(randomChaincode))
	if err != nil {
		fmt.Printf("start chaincode err: %v\n", err)
	}
}
