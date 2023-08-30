package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zenorachi/dynamic-user-segmentation/internal/entity"
)

func TestSegments_Create(t *testing.T) {
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
