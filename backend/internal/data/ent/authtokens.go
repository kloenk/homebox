// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/authroles"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/authtokens"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/user"
)

// AuthTokens is the model entity for the AuthTokens schema.
type AuthTokens struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Token holds the value of the "token" field.
	Token []byte `json:"token,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthTokensQuery when eager-loading is set.
	Edges            AuthTokensEdges `json:"edges"`
	user_auth_tokens *uuid.UUID
	selectValues     sql.SelectValues
}

// AuthTokensEdges holds the relations/edges for other nodes in the graph.
type AuthTokensEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Roles holds the value of the roles edge.
	Roles *AuthRoles `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthTokensEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthTokensEdges) RolesOrErr() (*AuthRoles, error) {
	if e.Roles != nil {
		return e.Roles, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: authroles.Label}
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthTokens) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authtokens.FieldToken:
			values[i] = new([]byte)
		case authtokens.FieldCreatedAt, authtokens.FieldUpdatedAt, authtokens.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case authtokens.FieldID:
			values[i] = new(uuid.UUID)
		case authtokens.ForeignKeys[0]: // user_auth_tokens
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthTokens fields.
func (at *AuthTokens) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authtokens.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				at.ID = *value
			}
		case authtokens.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				at.CreatedAt = value.Time
			}
		case authtokens.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				at.UpdatedAt = value.Time
			}
		case authtokens.FieldToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value != nil {
				at.Token = *value
			}
		case authtokens.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				at.ExpiresAt = value.Time
			}
		case authtokens.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_auth_tokens", values[i])
			} else if value.Valid {
				at.user_auth_tokens = new(uuid.UUID)
				*at.user_auth_tokens = *value.S.(*uuid.UUID)
			}
		default:
			at.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AuthTokens.
// This includes values selected through modifiers, order, etc.
func (at *AuthTokens) Value(name string) (ent.Value, error) {
	return at.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the AuthTokens entity.
func (at *AuthTokens) QueryUser() *UserQuery {
	return NewAuthTokensClient(at.config).QueryUser(at)
}

// QueryRoles queries the "roles" edge of the AuthTokens entity.
func (at *AuthTokens) QueryRoles() *AuthRolesQuery {
	return NewAuthTokensClient(at.config).QueryRoles(at)
}

// Update returns a builder for updating this AuthTokens.
// Note that you need to call AuthTokens.Unwrap() before calling this method if this AuthTokens
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *AuthTokens) Update() *AuthTokensUpdateOne {
	return NewAuthTokensClient(at.config).UpdateOne(at)
}

// Unwrap unwraps the AuthTokens entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *AuthTokens) Unwrap() *AuthTokens {
	_tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthTokens is not a transactional entity")
	}
	at.config.driver = _tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *AuthTokens) String() string {
	var builder strings.Builder
	builder.WriteString("AuthTokens(")
	builder.WriteString(fmt.Sprintf("id=%v, ", at.ID))
	builder.WriteString("created_at=")
	builder.WriteString(at.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(at.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(fmt.Sprintf("%v", at.Token))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(at.ExpiresAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AuthTokensSlice is a parsable slice of AuthTokens.
type AuthTokensSlice []*AuthTokens
