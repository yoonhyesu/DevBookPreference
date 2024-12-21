// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "USER_ID"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "PASSWORD"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "USER_NAME"
	// FieldUserStatus holds the string denoting the userstatus field in the database.
	FieldUserStatus = "USER_STATUS"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "EMAIL"
	// FieldPhoneNumber holds the string denoting the phonenumber field in the database.
	FieldPhoneNumber = "PHONE_NUMBER"
	// FieldJobCd holds the string denoting the jobcd field in the database.
	FieldJobCd = "JOB_CODE"
	// FieldProfileImage holds the string denoting the profileimage field in the database.
	FieldProfileImage = "PROFILE_IMAGE"
	// FieldGithubLink holds the string denoting the githublink field in the database.
	FieldGithubLink = "GITHUB_LINK"
	// FieldBlogLink holds the string denoting the bloglink field in the database.
	FieldBlogLink = "BLOG_LINK"
	// FieldUserText holds the string denoting the usertext field in the database.
	FieldUserText = "USER_TEXT"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "COMPANY"
	// FieldSkill holds the string denoting the skill field in the database.
	FieldSkill = "SKILL"
	// FieldCreateDate holds the string denoting the createdate field in the database.
	FieldCreateDate = "CREATE_DATE"
	// FieldUpdateDate holds the string denoting the updatedate field in the database.
	FieldUpdateDate = "UPDATE_DATE"
	// FieldLastLoginDate holds the string denoting the lastlogindate field in the database.
	FieldLastLoginDate = "LAST_LOGIN_DATE"
	// FieldIsAdmin holds the string denoting the isadmin field in the database.
	FieldIsAdmin = "IS_ADMIN"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldPassword,
	FieldUserName,
	FieldUserStatus,
	FieldEmail,
	FieldPhoneNumber,
	FieldJobCd,
	FieldProfileImage,
	FieldGithubLink,
	FieldBlogLink,
	FieldUserText,
	FieldCompany,
	FieldSkill,
	FieldCreateDate,
	FieldUpdateDate,
	FieldLastLoginDate,
	FieldIsAdmin,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UserIdValidator is a validator for the "userId" field. It is called by the builders before save.
	UserIdValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// UserNameValidator is a validator for the "userName" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// DefaultUserStatus holds the default value on creation for the "userStatus" field.
	DefaultUserStatus bool
	// PhoneNumberValidator is a validator for the "phoneNumber" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// DefaultCreateDate holds the default value on creation for the "createDate" field.
	DefaultCreateDate time.Time
	// DefaultUpdateDate holds the default value on creation for the "updateDate" field.
	DefaultUpdateDate time.Time
	// DefaultIsAdmin holds the default value on creation for the "isAdmin" field.
	DefaultIsAdmin bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserId orders the results by the userId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByUserName orders the results by the userName field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByUserStatus orders the results by the userStatus field.
func ByUserStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserStatus, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phoneNumber field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByJobCd orders the results by the jobCd field.
func ByJobCd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJobCd, opts...).ToFunc()
}

// ByProfileImage orders the results by the profileImage field.
func ByProfileImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfileImage, opts...).ToFunc()
}

// ByGithubLink orders the results by the githubLink field.
func ByGithubLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGithubLink, opts...).ToFunc()
}

// ByBlogLink orders the results by the blogLink field.
func ByBlogLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlogLink, opts...).ToFunc()
}

// ByUserText orders the results by the userText field.
func ByUserText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserText, opts...).ToFunc()
}

// ByCompany orders the results by the company field.
func ByCompany(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompany, opts...).ToFunc()
}

// BySkill orders the results by the skill field.
func BySkill(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSkill, opts...).ToFunc()
}

// ByCreateDate orders the results by the createDate field.
func ByCreateDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateDate, opts...).ToFunc()
}

// ByUpdateDate orders the results by the updateDate field.
func ByUpdateDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateDate, opts...).ToFunc()
}

// ByLastLoginDate orders the results by the lastLoginDate field.
func ByLastLoginDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginDate, opts...).ToFunc()
}

// ByIsAdmin orders the results by the isAdmin field.
func ByIsAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAdmin, opts...).ToFunc()
}
