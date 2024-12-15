package main

/*
#include <memory.h>
#include <context.h>
#include <protobuf.h>
*/
import "C"
import (
	"cosmossdk.io/depinject"

	"github.com/pokt-network/libpoktroll-clients/memory"
	multiclient "github.com/pokt-network/libpoktroll-clients/multi_client"
)

// TODO_IN_THIS_COMMIT: godoc...
//
//export NewQueryClient
func NewQueryClient(selfRef C.go_ref, queryNodeWebsocketURL *C.char, cErr **C.char) C.go_ref {
	deps, err := memory.GetGoMem[depinject.Config](memory.GoRef(selfRef))
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(memory.NilGoRef)
	}

	multiClient, err := multiclient.NewMultiQueryClient(deps, C.GoString(queryNodeWebsocketURL))
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(memory.NilGoRef)
	}

	return C.go_ref(memory.SetGoMem(multiClient))
}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetSharedParams
func QueryClient_GetSharedParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetApplicationParams
func QueryClient_GetApplicationParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetGetewayParams
func QueryClient_GetGetewayParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetSupplierParams
func QueryClient_GetSupplierParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetSessionParams
func QueryClient_GetSessionParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetServiceParams
func QueryClient_GetServiceParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetProofParams
func QueryClient_GetProofParams(selfRef C.go_ref, op *C.AsyncOperation) {

}

// TODO_IN_THIS_COMMIT: godoc...
//
//export QueryClient_GetTokenomicsParams
func QueryClient_GetTokenomicsParams(selfRef C.go_ref, op *C.AsyncOperation) {

}
