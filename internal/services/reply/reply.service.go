package reply

import (
	"fmt"
	"markdown-note/internal/lib"
	"net/http"
)

//
// =======================
// == Struct Definitions ==
// =======================
//

// Meta represents metadata information of the API reply.
type Meta struct {
	Status string `json:"status"` // Overall reply status (e.g. "SUCCESS" or "ERROR")
}

// ReplyEnvelope wraps the general structure of an API reply.
type ReplyEnvelope struct {
	Meta Meta `json:"meta"` // Contains status metadata
	Data any  `json:"data"` // Holds the actual reply payload
}

// ReplyError defines the structure of an error reply payload.
type ReplyError struct {
	Code    string `json:"code"`              // Short machine-readable error code
	Message string `json:"message"`           // Human-readable message describing the error
	Details string `json:"details,omitempty"` // Optional detailed context or debug info
}

// Reply represents a unified HTTP reply writer with utility methods.
type Reply struct {
	w http.ResponseWriter
	ReplyEnvelope
}

//
// =======================
// == Constructor & Utils ==
// =======================
//

// New creates a new Reply instance with the JSON content-type header already set.
//
// @param w http.ResponseWriter - the HTTP writer for the reply
// @return *Reply - initialized Reply instance
func New(w http.ResponseWriter) *Reply {
	r := &Reply{w: w}
	r.w.Header().Set("Content-Type", "application/json")
	return r
}

func (r *Reply) Header() http.Header {
	return r.w.Header()
}

//
// =======================
// == Setter Methods ==
// =======================
//

// setStatusCode sets the HTTP status code for the reply.
// Returns the Reply for chaining.
func (r *Reply) setStatusCode(status int) *Reply {
	r.w.WriteHeader(status)
	return r
}

// SetStatus sets the "meta.status" field value.
// Returns the Reply for chaining.
func (r *Reply) SetStatus(status string) *Reply {
	r.Meta.Status = status
	return r
}

// SetData assigns data to the "data" field.
// Returns the Reply for chaining.
func (r *Reply) SetData(data any) *Reply {
	r.Data = data
	return r
}

//
// =======================
// == High-level Helpers ==
// =======================
//

// Success marks the reply as successful and attaches data.
func (r *Reply) Success(data any) *Reply {
	r.SetStatus("SUCCESS")
	r.SetData(data)
	return r
}

// Error sets reply status to "ERROR" and attaches an error payload.
func (r *Reply) Error(code, message string, details ...string) *Reply {
	r.SetStatus("ERROR")
	d := ""
	if len(details) > 0 {
		d = details[0]
	}
	r.SetData(ReplyError{code, message, d})
	return r
}

//
// =======================
// == Senders ==
// =======================
//

// Reply writes the full JSON reply to the client.
func (r *Reply) Reply(code int) {
	r.setStatusCode(code)
	fmt.Fprint(r.w, lib.Json.Stringify(r))
}

func (r *Reply) RawReply(code int) {
	r.setStatusCode(code)
	fmt.Fprint(r.w, r.Data)
}

// Ok sends a 200 OK reply.
func (r *Reply) Ok() {
	r.Reply(http.StatusOK)
}

// NoContent sends a 204 No Content reply.
func (r *Reply) NoContent() {
	r.Reply(http.StatusNoContent)
}

// Created sends a 201 Created reply.
func (r *Reply) Created() {
	r.Reply(http.StatusCreated)
}

// Fail sends a failure reply with a specific HTTP status code.
func (r *Reply) Fail(code int) {
	r.Reply(code)
}

//
// =======================
// == Templates ==
// =======================
//

var (
	CodeNotFound    = "NOT_FOUND"
	CodeServerError = "SERVER_ERROR"
	CodeBadRequest  = "CLIENT_ERROR"
	CodeBadGateWay  = "BAD_GATEWAY"
)
