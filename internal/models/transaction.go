package models

import "github.com/jackc/pgx/v5"

type Tx interface {
	pgx.Tx
}

type TxIsoLevel string
type TxAccessMode string
type DeferrableMode string

const (
	Serializable    TxIsoLevel = "serializable"
	RepeatableRead  TxIsoLevel = "repeatable read"
	ReadCommitted   TxIsoLevel = "read committed"
	ReadUncommitted TxIsoLevel = "read uncommitted"
)

const (
	ReadWrite TxAccessMode = "read write"
	ReadOnly  TxAccessMode = "read only"
)

type TxOptions struct {
	IsoLevel   TxIsoLevel
	AccessMode TxAccessMode
}
