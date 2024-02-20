package query

import (
	"bytes"
	"reflect"
	"testing"
	"zbackend/data/models"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test read file string",
			input:    "test read file string",
			expected: "test read file string",
		},
		{
			name:     "empty file",
			input:    "",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			buffer.WriteString(tc.input)
			content, err := readFile(&buffer)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, string(content))
		})
	}
}

func TestReadFilePanic(t *testing.T) {
	assert.Panics(t, func() {
		readFile(nil)
	})
}

func Test_parseDonations(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []models.TransactionWithDonationObject
		wantErr bool
	}{
		{
			name: "test parse donations",
			args: args{
				bytes: []byte(`[
					{
						"id": null,
						"type": null,
						"refundedAmount": null,
						"donation": {
							"id": null,
							"firstName": null,
							"lastName": null,
							"createdAtUtc": null,
							"amount": null,
							"thankYouComment": null,
							"isAnonymous": null,
							"companyName": null,
							"__typename": null
						},
						"__typename": null
					}]`),
			},
			want: []models.TransactionWithDonationObject{
				{
					Id:             nil,
					Type:           nil,
					RefundedAmount: nil,
					Donation:       &models.DonationObject{},
					Typename:       nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDonations(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDonations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDonations() = %v, want %v", got, tt.want)
			}

			sliced := sliceDonationsPaginated(got, 1, 0)
			if !reflect.DeepEqual(sliced, tt.want) {
				t.Errorf("parseDonations() = %v, want %v", sliced, tt.want)
			}
		})
	}
}
