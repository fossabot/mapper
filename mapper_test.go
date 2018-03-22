package mapper_test

import (
	"testing"

	"github.com/davidsbond/mapper"
	"github.com/stretchr/testify/assert"
)

type (
	Source struct {
		AField string            `map:"Target:A;Target2:A"`
		BField int               `map:"Target:B"`
		CField bool              `map:"Target:C"`
		DField float32           `map:"Target:D"`
		EField float64           `map:"Target:E"`
		FField map[string]string `map:"Target:F"`
		GField string
	}

	Target struct {
		A string
		B int
		C bool
		D float32
		E float64
		F map[string]string
	}
)

func ExampleMap() {
	source := Source{
		AField: "some data",
	}

	target := Target{}

	mapper.Map(source, &target)
	// Output:
}

func BenchmarkMapper_Map(b *testing.B) {
	in := Source{
		AField: "test",
		BField: 1,
		CField: true,
		DField: 1.1,
		EField: 1.2,
		FField: map[string]string{
			"test": "test",
		},
	}

	out := Target{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mapper.Map(in, &out)
	}
}

func TestMapper_Map(t *testing.T) {

	tt := []struct {
		Source        Source
		Target        *Target
		ExpectedError string
	}{
		{
			Source: Source{
				AField: "test",
				BField: 1,
				CField: true,
				DField: 1.1,
				EField: 1.2,
				FField: map[string]string{
					"test": "test",
				},
			},

			Target: &Target{},
		},
	}

	for _, tc := range tt {
		if err := mapper.Map(tc.Source, tc.Target); err != nil && err.Error() != tc.ExpectedError {
			assert.Fail(t, err.Error())
		}

		assert.Equal(t, tc.Source.AField, tc.Target.A)
		assert.Equal(t, tc.Source.BField, tc.Target.B)
		assert.Equal(t, tc.Source.CField, tc.Target.C)
		assert.Equal(t, tc.Source.DField, tc.Target.D)
		assert.Equal(t, tc.Source.EField, tc.Target.E)

		for key, value := range tc.Source.FField {
			act := tc.Target.F[key]

			assert.Equal(t, value, act)
		}
	}
}
