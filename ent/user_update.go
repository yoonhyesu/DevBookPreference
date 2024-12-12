// Code generated by ent, DO NOT EDIT.

package ent

import (
	"DBP/ent/predicate"
	"DBP/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserId sets the "userId" field.
func (uu *UserUpdate) SetUserId(s string) *UserUpdate {
	uu.mutation.SetUserId(s)
	return uu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserId(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserId(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetUserName sets the "userName" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserName(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserName(*s)
	}
	return uu
}

// SetUserStatus sets the "userStatus" field.
func (uu *UserUpdate) SetUserStatus(b bool) *UserUpdate {
	uu.mutation.SetUserStatus(b)
	return uu
}

// SetNillableUserStatus sets the "userStatus" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserStatus(b *bool) *UserUpdate {
	if b != nil {
		uu.SetUserStatus(*b)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPhoneNumber sets the "phoneNumber" field.
func (uu *UserUpdate) SetPhoneNumber(s string) *UserUpdate {
	uu.mutation.SetPhoneNumber(s)
	return uu
}

// SetNillablePhoneNumber sets the "phoneNumber" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhoneNumber(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhoneNumber(*s)
	}
	return uu
}

// SetJobCd sets the "jobCd" field.
func (uu *UserUpdate) SetJobCd(i int) *UserUpdate {
	uu.mutation.ResetJobCd()
	uu.mutation.SetJobCd(i)
	return uu
}

// SetNillableJobCd sets the "jobCd" field if the given value is not nil.
func (uu *UserUpdate) SetNillableJobCd(i *int) *UserUpdate {
	if i != nil {
		uu.SetJobCd(*i)
	}
	return uu
}

// AddJobCd adds i to the "jobCd" field.
func (uu *UserUpdate) AddJobCd(i int) *UserUpdate {
	uu.mutation.AddJobCd(i)
	return uu
}

// ClearJobCd clears the value of the "jobCd" field.
func (uu *UserUpdate) ClearJobCd() *UserUpdate {
	uu.mutation.ClearJobCd()
	return uu
}

// SetProfileImage sets the "profileImage" field.
func (uu *UserUpdate) SetProfileImage(s string) *UserUpdate {
	uu.mutation.SetProfileImage(s)
	return uu
}

// SetNillableProfileImage sets the "profileImage" field if the given value is not nil.
func (uu *UserUpdate) SetNillableProfileImage(s *string) *UserUpdate {
	if s != nil {
		uu.SetProfileImage(*s)
	}
	return uu
}

// ClearProfileImage clears the value of the "profileImage" field.
func (uu *UserUpdate) ClearProfileImage() *UserUpdate {
	uu.mutation.ClearProfileImage()
	return uu
}

// SetGithubLink sets the "githubLink" field.
func (uu *UserUpdate) SetGithubLink(s string) *UserUpdate {
	uu.mutation.SetGithubLink(s)
	return uu
}

// SetNillableGithubLink sets the "githubLink" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGithubLink(s *string) *UserUpdate {
	if s != nil {
		uu.SetGithubLink(*s)
	}
	return uu
}

// ClearGithubLink clears the value of the "githubLink" field.
func (uu *UserUpdate) ClearGithubLink() *UserUpdate {
	uu.mutation.ClearGithubLink()
	return uu
}

// SetBlogLink sets the "blogLink" field.
func (uu *UserUpdate) SetBlogLink(s string) *UserUpdate {
	uu.mutation.SetBlogLink(s)
	return uu
}

// SetNillableBlogLink sets the "blogLink" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBlogLink(s *string) *UserUpdate {
	if s != nil {
		uu.SetBlogLink(*s)
	}
	return uu
}

// ClearBlogLink clears the value of the "blogLink" field.
func (uu *UserUpdate) ClearBlogLink() *UserUpdate {
	uu.mutation.ClearBlogLink()
	return uu
}

// SetUserText sets the "userText" field.
func (uu *UserUpdate) SetUserText(s string) *UserUpdate {
	uu.mutation.SetUserText(s)
	return uu
}

