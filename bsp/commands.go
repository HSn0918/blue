// Code generated by go generate; DO NOT EDIT.
// Code generated by go generate; DO NOT EDIT.
// Code generated by go generate; DO NOT EDIT.

package bsp

import "blue/commands"

const cmdLen = 18

// db -----------------------------
const (
	DEL Header = 1 + TypeDB
	EXPIRE Header = 2 + TypeDB
	SELECT Header = 3 + TypeDB
)

// number -----------------------------
const (
	INCR Header = 1 + TypeNumber
	NGET Header = 2 + TypeNumber
	NSET Header = 3 + TypeNumber
)

// string -----------------------------
const (
	GET Header = 1 + TypeString
	LEN Header = 2 + TypeString
	SET Header = 3 + TypeString
)

// list -----------------------------
const (
	LGET Header = 1 + TypeList
	LLEN Header = 2 + TypeList
	LPOP Header = 3 + TypeList
	LPUSH Header = 4 + TypeList
	LSET Header = 5 + TypeList
	RPOP Header = 6 + TypeList
	RPUSH Header = 7 + TypeList
)

// set -----------------------------
const (
)

// json -----------------------------
const (
)

// system -----------------------------
const (
	KVS Header = 1 + TypeSystem
	VERSION Header = 2 + TypeSystem
)

var HandleMap = [...]string{
	DEL: "DEL",
	EXPIRE: "EXPIRE",
	GET: "GET",
	INCR: "INCR",
	KVS: "KVS",
	LEN: "LEN",
	LGET: "LGET",
	LLEN: "LLEN",
	LPOP: "LPOP",
	LPUSH: "LPUSH",
	LSET: "LSET",
	NGET: "NGET",
	NSET: "NSET",
	RPOP: "RPOP",
	RPUSH: "RPUSH",
	SELECT: "SELECT",
	SET: "SET",
	VERSION: "VERSION",
}

var HandleMap2 = map[string]Header{
	"DEL": DEL,
	"EXPIRE": EXPIRE,
	"GET": GET,
	"INCR": INCR,
	"KVS": KVS,
	"LEN": LEN,
	"LGET": LGET,
	"LLEN": LLEN,
	"LPOP": LPOP,
	"LPUSH": LPUSH,
	"LSET": LSET,
	"NGET": NGET,
	"NSET": NSET,
	"RPOP": RPOP,
	"RPUSH": RPUSH,
	"SELECT": SELECT,
	"SET": SET,
	"VERSION": VERSION,
}

var CommandsMap = [...]commands.Cmd{
	DEL: {Name:"DEL",Summary: "Remove the specified keys", Group: "db", Arity: 1, Key: "list", Value: "", Arguments: []string{}},
	EXPIRE: {Name:"EXPIRE",Summary: "Set a key's time to live in seconds", Group: "db", Arity: 2, Key: "string", Value: "number", Arguments: []string{}},
	GET: {Name:"GET",Summary: "Returns the string value of a key.", Group: "string", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	INCR: {Name:"INCR",Summary: "Increment the integer value of a key by the given amount", Group: "number", Arity: 1, Key: "string", Value: "number", Arguments: []string{}},
	KVS: {Name:"KVS",Summary: "Returns all key-value pairs in the database", Group: "system", Arity: 0, Key: "", Value: "", Arguments: []string{}},
	LEN: {Name:"LEN",Summary: "Returns the length of a string", Group: "string", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	LGET: {Name:"LGET",Summary: "Gets all values for the list of given key", Group: "list", Arity: 1, Key: "string", Value: "list", Arguments: []string{}},
	LLEN: {Name:"LLEN",Summary: "Returns the length of the list stored at key.", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	LPOP: {Name:"LPOP",Summary: "Remove and get the first element in a list", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	LPUSH: {Name:"LPUSH",Summary: "Insert values at the head of the list stored at key.", Group: "list", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	LSET: {Name:"LSET",Summary: "Set the value of a list", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	NGET: {Name:"NGET",Summary: "Returns the number value of a key.", Group: "number", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	NSET: {Name:"NSET",Summary: "Set the value of a number", Group: "number", Arity: 2, Key: "string", Value: "number", Arguments: []string{"expire"}},
	RPOP: {Name:"RPOP",Summary: "Remove and get the last element in a list", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	RPUSH: {Name:"RPUSH",Summary: "Add the value to the end of the list stored at key", Group: "list", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	SELECT: {Name:"SELECT",Summary: "Select a db.", Group: "db", Arity: 1, Key: "number", Value: "", Arguments: []string{}},
	SET: {Name:"SET",Summary: "Set the value of a key", Group: "string", Arity: 2, Key: "string", Value: "string", Arguments: []string{"expire"}},
	VERSION: {Name:"VERSION",Summary: "Get the version of the system.", Group: "system", Arity: 0, Key: "", Value: "", Arguments: []string{}},
}

