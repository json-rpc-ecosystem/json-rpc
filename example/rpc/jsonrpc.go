// Code generated by jsonrpcgen. DO NOT EDIT.

package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      int             `json:"id"`
}

type jsonRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *jsonRPCError   `json:"error,omitempty"`
	ID      int             `json:"id"`
}

type jsonRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Client struct {
	httpClient *http.Client
	endpoint string

	
	Arith Arith
	
	Greeter Greeter
	
}

func NewClient(httpClient *http.Client, endpoint string) *Client {
	c := Client{
		httpClient: httpClient,
		endpoint: endpoint,
	}

	
	c.Arith = &internalArithClient{client: &c, endpoint: "/arith"}
	
	c.Greeter = &internalGreeterClient{client: &c, endpoint: "/greeter"}
	

	return &c
}


type internalArithClient struct {
	client *Client
	endpoint string
}



func (c *internalArithClient) Add(params *ArithAddParams) (*ArithAddResult, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ArithAddParams: %v", err)
	}

	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "Add",
		Params:  data,
	}

	rpcreq := bytes.NewBuffer(nil)
	err = json.NewEncoder(rpcreq).Encode(req)
	if err != nil {
		return nil, fmt.Errorf("error encoding Arith Add RPC request: %v", err)
	}

	res, err := c.client.httpClient.Post(c.client.endpoint+c.endpoint, "application/json", rpcreq)
	if err != nil {
		return nil, fmt.Errorf("error POSTing Arith Add RPC request: %v", err)
	}

	var rpcres *jsonRPCResponse
	err = json.NewDecoder(res.Body).Decode(&rpcres)
	if err != nil {
		return nil, fmt.Errorf("error decoding Arith.Add RPC response: %v", err)
	}

	var result *ArithAddResult
	err = json.Unmarshal(rpcres.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling ArithAddResult: %v", err)
	}

	return result, nil
}

func (c *internalArithClient) Pow(params *ArithPowParams) (*ArithPowResult, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ArithPowParams: %v", err)
	}

	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "Pow",
		Params:  data,
	}

	rpcreq := bytes.NewBuffer(nil)
	err = json.NewEncoder(rpcreq).Encode(req)
	if err != nil {
		return nil, fmt.Errorf("error encoding Arith Pow RPC request: %v", err)
	}

	res, err := c.client.httpClient.Post(c.client.endpoint+c.endpoint, "application/json", rpcreq)
	if err != nil {
		return nil, fmt.Errorf("error POSTing Arith Pow RPC request: %v", err)
	}

	var rpcres *jsonRPCResponse
	err = json.NewDecoder(res.Body).Decode(&rpcres)
	if err != nil {
		return nil, fmt.Errorf("error decoding Arith.Pow RPC response: %v", err)
	}

	var result *ArithPowResult
	err = json.Unmarshal(rpcres.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling ArithPowResult: %v", err)
	}

	return result, nil
}

func (c *internalArithClient) IsNegative(params *ArithIsNegativeParams) (*ArithIsNegativeResult, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ArithIsNegativeParams: %v", err)
	}

	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "IsNegative",
		Params:  data,
	}

	rpcreq := bytes.NewBuffer(nil)
	err = json.NewEncoder(rpcreq).Encode(req)
	if err != nil {
		return nil, fmt.Errorf("error encoding Arith IsNegative RPC request: %v", err)
	}

	res, err := c.client.httpClient.Post(c.client.endpoint+c.endpoint, "application/json", rpcreq)
	if err != nil {
		return nil, fmt.Errorf("error POSTing Arith IsNegative RPC request: %v", err)
	}

	var rpcres *jsonRPCResponse
	err = json.NewDecoder(res.Body).Decode(&rpcres)
	if err != nil {
		return nil, fmt.Errorf("error decoding Arith.IsNegative RPC response: %v", err)
	}

	var result *ArithIsNegativeResult
	err = json.Unmarshal(rpcres.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling ArithIsNegativeResult: %v", err)
	}

	return result, nil
}


type internalGreeterClient struct {
	client *Client
	endpoint string
}



