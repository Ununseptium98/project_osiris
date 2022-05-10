package regmanip

import(
	"golang.org/x/sys/windows/registry"
	"fmt"
	"log"
	"errors"
)

func requestKey(requestedKey string){

	log.SetPrefix("RequestKet: ")
	log.SetFlags(0)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, requestedKey, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	
}



func TryAccess(requestedKey string, requestedRootKey int ){ 
	/*
		Takes the request Key as string then returns its value

	*/

	k, err := registry.OpenKey(getRootKey(requestedRootKey), requestedKey, registry.QUERY_VALUE)
	//k, err := registry.OpenKey(registry.LOCAL_MACHINE,  ,registry.QUERY_VALUE)
	
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	s, _, err := k.GetStringValue("SystemRoot")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows system root is %q\n", s)

}


func getRootKey(codifiedKey int) (registry.Key) { 
	/*
		Associate an integer in input to a Windows registry root key
		1 : LOCAL_MACHINE
		2 : USERS
		3 : CURRENT_USER
		default : LOCAL_MACHINE
		

	*/

	//keyMapping := make(map[int]string)

	switch codifiedKey {
	case 1:
		return registry.LOCAL_MACHINE 
	case 2:
		return registry.USERS
	case 3: 
		return registry.CURRENT_USER
	default:
		log.Fatal(errors.New("Input int unassociated, default = 1 (1,2 or 3 only)"))
		return registry.LOCAL_MACHINE
	}
	
}

