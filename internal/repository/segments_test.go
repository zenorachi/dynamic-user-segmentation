package repository

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

func TestSegmentsRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)

	type args struct {
		segment entity.Segment
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		args          args
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				segment: entity.Segment{
					ID:            1,
					Name:          "segment",
					AssignPercent: 1,
				},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedExec := "INSERT INTO segments (name, assign_percent) VALUES ($1, $2) RETURNING id"
				mock.ExpectQuery(regexp.QuoteMeta(expectedExec)).
					WithArgs(args.segment.Name, args.segment.AssignPercent).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "ERROR",
			args: args{
				segment: entity.Segment{},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedExec := "INSERT INTO segments (name, assign_percent) VALUES ($1, $2) RETURNING id"
				mock.ExpectExec(regexp.QuoteMeta(expectedExec)).
					WithArgs(args.segment.Name, args.segment.AssignPercent).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			_, err := repo.Create(context.Background(), tt.args.segment)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSegmentsRepository_GetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)

	type args struct {
		name string
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		segment       entity.Segment
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				name: "name",
			},
			segment: entity.Segment{
				ID:            1,
				Name:          "name",
				AssignPercent: 1,
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "assign_percent"}).
					AddRow(1, "name", 1.0)

				expectedQuery := "SELECT * FROM segments WHERE name = $1"
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(args.name).
					WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name: "ERROR",
			args: args{
				name: "name",
			},
			segment: entity.Segment{},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedQuery := "SELECT * FROM segments WHERE name = $1"
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(args.name).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)

			segment, err := repo.GetByName(context.Background(), tt.args.name)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.segment, segment)
			}
		})
	}
}

func TestSegmentsRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)

	type args struct {
		id int
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		segment       entity.Segment
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				id: 1,
			},
			segment: entity.Segment{
				ID:            1,
				Name:          "name",
				AssignPercent: 1,
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "assign_percent"}).
					AddRow(1, "name", 1.0)

				expectedQuery := "SELECT * FROM segments WHERE id = $1"
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(args.id).
					WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name: "ERROR",
			args: args{
				id: 1,
			},
			segment: entity.Segment{},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedQuery := "SELECT * FROM segments WHERE id = $1"
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(args.id).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)

			segment, err := repo.GetByID(context.Background(), tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.segment, segment)
			}
		})
	}
}

func TestSegmentsRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)
	type mockBehaviour func()

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		segments      []entity.Segment
		wantErr       bool
	}{
		{
			name: "OK",
			segments: []entity.Segment{
				{
					ID:            1,
					Name:          "name1",
					AssignPercent: 1,
				},
				{
					ID:            2,
					Name:          "name2",
					AssignPercent: 1,
				},
			},
			mockBehaviour: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "name", "assign_percent"}).
					AddRow(1, "name1", 1.0).
					AddRow(2, "name2", 1.0)

				expectedQuery := "SELECT * FROM segments"
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name:     "ERROR",
			segments: []entity.Segment{},
			mockBehaviour: func() {
				mock.ExpectBegin()

				expectedQuery := "SELECT * FROM segments"
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour()

			segments, err := repo.GetAll(context.Background())

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.segments, segments)
			}
		})
	}
}

func TestSegmentsRepository_GetActiveUsersBySegmentID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)

	type args struct {
		id int
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		users         []entity.User
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				id: 1,
			},
			users: []entity.User{
				{ID: 1, Login: "user1"},
				{ID: 2, Login: "user2"},
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "login", "registered_at"}).
					AddRow(1, "user1", time.Time{}).
					AddRow(2, "user2", time.Time{})

				expectedQuery := fmt.Sprintf(
					"SELECT %s.id, %s.login, %s.registered_at FROM %s JOIN %s ON %s.user_id = %s.id WHERE %s.segment_id = $1",
					collectionUsers, collectionUsers, collectionUsers, collectionUsers, collectionRelations, collectionRelations, collectionUsers, collectionRelations)
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(args.id).
					WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name: "ERROR",
			args: args{
				id: 2,
			},
			users: []entity.User{},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedQuery := fmt.Sprintf(
					"SELECT %s.id, %s.login, %s.registered_at FROM %s JOIN %s ON %s.user_id = %s.id WHERE %s.segment_id = $1",
					collectionUsers, collectionUsers, collectionUsers, collectionUsers, collectionRelations, collectionRelations, collectionUsers, collectionRelations)
				mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(args.id).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			segments, err := repo.GetActiveUsersBySegmentID(context.Background(), tt.args.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.users, segments)
			}
		})
	}
}

func TestSegmentsRepository_DeleteByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)

	type args struct {
		segmentName string
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				segmentName: "segment",
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedDeleteQuery := fmt.Sprintf("DELETE FROM %s WHERE name = $1", collectionSegments)
				mock.ExpectExec(regexp.QuoteMeta(expectedDeleteQuery)).WithArgs(args.segmentName).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "ERROR",
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedDeleteQuery := fmt.Sprintf("DELETE FROM %s WHERE name = $1", collectionSegments)
				mock.ExpectExec(regexp.QuoteMeta(expectedDeleteQuery)).WithArgs(args.segmentName).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			err := repo.DeleteByName(context.Background(), tt.args.segmentName)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSegmentsRepository_DeleteByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating database connection: %v\n", err)
	}
	defer db.Close()

	repo := NewSegments(db)

	type args struct {
		segmentId int
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				segmentId: 1,
			},
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedDeleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", collectionSegments)
				mock.ExpectExec(regexp.QuoteMeta(expectedDeleteQuery)).WithArgs(args.segmentId).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "ERROR",
			mockBehaviour: func(args args) {
				mock.ExpectBegin()

				expectedDeleteQuery := fmt.Sprintf("DELETE FROM %s WHERE name = $1", collectionSegments)
				mock.ExpectExec(regexp.QuoteMeta(expectedDeleteQuery)).WithArgs(args.segmentId).
					WillReturnError(errors.New("test error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			err := repo.DeleteByID(context.Background(), tt.args.segmentId)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
