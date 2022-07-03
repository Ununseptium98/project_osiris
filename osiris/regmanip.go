package osiris

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"golang.org/x/sys/windows/registry"
)

func TestRequestKey(requestedKey string) error {

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, requestedKey, registry.QUERY_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	return nil
}

func TryAccess(requestedKey string, requestedRootKey int) {
	/*
		Takes the request Key as string then returns its value

	*/

	k, err := registry.OpenKey(getRootKey(requestedRootKey), requestedKey, registry.QUERY_VALUE)

	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	values, err := k.ReadValueNames(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range values {
		fmt.Printf("Name is : %s \n", name)
	}

	s, _, err := k.GetStringValue("SystemRoot")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Windows system root is %q\n", s)

}

func getRootKey(codifiedKey int) registry.Key {
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
		log.Fatal(errors.New("input int unassociated, default = 1 (1,2 or 3 only)"))
		return registry.LOCAL_MACHINE
	}

}

func TestKey(requestedKey string, requestedRootKey int) {

	k, err := registry.OpenKey(getRootKey(requestedRootKey), requestedKey, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)

	if err != nil {
		log.Fatal()
	}
	defer k.Close()

	subkeysTable, err := k.ReadSubKeyNames(-1)

	if err != nil {

		log.Fatal(err)
	}

	for iterator, subkey := range subkeysTable {

		fmt.Printf("subkey %d => %s \n ", iterator, subkey)
		s, _, err := k.GetStringValue(subkey)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("subkey %d => %s value => %s \n ", iterator, subkey, s)

	}

}

func GetAllKeyValuesJSON(requestedKey string, requestedRootKey int, printJSON bool, requestID string) map[string]string {
	/*
		List all values of a given key or subkey

	*/

	k, err := registry.OpenKey(getRootKey(requestedRootKey), requestedKey, registry.QUERY_VALUE)

	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	keyValueMap := make(map[string]string)

	valueNames, err := k.ReadValueNames(-1) // reads all names in the key
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range valueNames {

		_, valtype, err := k.GetValue(name, nil)

		if err != nil {
			log.Fatal(err)
		}

		switch valtype {
		case registry.NONE:
		case registry.SZ:
			value, _, err := k.GetStringValue(name)
			if err != nil {
				log.Fatal(err)
			}

			keyValueMap[name] = value
			fmt.Printf("----\nName > %s\n value = %s \n", name, value)

		case registry.EXPAND_SZ:
			s, _, err := k.GetStringValue(name)
			if err != nil {
				log.Fatal(err)
			}

			value, err := registry.ExpandString(s)
			if err != nil {
				log.Fatal(err)
			}

			keyValueMap[name] = value
			fmt.Printf("----\nName > %s\n value =  %s \n", name, value)

		case registry.DWORD, registry.QWORD:
			value, _, err := k.GetIntegerValue(name)
			if err != nil {
				log.Fatal(err)
			}

			keyValueMap[name] = strconv.FormatInt(int64(value), 10)
			fmt.Printf("----\nName > %s\n value = %s \n", name, strconv.FormatInt(int64(value), 10))

		case registry.BINARY:
			value, _, err := k.GetBinaryValue(name)
			if err != nil {
				log.Fatal(err)
			}

			keyValueMap[name] = string(value)
			fmt.Printf("----\nName > %s\n value = %x \n", name, value)

		case registry.MULTI_SZ:
			value, _, err := k.GetStringsValue(name)
			if err != nil {
				log.Fatal(err)
			}
			keyValueMap[name] = value[0]
			fmt.Printf("----\nName > %s\n value =  %s \n", name, value)

		default:
			value := "Unhandled Value type"
			keyValueMap[name] = value
			fmt.Printf("----\nName is : %s  = %s \n", name, value)

		}

		value, _, err := k.GetStringValue(name)

		if err != nil {

			continue
		}

		keyValueMap[name] = value //associates ValueName name with its value

	}

	if printJSON {
		json, _ := json.Marshal(keyValueMap)

		ioutil.WriteFile(requestID+"_KeyValues.json", json, os.ModePerm)

		fmt.Println("OUTPUT : " + requestID + "_KeyValues.json")

	}

	return keyValueMap

}

func GetSubKeyValues(requestedKey string, requestedRootKey int, subkeyNumber int) {
	/*
		Outputs values from subkeys in the given requestedKey
	*/

	k, err := registry.OpenKey(getRootKey(requestedRootKey), requestedKey, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)

	if err != nil {
		log.Fatal()
	}
	defer k.Close()

	subkeysTable, err := k.ReadSubKeyNames(subkeyNumber) // fetches subkeys

	if err != nil {

		log.Fatal(err)
	}

	for iterator, subkey := range subkeysTable {

		fmt.Printf("Iteration number %d with %s \n", iterator, requestedKey+`\`+subkey)

		GetAllKeyValuesJSON(requestedKey+`\`+subkey, requestedRootKey, false, "")

	}

}
