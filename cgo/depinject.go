package main

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#include <memory.h>
*/
import "C"
import (
	"unsafe"

	"cosmossdk.io/depinject"

	"github.com/pokt-network/libpoktroll-clients/memory"
)

//export Supply
func Supply(goRef C.GoRef, cErr **C.char) C.go_ref {
	toSupply, err := memory.GetGoMem[any](goRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return 0
	}

	return C.go_ref(memory.SetGoMem(depinject.Supply(toSupply)))
}

//export SupplyMany
func SupplyMany(goRefs *C.go_ref, numGoRefs C.int, cErr **C.char) C.go_ref {
	refs := unsafe.Slice(goRefs, numGoRefs)

	//fmt.Printf(">>> refs: %+v\n", refs)
	//val, err := GetGoMem[any](refs[0])
	//if err != nil {
	//	*cErr = C.CString(err.Error())
	//	return 0
	//}
	//fmt.Printf(">>> ref: %+v\n", val)

	var toSupply []any
	for _, ref := range refs {
		valueToSupply, err := memory.GetGoMem[any](memory.GoRef(ref))
		if err != nil {
			*cErr = C.CString(err.Error())
			//*cErr = C.CString(fmt.Sprintf("%+v", err))
			return C.go_ref(memory.NilGoRef)
		}

		toSupply = append(toSupply, valueToSupply)
	}

	return C.go_ref(memory.SetGoMem(depinject.Supply(toSupply...)))
}

//export Config
func Config(goRefs *C.go_ref, numGoRefs C.int, cErr **C.char) C.go_ref {
	refs := unsafe.Slice(goRefs, numGoRefs)

	var configs []depinject.Config
	for _, ref := range refs {
		cfg, err := memory.GetGoMem[depinject.Config](memory.GoRef(ref))
		if err != nil {
			*cErr = C.CString(err.Error())
			return 0
		}

		configs = append(configs, cfg)
	}

	return C.go_ref(memory.SetGoMem(depinject.Configs(configs...)))
}
