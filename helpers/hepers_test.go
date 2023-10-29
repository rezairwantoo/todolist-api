package helpers

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestConvTimeExpired(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		g := gomega.NewWithT(t)
		now, _ := GetNow()
		g.Expect(now).ToNot(gomega.BeZero())
	})
}
