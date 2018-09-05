package web_definitions

import (
	"log"
	"net/http"
	"reflect"
)

// ------------------------------------------- Types ------------------------------------------- //

//
// Commonly used reflect.Type objects
//

var InterfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
var InterfaceTypeArray = []reflect.Type{
	InterfaceType,
}

var HandlerType1 = reflect.TypeOf((*http.Handler)(nil)).Elem()
var HandlerType2 = reflect.TypeOf((*func(http.ResponseWriter, *http.Request))(nil)).Elem()
var HandlerTypeArray = []reflect.Type{
	HandlerType1,
	HandlerType2,
}

var StringType = reflect.TypeOf((*string)(nil)).Elem()
var StringTypeArray = []reflect.Type{
	StringType,
}

// ------------------------------------------- Public ------------------------------------------- //

//
// Takes a struct or its pointer. Looks for field with matching name type
// Requires caller to know exact field name and set of possible types
// It is up to the caller to check which of the requested types the returned value is
// To use the returned value, use returnedValue.(<type>) to extract data
// This is UNSAFE. Only use if necessary. This can cause run-time erros (errors compiler will not catch)
//

func GetData(p WebPageInterface, name string, requestedTypes []reflect.Type) interface{} {

	// Search for field by name. Check if valid
	data := reflect.ValueOf(p).Elem().FieldByName(name)
	if !data.IsValid() {
		log.Fatal("Requested named data \"" + name + "\" does not exist.")
	}

	// If field has interface{} type, get the underlying type
	dataType := data.Type()
	if dataType == InterfaceType {
		dataType = data.Elem().Type()
	}

	// Check if field's type matches any of the requested types
	flag := false
	for _, requestType := range requestedTypes {
		if dataType == requestType {
			flag = true
			break
		}
	}
	if flag == false {
		log.Fatal("Requested wrong data type. Data named \"" + name + "\" had type: " + dataType.String())
	}

	return data.Interface()
}
