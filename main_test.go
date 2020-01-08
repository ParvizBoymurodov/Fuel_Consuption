package main

import "testing"

func Test_fuelConsuption(t *testing.T) {
	test := [] struct {
		name       string
		consuption int
		fuel       int
		want       int
	}{
		{name: "How many kilometers is enough fuel", consuption: 20, fuel: 9, want: 44},
	}
	for _, testes := range test {
		got := fuelConsuption(testes.consuption, testes.fuel)
		if got != testes.want {
			t.Error("For fuel consuption  test:", "got:", got, "want:", testes.want)
		}

	}
}
