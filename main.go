package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// ATS is a high level smart contract that ATSs together business artifact based smart contracts
type ATS struct {

}

// AssetDetails is for storing Asset Details

type AssetDetails struct{	
	assetId string `json:"assetId"`
	Category string `json:"category"`
	CType string `json:"ctype"`
	CName1 string `json:"cName1"`
	CName2 string `json:"cName2"`
	Doc string `json:"doc"`
	MafCode string `json:"mafcode"`
	Country string `json:"country"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	CreatedBy string `json:"createdBy"`
	TotalItems string `json:"totalItems"`
}

// Transaction is for storing transaction Details

type Transaction struct{	
	TrxId string `json:"trxId"`
	TimeStamp string `json:"timeStamp"`
	assetId string `json:"assetId"`
	Source string `json:"source"`
	Items string `json:"items"`
	Trxntype string `json:"trxntype"`
	TrxnSubType string `json:"trxnSubType"`
	Event string `json:"event"`
}


// GetItem is for storing retreived Get the total Items

type GetLoadCarried struct{	
	TotalItem string `json:"totalItem"`
}

// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	



// Init initializes the smart contracts
func (t *CTS) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("AssetDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("AssetDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "assetId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "category", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ctype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "cName1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "cName2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "doc", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "mafcode", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "totalItem", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating AssetDetails.")
	}
	


	// Check if table already exists
	_, err = stub.GetTable("Transaction")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("Transaction", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assetId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Items", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxntype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxnSubType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "event", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
		
	// setting up the Asset role
	stub.PutState("Asset_type1_1", []byte("ABC"))
	stub.PutState("Asset_type1_2", []byte("XYZ"))
	stub.PutState("Asset_type1_3", []byte("PQR"))
	stub.PutState("Asset_type1_4", []byte("IJK"))	
	
	return nil, nil
}
	

	
//registerasset to register a asset
func (t *FFP) registerAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 12 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
		}
		
		assetId:=args[0]
		category:=args[1]
		ctype:=args[2]
		cName1:=args[3]
		cName2:=args[4]
		doc:=args[5]
		mafcode:=args[6]
		country:=args[7]
		address:=args[8]
		city:=args[9]
		zip:=args[10]
		
		assignerOrg1, err := stub.GetState(args[11])
		assignerOrg := string(assignerOrg1)
		
		createdBy:=assignerOrg
		totalItem:="0"


		// Insert a row
		ok, err := stub.InsertRow("AssetDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: assetId}},
				&shim.Column{Value: &shim.Column_String_{String_: category}},
				&shim.Column{Value: &shim.Column_String_{String_: ctype}},
				&shim.Column{Value: &shim.Column_String_{String_: cName1}},
				&shim.Column{Value: &shim.Column_String_{String_: cName2}},
				&shim.Column{Value: &shim.Column_String_{String_: doc}},
				&shim.Column{Value: &shim.Column_String_{String_: mafcode}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: totalItem}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}



// add or delete Items and insert the transaction(irrespective of org)
func (t *FFP) addDeleteLoadCarried(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	trxId := args[0]
	timeStamp:=args[1]
	assetId := args[2]
	
	assignerOrg1, err := stub.GetState(args[3])
	assignerOrg := string(assignerOrg1)
	
	source := assignerOrg
	items := args[4]
	trxntype := args[5]
	trxnSubType := args[6]
	event := args[7]
	
	newItems, _ := strconv.ParseInt(items, 10, 0)
	
	//whether ADD_PENDING, DELETE_PENDING 
	if trxnSubType == "ADD_PENDING" || trxnSubType == "DELETE_PENDING"{
		newItems = 0
	}
	

	// Get the row pertaining to this assetId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assetId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving Asset with assetId %s. Error %s", assetId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	newLCItem := row.Columns[12].GetString_()
	
	if trxntype=="add"{
		earlierLoadCarried:=row.Columns[12].GetString_()
		earlierLC, _:=strconv.ParseInt(earlierLoadCarried, 10, 0)
		newLCItem = strconv.Itoa(int(earlierLC) + int(newItems))
	}else if trxntype=="delete"{
	
		earlierLoadCarried:=row.Columns[12].GetString_()
		earlierLC, _:=strconv.ParseInt(earlierLoadCarried, 10, 0)
		newLCItemtoTest := int(earlierLC) - int(newItems)
		
		if newLCItemtoTest < 0 {
			return nil, errors.New("can't deduct as the resulting LC becoming less than zero.")
		}
		newLCItem = strconv.Itoa(int(earlierLC) - int(newItems))
	}else{
		return nil, fmt.Errorf("Error: Failed retrieving Asset with assetId %s. Error %s", assetId, err.Error())
	}
	
	
	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this assetId
	err = stub.DeleteRow(
		"AssetDetails",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	
	//assetId := row.Columns[0].GetString_()
	
	category := row.Columns[1].GetString_()
	ctype := row.Columns[2].GetString_()
	cName1 := row.Columns[3].GetString_()
	cName2 := row.Columns[4].GetString_()
	doc := row.Columns[5].GetString_()
	mafcode := row.Columns[6].GetString_()
	country := row.Columns[7].GetString_()
	address := row.Columns[8].GetString_()
	city := row.Columns[9].GetString_()
	zip := row.Columns[10].GetString_()
	createdBy := row.Columns[11].GetString_()
	totalItem := newLCItem


		// Insert a row
		ok, err := stub.InsertRow("AssetDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: assetId}},
				&shim.Column{Value: &shim.Column_String_{String_: category}},
				&shim.Column{Value: &shim.Column_String_{String_: ctype}},
				&shim.Column{Value: &shim.Column_String_{String_: cName1}},
				&shim.Column{Value: &shim.Column_String_{String_: cName2}},
				&shim.Column{Value: &shim.Column_String_{String_: doc}},
				&shim.Column{Value: &shim.Column_String_{String_: mafcode}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: totalItem}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		
		//inserting the transaction
		
		// Insert a row
		ok, err = stub.InsertRow("Transaction", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: assetId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: Items}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: event}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}		
	return nil, nil

}


//get the LoadCarrieds against the assetId (irrespective of org)
func (t *FFP) getLoadCarried(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting assetId to query")
	}

	assetId := args[0]
	

	// Get the row pertaining to this assetId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assetId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the assetId " + assetId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the assetId " + assetId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := GetLoadCarried{}
	
	res2E.totalItem = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



//get all transaction against the assetId (depends on org)
func (t *FFP) getTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting assetId to query")
	}

	assetId := args[0]
	assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.assetId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Items = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.event = row.Columns[7].GetString_()
		
		if newApp.assetId == assetId && newApp.Source == assignerOrg{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}




//get All transaction against assetId (irrespective of org)
func (t *FFP) getAllTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting assetId to query")
	}

	assetId := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.assetId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Items = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.event = row.Columns[7].GetString_()
		
		if newApp.assetId == assetId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// to get the deatils of a Asset against assetId (for internal testing, irrespective of org)
func (t *FFP) getAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting assetId to query")
	}

	assetId := args[0]
	

	// Get the row pertaining to this assetId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assetId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + assetId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + assetId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	res2E := AssetDetails{}
	
	res2E.assetId = row.Columns[0].GetString_()
	res2E.category = row.Columns[1].GetString_()
	res2E.ctype = row.Columns[2].GetString_()
	res2E.cName1 = row.Columns[3].GetString_()
	res2E.cName2 = row.Columns[4].GetString_()
	res2E.doc = row.Columns[5].GetString_()
	res2E.mafcode = row.Columns[6].GetString_()
	res2E.Country = row.Columns[7].GetString_()
	res2E.Address = row.Columns[8].GetString_()
	res2E.City = row.Columns[9].GetString_()
	res2E.Zip = row.Columns[10].GetString_()
	res2E.CreatedBy = row.Columns[11].GetString_()
	res2E.totalItem = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// verify the Asset is present or not (for internal testing, irrespective of org)
func (t *FFP) verifyAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting assetId to query")
	}

	assetId := args[0]
	doc := args[1]
	

	// Get the row pertaining to this assetId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: assetId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + assetId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + assetId + "\"}"
		return nil, errors.New(jsonResp)
	}

	Assetdoc := row.Columns[5].GetString_()
	
	res2E := VerifyU{}
	
	if doc == Assetdoc{
		res2E.Result="success"
	}else{
		res2E.Result="failed"
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



// Invoke invokes the chaincode
func (t *FFP) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerasset" {
		t := FFP{}
		return t.registerasset(stub, args)	
	} else if function == "addDeleteLoadCarried" { 
		t := FFP{}
		return t.addDeleteLoadCarried(stub, args)
	}

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *FFP) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getLoadCarried" {
		t := FFP{}
		return t.getLoadCarried(stub, args)		
	} else if function == "getTransaction" { 
		t := FFP{}
		return t.getTransaction(stub, args)
	}else if function == "getAllTransaction" { 
		t := FFP{}
		return t.getAllTransaction(stub, args)
	} else if function == "getAsset" { 
		t := FFP{}
		return t.getasset(stub, args)
	}else if function == "verifyAsset" { 
		t := FFP{}
		return t.verifyasset(stub, args)
	}
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(FFP))
	if err != nil {
		fmt.Printf("Error starting FFP: %s", err)
	}
} 
