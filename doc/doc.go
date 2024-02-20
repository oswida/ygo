package doc

import (
	"ygo/basic"

	"github.com/google/uuid"
)

// Should be observable
type YDoc struct {
	Guid     uuid.UUID
	ClientID *basic.Id
	// CollectionID
	// GC
	// GCFilter
	// Share
	// Store
	// Transaction
	// this._transactionCleanups = []
	// this.subdocs = new Set()
	/**
	 * If this document is a subdocument - a document integrated into another document - then _item is defined.
	 * @type {Item?}
	 */
	// this._item = null
	// this.shouldLoad = shouldLoad
	// this.autoLoad = autoLoad
	// this.meta = meta
	/**
	 * This is set to true when the persistence provider loaded the document from the database or when the `sync` event fires.
	 * Note that not all providers implement this feature. Provider authors are encouraged to fire the `load` event when the doc content is loaded from the database.
	 *
	 * @type {boolean}
	 */
	// this.isLoaded = false
	/**
	 * This is set to true when the connection provider has successfully synced with a backend.
	 * Note that when using peer-to-peer providers this event may not provide very useful.
	 * Also note that not all providers implement this feature. Provider authors are encouraged to fire
	 * the `sync` event when the doc has been synced (with `true` as a parameter) or if connection is
	 * lost (with false as a parameter).
	 */
}

func NewYDoc() (*YDoc, error) {
	id, err := basic.GenerateId()
	if err != nil {
		return nil, err
	}
	return &YDoc{
		Guid:     uuid.New(),
		ClientID: id,
	}, nil
}