// SetNillableUserText sets the "userText" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUserText(s *string) *UserUpdate {
	if s != nil {
		uu.SetUserText(*s)
	}
	return uu
}

// ClearUserText clears the value of the "userText" field.
func (uu *UserUpdate) ClearUserText() *UserUpdate {
	uu.mutation.ClearUserText()
	return uu
}

// SetCompany sets the "company" field.
func (uu *UserUpdate) SetCompany(s string) *UserUpdate {
	uu.mutation.SetCompany(s)
	return uu
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCompany(s *string) *UserUpdate {
	if s != nil {
		uu.SetCompany(*s)
	}
	return uu
}

// ClearCompany clears the value of the "company" field.
func (uu *UserUpdate) ClearCompany() *UserUpdate {
	uu.mutation.ClearCompany()
	return uu
}

// SetSkill sets the "skill" field.
func (uu *UserUpdate) SetSkill(s string) *UserUpdate {
	uu.mutation.SetSkill(s)
	return uu
}

// SetNillableSkill sets the "skill" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSkill(s *string) *UserUpdate {
	if s != nil {
		uu.SetSkill(*s)
	}
	return uu
}

// ClearSkill clears the value of the "skill" field.
func (uu *UserUpdate) ClearSkill() *UserUpdate {
	uu.mutation.ClearSkill()
	return uu
}

// SetCreateDate sets the "createDate" field.
func (uu *UserUpdate) SetCreateDate(t time.Time) *UserUpdate {
	uu.mutation.SetCreateDate(t)
	return uu
}

// SetNillableCreateDate sets the "createDate" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateDate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreateDate(*t)
	}
	return uu
}

// SetUpdateDate sets the "updateDate" field.
func (uu *UserUpdate) SetUpdateDate(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateDate(t)
	return uu
}

// SetNillableUpdateDate sets the "updateDate" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdateDate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdateDate(*t)
	}
	return uu
}

// SetSessionToken sets the "sessionToken" field.
func (uu *UserUpdate) SetSessionToken(s string) *UserUpdate {
	uu.mutation.SetSessionToken(s)
	return uu
}

// SetNillableSessionToken sets the "sessionToken" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSessionToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetSessionToken(*s)
	}
	return uu
}

// ClearSessionToken clears the value of the "sessionToken" field.
func (uu *UserUpdate) ClearSessionToken() *UserUpdate {
	uu.mutation.ClearSessionToken()
	return uu
}

// SetSessionExpiry sets the "sessionExpiry" field.
func (uu *UserUpdate) SetSessionExpiry(t time.Time) *UserUpdate {
	uu.mutation.SetSessionExpiry(t)
	return uu
}

// SetNillableSessionExpiry sets the "sessionExpiry" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSessionExpiry(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetSessionExpiry(*t)
	}
	return uu
}

