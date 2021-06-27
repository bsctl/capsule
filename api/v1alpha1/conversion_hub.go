// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/clastix/capsule/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (t *Tenant) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Tenant)

	println(dst)

	return nil
}

func (t *Tenant) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Tenant)

	println(src)

	return nil
}
