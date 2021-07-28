package rpc

import (
	"encoding/json"
	"log"
	"net/http"
    "strings"
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

type Server struct {
    
    ArithService Arith
    
    GreeterService Greeter
    
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request jsonRPCRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	methodSegments := strings.Split(request.Method, ".")
	serviceName := methodSegments[0]
	methodName := methodSegments[1]

	response := jsonRPCResponse{
		JSONRPC: request.JSONRPC,
		ID:      request.ID,
	}

	var resultData json.RawMessage

	switch serviceName {
    
    
    case "Arith":
        switch methodName {
        
        case "Add":
            var params ArithAddParams
			var result ArithAddResult

			err := s.ArithService.Add(&params, &result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}

			resultData, err = json.Marshal(result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}
        
        case "Pow":
            var params ArithPowParams
			var result ArithPowResult

			err := s.ArithService.Pow(&params, &result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}

			resultData, err = json.Marshal(result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}
        
        case "IsNegative":
            var params ArithIsNegativeParams
			var result ArithIsNegativeResult

			err := s.ArithService.IsNegative(&params, &result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}

			resultData, err = json.Marshal(result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}
        
        }
    
    
    case "Greeter":
        switch methodName {
        
        case "SayHello":
            var params GreeterSayHelloParams
			var result GreeterSayHelloResult

			err := s.GreeterService.SayHello(&params, &result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}

			resultData, err = json.Marshal(result)
			if err != nil {
				response.Error = &jsonRPCError{Message: err.Error()}
			}
        
        }
    
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
    
    Add(*ArithAddParams, *ArithAddResult) error
    
    Pow(*ArithPowParams, *ArithPowResult) error
    
    IsNegative(*ArithIsNegativeParams, *ArithIsNegativeResult) error
    
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
    
    SayHello(*GreeterSayHelloParams, *GreeterSayHelloResult) error
    
}