// ClearSessionExpiry clears the value of the "sessionExpiry" field.
func (uu *UserUpdate) ClearSessionExpiry() *UserUpdate {
	uu.mutation.ClearSessionExpiry()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.UserId(); ok {
		if err := user.UserIdValidator(v); err != nil {
			return &ValidationError{Name: "userId", err: fmt.Errorf(`ent: validator failed for field "User.userId": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "userName", err: fmt.Errorf(`ent: validator failed for field "User.userName": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PhoneNumber(); ok {
		if err := user.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phoneNumber", err: fmt.Errorf(`ent: validator failed for field "User.phoneNumber": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserId(); ok {
		_spec.SetField(user.FieldUserId, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserStatus(); ok {
		_spec.SetField(user.FieldUserStatus, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.PhoneNumber(); ok {
		_spec.SetField(user.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := uu.mutation.JobCd(); ok {
		_spec.SetField(user.FieldJobCd, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedJobCd(); ok {
		_spec.AddField(user.FieldJobCd, field.TypeInt, value)
	}
	if uu.mutation.JobCdCleared() {
		_spec.ClearField(user.FieldJobCd, field.TypeInt)
	}
	if value, ok := uu.mutation.ProfileImage(); ok {
		_spec.SetField(user.FieldProfileImage, field.TypeString, value)
	}
	if uu.mutation.ProfileImageCleared() {
		_spec.ClearField(user.FieldProfileImage, field.TypeString)
	}
	if value, ok := uu.mutation.GithubLink(); ok {
		_spec.SetField(user.FieldGithubLink, field.TypeString, value)
	}
	if uu.mutation.GithubLinkCleared() {
		_spec.ClearField(user.FieldGithubLink, field.TypeString)
	}
	if value, ok := uu.mutation.BlogLink(); ok {
		_spec.SetField(user.FieldBlogLink, field.TypeString, value)
	}
	if uu.mutation.BlogLinkCleared() {
		_spec.ClearField(user.FieldBlogLink, field.TypeString)
	}
	if value, ok := uu.mutation.UserText(); ok {
		_spec.SetField(user.FieldUserText, field.TypeString, value)
	}
	if uu.mutation.UserTextCleared() {
		_spec.ClearField(user.FieldUserText, field.TypeString)
	}
	if value, ok := uu.mutation.Company(); ok {
		_spec.SetField(user.FieldCompany, field.TypeString, value)
	}
	if uu.mutation.CompanyCleared() {
		_spec.ClearField(user.FieldCompany, field.TypeString)
	}
	if value, ok := uu.mutation.Skill(); ok {
		_spec.SetField(user.FieldSkill, field.TypeString, value)
	}
	if uu.mutation.SkillCleared() {
		_spec.ClearField(user.FieldSkill, field.TypeString)
	}
	if value, ok := uu.mutation.CreateDate(); ok {
		_spec.SetField(user.FieldCreateDate, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdateDate(); ok {
		_spec.SetField(user.FieldUpdateDate, field.TypeTime, value)
	}
	if value, ok := uu.mutation.SessionToken(); ok {
		_spec.SetField(user.FieldSessionToken, field.TypeString, value)
	}
	if uu.mutation.SessionTokenCleared() {
		_spec.ClearField(user.FieldSessionToken, field.TypeString)
	}
	if value, ok := uu.mutation.SessionExpiry(); ok {
		_spec.SetField(user.FieldSessionExpiry, field.TypeTime, value)
	}
	if uu.mutation.SessionExpiryCleared() {
		_spec.ClearField(user.FieldSessionExpiry, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserId sets the "userId" field.
func (uuo *UserUpdateOne) SetUserId(s string) *UserUpdateOne {
	uuo.mutation.SetUserId(s)
	return uuo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserId(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserId(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetUserName sets the "userName" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserName(*s)
	}
	return uuo
}

// SetUserStatus sets the "userStatus" field.
func (uuo *UserUpdateOne) SetUserStatus(b bool) *UserUpdateOne {
	uuo.mutation.SetUserStatus(b)
	return uuo
}

// SetNillableUserStatus sets the "userStatus" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserStatus(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetUserStatus(*b)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPhoneNumber sets the "phoneNumber" field.
func (uuo *UserUpdateOne) SetPhoneNumber(s string) *UserUpdateOne {
	uuo.mutation.SetPhoneNumber(s)
	return uuo
}

// SetNillablePhoneNumber sets the "phoneNumber" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhoneNumber(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhoneNumber(*s)
	}
	return uuo
}

// SetJobCd sets the "jobCd" field.
func (uuo *UserUpdateOne) SetJobCd(i int) *UserUpdateOne {
	uuo.mutation.ResetJobCd()
	uuo.mutation.SetJobCd(i)
	return uuo
}

// SetNillableJobCd sets the "jobCd" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableJobCd(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetJobCd(*i)
	}
	return uuo
}

// AddJobCd adds i to the "jobCd" field.
func (uuo *UserUpdateOne) AddJobCd(i int) *UserUpdateOne {
	uuo.mutation.AddJobCd(i)
	return uuo
}

// ClearJobCd clears the value of the "jobCd" field.
func (uuo *UserUpdateOne) ClearJobCd() *UserUpdateOne {
	uuo.mutation.ClearJobCd()
	return uuo
}

// SetProfileImage sets the "profileImage" field.
func (uuo *UserUpdateOne) SetProfileImage(s string) *UserUpdateOne {
	uuo.mutation.SetProfileImage(s)
	return uuo
}

// SetNillableProfileImage sets the "profileImage" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfileImage(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetProfileImage(*s)
	}
	return uuo
}

// ClearProfileImage clears the value of the "profileImage" field.
func (uuo *UserUpdateOne) ClearProfileImage() *UserUpdateOne {
	uuo.mutation.ClearProfileImage()
	return uuo
}

// SetGithubLink sets the "githubLink" field.
func (uuo *UserUpdateOne) SetGithubLink(s string) *UserUpdateOne {
	uuo.mutation.SetGithubLink(s)
	return uuo
}

// SetNillableGithubLink sets the "githubLink" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGithubLink(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGithubLink(*s)
	}
	return uuo
}

// ClearGithubLink clears the value of the "githubLink" field.
func (uuo *UserUpdateOne) ClearGithubLink() *UserUpdateOne {
	uuo.mutation.ClearGithubLink()
	return uuo
}

// SetBlogLink sets the "blogLink" field.
func (uuo *UserUpdateOne) SetBlogLink(s string) *UserUpdateOne {
	uuo.mutation.SetBlogLink(s)
	return uuo
}

// SetNillableBlogLink sets the "blogLink" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBlogLink(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBlogLink(*s)
	}
	return uuo
}

// ClearBlogLink clears the value of the "blogLink" field.
func (uuo *UserUpdateOne) ClearBlogLink() *UserUpdateOne {
	uuo.mutation.ClearBlogLink()
	return uuo
}

// SetUserText sets the "userText" field.
func (uuo *UserUpdateOne) SetUserText(s string) *UserUpdateOne {
	uuo.mutation.SetUserText(s)
	return uuo
}

// SetNillableUserText sets the "userText" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUserText(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUserText(*s)
	}
	return uuo
}

// ClearUserText clears the value of the "userText" field.
func (uuo *UserUpdateOne) ClearUserText() *UserUpdateOne {
	uuo.mutation.ClearUserText()
	return uuo
}

// SetCompany sets the "company" field.
func (uuo *UserUpdateOne) SetCompany(s string) *UserUpdateOne {
	uuo.mutation.SetCompany(s)
	return uuo
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCompany(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCompany(*s)
	}
	return uuo
}

// ClearCompany clears the value of the "company" field.
func (uuo *UserUpdateOne) ClearCompany() *UserUpdateOne {
	uuo.mutation.ClearCompany()
	return uuo
}

// SetSkill sets the "skill" field.
func (uuo *UserUpdateOne) SetSkill(s string) *UserUpdateOne {
	uuo.mutation.SetSkill(s)
	return uuo
}

// SetNillableSkill sets the "skill" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSkill(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSkill(*s)
	}
	return uuo
}

// ClearSkill clears the value of the "skill" field.
func (uuo *UserUpdateOne) ClearSkill() *UserUpdateOne {
	uuo.mutation.ClearSkill()
	return uuo
}

// SetCreateDate sets the "createDate" field.
func (uuo *UserUpdateOne) SetCreateDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreateDate(t)
	return uuo
}

// SetNillableCreateDate sets the "createDate" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateDate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreateDate(*t)
	}
	return uuo
}

// SetUpdateDate sets the "updateDate" field.
func (uuo *UserUpdateOne) SetUpdateDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateDate(t)
	return uuo
}

