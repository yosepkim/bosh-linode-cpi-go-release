package dispatcher

import (
	"bytes"
	"encoding/json"

	blcaction "bosh-linode-cpi/action"
	blcapi "bosh-linode-cpi/api"
)

const (
	jsonLogTag = "json"

	jsonCloudErrorType          = "Bosh::Clouds::CloudError"
	jsonCpiErrorType            = "Bosh::Clouds::CpiError"
	jsonNotImplementedErrorType = "Bosh::Clouds::NotImplemented"
)

type Request struct {
	Method    string        `json:"method"`
	Arguments []interface{} `json:"arguments"`
}

type Response struct {
	Result interface{}    `json:"result"`
	Error  *ResponseError `json:"error"`

	Log string `json:"log"`
}

type ResponseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (r ResponseError) Error() string {
	return r.Message
}

type JSON struct {
	actionFactory blcaction.Factory
	caller        Caller
	logger        blcapi.MultiLogger
}

func NewJSON(
	actionFactory blcaction.Factory,
	caller Caller,
	logger blcapi.MultiLogger,
) JSON {
	return JSON{
		actionFactory: actionFactory,
		caller:        caller,
		logger:        logger,
	}
}

func (c JSON) Dispatch(reqBytes []byte) []byte {
	var req Request

	c.logger.DebugWithDetails(jsonLogTag, "Request bytes", string(reqBytes))

	decoder := json.NewDecoder(bytes.NewReader(reqBytes))
	decoder.UseNumber()
	if err := decoder.Decode(&req); err != nil {
		return c.buildCpiError("Must provide valid JSON payload")
	}

	c.logger.DebugWithDetails(jsonLogTag, "Deserialized request", req)

	if req.Method == "" {
		return c.buildCpiError("Must provide method key")
	}

	if req.Arguments == nil {
		return c.buildCpiError("Must provide arguments key")
	}

	action, err := c.actionFactory.Create(req.Method)
	if err != nil {
		return c.buildNotImplementedError()
	}

	result, err := c.caller.Call(action, req.Arguments)
	if err != nil {
		return c.buildCloudError(err)
	}

	resp := Response{
		Result: result,
		Log:    c.logger.LogBuff.String(),
	}

	c.logger.DebugWithDetails(jsonLogTag, "Deserialized response", resp)

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return c.buildCpiError("Failed to serialize result")
	}

	c.logger.DebugWithDetails(jsonLogTag, "Response bytes", string(respBytes))

	return respBytes
}

func (c JSON) buildCloudError(err error) []byte {
	respErr := Response{
		Log:   c.logger.LogBuff.String(),
		Error: &ResponseError{},
	}

	respErr.Log = c.logger.LogBuff.String()

	if typedErr, ok := err.(blcapi.CloudError); ok {
		respErr.Error.Type = typedErr.Type()
	} else {
		respErr.Error.Type = jsonCloudErrorType
	}

	respErr.Error.Message = err.Error()

	respErrBytes, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	c.logger.DebugWithDetails(jsonLogTag, "CloudError response bytes", string(respErrBytes))

	return respErrBytes
}

func (c JSON) buildCpiError(message string) []byte {
	respErr := Response{
		Log: c.logger.LogBuff.String(),
		Error: &ResponseError{
			Type:    jsonCpiErrorType,
			Message: message,
		},
	}

	respErrBytes, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	c.logger.DebugWithDetails(jsonLogTag, "CpiError response bytes", string(respErrBytes))

	return respErrBytes
}

func (c JSON) buildNotImplementedError() []byte {
	respErr := Response{
		Log: c.logger.LogBuff.String(),
		Error: &ResponseError{
			Type:    jsonNotImplementedErrorType,
			Message: "Must call implemented method",
		},
	}

	respErrBytes, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	c.logger.DebugWithDetails(jsonLogTag, "NotImplementedError response bytes", string(respErrBytes))

	return respErrBytes
}
