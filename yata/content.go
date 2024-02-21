package yata

type Content interface {
	MergeWith(other Content) bool
	Length() uint64
	Delete(Transaction)
}
