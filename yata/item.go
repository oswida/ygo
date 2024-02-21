package yata

import "ygo/common"

type Item[T Content] struct {
	Id          *common.Id
	Parent      *Item[T]
	Left        *Item[T]
	Origin      *Item[T]
	Right       *Item[T]
	RightOrigin *Item[T]
	Content     T
	Keep        bool
	Countable   bool
	Deleted     bool
	Redone      *common.Id
	Length      uint64
	Marked      bool
}

func NewItem[T Content](id *common.Id, parent, left, origin, right, rightOrigin *Item[T], content T) *Item[T] {
	return &Item[T]{
		Id:          id,
		Parent:      parent,
		Left:        left,
		Right:       right,
		Origin:      origin,
		RightOrigin: rightOrigin,
		Content:     content,
		Keep:        false,
		Countable:   false,
		Deleted:     false,
		Redone:      nil,
		Length:      content.Length(),
		Marked:      false,
	}
}

// getMissing - Return the creator clientID of the missing op or define missing items and return null.

// Return the creator clientID of the missing op or define missing items and return null.
func (it *Item[T]) GetMissing() uint64 {

}

// integrate (transaction, offset)

// Computes the last content address of this Item
func (it *Item[T]) LastId() *common.Id {
	if it.Length == 1 {
		return it.Id
	}
	return common.NewId(it.Id.Client, it.Id.Clock+it.Length-1)
}

// Get next non-deleted item
func (it *Item[T]) Next() *Item[T] {
	n := it
	for n != nil && n.Deleted {
		n = n.Right
	}
	return n
}

// Get previous non-deleted item
func (it *Item[T]) Prev() *Item[T] {
	n := it
	for n != nil && n.Deleted {
		n = n.Left
	}
	return n
}

// Try to merge two items
func (it *Item[T]) MergeWith(other *Item[T]) bool {
	// this.constructor === right.constructor ?
	canMerge := other.Origin.Id.Equal(it.LastId()) &&
		it.Right == other &&
		it.RightOrigin.Id.Equal(other.RightOrigin.Id) &&
		it.Id.Client == other.Id.Client &&
		it.Id.Clock+it.Length == other.Id.Clock &&
		it.Deleted == other.Deleted &&
		it.Redone == nil &&
		other.Redone == nil &&
		it.Content.MergeWith(other.Content)
	// Item.js 582 - update search markers
	if canMerge {
		if other.Keep {
			it.Keep = true
		}
		it.Right = other.Right
		if it.Right != nil {
			it.Right.Left = it
		}
		it.Length += other.Length
		return true
	}
	return false
}

// Mark this Item as deleted.
func (it *Item[T]) Delete(tx Transaction) {
	if it.Deleted {
		return
	}
	// const parent = /** @type {AbstractType<any>} */ (this.parent)
	// adjust the length of parent
	// if (this.countable && this.parentSub === null) {
	if it.Countable /* && this.parentSub === null */ {
		it.Parent.Length -= it.Length
	}
	it.Deleted = true
	// addToDeleteSet(transaction.deleteSet, this.id.client, this.id.clock, this.length)
	// addChangedTypeToTransaction(transaction, parent, this.parentSub)
	it.Content.Delete(tx)
}
