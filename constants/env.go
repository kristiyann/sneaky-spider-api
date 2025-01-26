package constants

const (
	EnvEnvironment                 = "ENVIRONMENT"
	EnvEnvironmentValueDevelopment = "development"
	EnvEnvironmentValueProduction  = "production"

	EnvOriginaAllowed = "ORIGIN_ALLOWED"
	EnvPostgresUrl    = "POSTGRES_URL"

	EnvSupabaseApiUrl = "SUPABASE_API_URL"
	EnvSupabaseApiKey = "SUPABASE_API_KEY"
	EnvSupabaseSecret = "SUPABASE_SECRET"

	EnvApiAddr = "URL"

	EnvSmtpEmail    = "SENDER_EMAIL"
	EnvSmtpPassword = "SENDER_EMAIL_PASSWORD"

	EnvStripeReturnUrl      = "STRIPE_RETURN_URL"
	EnvStripeErrorReturnUrl = "STRIPE_ERROR_RETURN_URL"
	EnvStripeKey            = "STRIPE_KEY"
	EnvStripeWebhookSecret  = "STRIPE_WEBHOOK_SECRET"

	EnvResendAPIKey = "RESEND_API_KEY"
)
