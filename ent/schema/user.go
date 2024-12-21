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
		field.String("userId").
			Unique().
			NotEmpty().
			StorageKey("USER_ID"),
		field.String("password").
			NotEmpty().
			StorageKey("PASSWORD").
			Sensitive(),
		field.String("userName").
			NotEmpty().
			StorageKey("USER_NAME"),
		field.Bool("userStatus").
			Default(false).
			StorageKey("USER_STATUS"),
		field.String("email").
			Unique().
			StorageKey("EMAIL"),
		field.String("phoneNumber").
			Unique().
			NotEmpty().
			StorageKey("PHONE_NUMBER"),
		field.Int("jobCd").
			Optional().
			StorageKey("JOB_CODE"),
		field.String("profileImage").
			Optional().
			StorageKey("PROFILE_IMAGE"),
		field.String("githubLink").
			Optional().
			StorageKey("GITHUB_LINK"),
		field.String("blogLink").
			Optional().
			StorageKey("BLOG_LINK"),
		field.String("userText").
			Optional().
			StorageKey("USER_TEXT"),
		field.String("company").
			Optional().
			StorageKey("COMPANY"),
		field.String("skill").
			Optional().
			StorageKey("SKILL"),
		field.Time("createDate").
			Default(time.Now()).
			StorageKey("CREATE_DATE"),
		field.Time("updateDate").
			Default(time.Now()).
			StorageKey("UPDATE_DATE"),
		field.Time("lastLoginDate").
			Optional().
			StorageKey("LAST_LOGIN_DATE"),
		field.Bool("isAdmin").
			Default(false).
			StorageKey("IS_ADMIN"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
