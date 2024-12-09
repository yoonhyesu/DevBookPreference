package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			StorageKey("ID"),
		field.String("user_id").
			Unique().
			NotEmpty().
			StorageKey("USER_ID"),
		field.String("password").
			NotEmpty().
			StorageKey("PASSWORD").
			Sensitive(),
		field.String("user_name").
			NotEmpty().
			StorageKey("USER_NAME"),
		field.Bool("user_status").
			Default(false).
			StorageKey("USER_STATUS"),
		field.String("email").
			Unique().
			StorageKey("EMAIL"),
		field.String("phone_number").
			Unique().
			NotEmpty().
			StorageKey("PHONE_NUMBER"),
		field.Int("job_cd").
			Optional().
			StorageKey("JOB_CODE"),
		field.String("profile_image").
			Optional().
			StorageKey("PROFILE_IMAGE"),
		field.String("github_link").
			Optional().
			StorageKey("GITHUB_LINK"),
		field.String("blog_link").
			Optional().
			StorageKey("BLOG_LINK"),
		field.String("user_text").
			Optional().
			StorageKey("USER_TEXT"),
		field.String("company").
			Optional().
			StorageKey("COMPANY"),
		field.String("skill").
			Optional().
			StorageKey("SKILL"),
		field.Time("create_date").
			Default(time.Now()).
			StorageKey("CREATE_DATE"),
		field.Time("update_date").
			Default(time.Now()).
			StorageKey("UPDATE_DATE"),
		field.String("session_token").
			Optional().
			Sensitive().
			StorageKey("SESSION_TOKEN"),
		field.Time("session_expiry").
			Optional().
			StorageKey("SESSION_EXPIRY"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
