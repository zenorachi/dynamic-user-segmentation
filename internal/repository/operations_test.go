package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
	"regexp"
	"testing"
)

func TestOperationsRepository_CreateRelationsBySegmentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewOperations(db)

	type args struct {
		userId     int
		segmentIDs []int
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantIDs       []int
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				userId:     1,
				segmentIDs: []int{1, 2},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				// First iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("segment1"))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				// Second iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[1]).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("segment2"))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, args.segmentIDs[1]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment2", entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(102))

				mock.ExpectCommit()
			},
			wantIDs: []int{101, 102},
		},
		{
			name: "ERROR1",
			args: args{
				userId:     1,
				segmentIDs: []int{1},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnError(errors.New("test error"))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR2",
			args: args{
				userId:     1,
				segmentIDs: []int{1},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnError(errors.New("test error"))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR3",
			args: args{
				userId:     1,
				segmentIDs: []int{1},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("segment1"))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeAdd).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			operationsIDs, err := repo.CreateRelationsBySegmentIDs(context.Background(), tt.args.userId, tt.args.segmentIDs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantIDs, operationsIDs)
			}
		})
	}
}

func TestOperationsRepository_CreateRelationsBySegmentNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewOperations(db)

	type args struct {
		userId       int
		segmentNames []string
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantIDs       []int
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1", "segment2"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				// First iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				// Second iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[1]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, 2).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[1], entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(102))

				mock.ExpectCommit()
			},
			wantIDs: []int{101, 102},
		},
		{
			name: "ERROR1",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnError(errors.New("test error"))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR2",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, 1).
					WillReturnError(errors.New("test error"))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeAdd).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR3",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO relations (user_id, segment_id) VALUES ($1, $2)")).
					WithArgs(args.userId, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeAdd).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			operationsIDs, err := repo.CreateRelationsBySegmentNames(context.Background(), tt.args.userId, tt.args.segmentNames)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantIDs, operationsIDs)
			}
		})
	}
}

func TestOperationsRepository_DeleteRelationsBySegmentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewOperations(db)

	type args struct {
		userId     int
		segmentIDs []int
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantIDs       []int
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				userId:     1,
				segmentIDs: []int{1, 2},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				// First iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("segment1"))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				// Second iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[1]).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("segment2"))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, args.segmentIDs[1]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment2", entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(102))

				mock.ExpectCommit()
			},
			wantIDs: []int{101, 102},
		},
		{
			name: "ERROR1",
			args: args{
				userId:     1,
				segmentIDs: []int{1},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnError(errors.New("test error"))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR2",
			args: args{
				userId:     1,
				segmentIDs: []int{1},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnError(errors.New("test error"))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR3",
			args: args{
				userId:     1,
				segmentIDs: []int{1},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM segments WHERE id = $1")).
					WithArgs(args.segmentIDs[0]).
					WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("segment1"))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, args.segmentIDs[0]).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, "segment1", entity.TypeDelete).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			operationsIDs, err := repo.DeleteRelationsBySegmentIDs(context.Background(), tt.args.userId, tt.args.segmentIDs)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantIDs, operationsIDs)
			}
		})
	}
}

func TestOperationsRepository_DeleteRelationsBySegmentNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewOperations(db)

	type args struct {
		userId       int
		segmentNames []string
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantIDs       []int
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1", "segment2"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				// First iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				// Second iteration
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[1]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, 2).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[1], entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(102))

				mock.ExpectCommit()
			},
			wantIDs: []int{101, 102},
		},
		{
			name: "ERROR1",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnError(errors.New("test error"))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR2",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, 1).
					WillReturnError(errors.New("test error"))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeDelete).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(101))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "ERROR3",
			args: args{
				userId:       1,
				segmentNames: []string{"segment1"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM segments WHERE name = $1")).
					WithArgs(args.segmentNames[0]).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectExec(regexp.QuoteMeta("DELETE FROM relations WHERE user_id = $1 AND segment_id = $2")).
					WithArgs(args.userId, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO operations (user_id, segment_name, type) VALUES ($1, $2, $3)")).
					WithArgs(1, args.segmentNames[0], entity.TypeDelete).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			operationsIDs, err := repo.DeleteRelationsBySegmentNames(context.Background(), tt.args.userId, tt.args.segmentNames)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantIDs, operationsIDs)
			}
		})
	}
}
