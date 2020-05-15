// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Severity string

const (
	SeverityWishlist              Severity = "Wishlist"
	SeverityMinor                 Severity = "Minor"
	SeverityNormal                Severity = "Normal"
	SeverityGrave                 Severity = "Grave"
	SeverityMajor                 Severity = "Major"
	SeverityCritical              Severity = "Critical"
	SeverityCrash                 Severity = "Crash"
	SeverityTask                  Severity = "Task"
	SeverityMinorSecurityIssue    Severity = "MinorSecurityIssue"
	SeverityMediumSecurityIssue   Severity = "MediumSecurityIssue"
	SeverityMajorSecurityIssue    Severity = "MajorSecurityIssue"
	SeverityCriticalSecurityIssue Severity = "CriticalSecurityIssue"
)

func (e *Severity) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Severity(s)
	case string:
		*e = Severity(s)
	default:
		return fmt.Errorf("unsupported scan type for Severity: %T", src)
	}
	return nil
}

type Status string

const (
	StatusReported  Status = "Reported"
	StatusConfirmed Status = "Confirmed"
	StatusAssigned  Status = "Assigned"
	StatusResolved  Status = "Resolved"
	StatusNeedInfo  Status = "NeedInfo"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type Bug struct {
	ID             int64          `json:"id"`
	Title          sql.NullString `json:"title"`
	CreatedAt      time.Time      `json:"created_at"`
	EditedAt       sql.NullTime   `json:"edited_at"`
	BugDescription sql.NullString `json:"bug_description"`
	BugSeverity    Severity       `json:"bug_severity"`
	BugStatus      Status         `json:"bug_status"`
	ComponentID    int64          `json:"component_id"`
	AuthorID       int64          `json:"author_id"`
}

type BugDependency struct {
	RequiredBy int64 `json:"required_by"`
	DependsOn  int64 `json:"depends_on"`
}

type Comment struct {
	ID          int64          `json:"id"`
	CommentText sql.NullString `json:"comment_text"`
	CreatedAt   time.Time      `json:"created_at"`
	EditedAt    sql.NullTime   `json:"edited_at"`
	Redacted    sql.NullBool   `json:"redacted"`
	BugID       int64          `json:"bug_id"`
	AuthorID    int64          `json:"author_id"`
}

type Component struct {
	ID                   int64          `json:"id"`
	Name                 sql.NullString `json:"name"`
	ComponentDescription sql.NullString `json:"component_description"`
	ProductID            int64          `json:"product_id"`
}

type Product struct {
	ID                 int64          `json:"id"`
	Name               sql.NullString `json:"name"`
	ProductDescription sql.NullString `json:"product_description"`
	Active             sql.NullBool   `json:"active"`
}

type User struct {
	ID       int64          `json:"id"`
	Email    sql.NullString `json:"email"`
	Fullname sql.NullString `json:"fullname"`
}
