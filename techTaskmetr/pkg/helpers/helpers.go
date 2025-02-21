package helpers

const (
	ServicePrefix          = "[ SERVICE ]"
	ServiceGetElementError = "FAILED_TO_GET_ELEMENT"
	ServiceUpdateError     = "FAILED_TO_UPDATE_ELEMENT"
	ServiceDeleteError     = "FAILED_TO_DELETE_ELEMENT"

	RepoPrefix                   = "[ REPO ]"
	RepoGetError                 = "FAILED_TO_GET_ELEMENT_FROM_DB"
	RepoCacheError               = "FAILED_TO_CACHE_ELEMENT_IN_REDIS"
	RepoUpdateError              = "FAILED_TO_UPDATE_ELEMENT"
	RepoFetchError               = "FAILED_TO_FETCH_UPDATED_ELEMENT"
	RepoRedisParseError          = "FAILED_TO_PARSE_CACHED_ELEMENT"
	RepoCreateTransactionError   = "FAILED_TO_CREATE_TRANSACTION"
	RepoCacheTransactionError    = "FAILED_TO_CACHE_TRANSACTION"
	RepoGetTransactionsError     = "FAILED_TO_GET_TRANSACTIONS"
	RepoScanTransactionError     = "FAILED_TO_SCAN_TRANSACTION"
	RepoIterateTransactionsError = "FAILED_TO_ITERATE_TRANSACTIONS"

	HandlerPrefix = "[ HANDLERS ]"

	AppPrefix = " [ APP ] "

	PgPrefix            = " [ POSTGRES ] "
	ReconnectDB         = "RECONNECTING TO DATABASE..."        // ReconnectDB contains reconnect db message
	DisconnectDB        = "DISCONNECTED FROM DATABASE"         // DisconnectDB contains disconnect db message
	PgConnectFailed     = "FAILED TO CONNECT TO DATABASE"      // PgConnectFailed contains error message for failed to connect to database
	PgConnectSuccess    = "SUCCESSFULLY CONNECTED TO POSTGRES" // PgConnectSuccess contains success message for successfully connected to database
	PgTransactionFailed = "FAILED TO FETCH TRANSACTION"        // PgTransactionFailed contains error message for failed to fetch transaction
	PgMigrateFailed     = "FAILED TO MIGRATE DATABASE"         // PgMigrateFailed contains error message for failed to migrate database
	NoRowsAffected      = "NO ROWS AFFECTED"                   // NoRowsAffected contains error message for no rows affected
	FailedToRollback    = "FAILED TO ROLLBACK"                 // FailedToRollback contains error message for failed to rollback
	FailedToClose       = "FAILED TO CLOSE"                    // FailedToClose contains error message for failed to close
)
