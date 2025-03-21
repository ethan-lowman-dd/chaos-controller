// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023 Datadog, Inc.

package calculations_test

import (
	. "github.com/DataDog/chaos-controller/grpc/calculations"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("get mapping from alterationConfigToQueryPercent to Alterations using FlattenAlterationMap", func() {
	var (
		altCfgs                        []AlterationConfiguration
		alterationConfigToQueryPercent map[AlterationConfiguration]QueryPercent
		alterations                    []AlterationConfiguration
	)

	altCfgs = []AlterationConfiguration{
		{
			ErrorToReturn:    "CANCELED",
			OverrideToReturn: "",
		},
		{
			ErrorToReturn:    "ALREADY_EXISTS",
			OverrideToReturn: "",
		},
		{
			ErrorToReturn:    "",
			OverrideToReturn: "{}",
		},
	}

	BeforeEach(func() {
		alterationConfigToQueryPercent = make(map[AlterationConfiguration]QueryPercent)
	})

	Context("with one alteration", func() {
		It("should create 15 (out of 100 potential) slots with alteration configurations", func() {
			alterationConfigToQueryPercent[altCfgs[0]] = QueryPercent(15)

			alterations = FlattenAlterationMap(alterationConfigToQueryPercent)

			Expect(alterations).To(HaveLen(15))

			altCfg := alterations[0]
			Expect(altCfg.ErrorToReturn).To(Equal("CANCELED"))
			Expect(altCfg.OverrideToReturn).To(BeEmpty())

			altCfg = alterations[14]
			Expect(altCfg.ErrorToReturn).To(Equal("CANCELED"))
			Expect(altCfg.OverrideToReturn).To(BeEmpty())
		})
	})

	Context("with three alterations that add up to less than 100 percent", func() {
		It("should create 70 (out of 100 potential) slots with alteration configurations", func() {
			alterationConfigToQueryPercent[altCfgs[0]] = QueryPercent(15)
			alterationConfigToQueryPercent[altCfgs[1]] = QueryPercent(20)
			alterationConfigToQueryPercent[altCfgs[2]] = QueryPercent(35)

			alterations = FlattenAlterationMap(alterationConfigToQueryPercent)

			By("having 70 entries", func() {
				Expect(alterations).To(HaveLen(70))
			})

			By("having three alteration types each with the right number of configurations", func() {
				canceled_cnt := 0
				already_exists_cnt := 0
				empty_cnt := 0

				for _, altCfg := range alterations {
					switch altCfg.ErrorToReturn {
					case "CANCELED":
						Expect(altCfg.OverrideToReturn).To(Equal(""))
						canceled_cnt++
					case "ALREADY_EXISTS":
						Expect(altCfg.OverrideToReturn).To(Equal(""))
						already_exists_cnt++
					default:
						Expect(altCfg.OverrideToReturn).To(Equal("{}"))
						Expect(altCfg.ErrorToReturn).To(Equal(""))
						empty_cnt++
					}
				}

				Expect(canceled_cnt).To(Equal(15))
				Expect(already_exists_cnt).To(Equal(20))
				Expect(empty_cnt).To(Equal(35))
			})
		})
	})

	Context("with three alterations that add to 100", func() {
		It("should create 100 (out of 100 potential) slots with alteration configurations", func() {
			alterationConfigToQueryPercent[altCfgs[0]] = QueryPercent(40)
			alterationConfigToQueryPercent[altCfgs[1]] = QueryPercent(40)
			alterationConfigToQueryPercent[altCfgs[2]] = QueryPercent(20)

			alterations = FlattenAlterationMap(alterationConfigToQueryPercent)

			By("having 100 entries", func() {
				Expect(alterations).To(HaveLen(100))
			})

			By("having three alteration types each with the right number of configurations", func() {
				canceled_cnt := 0
				already_exists_cnt := 0
				empty_cnt := 0

				for _, altCfg := range alterations {
					switch altCfg.ErrorToReturn {
					case "CANCELED":
						Expect(altCfg.OverrideToReturn).To(Equal(""))
						canceled_cnt++
					case "ALREADY_EXISTS":
						Expect(altCfg.OverrideToReturn).To(Equal(""))
						already_exists_cnt++
					default:
						Expect(altCfg.OverrideToReturn).To(Equal("{}"))
						Expect(altCfg.ErrorToReturn).To(Equal(""))
						empty_cnt++
					}
				}

				Expect(canceled_cnt).To(Equal(40))
				Expect(already_exists_cnt).To(Equal(40))
				Expect(empty_cnt).To(Equal(20))
			})
		})
	})
})
