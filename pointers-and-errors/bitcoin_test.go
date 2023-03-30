package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {
	t.Parallel()

	type args struct {
		amount Bitcoin
	}

	testCases := []struct {
		name    string
		args    args
		want    Bitcoin
		wantStr string
	}{
		{
			name:    "Deposit 10",
			args:    args{amount: Bitcoin(10)},
			want:    Bitcoin(10),
			wantStr: "10 BTC",
		},
		{
			name:    "Deposit 0",
			args:    args{amount: Bitcoin(0)},
			want:    Bitcoin(0),
			wantStr: "0 BTC",
		},
		{
			name:    "Deposit -5",
			args:    args{amount: Bitcoin(-5)},
			want:    Bitcoin(-5),
			wantStr: "-5 BTC",
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				wallet := Wallet{}
				wallet.Deposit(tc.args.amount)

				assertBalance(t, wallet, tc.want)
			},
		)
	}

	t.Run(
		"withdraw with insufficient funds",
		func(t *testing.T) {
			wallet := Wallet{balance: Bitcoin(20)}
			withdrawAmount := Bitcoin(100)

			err := wallet.Withdraw(withdrawAmount)

			assertError(t, err, ErrInsufficientFunds)

			assertBalance(t, wallet, 20)
		},
	)

	t.Run(
		"withdraw with sufficient funds",
		func(t *testing.T) {
			wallet := Wallet{balance: Bitcoin(20)}
			withdrawAmount := Bitcoin(10)

			err := wallet.Withdraw(withdrawAmount)

			assertNoError(t, err)

			assertBalance(t, wallet, 10)
		},
	)
}

var assertBalance = func(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

var assertError = func(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		// Call Fatal() so that the test stops here
		t.Fatal("wanted error but got none")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q but wanted %q", got.Error(), want.Error())
	}
}

var assertNoError = func(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("got error but expected none %q", got.Error())
	}
}
