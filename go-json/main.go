/**
	Marshalling and Unmarshalling Data
	Strutured and Unstructured JSON data
	JSON STruct Tags and field name
**/

// Marshal -> enode
// Unmarshal -> decode

/**
	Unmarshalling (Parsing) Raw JSON strings.

	The "Unmarshall" function provided by Go's JSON standard library
	let us parse rw JSON data in the form of []byte variables.


	We can convert JSON strings into bytes and unmarshall the data into a variables address

**/

/**
	There are 2 key terminlogies to note when working with JSON in GO:
	1. Marhsalling: the act of converting a Go data structure into valid JSON.
	2. UNmarshalling: the act of parsing a valid JSON string into a data structure in GO
**/

// package main

// import "encoding/json"

// myJsonString := `{"some":"json"}`

// json.Unmarshal([]byte(myJsonString), &myStoredVariable)

/**
	Lets look at the different variable types for myStoredVariable, and when you should
	use them.

	There are 2 types of data you will enounter when working with JSON:
	1.Structured data
	2. Unstructured data

	Structured data (Decoding JSON Into structs)

	"Structured data" referes to data where you know the format beforehand.
	For example, lets say you have a bird object, where each bird has a species field and
	a description field:

	{
		"species": "pigeon",
		"description": "likes to perch on rocks"
	}

	To work with this kinnd of data, create a struct that mirrors the data you want to parse.
	In our case we will create a bird struct whihc has a Species and Description attribute:


		type Bird struct {
			Species string
			Description string
		}

		ANd unmarshall it as follows:

		birdJson := `{
			"species": "pigeon",
			"description": "likes to perch on rocks"
		}`

		var bird Bird

		json.Unmarshall([]byte(birdJson), &bird)
		fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)

**/

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Bird struct {
// 	Species     string `json:"species"`
// 	Description string `json:"description"`
// }

// func main() {
// 	birdJSON := `{
// 		"species": "pigeon",
// 		"description": "likes to perch on rocks"
// 	}`

// 	var bird Bird

// 	err := json.Unmarshal([]byte(birdJSON), &bird)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Bird", bird)
// 	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)
// }

/**
	By convention, GO uses the same title cased attribute names as are pressent in the case
	insensitive JSON properties. So the Species attribute in our Bird struct will map
	to the species or Species or sPeIEs JSON property.
**/

/**
	JSON Arrays
	birdJson := `[
		{"species":"pigeon","description":"likes to perch on rocks"},
		{"species":"eagle","description":"bird of prey"}
		]`

	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds: %v\n", birds)


	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]

**/

/**
	Nested objects

	NOw consider the case when we have a property called Dimensions,
	that measures the Height and Length of the bird in question:

	{
	"species": "pigeon",
	"description": "likes to perch on rocks"
	"dimensions": {
		"height": 24,
		"width": 10
    }

	As with out previous examples, we need to mirror the structure of the
	objet in our GO code. TO add a nested dimensions object, we can create a dimensions
	struct and add it to our Bird struct as follows:

	type DImensions struct {
		Height int
		Width int
	}

	type Bird struct {
		Species string
		Description string
		Dimensions Dimensions
	}

	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`

	var birds Bird

	json.Unmarshall([]byte(birdJson), &birds)
	fmt.Printf("Birds: %v\n", birds)
}


**/

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Dimensions struct {
// 	Height int
// 	Width  int
// }

// type Bird struct {
// 	Species     string
// 	Description string
// 	Dimensions  Dimensions
// }

// func main() {

// 	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`

// 	var birds Bird

// 	json.Unmarshal([]byte(birdJson), &birds)
// 	fmt.Println(birds)

// }

/**
	Primitive types

	We mostly deal with complex objects or arrays when working with JSON, but data like 3, 3.1412 and "birds" are also valid JSON strings.

	We can unmarshal these values to their corresponding data type in Go by using primitive types:

	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n) //3

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi) //3.1412

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str) //bird

**/

/**
	Custom parsing logic

	{
		"species": "pigeon",
		"description": "likes to perch on rocks",
		"dimensions": "24x10"
	}

	We can modify the Dimensions type to implement the Unmarshaler interface whih will have
	custom parsing logic for out data:

	type Dimensions struct {
	Height int
	Width  int
}

// unmarshals a JSON string with format
// "heightxwidth" into a Dimensions struct
func (d *Dimensions) UnmarshalJSON(data []byte) error {
	// the "data" parameter is expected to be JSON string as a byte slice
	// for example, `"20x30"`

	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}
	// remove the quotes
	s := string(data)[1 : len(data)-1]
	// split the string into its two parts
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}
	// convert the two parts into ints
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("dimension height must be an int")
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("dimension width must be an int")
	}
	// assign the two ints to the Dimensions struct
	d.Height = height
	d.Width = width
	return nil
}


**/

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// type Dimensions struct {
// 	Height int
// 	Width  int
// }

// type Bird struct {
// 	Species     string
// 	Description string
// 	Dimensions  Dimensions
// }

// // unmarshals a JSON string with format
// // "heightxwidth" into a Dimensions struct
// func (d *Dimensions) UnmarshalJSON(data []byte) error {
// 	// the "data" parameter is expected to be JSON string as a byte slice
// 	// for example, `"20x30"`

// 	if len(data) < 2 {
// 		return fmt.Errorf("dimensions string too short")
// 	}
// 	// remove the quotes
// 	s := string(data)[1 : len(data)-1]
// 	// split the string into its two parts
// 	parts := strings.Split(s, "x")
// 	if len(parts) != 2 {
// 		return fmt.Errorf("dimensions string must contain two parts")
// 	}
// 	// convert the two parts into ints
// 	height, err := strconv.Atoi(parts[0])
// 	if err != nil {
// 		return fmt.Errorf("dimension height must be an int")
// 	}
// 	width, err := strconv.Atoi(parts[1])
// 	if err != nil {
// 		return fmt.Errorf("dimension width must be an int")
// 	}
// 	// assign the two ints to the Dimensions struct
// 	d.Height = height
// 	d.Width = width
// 	return nil
// }

// func main() {

// 	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":"24x10"}`

// 	var bird Bird

// 	json.Unmarshal([]byte(birdJson), &bird)
// 	fmt.Println(bird)
// }

/**
	JSON Struct tags - Custom field names

	We sometimes want a different attribute name than the one provided in your JSON data.
	For example, consider the below data:

	{
		"birdType": "pigeon",
		"what it does": "likes to perch on rocks"
	}

	Here, we would prefer birdType to remain as the Species attribute in our GO code.
	It is also not possible for us to provide a suitable attribute name for a key like "what it does".

	To solve this, we can use struct filed tags:

	type Bird struct {
		Species string `json:"birdType"`
		Description string `json:"what it does"`
	}

	Now we an explicitly tell our ode whih JSON property to map to whihc attribute

	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks}

**/

/**
	Decoding JSON to Maps - Unstructured data

	If we dont know the structure of our JSON properities beforehand, we cannot use structs
	to unmarshal our data

	INsteaad we can use "mapps". COnsider some JSON of the form:

	{
		"birds":{
			"pigeon":"likes to perch on rocks",
			"eagle":"bird of prey"
		},
		"animals":"none"
	}

	There is no struct we can build to represent the above data for all cases
	sine the keys orresponding to the birds can change, which will change the structure.

	To deal with this case we reate a map of strings to empty interfaces:

	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	birds := result["birds"].(map[string]any)

	for key, value := range birds {
	// Each value is an `any` type, that is type asserted as a string
	fmt.Println(key, value.(string))
	}

**/

/**
	Validating JSON Data

	In real-world applications, we may sometimes get invalid(or incomplete) JSON data.
	Eg. where some of the data us ut off, and he resulting JSON string is invalid.

		{
			"birds": {
				"pigeon":"likes to perch on rocks",
				"eagle":"bird of prey"

	In actual applications, this may happen due to network errors or incomplete data written to files

	If we try to unmarshal this, our ode will panic:

	panic: interface conversion: interface {} is nil, not map[string]interface {}

	We an of ourse handle the panic and recover from our code, but this would not be idiiomatic or readable

	Instead we can use the json.Valid function to checck the validity of our JSON data.

		if !json.Valid([]byte(birdJson)) {
			fmt.Println("JSON is not valid")
			return
		}

	Now our ode will return earlt and give the output:

	invalid JSON string: {"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"
**/

/**
	Marshalling JSON data

	Marshalling is the process of transforming structured data into a serializable JSON string.
	SImilar to unmarshalling, we can marsahl data into structs, maps and slices.

**/

package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// data, _ := json.Marshal(pigeon)
	data, _ := json.Marshal([]*Bird{pigeon, pigeon})

	fmt.Println(string(data))
	// fmt.Println(data)
}

/**
	If we want to always ignore a field,, we an use the json: "-" strut tag to denote that
	we never want this field included

	type Bird struct {
		Species     string `json:"-"`
	}
**/

/** Marshalling Slices

	This isnt much different from strcuts. We just need to pas the slie or array
	to the json.Marshal function, and it will encode data like you expect:

	data, _ := json.Marshal([]*Bird{pigeon, pigeon})

**/

/** Marshalling Maps

	We can use maps to encode unstructured data.

	The keys of the map need to be strings, or a type that an onvert to strings. The
	values can be any serializable type.

**/

/**
	Custom encoding logic

	type Dimensions struct {
		Height int
		Width  int
	}

	// marshals a Dimensions struct into a JSON string
	// with format "heightxwidth"
	func (d Dimensions) MarshalJSON() ([]byte, error) {
		return []byte(fmt.Sprintf(`"%dx%d"`, d.Height, d.Width)), nil
	}

	func main() {
		bird := Bird{
			Species: "pigeon",
			Dimensions: Dimensions{
				Height: 24,
				Width:  10,
			},
		}
		birdJson, _ := json.Marshal(bird)
		fmt.Println(string(birdJson))
		// {"Species":"pigeon","Dimensions":"24x10"}
	}

**/

/**
	Printing formatted (pretty-printed) JSON

	BY default the JSON enoder will not add any whitespace to the encoded JSON string.
	This is done to reduce the size of the JSON string, and is useful
	when ending data over the network.

	But if you want to print the JSON string to the consile, or write to a file, you
	may want to add whitespace to make it more readble We can do this by using
	the json.MarshalIndent function.

	bird := Bird{
	Species: "pigeon",
	Description: "likes to eat seed",
}

	// The second parameter is the prefix of each line, and the third parameter
	// is the indentation to use for each level
	data, _ := json.MarshalIndent(bird, "", "  ")
	fmt.Println(string(data))

	Output:
		{
			"Species": "pigeon",
			"Description": "likes to eat seed"
		}
**/

/**
		JSON -> String: Marshal
		String -> JSON: Unmarshal
		JSON -> Stream: Encode
		Stream -> JSON: Decode

		Marshaling and Encoding are of course different concepts, better addressed on Wikipedia (or elsewhere). But in short, objects are marshaled into JSON encoded strings.

		Also don't let the Golang json.NewEncoder / (json.Encoder).Encode and json.Marshal methods confuse you. They both marshal objects into JSON encoded strings. The difference being the Encoder, first marshals the object to a JSON encoded string, then writes that data to a buffer stream (or Data Buffer on Wikipedia). The Encoder therefore, uses more code and memory overhead than the simpler json.Marshal.

		You can also see this in the Golang source code:

			Marshal: https://golang.org/src/encoding/json/encode.go?s=6458:6501#L148
			Encode: https://golang.org/src/encoding/json/stream.go?s=5070:5117#L191

		Typically, if you need to send the JSON encoded string to a file system, or as an HTTP response, you may need the use of a buffer stream. However, you can also send this JSON encoded string without a buffer stream using a pipe.
**/
