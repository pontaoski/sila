// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addCommentStmt, err = db.PrepareContext(ctx, addComment); err != nil {
		return nil, fmt.Errorf("error preparing query AddComment: %w", err)
	}
	if q.addDependencyStmt, err = db.PrepareContext(ctx, addDependency); err != nil {
		return nil, fmt.Errorf("error preparing query AddDependency: %w", err)
	}
	if q.clearDependenciesForBugStmt, err = db.PrepareContext(ctx, clearDependenciesForBug); err != nil {
		return nil, fmt.Errorf("error preparing query ClearDependenciesForBug: %w", err)
	}
	if q.clearDependenciesOnBugStmt, err = db.PrepareContext(ctx, clearDependenciesOnBug); err != nil {
		return nil, fmt.Errorf("error preparing query ClearDependenciesOnBug: %w", err)
	}
	if q.countBugsStmt, err = db.PrepareContext(ctx, countBugs); err != nil {
		return nil, fmt.Errorf("error preparing query CountBugs: %w", err)
	}
	if q.countComponentBugsStmt, err = db.PrepareContext(ctx, countComponentBugs); err != nil {
		return nil, fmt.Errorf("error preparing query CountComponentBugs: %w", err)
	}
	if q.createBugStmt, err = db.PrepareContext(ctx, createBug); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBug: %w", err)
	}
	if q.createComponentStmt, err = db.PrepareContext(ctx, createComponent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateComponent: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.deleteBugStmt, err = db.PrepareContext(ctx, deleteBug); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteBug: %w", err)
	}
	if q.deleteComponentStmt, err = db.PrepareContext(ctx, deleteComponent); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteComponent: %w", err)
	}
	if q.deleteProductStmt, err = db.PrepareContext(ctx, deleteProduct); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProduct: %w", err)
	}
	if q.editCommentStmt, err = db.PrepareContext(ctx, editComment); err != nil {
		return nil, fmt.Errorf("error preparing query EditComment: %w", err)
	}
	if q.editComponentStmt, err = db.PrepareContext(ctx, editComponent); err != nil {
		return nil, fmt.Errorf("error preparing query EditComponent: %w", err)
	}
	if q.editDescriptionStmt, err = db.PrepareContext(ctx, editDescription); err != nil {
		return nil, fmt.Errorf("error preparing query EditDescription: %w", err)
	}
	if q.editProductStmt, err = db.PrepareContext(ctx, editProduct); err != nil {
		return nil, fmt.Errorf("error preparing query EditProduct: %w", err)
	}
	if q.editSeverityStmt, err = db.PrepareContext(ctx, editSeverity); err != nil {
		return nil, fmt.Errorf("error preparing query EditSeverity: %w", err)
	}
	if q.editStatusStmt, err = db.PrepareContext(ctx, editStatus); err != nil {
		return nil, fmt.Errorf("error preparing query EditStatus: %w", err)
	}
	if q.listBugsStmt, err = db.PrepareContext(ctx, listBugs); err != nil {
		return nil, fmt.Errorf("error preparing query ListBugs: %w", err)
	}
	if q.redactCommentStmt, err = db.PrepareContext(ctx, redactComment); err != nil {
		return nil, fmt.Errorf("error preparing query RedactComment: %w", err)
	}
	if q.relocateBugStmt, err = db.PrepareContext(ctx, relocateBug); err != nil {
		return nil, fmt.Errorf("error preparing query RelocateBug: %w", err)
	}
	if q.removeDependencyStmt, err = db.PrepareContext(ctx, removeDependency); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveDependency: %w", err)
	}
	if q.requiresStmt, err = db.PrepareContext(ctx, requires); err != nil {
		return nil, fmt.Errorf("error preparing query Requires: %w", err)
	}
	if q.whatRequiresStmt, err = db.PrepareContext(ctx, whatRequires); err != nil {
		return nil, fmt.Errorf("error preparing query WhatRequires: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addCommentStmt != nil {
		if cerr := q.addCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addCommentStmt: %w", cerr)
		}
	}
	if q.addDependencyStmt != nil {
		if cerr := q.addDependencyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addDependencyStmt: %w", cerr)
		}
	}
	if q.clearDependenciesForBugStmt != nil {
		if cerr := q.clearDependenciesForBugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing clearDependenciesForBugStmt: %w", cerr)
		}
	}
	if q.clearDependenciesOnBugStmt != nil {
		if cerr := q.clearDependenciesOnBugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing clearDependenciesOnBugStmt: %w", cerr)
		}
	}
	if q.countBugsStmt != nil {
		if cerr := q.countBugsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countBugsStmt: %w", cerr)
		}
	}
	if q.countComponentBugsStmt != nil {
		if cerr := q.countComponentBugsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countComponentBugsStmt: %w", cerr)
		}
	}
	if q.createBugStmt != nil {
		if cerr := q.createBugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBugStmt: %w", cerr)
		}
	}
	if q.createComponentStmt != nil {
		if cerr := q.createComponentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createComponentStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.deleteBugStmt != nil {
		if cerr := q.deleteBugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteBugStmt: %w", cerr)
		}
	}
	if q.deleteComponentStmt != nil {
		if cerr := q.deleteComponentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteComponentStmt: %w", cerr)
		}
	}
	if q.deleteProductStmt != nil {
		if cerr := q.deleteProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductStmt: %w", cerr)
		}
	}
	if q.editCommentStmt != nil {
		if cerr := q.editCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editCommentStmt: %w", cerr)
		}
	}
	if q.editComponentStmt != nil {
		if cerr := q.editComponentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editComponentStmt: %w", cerr)
		}
	}
	if q.editDescriptionStmt != nil {
		if cerr := q.editDescriptionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editDescriptionStmt: %w", cerr)
		}
	}
	if q.editProductStmt != nil {
		if cerr := q.editProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editProductStmt: %w", cerr)
		}
	}
	if q.editSeverityStmt != nil {
		if cerr := q.editSeverityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editSeverityStmt: %w", cerr)
		}
	}
	if q.editStatusStmt != nil {
		if cerr := q.editStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing editStatusStmt: %w", cerr)
		}
	}
	if q.listBugsStmt != nil {
		if cerr := q.listBugsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listBugsStmt: %w", cerr)
		}
	}
	if q.redactCommentStmt != nil {
		if cerr := q.redactCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing redactCommentStmt: %w", cerr)
		}
	}
	if q.relocateBugStmt != nil {
		if cerr := q.relocateBugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing relocateBugStmt: %w", cerr)
		}
	}
	if q.removeDependencyStmt != nil {
		if cerr := q.removeDependencyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeDependencyStmt: %w", cerr)
		}
	}
	if q.requiresStmt != nil {
		if cerr := q.requiresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing requiresStmt: %w", cerr)
		}
	}
	if q.whatRequiresStmt != nil {
		if cerr := q.whatRequiresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing whatRequiresStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                          DBTX
	tx                          *sql.Tx
	addCommentStmt              *sql.Stmt
	addDependencyStmt           *sql.Stmt
	clearDependenciesForBugStmt *sql.Stmt
	clearDependenciesOnBugStmt  *sql.Stmt
	countBugsStmt               *sql.Stmt
	countComponentBugsStmt      *sql.Stmt
	createBugStmt               *sql.Stmt
	createComponentStmt         *sql.Stmt
	createProductStmt           *sql.Stmt
	deleteBugStmt               *sql.Stmt
	deleteComponentStmt         *sql.Stmt
	deleteProductStmt           *sql.Stmt
	editCommentStmt             *sql.Stmt
	editComponentStmt           *sql.Stmt
	editDescriptionStmt         *sql.Stmt
	editProductStmt             *sql.Stmt
	editSeverityStmt            *sql.Stmt
	editStatusStmt              *sql.Stmt
	listBugsStmt                *sql.Stmt
	redactCommentStmt           *sql.Stmt
	relocateBugStmt             *sql.Stmt
	removeDependencyStmt        *sql.Stmt
	requiresStmt                *sql.Stmt
	whatRequiresStmt            *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                          tx,
		tx:                          tx,
		addCommentStmt:              q.addCommentStmt,
		addDependencyStmt:           q.addDependencyStmt,
		clearDependenciesForBugStmt: q.clearDependenciesForBugStmt,
		clearDependenciesOnBugStmt:  q.clearDependenciesOnBugStmt,
		countBugsStmt:               q.countBugsStmt,
		countComponentBugsStmt:      q.countComponentBugsStmt,
		createBugStmt:               q.createBugStmt,
		createComponentStmt:         q.createComponentStmt,
		createProductStmt:           q.createProductStmt,
		deleteBugStmt:               q.deleteBugStmt,
		deleteComponentStmt:         q.deleteComponentStmt,
		deleteProductStmt:           q.deleteProductStmt,
		editCommentStmt:             q.editCommentStmt,
		editComponentStmt:           q.editComponentStmt,
		editDescriptionStmt:         q.editDescriptionStmt,
		editProductStmt:             q.editProductStmt,
		editSeverityStmt:            q.editSeverityStmt,
		editStatusStmt:              q.editStatusStmt,
		listBugsStmt:                q.listBugsStmt,
		redactCommentStmt:           q.redactCommentStmt,
		relocateBugStmt:             q.relocateBugStmt,
		removeDependencyStmt:        q.removeDependencyStmt,
		requiresStmt:                q.requiresStmt,
		whatRequiresStmt:            q.whatRequiresStmt,
	}
}
