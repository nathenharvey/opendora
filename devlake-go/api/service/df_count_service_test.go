package service

import (
	"fmt"
	"github.com/devoteamnl/opendora/api/models"
	"github.com/devoteamnl/opendora/api/sql_client"
	"github.com/devoteamnl/opendora/api/sql_client/sql_queries"
	"reflect"
	"testing"
)

func TestDfCountService_ServeRequest(t *testing.T) {

	exampleWeeklyDataPoints := []models.DataPoint{{Key: "202338", Value: 0}, {Key: "202337", Value: 1}, {Key: "202336", Value: 2}}
	exampleMonthlyDataPoints := []models.DataPoint{{Key: "23/04", Value: 6}, {Key: "23/03", Value: 5}, {Key: "23/02", Value: 4}}
	exampleQuarterlyDataPoints := []models.DataPoint{{Key: "2024-01-01", Value: 6}, {Key: "2023-10-01", Value: 5}, {Key: "2023-07-01", Value: 4}}

	tests := []struct {
		name           string
		params         ServiceParameters
		mockClient     sql_client.MockClient
		expectResponse models.Response
		expectError    string
	}{
		{
			name:   "should return an error with an unexpected error from the database",
			params: ServiceParameters{TypeQuery: "df_count", Aggregation: "weekly", Project: "", To: 0, From: 0},
			mockClient: sql_client.MockClient{
				MockDataMap: map[string]sql_client.MockDataReturn{
					sql_queries.WeeklyDeploymentSql:    {Err: fmt.Errorf("error from weekly query")},
					sql_queries.MonthlyDeploymentSql:   {Err: fmt.Errorf("error from monthly query")},
					sql_queries.QuarterlyDeploymentSql: {Err: fmt.Errorf("error from quarterly query")},
				},
			},
			expectResponse: models.Response{Aggregation: "weekly", DataPoints: nil},
			expectError:    "error from weekly query",
		},
		{
			name:   "should return an error with an unexpected error from the database",
			params: ServiceParameters{TypeQuery: "df_count", Aggregation: "monthly", Project: "", To: 0, From: 0},
			mockClient: sql_client.MockClient{
				MockDataMap: map[string]sql_client.MockDataReturn{
					sql_queries.WeeklyDeploymentSql:    {Err: fmt.Errorf("error from weekly query")},
					sql_queries.MonthlyDeploymentSql:   {Err: fmt.Errorf("error from monthly query")},
					sql_queries.QuarterlyDeploymentSql: {Err: fmt.Errorf("error from quarterly query")},
				},
			},
			expectResponse: models.Response{Aggregation: "monthly", DataPoints: nil},
			expectError:    "error from monthly query",
		},
		{
			name:   "should return an error with an unexpected error from the database",
			params: ServiceParameters{TypeQuery: "df_count", Aggregation: "quarterly", Project: "", To: 0, From: 0},
			mockClient: sql_client.MockClient{
				MockDataMap: map[string]sql_client.MockDataReturn{
					sql_queries.WeeklyDeploymentSql:    {Err: fmt.Errorf("error from weekly query")},
					sql_queries.MonthlyDeploymentSql:   {Err: fmt.Errorf("error from monthly query")},
					sql_queries.QuarterlyDeploymentSql: {Err: fmt.Errorf("error from quarterly query")},
				},
			},
			expectResponse: models.Response{Aggregation: "quarterly", DataPoints: nil},
			expectError:    "error from quarterly query",
		},
		{
			name:   "should return weekly data points from the database",
			params: ServiceParameters{TypeQuery: "df_count", Aggregation: "weekly", Project: "", To: 0, From: 0},
			mockClient: sql_client.MockClient{
				MockDataMap: map[string]sql_client.MockDataReturn{
					sql_queries.WeeklyDeploymentSql:    {Data: exampleWeeklyDataPoints},
					sql_queries.MonthlyDeploymentSql:   {Data: exampleMonthlyDataPoints},
					sql_queries.QuarterlyDeploymentSql: {Data: exampleQuarterlyDataPoints},
				},
			},
			expectResponse: models.Response{Aggregation: "weekly", DataPoints: exampleWeeklyDataPoints},
			expectError:    "",
		},
		{
			name:   "should return monthly data points from the database",
			params: ServiceParameters{TypeQuery: "df_count", Aggregation: "monthly", Project: "", To: 0, From: 0},
			mockClient: sql_client.MockClient{
				MockDataMap: map[string]sql_client.MockDataReturn{
					sql_queries.WeeklyDeploymentSql:    {Data: exampleWeeklyDataPoints},
					sql_queries.MonthlyDeploymentSql:   {Data: exampleMonthlyDataPoints},
					sql_queries.QuarterlyDeploymentSql: {Data: exampleQuarterlyDataPoints},
				},
			},
			expectResponse: models.Response{Aggregation: "monthly", DataPoints: exampleMonthlyDataPoints},
			expectError:    "",
		},
		{
			name:   "should return quarterly data points from the database",
			params: ServiceParameters{TypeQuery: "df_count", Aggregation: "quarterly", Project: "", To: 0, From: 0},
			mockClient: sql_client.MockClient{
				MockDataMap: map[string]sql_client.MockDataReturn{
					sql_queries.WeeklyDeploymentSql:    {Data: exampleWeeklyDataPoints},
					sql_queries.MonthlyDeploymentSql:   {Data: exampleMonthlyDataPoints},
					sql_queries.QuarterlyDeploymentSql: {Data: exampleQuarterlyDataPoints},
				},
			},
			expectResponse: models.Response{Aggregation: "quarterly", DataPoints: exampleQuarterlyDataPoints},
			expectError:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfCountService := DfCountService{Client: tt.mockClient}
			got, err := dfCountService.ServeRequest(tt.params)

			if err == nil && tt.expectError != "" {
				t.Errorf("expected '%v' got no error", tt.expectError)
			}
			if err != nil && err.Error() != tt.expectError {
				t.Errorf("expected '%v' got '%v'", tt.expectError, err)
			}
			if !reflect.DeepEqual(got, tt.expectResponse) {
				t.Errorf("DfCountService.ServeRequest() = %v, want %v", got, tt.expectResponse)
			}
		})
	}
}