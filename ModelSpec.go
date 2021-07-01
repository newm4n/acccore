package acccore

import (
	"time"
)

const (
	// DEBIT is enum transaction type DEBIT
	DEBIT TransactionType = iota
	// CREDIT is enum transaction type CREDIT
	CREDIT
)

// TransactionType is the enum type of transaction type, DEBIT and CREDIT
type TransactionType int

// Journal interface define a base Journal structure.
// A journal depict an event where transactions is happening.
// Important to understand, that Journal don't have update or delete function, its due to accountability reason.
// To delete a journal, one should create a reversal journal.
// To update a journal, one should create a reversal journal and then followed with a correction journal.
// If your implementation database do not support 2 phased commit, you should maintain your own committed flag in
// this journal table. When you want to select those journal, you only select those  that have committed flag status on.
// Committing this journal, will propagate to commit the child Transactions
type Journal interface {
	// GetJournalID would return the journal unique ID
	GetJournalID() string
	// SetJournalID will set a new JournalID
	SetJournalID(newId string) Journal

	// GetJournalingTime will return the timestamp of when this journal entry is created
	GetJournalingTime() time.Time
	// SetJournalingTime will set new JournalTime
	SetJournalingTime(newTime time.Time) Journal

	// GetDescription returns description about this journal entry
	GetDescription() string
	// SetDescription will set new description
	SetDescription(newDesc string) Journal

	// IsReversal return an indicator if this journal entry is a reversal of other journal
	IsReversal() bool
	// SetReversal will set new reversal status
	SetReversal(rev bool) Journal

	// GetReversedJournal should returned the Journal that is reversed IF `IsReverse()` function returned `true`
	GetReversedJournal() Journal
	// SetReversedJournal will set the reversed journal
	SetReversedJournal(journal Journal) Journal

	// GetAmount should return the current amount of total transaction values
	GetAmount() int64
	// SetAmount will set new total transaction amount
	SetAmount(newAmount int64) Journal

	// GetTransactions should returns all transaction information that being part of this journal entry.
	GetTransactions() []Transaction
	// SetTransactions will set new list of transaction under this journal
	SetTransactions(transactions []Transaction) Journal

	// GetCreateTime function should return the time when this entry is created/recorded. Logically it the same as `GetTime()` function
	// this function serves as audit trail.
	GetCreateTime() time.Time
	// SetCreateTime will set the creation time
	SetCreateTime(newTime time.Time) Journal

	// GetCreateBy function should return the user accountNumber or some identification of who is creating this journal.
	// this function serves as audit trail.
	GetCreateBy() string
	// SetCreateBy will set the creator name
	SetCreateBy(creator string) Journal
}

// Transaction interface define a base Transaction structure
// A transaction is a unit of transaction element that involved within a journal.
// A transaction must include reference to the journal that binds the transaction with other transaction and
// also must state the Account tha doing the transaction
// If your implementation database do not support 2 phased commit, you should maintain your own committed flag in
// this transaction table. When you want to select those transaction, you only select those  that have committed flag status on.
type Transaction interface {
	// GetTransactionID returns the unique ID of this transaction
	GetTransactionID() string
	// SetTransactionID will set new transaction ID
	SetTransactionID(newId string) Transaction

	// GetTransactionTime returns the timestamp of this transaction
	GetTransactionTime() time.Time
	// SetTransactionTime will set new transaction time
	SetTransactionTime(newTime time.Time) Transaction

	// GetAccountNumber return the account number of account ID who owns this transaction
	GetAccountNumber() string
	// SetAccountNumber will set new account number who own this transaction
	SetAccountNumber(number string) Transaction

	// GetJournal returns the journal information where this transaction is recorded.
	GetJournalID() string
	// SetJournal will set the journal to which this transaction is recorded
	SetJournalID(journalID string) Transaction

	// GetDescription return the description of this Transaction.
	GetDescription() string
	// SetDescription will set the transaction description
	SetDescription(desc string) Transaction

	// GetTransactionType get the transaction type DEBIT or CREDIT
	GetTransactionType() TransactionType
	// SetTransactionType will set the transaction type
	SetTransactionType(txType TransactionType) Transaction

	// GetAmount return the transaction amount
	GetAmount() int64
	// SetAmount will set the amount
	SetAmount(newAmount int64) Transaction

	// GetBookBalance return the balance of the account at the time when this transaction has been written.
	GetAccountBalance() int64
	// SetAccountBalance will set new account balance
	SetAccountBalance(newBalance int64) Transaction

	// GetCreateTime function should return the time when this transaction is created/recorded.
	// this function serves as audit trail.
	GetCreateTime() time.Time
	// SetCreateTime will set new creation time
	SetCreateTime(newTime time.Time) Transaction

	// GetCreateBy function should return the user accountNumber or some identification of who is creating this transaction.
	// this function serves as audit trail.
	GetCreateBy() string
	// SetCreateBy will set new creator name
	SetCreateBy(creator string) Transaction
}

// Account interface provides base structure of Account
type Account interface {
	// GetCurrency returns the currency identifier such as `GOLD` or `POINT` or `IDR`
	GetCurrency() string
	// SetCurrency will set the account currency
	SetCurrency(newCurrency string) Account

	// GetAccountNumber returns the unique account number
	GetAccountNumber() string
	// SetAccountNumber will set new account ID
	SetAccountNumber(newNumber string) Account

	// GetName returns the account name
	GetName() string
	// SetName will set the new account name
	SetName(newName string) Account

	// GetDescription returns some description text about this account
	GetDescription() string
	// SetDescription will set new description
	SetDescription(newDesc string) Account

	// GetBaseTransactionType returns the base transaction type of this account,
	// 1. Asset based should be DEBIT
	// 2. Equity or Liability based should be CREDIT
	GetBaseTransactionType() TransactionType
	// SetBaseTransactionType will set new base transaction type
	SetBaseTransactionType(newType TransactionType) Account

	// GetBalance returns the current balance of this account.
	// for each transaction created for this account, this balance MUST BE UPDATED
	GetBalance() int64
	// SetBalance will set new transaction balance
	SetBalance(newBalance int64) Account

	// GetCOA returns the COA code for this account, used for categorization of account.
	GetCOA() string
	// SetCOA Will set new COA code
	SetCOA(newCoa string) Account

	// GetCreateTime function should return the time when this account is created/recorded.
	// this function serves as audit trail.
	GetCreateTime() time.Time
	// SetCreateTime will set new creation time
	SetCreateTime(newTime time.Time) Account

	// GetCreateBy function should return the user accountNumber or some identification of who is creating this account.
	// this function serves as audit trail.
	GetCreateBy() string
	// SetCreateBy will set the creator name
	SetCreateBy(creator string) Account

	// GetUpdateTime function should return the time when this account is updated.
	// this function serves as audit trail.
	GetUpdateTime() time.Time
	// SetUpdateTime will set the last update time.
	SetUpdateTime(newTime time.Time) Account

	// GetUpdateBy function should return the user accountNumber or some identification of who is updating this account.
	// this function serves as audit trail.
	GetUpdateBy() string
	// SetUpdateBy will set the updater name
	SetUpdateBy(editor string) Account
}