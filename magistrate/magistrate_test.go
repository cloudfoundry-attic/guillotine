package magistrate_test

import (
	"errors"

	"github.com/cloudfoundry-incubator/guillotine/inquisitor/fakes"
	"github.com/cloudfoundry-incubator/guillotine/magistrate"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Magistrate", func() {
	Context("when there are multiple nodes with open connections", func(){
		var iq fakes.FakeInquisitor

		BeforeEach(func(){
			iq = fakes.FakeInquisitor{}
			iq.ReadCsvReturns([][]string{}, nil)
			iq.ConnectionArrayFromCsvReturns([]string{"1","2","0"})
		})

		It("Does not cut the connections", func(){
			headsman := FakeHeadsman{}
			magistrate := magistrate.NewMagistrate(&iq, &headsman)

			magistrate.DeliberateAndExecute()
			Expect(headsman.Chopped).To(BeTrue())
		})
	})

	Context("when there is only one node with open connections", func(){
		var iq fakes.FakeInquisitor

		BeforeEach(func(){
			iq = fakes.FakeInquisitor{}
			iq.ReadCsvReturns([][]string{}, nil)
			iq.ConnectionArrayFromCsvReturns([]string{"1","0","0"})
		})

		It("Does not cut the connections", func(){
			headsman := FakeHeadsman{}
			magistrate := magistrate.NewMagistrate(&iq, &headsman)

			magistrate.DeliberateAndExecute()
			Expect(headsman.Chopped).To(BeFalse())
		})
	})

	Context("When the HAProxy is unreachable or gives a bad response", func(){
		var iq fakes.FakeInquisitor

		BeforeEach(func(){
			iq = fakes.FakeInquisitor{}
			iq.ReadCsvReturns(nil, errors.New("something went wrong"))
		})

		It("Cuts the connections", func(){
			headsman := FakeHeadsman{}
			magistrate := magistrate.NewMagistrate(&iq, &headsman)

			magistrate.DeliberateAndExecute()
			Expect(headsman.Chopped).To(BeTrue())
		})
	})
})

type FakeHeadsman struct{
	Chopped bool
}
func (f *FakeHeadsman) Chop() {
	f.Chopped = true
}
