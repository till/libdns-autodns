package autodns_test

import (
	"context"
	"os"
	"testing"

	"github.com/libdns/autodns"
)

func TestProvider(t *testing.T) {
	if os.Getenv("AUTODNS_USERNAME") == "" || os.Getenv("AUTODNS_PASSWORD") == "" {
		t.Skip()
	}

	provider := autodns.NewWithDefaults(os.Getenv("AUTODNS_USERNAME"), os.Getenv("AUTODNS_PASSWORD"))

	t.Run("GetRecords", func(t *testing.T) {
		if os.Getenv("TEST_ZONE") == "" {
			t.Skip()
		}

		records, err := provider.GetRecords(context.TODO(), os.Getenv("TEST_ZONE"))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if len(records) == 0 {
			t.Fatalf("expected at least one record: %#v", records)
		}
	})

}
