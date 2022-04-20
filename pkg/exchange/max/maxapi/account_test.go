package max

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountService_GetAccountsRequest(t *testing.T) {
	key, secret, ok := integrationTestConfigured(t, "MAX")
	if !ok {
		t.SkipNow()
	}

	ctx := context.Background()

	client := NewRestClient(ProductionAPIURL)
	client.Auth(key, secret)

	req := client.AccountService.NewGetAccountsRequest()
	accounts, err := req.Do(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, accounts)
	assert.NotEmpty(t, accounts)

	t.Logf("accounts: %+v", accounts)
}

func TestAccountService_GetAccountRequest(t *testing.T) {
	key, secret, ok := integrationTestConfigured(t, "MAX")
	if !ok {
		t.SkipNow()
	}

	ctx := context.Background()

	client := NewRestClient(ProductionAPIURL)
	client.Auth(key, secret)

	req := client.AccountService.NewGetAccountRequest()
	req.Currency("twd")
	account, err := req.Do(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	t.Logf("account: %+v", account)

	req2 := client.AccountService.NewGetAccountRequest()
	req2.Currency("usdt")
	account, err = req.Do(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	t.Logf("account: %+v", account)
}
