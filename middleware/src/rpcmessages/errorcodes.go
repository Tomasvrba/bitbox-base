package rpcmessages

// ErrorCode is a unique and short code represeting an Error
type ErrorCode string

// Implementation error codes.
// These should only occur on incorrect implementation.
const (
	// If either the bbb-cmd.sh or bbb-config.sh script are not found
	ErrorScriptNotFound ErrorCode = "SCRIPT_NOT_FOUND"
	// The bbb-cmd.sh or bbb-config.sh scripts need to be run as superuser.
	ErrorScriptNotSuperuser ErrorCode = "SCRIPT_NOT_RUN_AS_SUPERUSER"
	// General Redis related error.
	// If there is an error with redis it's proably only important for a developer.
	// There is no differentiation of Redis errors because the front-end most likely handles them similar.
	ErrorRedisError ErrorCode = "REDIS_ERROR"
	// An unknown error occurred.
	ErrorUnhandled ErrorCode = "UNKNOWN"
)

// ErrorCodes for the cmd script
const (
	// The argument for the bbb-cmd.sh script is not known.
	ErrorCmdScriptInvalidArg ErrorCode = "CMD_SCRIPT_INVALID_ARG"

	/* bbb-cmd.sh flashdrive check */
	// Multiple USB flashdrives are found. Need exacly one.
	ErrorFlashdriveCheckMultiple ErrorCode = "FLASHDRIVE_CHECK_MULTI"
	// No USB flashdrive found.
	ErrorFlashdriveCheckNone ErrorCode = "FLASHDRIVE_CHECK_NONE"

	/* bbb-cmd.sh flashdrive mount <path> */
	// No flashdrive found on the passed <path>.
	ErrorFlashdriveMountNotFound ErrorCode = "FLASHDRIVE_MOUNT_NOT_FOUND"
	// The passed <path> does not uniquely identify a flashdrive.
	ErrorFlashdriveMountNotUnique ErrorCode = "FLASHDRIVE_MOUNT_NOT_UNIQUE"
	// The flashdrive is either bigger than 64GB or the filesystem is not supported.
	ErrorFlashdriveMountNotSupported ErrorCode = "FLASHDRIVE_MOUNT_NOT_SUPPORTED"

	/* bbb-cmd.sh flashdrive unmount */
	// There is no flashdrive to unmount at /mnt/backup.
	ErrorFlashdriveUnmountNotMounted ErrorCode = "FLASHDRIVE_UNMOUNT_NOT_MOUNTED"

	/* bbb-cmd.sh backup sysconfig */
	// /mnt/backup needs to be a mountpoint to backup the sysconfig.
	ErrorBackupSysconfigNotAMountpoint ErrorCode = "BACKUP_SYSCONFIG_NOT_A_MOUNTPOINT"

	/* bbb-cmd.sh restore sysconfig */
	// Backup file /mnt/backup/bbb-backup.rdb not found.
	ErrorRestoreSysconfigBackupNotFound ErrorCode = "RESTORE_SYSCONFIG_BACKUP_NOT_FOUND"

	/* bbb-cmd.sh mender-update */
	// Mender application is not enabled in diskimage.
	ErrorMenderUpdateImageNotMenderEnabled ErrorCode = "MENDER_UPDATE_IMAGE_NOT_MENDER_ENABLED"

	/* bbb-cmd.sh mender-update install*/
	// mender -install failed, see error code in stdout.
	ErrorMenderUpdateInstallFailed ErrorCode = "MENDER_UPDATE_INSTALL_FAILED"

	/* bbb-cmd.sh mender-update install*/
	// no firmware version supplied.
	ErrorMenderUpdateNoVersion ErrorCode = "MENDER_UPDATE_NO_VERSION"

	/* bbb-cmd.sh mender-update install*/
	// invalid firmware version supplied.
	ErrorMenderUpdateInvalidVersion ErrorCode = "MENDER_UPDATE_INVALID_VERSION"

	/* bbb-cmd.sh mender-update commit*/
	// mender -commit failed, see error code in stdout.
	ErrorMenderUpdateCommitFailed ErrorCode = "MENDER_UPDATE_COMMIT_FAILED"
)

// ErrorCodes for the config script
const (
	// The argument for the bbb-config.sh script is not known.
	ErrorConfigScriptInvalidArg ErrorCode = "CONFIG_SCRIPT_INVALID_ARG"

	/* bbb-config.sh enable bitcoin_ibd_clearnet */
	// Tor service is already disabled for the whole system, cannot enable IBD over clearnet.
	ErrorEnableClearnetIBDTorAlreadyDisabled ErrorCode = "ENABLE_CLEARNETIBD_TOR_ALREADY_DISABLED"

	/* bbb-config.sh set <key> <value> */
	// Set needs two arguments.
	ErrorSetNeedsTwoArguments ErrorCode = "SET_NEEDS_TWO_ARGUMENTS"

	/* bbb-config.sh set bitcoin_network <value> */
	// <value> has to be either testnet or mainnet
	ErrorSetBitcoinNetworkInvalidValue ErrorCode = "SET_BITCOINETWORK_INVALID_VALUE"

	/* bbb-config.sh set bitcoin_ibd <value> */
	// <value> has to be either true or false
	ErrorSetBitcoinIBDInvalidValue ErrorCode = "SET_BITCOINIBD_INVALID_VALUE"

	/* bbb-config.sh set bitcoin_ibd_clearnet <value> */
	// <value> has to be either true or false
	ErrorSetBitcoinIBDClearnetInvalidValue ErrorCode = "SET_BITCOINIBD_CLEARNET_INVALID_VALUE"

	/* bbb-config.sh set bitcoin_dbcache <value> */
	// <value> has to be an integer in MB between 50 and 3000.
	ErrorSetBitcoinDBCacheInvalidValue ErrorCode = "SET_BITCOINDBCACHE_INVALID_VALUE"

	/* bbb-config.sh set hostname <value> */
	// <value> has to be a valid hostname according to this regex '^[a-z][a-z0-9-]{0,22}[a-z0-9]$'.
	ErrorSetHostnameInvalidValue ErrorCode = "SET_HOSTNAME_INVALID_VALUE"
)

// Middleware Error Codes
const (
	// is returned when the dummy autentication is not successful
	ErrorDummyAuthenticationNotSuccessful ErrorCode = "DUMMY_AUTHENTICATION_NOT_SUCCESSFUL"

	// The provided password is too short.
	ErrorDummyPasswordTooShort ErrorCode = "DUMMY_CHANGEPASSWORD_TOO_SHORT"

	// The provided root password is to short.
	ErrorSetRootPasswordToShort ErrorCode = "SET_ROOTPASSWORD_PASSWORD_TO_SHORT"
)