func (c *internalGreeterClient) SayHello(params *GreeterSayHelloParams) (*GreeterSayHelloResult, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("error marshalling GreeterSayHelloParams: %v", err)
	}

	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "SayHello",
		Params:  data,
	}

	rpcreq := bytes.NewBuffer(nil)
	err = json.NewEncoder(rpcreq).Encode(req)
	if err != nil {
		return nil, fmt.Errorf("error encoding Greeter SayHello RPC request: %v", err)
	}

	res, err := c.client.httpClient.Post(c.client.endpoint+c.endpoint, "application/json", rpcreq)
	if err != nil {
		return nil, fmt.Errorf("error POSTing Greeter SayHello RPC request: %v", err)
	}

	var rpcres *jsonRPCResponse
	err = json.NewDecoder(res.Body).Decode(&rpcres)
	if err != nil {
		return nil, fmt.Errorf("error decoding Greeter.SayHello RPC response: %v", err)
	}

	var result *GreeterSayHelloResult
	err = json.Unmarshal(rpcres.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling GreeterSayHelloResult: %v", err)
	}

	return result, nil
}



type Server struct {
    
    Arith Arith
    
    Greeter Greeter
    
}

func (s *Server) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/arith", s.ArithHandler)
	
	mux.HandleFunc("/greeter", s.GreeterHandler)
	
	return mux
}



func (s *Server) ArithHandler(w http.ResponseWriter, r *http.Request) {
	var request jsonRPCRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	response := jsonRPCResponse{
		JSONRPC: request.JSONRPC,
		ID:      request.ID,
	}

	var resultData json.RawMessage


	switch request.Method {
	
	case "Add":
		var params ArithAddParams
		err = json.Unmarshal(request.Params, &params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		result, err := s.Arith.Add(&params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		resultData, err = json.Marshal(result)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}
	
	case "Pow":
		var params ArithPowParams
		err = json.Unmarshal(request.Params, &params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		result, err := s.Arith.Pow(&params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		resultData, err = json.Marshal(result)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}
	
	case "IsNegative":
		var params ArithIsNegativeParams
		err = json.Unmarshal(request.Params, &params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		result, err := s.Arith.IsNegative(&params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		resultData, err = json.Marshal(result)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}
	
	default:
		response.Error = &jsonRPCError{Code: -32601, Message: "Method not found"}
	}

	response.Result = resultData

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}


func (s *Server) GreeterHandler(w http.ResponseWriter, r *http.Request) {
	var request jsonRPCRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	response := jsonRPCResponse{
		JSONRPC: request.JSONRPC,
		ID:      request.ID,
	}

	var resultData json.RawMessage


	switch request.Method {
	
	case "SayHello":
		var params GreeterSayHelloParams
		err = json.Unmarshal(request.Params, &params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		result, err := s.Greeter.SayHello(&params)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}

		resultData, err = json.Marshal(result)
		if err != nil {
			response.Error = &jsonRPCError{Message: err.Error()}
		}
	
	default:
		response.Error = &jsonRPCError{Code: -32601, Message: "Method not found"}
	}

	response.Result = resultData

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}





// ArithAddParams...
type ArithAddParams struct {
    
    Nums []float64 `json:"Nums"`
    
}

// ArithAddResult...
type ArithAddResult struct {
    
    Sum float64 `json:"Sum"`
    
}

// ArithPowParams...
type ArithPowParams struct {
    
    Base float64 `json:"Base"`
    
    Pow float64 `json:"Pow"`
    
}

// ArithPowResult...
type ArithPowResult struct {
    
    Num float64 `json:"Num"`
    
}

// ArithIsNegativeParams...
type ArithIsNegativeParams struct {
    
    Num float64 `json:"Num"`
    
}

// ArithIsNegativeResult...
type ArithIsNegativeResult struct {
    
    Negative bool `json:"Negative"`
    
}


// Arith exposes basic arithmatic functions
type Arith interface {
    
    Add(*ArithAddParams) (*ArithAddResult, error)
    
    Pow(*ArithPowParams) (*ArithPowResult, error)
    
    IsNegative(*ArithIsNegativeParams) (*ArithIsNegativeResult, error)
    
}



// GreeterSayHelloParams...
type GreeterSayHelloParams struct {
    
    From string `json:"From"`
    
    To string `json:"To"`
    
}

// GreeterSayHelloResult...
type GreeterSayHelloResult struct {
    
    Message string `json:"Message"`
    
}


// Greeter does lots of greetings
type Greeter interface {
    
    SayHello(*GreeterSayHelloParams) (*GreeterSayHelloResult, error)
    
}