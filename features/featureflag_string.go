// Code generated by "stringer -type=FeatureFlag"; DO NOT EDIT.

package features

import "strconv"

const _FeatureFlag_name = "unusedUseAIAIssuerURLAllowTLS02ChallengesReusePendingAuthzCountCertificatesExactIPv6FirstAllowRenewalFirstRLWildcardDomainsEnforceChallengeDisableTLSSNIRevalidationCancelCTSubmissions"

var _FeatureFlag_index = [...]uint8{0, 6, 21, 41, 58, 80, 89, 108, 123, 146, 164, 183}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return "FeatureFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
