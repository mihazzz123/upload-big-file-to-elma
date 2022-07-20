package mw

const (
	// CompanyHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) имя компании
	CompanyHTTPHeader = "X-Company"
	// CompanyGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) имя компании
	CompanyGRPCHeader = "company"
	// CompanyAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) имя компании
	CompanyAMQPHeader = "company"
	// CompanyLogEntry название поля в логах и трейсах
	CompanyLogEntry = "company"

	// CompanyAliasHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) алиас компании
	CompanyAliasHTTPHeader = "X-Company-Alias"
	// CompanyAliasGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) алиас компании
	CompanyAliasGRPCHeader = "company_alias"
	// CompanyAliasAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) алиас компании
	CompanyAliasAMQPHeader = "company_alias"
	// CompanyAliasLogEntry название поля в логах и трейсах
	CompanyAliasLogEntry = "company_alias"

	// CompanyLangHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) язык компании
	CompanyLangHTTPHeader = "X-Company-language"
	// CompanyLangGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) язык компании
	CompanyLangGRPCHeader = "company_language"
	// CompanyLangAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) язык компании
	CompanyLangAMQPHeader = "company_language"
	// CompanyLangLogEntry название поля в логах и трейсах
	CompanyLangLogEntry = "company_language"

	// TimestampHTTPHeader заголовок HTTP-запрос, из которого (и в который) будет извлекаться (записываться) временная метка
	TimestampHTTPHeader = "X-Timestamp"
	// TimestampLogEntry название поля в логах и трейсах
	TimestampLogEntry = "timestamp"

	// IsPortalUserHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) что пользователь портальный
	IsPortalUserHTTPHeader = "X-Is-Portal-User"
	// IsPortalUserGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) что пользователь портальный
	IsPortalUserGRPCHeader = "is_portal_user"
	// IsPortalUserAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) что пользователь портальный
	IsPortalUserAMQPHeader = "is_portal_user"

	// UserIDHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) id пользователя
	UserIDHTTPHeader = "X-User-ID"
	// UserIDGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) id пользователя
	UserIDGRPCHeader = "user_id"
	// UserIDAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) id пользователя
	UserIDAMQPHeader = "user_id"
	// UserIDLogEntry название поля в логах и трейсах
	UserIDLogEntry = "user_id"

	// IsAdminHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) признак того, что пользователь администратор
	IsAdminHTTPHeader = "X-Is-Admin"
	// IsAdminGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) признак того, что пользователь администратор
	IsAdminGRPCHeader = "is_admin"
	// IsAdminAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) признак того, что пользователь администратор
	IsAdminAMQPHeader = "is_admin"
	// IsAdminLogEntry название поля в логах и трейсах
	IsAdminLogEntry = "is_admin"

	// TraceIDHeader название заголовка ответа, в которое будет положен id трейса
	TraceIDHeader = "X-Trace-ID"

	// LangHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) текущий язык
	LangHTTPHeader = "X-Language"
	// LangGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) текущий язык
	LangGRPCHeader = "language"
	// LangAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извелекаться (записываться) значение языка контекста вызова
	LangAMQPHeader = "language"

	// GatewayHTTPHeader заголовок для списка шлюзов, через которые прошёл запрос
	GatewayHTTPHeader = "X-Gateway"
	// GatewayGRPCHeader заголовок для списка шлюзов, через которые прошёл запрос
	GatewayGRPCHeader = "gateway"
	// GatewayAMQPHeader заголовок для списка шлюзов, через которые прошёл запрос
	GatewayAMQPHeader = "gateway"

	// KVHTTPHeaderPrefix префикс для заголовков дополнительной информации запроса
	KVHTTPHeaderPrefix = "X-KV-"
	// KVGRPCHeaderPrefix префикс для заголовков дополнительной информации запроса
	KVGRPCHeaderPrefix = "kv-"
	// KVAMQPHeaderPrefix префикс для заголовков дополнительной информации запроса
	KVAMQPHeaderPrefix = "kv_"

	// EditionHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) edition компании
	EditionHTTPHeader = "X-Edition"
	// EditionGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) edition компании
	EditionGRPCHeader = "edition"
)
