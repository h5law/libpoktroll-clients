package main

import "C"
import (
	"context"

	"cosmossdk.io/depinject"

	"github.com/pokt-network/poktroll/pkg/client/block"

	"github.com/pokt-network/libpoktroll-clients/memory"
)

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include <memory.h>
*/
import "C"

//export NewBlockClient
func NewBlockClient(depsRef C.go_ref, cErr **C.char) C.go_ref {
	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	deps, err := memory.GetGoMem[depinject.Config](memory.GoRef(depsRef))
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(memory.NilGoRef)
	}

	blockClient, err := block.NewBlockClient(ctx, deps)
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.go_ref(memory.NilGoRef)
	}

	return C.go_ref(memory.SetGoMem(blockClient))
}
