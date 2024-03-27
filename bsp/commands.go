// Code generated by go generate; DO NOT EDIT.
// Code generated by go generate; DO NOT EDIT.
// Code generated by go generate; DO NOT EDIT.

package bsp

import "blue/commands"

const cmdLen = 26

// db -----------------------------
const (
	DEL Header = 1 + TypeDB
	EXPIRE Header = 2 + TypeDB
	KVS Header = 3 + TypeDB
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
	SADD Header = 1 + TypeSet
	SDEL Header = 2 + TypeSet
	SGET Header = 3 + TypeSet
	SIN Header = 4 + TypeSet
	SPOP Header = 5 + TypeSet
)

// json -----------------------------
const (
)

// system -----------------------------
const (
	EXIT Header = 1 + TypeSystem
	HELP Header = 2 + TypeSystem
	PING Header = 3 + TypeSystem
	SELECT Header = 4 + TypeSystem
	VERSION Header = 5 + TypeSystem
)

var HandleMap = [...]string{
	DEL: "DEL",
	EXIT: "EXIT",
	EXPIRE: "EXPIRE",
	GET: "GET",
	HELP: "HELP",
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
	PING: "PING",
	RPOP: "RPOP",
	RPUSH: "RPUSH",
	SADD: "SADD",
	SDEL: "SDEL",
	SELECT: "SELECT",
	SET: "SET",
	SGET: "SGET",
	SIN: "SIN",
	SPOP: "SPOP",
	VERSION: "VERSION",
}

var HandleMap2 = map[string]Header{
	"DEL": DEL,
	"EXIT": EXIT,
	"EXPIRE": EXPIRE,
	"GET": GET,
	"HELP": HELP,
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
	"PING": PING,
	"RPOP": RPOP,
	"RPUSH": RPUSH,
	"SADD": SADD,
	"SDEL": SDEL,
	"SELECT": SELECT,
	"SET": SET,
	"SGET": SGET,
	"SIN": SIN,
	"SPOP": SPOP,
	"VERSION": VERSION,
}

var CommandsMap = [...]commands.Cmd{
	DEL: {Name:"DEL",Summary: "Remove the specified keys", Group: "db", Arity: 1, Key: "list", Value: "", Arguments: []string{}},
	EXIT: {Name:"EXIT",Summary: "Exit the blue", Group: "system", Arity: 0, Key: "", Value: "", Arguments: []string{}},
	EXPIRE: {Name:"EXPIRE",Summary: "Set a key's time to live in seconds", Group: "db", Arity: 2, Key: "string", Value: "number", Arguments: []string{}},
	GET: {Name:"GET",Summary: "Returns the string value of a key.", Group: "string", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	HELP: {Name:"HELP",Summary: "Returns the action of the given command", Group: "system", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	INCR: {Name:"INCR",Summary: "Increment the integer value of a key by the given amount", Group: "number", Arity: 1, Key: "string", Value: "number", Arguments: []string{}},
	KVS: {Name:"KVS",Summary: "Returns all key-value pairs in the database", Group: "db", Arity: 0, Key: "", Value: "", Arguments: []string{}},
	LEN: {Name:"LEN",Summary: "Returns the length of a string", Group: "string", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	LGET: {Name:"LGET",Summary: "Gets all values for the list of given key", Group: "list", Arity: 1, Key: "string", Value: "list", Arguments: []string{}},
	LLEN: {Name:"LLEN",Summary: "Returns the length of the list stored at key.", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	LPOP: {Name:"LPOP",Summary: "Remove and get the first element in a list", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	LPUSH: {Name:"LPUSH",Summary: "Insert values at the head of the list stored at key.", Group: "list", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	LSET: {Name:"LSET",Summary: "Set the value of a list", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	NGET: {Name:"NGET",Summary: "Returns the number value of a key.", Group: "number", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	NSET: {Name:"NSET",Summary: "Set the value of a number", Group: "number", Arity: 2, Key: "string", Value: "number", Arguments: []string{"expire"}},
	PING: {Name:"PING",Summary: "Pings the bot to check if it is online.", Group: "system", Arity: 0, Key: "", Value: "", Arguments: []string{}},
	RPOP: {Name:"RPOP",Summary: "Remove and get the last element in a list", Group: "list", Arity: 1, Key: "string", Value: "string", Arguments: []string{}},
	RPUSH: {Name:"RPUSH",Summary: "Add the value to the end of the list stored at key", Group: "list", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	SADD: {Name:"SADD",Summary: "Add the specified members to the set stored at key. Specified members that are already a member of this set are ignored. If key does not exist, a new set is created before adding the specified members.", Group: "set", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	SDEL: {Name:"SDEL",Summary: "Removes a given value from a set if it exists", Group: "set", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	SELECT: {Name:"SELECT",Summary: "Select a db.", Group: "system", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	SET: {Name:"SET",Summary: "Set the value of a key", Group: "string", Arity: 2, Key: "string", Value: "string", Arguments: []string{"expire"}},
	SGET: {Name:"SGET",Summary: "Gets all the values in the set, if the set exists", Group: "set", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	SIN: {Name:"SIN",Summary: "Checks whether the given value is in the given set", Group: "set", Arity: 2, Key: "string", Value: "string", Arguments: []string{}},
	SPOP: {Name:"SPOP",Summary: "Remove and return a random member from a set", Group: "set", Arity: 1, Key: "string", Value: "", Arguments: []string{}},
	VERSION: {Name:"VERSION",Summary: "Get the version of the system.", Group: "system", Arity: 0, Key: "", Value: "", Arguments: []string{}},
}