// SetNillableUpdateDate sets the "updateDate" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdateDate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdateDate(*t)
	}
	return uuo
}

// SetSessionToken sets the "sessionToken" field.
func (uuo *UserUpdateOne) SetSessionToken(s string) *UserUpdateOne {
	uuo.mutation.SetSessionToken(s)
	return uuo
}

// SetNillableSessionToken sets the "sessionToken" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSessionToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSessionToken(*s)
	}
	return uuo
}

// ClearSessionToken clears the value of the "sessionToken" field.
func (uuo *UserUpdateOne) ClearSessionToken() *UserUpdateOne {
	uuo.mutation.ClearSessionToken()
	return uuo
}

// SetSessionExpiry sets the "sessionExpiry" field.
func (uuo *UserUpdateOne) SetSessionExpiry(t time.Time) *UserUpdateOne {
	uuo.mutation.SetSessionExpiry(t)
	return uuo
}

// SetNillableSessionExpiry sets the "sessionExpiry" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSessionExpiry(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetSessionExpiry(*t)
	}
	return uuo
}

// ClearSessionExpiry clears the value of the "sessionExpiry" field.
func (uuo *UserUpdateOne) ClearSessionExpiry() *UserUpdateOne {
	uuo.mutation.ClearSessionExpiry()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.UserId(); ok {
		if err := user.UserIdValidator(v); err != nil {
			return &ValidationError{Name: "userId", err: fmt.Errorf(`ent: validator failed for field "User.userId": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.UserName(); ok {
		if err := user.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "userName", err: fmt.Errorf(`ent: validator failed for field "User.userName": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PhoneNumber(); ok {
		if err := user.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phoneNumber", err: fmt.Errorf(`ent: validator failed for field "User.phoneNumber": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserId(); ok {
		_spec.SetField(user.FieldUserId, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserStatus(); ok {
		_spec.SetField(user.FieldUserStatus, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PhoneNumber(); ok {
		_spec.SetField(user.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := uuo.mutation.JobCd(); ok {
		_spec.SetField(user.FieldJobCd, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedJobCd(); ok {
		_spec.AddField(user.FieldJobCd, field.TypeInt, value)
	}
	if uuo.mutation.JobCdCleared() {
		_spec.ClearField(user.FieldJobCd, field.TypeInt)
	}
	if value, ok := uuo.mutation.ProfileImage(); ok {
		_spec.SetField(user.FieldProfileImage, field.TypeString, value)
	}
	if uuo.mutation.ProfileImageCleared() {
		_spec.ClearField(user.FieldProfileImage, field.TypeString)
	}
	if value, ok := uuo.mutation.GithubLink(); ok {
		_spec.SetField(user.FieldGithubLink, field.TypeString, value)
	}
	if uuo.mutation.GithubLinkCleared() {
		_spec.ClearField(user.FieldGithubLink, field.TypeString)
	}
	if value, ok := uuo.mutation.BlogLink(); ok {
		_spec.SetField(user.FieldBlogLink, field.TypeString, value)
	}
	if uuo.mutation.BlogLinkCleared() {
		_spec.ClearField(user.FieldBlogLink, field.TypeString)
	}
	if value, ok := uuo.mutation.UserText(); ok {
		_spec.SetField(user.FieldUserText, field.TypeString, value)
	}
	if uuo.mutation.UserTextCleared() {
		_spec.ClearField(user.FieldUserText, field.TypeString)
	}
	if value, ok := uuo.mutation.Company(); ok {
		_spec.SetField(user.FieldCompany, field.TypeString, value)
	}
	if uuo.mutation.CompanyCleared() {
		_spec.ClearField(user.FieldCompany, field.TypeString)
	}
	if value, ok := uuo.mutation.Skill(); ok {
		_spec.SetField(user.FieldSkill, field.TypeString, value)
	}
	if uuo.mutation.SkillCleared() {
		_spec.ClearField(user.FieldSkill, field.TypeString)
	}
	if value, ok := uuo.mutation.CreateDate(); ok {
		_spec.SetField(user.FieldCreateDate, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdateDate(); ok {
		_spec.SetField(user.FieldUpdateDate, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.SessionToken(); ok {
		_spec.SetField(user.FieldSessionToken, field.TypeString, value)
	}
	if uuo.mutation.SessionTokenCleared() {
		_spec.ClearField(user.FieldSessionToken, field.TypeString)
	}
	if value, ok := uuo.mutation.SessionExpiry(); ok {
		_spec.SetField(user.FieldSessionExpiry, field.TypeTime, value)
	}
	if uuo.mutation.SessionExpiryCleared() {
		_spec.ClearField(user.FieldSessionExpiry, field.TypeTime)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
